package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

func main() {
	ctx := context.Background()
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
		return
	}
// make sure u do docker pull ubuntu 
	config := &container.Config{
		Image: "ubuntu",
		Cmd:   []string{"echo","hi"},
	}

	container, err := client.ContainerCreate(ctx, config, nil, nil, &ocispec.Platform{
		OS: "Linux",
	}, "casual")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.ContainerStart(ctx, container.ID, types.ContainerStartOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	statusCh, errCh := client.ContainerWait(ctx, container.ID, "") //not-running is also good option
	select {
	case err := <-errCh:
		if err != nil {
			fmt.Println(err)
			return
		}
	case status := <-statusCh:
		fmt.Println("Container exited with status:", status)
	}
}
