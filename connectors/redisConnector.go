package connectors

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/go-redis/redis"
	"github.com/mati23/binocular/domain"
)

func RedisConnector() redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return *redisClient
}

func RedisPing(redisClient *redis.Client) {
	pong, pongErr := redisClient.Ping().Result()
	fmt.Println("Error: ", pong, pongErr)
}

func SetImage(redisClient redis.Client, dockerImage types.ImageSummary) {
	image := domain.Image{
		Prefix:     "image",
		ID:         dockerImage.ID,
		Repository: dockerImage.RepoDigests[0],
		Tag:        dockerImage.RepoTags[0],
		Created:    strconv.FormatInt(dockerImage.Created, 10),
		Size:       dockerImage.Size,
	}

	jsonObject, err := json.Marshal(image)
	if err != nil {
		panic(err)
	}

	redisClient.Set(image.ImageKey(), jsonObject, 0)
	fmt.Println(image.ImageKey())
}

func SetImageList(redisClient redis.Client, dockerImages []types.ImageSummary) {
	for _, dockerImage := range dockerImages {
		SetImage(redisClient, dockerImage)
	}
}
