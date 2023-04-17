package redis

import (
	"bytes"
	"encoding/gob"
	"github.com/go-redis/redis"
	"redismemcache/model"
	"time"
)

type Client struct {
	Client *redis.Client
}

func New() (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		Password:    "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:          0,
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return &Client{
		Client: client,
	}, nil
}

func (c *Client) GetDevice(key string) (model.Device, error) {
	item, err := c.Client.Get(key).Bytes()
	if err != nil {
		return model.Device{}, err
	}

	b := bytes.NewReader(item)

	var res model.Device

	if err := gob.NewDecoder(b).Decode(&res); err != nil {
		return model.Device{}, err
	}

	return res, nil
}

func (c *Client) SetDevice(d model.Device) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(d); err != nil {
		return err
	}

	return c.Client.Set(d.ID.String(), b.Bytes(), 24*time.Hour).Err()
}

func (c *Client) PublishMessage(channel string, message string) error {
	return c.Client.Publish(channel, message).Err()
}

func (c *Client) SubscribeMessage(channel string) (*redis.PubSub, error) {
	return c.Client.Subscribe(channel), nil
}
