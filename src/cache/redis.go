package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"realtime-chat/src/config"
)

type RedisConfig struct {
	Host string
	Port string
}

var redisConfig RedisConfig
var RedisClient *redis.Client
var PubSubConnection *redis.PubSub

/*
	A single go server instance have a single redis connection pool
	From the redis connection pool, one connection is used for PubSub connection
	That connection is used to subscribe to multiple channels
	A single go-routine is used to listen for messages on all subscribed channels

	The server also maintains a local map of all connections
	When a new connection is established, it is added to the map

	Each room is a channel in redis
	When someone wants to join a room, we add them to redis set and check if the server is already subscribed to that room
	If not, we subscribe to that room using the dedicated PubSub connection
	The same pubsub connection is used to subscribe to multiple channels
*/

func init() {
	redisConfig = RedisConfig{
		Host: config.Config.RedisHost,
		Port: config.Config.RedisPort,
	}
}

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
	})

	// Initialize PubSub connection
	ctx := context.Background()
	PubSubConnection = RedisClient.Subscribe(ctx)

	// Ping Redis to check if connection is established
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		panic(err)
	}
}
