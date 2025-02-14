package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type TagValue struct {
	Value     interface{} `json:"value"`
	Timestamp time.Time   `json:"timestamp"`
	Quality   int         `json:"quality"`
}

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(host string, port int, password string) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       0,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisCache{
		client: client,
		ctx:    ctx,
	}, nil
}

func (r *RedisCache) Close() error {
	return r.client.Close()
}

func (r *RedisCache) SetTagValue(plcID, tagID int, value interface{}) error {
	tagValue := TagValue{
		Value:     value,
		Timestamp: time.Now(),
		Quality:   100,
	}

	data, err := json.Marshal(tagValue)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("plc:%d:tag:%d", plcID, tagID)
	return r.client.Set(r.ctx, key, data, 0).Err()
}

func (r *RedisCache) GetTagValue(plcID, tagID int) (*TagValue, error) {
	key := fmt.Sprintf("plc:%d:tag:%d", plcID, tagID)

	data, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var tagValue TagValue
	if err := json.Unmarshal([]byte(data), &tagValue); err != nil {
		return nil, err
	}

	return &tagValue, nil
}
