package memcache

import (
	"bytes"
	"encoding/gob"
	"github.com/bradfitz/gomemcache/memcache"
	"redismemcache/model"
	"time"
)

type Client struct {
	client *memcache.Client
}

func New() (*Client, error) {
	client := memcache.New("127.0.0.1:11211")

	if err := client.Ping(); err != nil {
		return nil, err
	}

	client.Timeout = 100 * time.Millisecond
	client.MaxIdleConns = 100

	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetDevice(id string) (model.Device, error) {
	item, err := c.client.Get(id)
	if err != nil {
		return model.Device{}, err
	}

	b := bytes.NewReader(item.Value)

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

	return c.client.Set(&memcache.Item{
		Key:        d.ID.String(),
		Value:      b.Bytes(),
		Expiration: int32(time.Now().Add(24 * time.Hour).Unix()),
	})
}
