package connectors

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func DockerClient() *client.Client {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	return dockerClient
}

func GetImageList(dockerClient *client.Client) []types.ImageSummary {
	dockerImages, err := dockerClient.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	return dockerImages
}

func GetContainersList(dockerClient *client.Client, filter filters.Args) []types.Container {
	containerOptions := types.ContainerListOptions{Filters: filter, All: true}
	containerList, err := dockerClient.ContainerList(context.Background(), containerOptions)
	if err != nil {
		panic(err)
	}
	return containerList
}

func GetContainerById(id string, dockerClient *client.Client) *types.Container {
	filter := filters.NewArgs(filters.KeyValuePair{"id", id})
	containerOptions := types.ContainerListOptions{Filters: filter, All: true}
	containerList, err := dockerClient.ContainerList(context.Background(), containerOptions)
	if err != nil {
		panic(err)
	}
	return &containerList[0]
}
