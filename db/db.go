package db

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	logger "github.com/sirupsen/logrus"
	"time"
)

type Client struct {
	Redis *redis.Client
}

func NewClient(redisUrl string, db int) (*Client, error) {
	opts := &redis.Options{
		Addr:         redisUrl,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
		DB:           db,
	}

	client := redis.NewClient(opts)

	_, err := client.Ping().Result()
	if err != nil {
		logger.WithFields(CouldntConnectRedis(err)).Error(err.Error())
		return nil, err
	}

	return &Client{Redis: client}, nil
}

func (c Client) Add(key, value string) error {
	err := c.Redis.Set(key, value, 0).Err()
	if err != nil {
		logger.WithFields(CouldntSetKeyValueData(err)).Error(err.Error())
		return err
	}
	return nil
}

func (c Client) Get(key string) (string, error) {
	val, err := c.Redis.Get(key).Result()
	if err == redis.Nil {
		return "", errors.New(fmt.Sprintf("%s key doesn't exist", key))
	} else if err != nil {
		logger.WithFields(CouldntReadKeyValueData(err)).Error(err.Error())
		return "", err
	}

	return val, nil
}
