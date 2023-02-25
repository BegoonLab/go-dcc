package module

import "sync"

// Railway represents a DCC device, usually a PWM driver for servos and lights.
type Railway struct {
	Name    string `json:"name"`
	Address uint8  `json:"address"`
	Enabled bool   `json:"enabled"`

	mux sync.Mutex
}
