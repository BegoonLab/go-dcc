package controller

import (
	"encoding/json"
	"sync"

	"github.com/alexbegoon/go-dcc/software/dccpi/internal/pb/build/go/controller"
	"google.golang.org/protobuf/proto"

	"go.uber.org/zap"

	"github.com/alexbegoon/go-dcc/software/dccpi/internal/config"
	"github.com/alexbegoon/go-dcc/software/dccpi/internal/locomotive"
	"github.com/alexbegoon/go-dcc/software/dccpi/internal/packet"
)

// CommandRepeat specifies how many times a single
// packet is sent.
var CommandRepeat = 10

// CommandMaxQueue specifies how many commands can
// queue before sending a new command blocks
// the sender
var CommandMaxQueue = 3

// Driver can be implemented by any module to allow using go-dcc
// on different platforms. dcc.Driver modules are in charge of
// producing an electrical signal output (i.e. on a GPIO Pin)
type Driver interface {
	// Low sets the output to low state.
	Low()
	// High sets the output to high.
	High()
	// TracksOn turns the tracks on. The exact procedure is left to the
	// implementation, but tracks should be ready to receive packets from
	// this point.
	TracksOn()
	// TracksOff disables the tracks. The exact procedure is left to the
	// implementation, but tracks should not carry any power and all
	// trains should stop after calling it.
	TracksOff()
}

// Controller represents a DCC Control Station. The
// controller keeps tracks of the DCC Locomotives and
// is in charge of sending DCC packets continuously to
// the tracks.
type Controller struct {
	locomotives map[string]*locomotive.Locomotive
	mux         sync.RWMutex
	driver      Driver

	started    bool
	doneCh     chan bool
	shutdownCh chan bool
	commandCh  chan *packet.Packet
	Log        *zap.Logger
}

type ControllerJSON struct {
	Locomotives map[string]*locomotive.Locomotive `json:"locomotives"`
	Started     bool                              `json:"started"`
	Reboot      bool                              `json:"reboot"`
	Poweroff    bool                              `json:"poweroff"`
}

// NewController builds a Controller.
func NewController(d Driver) *Controller {
	d.TracksOff()

	return &Controller{
		driver:      d,
		locomotives: make(map[string]*locomotive.Locomotive),
		doneCh:      make(chan bool),
		shutdownCh:  make(chan bool),
		commandCh:   make(chan *packet.Packet, CommandMaxQueue),
	}
}

// ToJSON
// Deprecated: use ToProto
func (c *Controller) ToJSON() []byte {
	cj := ControllerJSON{
		Locomotives: c.locomotives,
		Started:     c.started,
	}

	d, err := json.Marshal(cj)
	if err != nil {
		c.Log.Error("Cannot marshal", zap.Error(err))
	}

	return d
}

func (c *Controller) ToProto() []byte {
	locos := make(map[string]*controller.Locomotive, len(c.locomotives))

	for _, l := range c.locomotives {
		locos[l.Name] = &controller.Locomotive{
			Name:      l.Name,
			Address:   uint32(l.Address),
			Speed:     uint32(l.Speed),
			Direction: controller.Locomotive_Direction(l.Direction),
			Enabled:   l.Enabled,
			Fl:        l.Fl,
			F1:        l.F1,
			F2:        l.F2,
			F3:        l.F3,
			F4:        l.F4,
			F5:        l.F5,
			F6:        l.F6,
			F7:        l.F7,
			F8:        l.F8,
			F9:        l.F9,
			F10:       l.F10,
			F11:       l.F11,
			F12:       l.F12,
			F13:       l.F13,
			F14:       l.F14,
			F15:       l.F15,
			F16:       l.F16,
			F17:       l.F17,
			F18:       l.F18,
			F19:       l.F19,
			F20:       l.F20,
			F21:       l.F21,
			F22:       l.F22,
			F23:       l.F23,
			F24:       l.F24,
			F25:       l.F25,
			F26:       l.F26,
			F27:       l.F27,
			F28:       l.F28,
		}
	}

	msg := &controller.Controller{
		Locomotives: locos,
		Started:     c.IsStarted(),
		Reboot:      false,
		Poweroff:    false,
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		c.Log.Error("Cannot marshal protobuf", zap.Error(err))
	}

	return data
}

// NewControllerWithConfig builds a new Controller using the
// given configuration.
func NewControllerWithConfig(d Driver, cfg *config.Config) *Controller {
	c := NewController(d)

	for _, loco := range cfg.Locomotives {
		c.AddLoco(&locomotive.Locomotive{
			Name:      loco.Name,
			Address:   loco.Address,
			Speed:     loco.Speed,
			Direction: loco.Direction,
			Fl:        loco.Fl,
		})
	}

	return c
}

// AddLoco adds a DCC device to the controller. The device
// will start receiving packets if the controller is running.
func (c *Controller) AddLoco(l *locomotive.Locomotive) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.locomotives[l.Name] = l
}

// RmLoco removes a DCC device from the controller. There
// will be no longer packets sent to it.
func (c *Controller) RmLoco(l *locomotive.Locomotive) {
	c.mux.Lock()
	defer c.mux.Unlock()
	delete(c.locomotives, l.Name)
}

// GetLoco retrieves a DCC device by its Name. The boolean is
// true if the Locomotive was found.
func (c *Controller) GetLoco(n string) (*locomotive.Locomotive, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	l, ok := c.locomotives[n]

	return l, ok
}

// Locos returns a list of all registered Locomotives.
func (c *Controller) Locos() []*locomotive.Locomotive {
	c.mux.RLock()
	defer c.mux.RUnlock()
	locos := make([]*locomotive.Locomotive, 0, len(c.locomotives))
	for _, l := range c.locomotives {
		locos = append(locos, l)
	}

	return locos
}

// Command allows to send a custom Packet to the tracks.
// The packet will be sent CommandRepeat times.
func (c *Controller) Command(p *packet.Packet) {
	c.commandCh <- p
}

// Start starts the controller: powers on the tracks
// and starts sending packets on them.
func (c *Controller) Start() {
	c.driver.TracksOn()
	go c.run()
	c.started = true
}

// Stop shuts down the controller by stopping to send
// packets and removing power from the tracks.
func (c *Controller) Stop() {
	if c.started {
		c.shutdownCh <- true
		<-c.doneCh
		c.started = false
	}
}

func (c *Controller) IsStarted() bool {
	return c.started
}

func (c *Controller) run() {
	idle := packet.NewBroadcastIdlePacket(c.driver)
	stop := packet.NewBroadcastStopPacket(c.driver, byte(locomotive.Forward), false, true)
	for {
		select {
		case <-c.shutdownCh:
			for i := 0; i < CommandRepeat; i++ {
				stop.Send()
			}
			c.driver.TracksOff()
			c.doneCh <- true

			return
		case p := <-c.commandCh:
			for i := 0; i < CommandRepeat; i++ {
				p.Send()
			}
		default:
			c.mux.RLock()
			{
				// Idle and retry later
				if len(c.locomotives) == 0 {
					c.commandCh <- idle
					c.mux.RUnlock()

					break // from the select
				}
				for _, loco := range c.locomotives {
					if !loco.Enabled {
						continue
					}
					for i := 0; i < CommandRepeat; i++ {
						loco.SendPackets(c.driver)
					}
				}
				idle.Send()
			}
			c.mux.RUnlock()
			idle.PacketPause()
		}
	}
}
