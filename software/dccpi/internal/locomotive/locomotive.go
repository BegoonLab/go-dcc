package locomotive

import (
	"fmt"
	"sync"

	"github.com/alexbegoon/go-dcc/internal/packet"
)

// Direction constants.
const (
	Backward Direction = 0
	Forward  Direction = 1
)

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

// Direction represents the locomotive direction and can be
// Forward or Backward.
type Direction byte

// Locomotive represents a DCC device, usually a locomotive.
// Locomotives are represented by their name and address and
// include certain properties like speed, direction or FL.
// Each locomotive produces two packets: one speed and direction
// packet and one Function Group One packet.
type Locomotive struct {
	Name      string    `json:"name"`
	Address   uint8     `json:"address"`
	Speed     uint8     `json:"speed"`
	Direction Direction `json:"direction"`
	Enabled   bool      `json:"enabled"`
	Fl        bool      `json:"fl"`
	F1        bool      `json:"f1"`
	F2        bool      `json:"f2"`
	F3        bool      `json:"f3"`
	F4        bool      `json:"f4"`
	F5        bool      `json:"f5"`
	F6        bool      `json:"f6"`
	F7        bool      `json:"f7"`
	F8        bool      `json:"f8"`
	F9        bool      `json:"f9"`
	F10       bool      `json:"f10"`
	F11       bool      `json:"f11"`
	F12       bool      `json:"f12"`
	F13       bool      `json:"f13"`
	F14       bool      `json:"f14"`
	F15       bool      `json:"f15"`
	F16       bool      `json:"f16"`
	F17       bool      `json:"f17"`
	F18       bool      `json:"f18"`
	F19       bool      `json:"f19"`
	F20       bool      `json:"f20"`
	F21       bool      `json:"f21"`
	F22       bool      `json:"f22"`
	F23       bool      `json:"f23"`
	F24       bool      `json:"f24"`
	F25       bool      `json:"f25"`
	F26       bool      `json:"f26"`
	F27       bool      `json:"f27"`
	F28       bool      `json:"f28"`

	mux sync.Mutex

	speedPacket *packet.Packet
	flPacket    *packet.Packet

	sendGroupTwo     bool
	fGroupTwoPacket0 *packet.Packet
	fGroupTwoPacket1 *packet.Packet

	sendExpansion0    bool
	fExpansionPacket0 *packet.Packet

	sendExpansion1    bool
	fExpansionPacket1 *packet.Packet
}

func (l *Locomotive) String() string {
	var dir string
	fl, f1, f2, f3, f4 := "off", "off", "off", "off", "off"
	if l.Direction == Forward {
		dir = ">"
	} else {
		dir = "<"
	}
	if l.Fl {
		fl = "on"
	}
	if l.F1 {
		f1 = "on"
	}
	if l.F2 {
		f2 = "on"
	}
	if l.F3 {
		f3 = "on"
	}
	if l.F4 {
		f4 = "on"
	}

	return fmt.Sprintf("%s:%d |%d%s| |%s| |%s|%s|%s|%s|",
		l.Name,
		l.Address,
		l.Speed,
		dir,
		fl,
		f1,
		f2,
		f3,
		f4)
}

func (l *Locomotive) SendPackets(d Driver) {
	l.mux.Lock()
	{
		if !l.sendGroupTwo {
			l.sendGroupTwo = l.F5 || l.F6 || l.F7 || l.F8 || l.F9 || l.F10 || l.F11 || l.F12
		}
		if !l.sendExpansion0 {
			l.sendExpansion0 = l.F13 || l.F14 || l.F15 || l.F16 || l.F17 || l.F18 || l.F19 || l.F20
		}
		if !l.sendExpansion1 {
			l.sendExpansion1 = l.F21 || l.F22 || l.F23 || l.F24 || l.F25 || l.F26 || l.F27 || l.F28
		}
		if l.speedPacket == nil {
			l.speedPacket = packet.NewSpeedAndDirectionPacket(d,
				l.Address, l.Speed, byte(l.Direction))
		}
		if l.flPacket == nil {
			l.flPacket = packet.NewFunctionGroupOnePacket(d,
				l.Address, l.Fl, l.F1, l.F2, l.F3, l.F4)
		}
		if l.fGroupTwoPacket0 == nil {
			l.fGroupTwoPacket0, l.fGroupTwoPacket1 = packet.NewFunctionGroupTwoPacket(d,
				l.Address, l.F5, l.F6, l.F7, l.F8, l.F9, l.F10, l.F11, l.F12)
		}
		if l.fExpansionPacket0 == nil {
			l.fExpansionPacket0, l.fExpansionPacket1 = packet.NewFunctionExpansionPacket(d,
				l.Address, l.F13, l.F14, l.F15, l.F16, l.F17, l.F18, l.F19, l.F20,
				l.F21, l.F22, l.F23, l.F24, l.F25, l.F26, l.F27, l.F28)
		}
		l.speedPacket.Send()
		l.flPacket.Send()
		if l.sendGroupTwo {
			l.fGroupTwoPacket0.Send()
			l.fGroupTwoPacket1.Send()
		}
		if l.sendExpansion0 {
			l.fExpansionPacket0.Send()
		}
		if l.sendExpansion1 {
			l.fExpansionPacket1.Send()
		}
	}
	l.mux.Unlock()
}

// Apply makes any changes to the Locomotive's properties
// to be reflected in the packets generated for it and,
// therefore, alter the behavior of the device on the tracks.
func (l *Locomotive) Apply() {
	l.mux.Lock()
	{
		l.speedPacket = nil
		l.flPacket = nil
		l.fGroupTwoPacket0 = nil
		l.fGroupTwoPacket1 = nil
		l.fExpansionPacket0 = nil
		l.fExpansionPacket1 = nil
	}
	l.mux.Unlock()
}
