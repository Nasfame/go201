import (
    "context"
    "fmt"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
)

func main() {
    ctx := context.Background()
    client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        fmt.Println(err)
        return
    }

    config := &container.Config{
        Image: "ubuntu:latest",
        Cmd:    []string{"bash"},
    }

    container, err := client.ContainerCreate(ctx, config, nil, nil, "")
    if err != nil {
        fmt.Println(err)
        return
    }

    err = client.ContainerStart(ctx, container.ID, types.ContainerStartOptions{})
    if err != nil {
        fmt.Println(err)
        return
    }

    statusCh, errCh := client.ContainerWait(ctx, container.ID, container.WaitConditionNotRunning)
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
