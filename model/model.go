package model

import (
	"github.com/google/uuid"
	"math/rand"
)

type Device struct {
	ID         uuid.UUID
	DeviceID   string
	DeviceType int
	Super      bool
	Stuff      []RandomStuff
}

type RandomStuff struct {
	StuffID   string
	StuffType int
}

// GenerateDevice generates random devices
func GenerateDevice(amount int) []Device {
	res := make([]Device, amount)
	for i := 0; i < amount; i++ {
		res[i] = Device{
			ID:         uuid.New(),
			DeviceID:   randStringBytes(10),
			DeviceType: rand.Intn(10),
			Super:      true,
			Stuff: []RandomStuff{
				{
					StuffID:   randStringBytes(10),
					StuffType: rand.Intn(10),
				},
			},
		}
	}
	return res
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
