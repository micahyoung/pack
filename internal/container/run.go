package container

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	dcontainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/pkg/errors"
)

func Run(ctx context.Context, docker client.CommonAPIClient, ctrID string, out, errOut io.Writer, execConfigs ...*types.ExecConfig) error {
	bodyChan, errChan := docker.ContainerWait(ctx, ctrID, dcontainer.WaitConditionNextExit)

	resp, err := docker.ContainerAttach(ctx, ctrID, types.ContainerAttachOptions{
		Stream: true,
		Stdout: true,
		Stderr: true,
	})
	if err != nil {
		return err
	}
	defer resp.Close()

	if err := docker.ContainerStart(ctx, ctrID, types.ContainerStartOptions{}); err != nil {
		return errors.Wrap(err, "container start")
	}

	for _, execConfig := range execConfigs {
		exec, err := docker.ContainerExecCreate(ctx, ctrID, *execConfig)
		if err != nil {
			return errors.Wrap(err, "container exec create")
		}

		if err := docker.ContainerExecStart(ctx, exec.ID, types.ExecStartCheck{}); err != nil {
			return errors.Wrap(err, "container exec start")
		}
	}

	copyErr := make(chan error)
	go func() {
		_, err := stdcopy.StdCopy(out, errOut, resp.Reader)
		defer optionallyCloseWriter(out)
		defer optionallyCloseWriter(errOut)

		copyErr <- err
	}()

	select {
	case body := <-bodyChan:
		if body.StatusCode != 0 {
			return fmt.Errorf("failed with status code: %d", body.StatusCode)
		}
	case err := <-errChan:
		return err
	}

	return <-copyErr
}

func optionallyCloseWriter(writer io.Writer) error {
	if closer, ok := writer.(io.Closer); ok {
		return closer.Close()
	}

	return nil
}
