//go:build ignore

package main

import (
	"context"
	"fmt"
	"io"
	"os"

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

	imageName := "ghcr.io/decenter-ai/compute:v1.5.5"

	reader, err := client.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Image pulled successfully!")

	// TODO: make sure u do docker pull , write go script to do docker pull
	config := &container.Config{
		Image: imageName,
		// Cmd:   []string{},
	}

	container, err := client.ContainerCreate(ctx, config, nil, nil, &ocispec.Platform{
		OS: "Linux",
	}, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("containerID", container.ID)

	err = client.ContainerStart(ctx, container.ID, types.ContainerStartOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	statusCh, errCh := client.ContainerWait(ctx, container.ID, "") // not-running is also good option
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
