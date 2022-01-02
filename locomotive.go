package dcc

import (
	"fmt"
	"sync"
)

// Direction constants.
const (
	Backward Direction = 0
	Forward  Direction = 1
)

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

	mux sync.Mutex

	speedPacket      *Packet
	flPacket         *Packet
	fGroupTwoPacket0 *Packet
	fGroupTwoPacket1 *Packet
}

func (l *Locomotive) String() string {
	var dir, fl, f1, f2, f3, f4 string = "", "off", "off", "off", "off", "off"
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

func (l *Locomotive) sendPackets(d Driver) {
	l.mux.Lock()
	{

		if l.speedPacket == nil {
			l.speedPacket = NewSpeedAndDirectionPacket(d,
				l.Address, l.Speed, l.Direction)
		}
		if l.flPacket == nil {
			l.flPacket = NewFunctionGroupOnePacket(d,
				l.Address, l.Fl, l.F1, l.F2, l.F3, l.F4)
		}
		if l.fGroupTwoPacket0 == nil && (l.F5 || l.F6 || l.F7 || l.F8 || l.F9 || l.F10 || l.F11 || l.F12) {
			l.fGroupTwoPacket0, l.fGroupTwoPacket1 = NewFunctionGroupTwoPacket(d,
				l.Address, l.F5, l.F6, l.F7, l.F8, l.F9, l.F10, l.F11, l.F12)
		}
		l.speedPacket.Send()
		l.flPacket.Send()
		if l.fGroupTwoPacket0 != nil {
			l.fGroupTwoPacket0.Send()
			l.fGroupTwoPacket1.Send()
		}
	}
	l.mux.Unlock()
}

// Apply makes any changes to the Locomotive's properties
// to be reflected in the packets generated for it and,
// therefore, alter the behaviour of the device on the tracks.
func (l *Locomotive) Apply() {
	l.mux.Lock()
	{
		l.speedPacket = nil
		l.flPacket = nil
		l.fGroupTwoPacket0 = nil
		l.fGroupTwoPacket1 = nil
	}
	l.mux.Unlock()
}
