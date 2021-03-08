package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
		containerJSON, _ := cli.ContainerInspect(ctx, container.ID)
		fmt.Println(containerJSON.GraphDriver.Data)
		for _, m := range containerJSON.Mounts {
			s, _ := json.MarshalIndent(m, "", "\t")
			fmt.Println(container.ID, "挂载点: ", string(s))
		}
	}

}
