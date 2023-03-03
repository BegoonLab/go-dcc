package module

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/alexbegoon/go-dcc/internal/packet"
	"github.com/gosimple/slug"
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

// Railway represents a DCC device, usually a PWM driver for servos and lights.
type Railway struct {
	Name    string `json:"name"`
	Address uint8  `json:"address"`
	Enabled bool   `json:"enabled"`
	RoutesData

	mux sync.Mutex

	Driver Driver

	flPacket *packet.Packet

	sendGroupTwo     bool
	fGroupTwoPacket0 *packet.Packet
	fGroupTwoPacket1 *packet.Packet

	sendExpansion0    bool
	fExpansionPacket0 *packet.Packet

	sendExpansion1    bool
	fExpansionPacket1 *packet.Packet
}

// Route represents a Route of a Railway Module.
type Route struct {
	Name string `json:"name"`

	F0  bool `json:"f0"`
	F1  bool `json:"f1"`
	F2  bool `json:"f2"`
	F3  bool `json:"f3"`
	F4  bool `json:"f4"`
	F5  bool `json:"f5"`
	F6  bool `json:"f6"`
	F7  bool `json:"f7"`
	F8  bool `json:"f8"`
	F9  bool `json:"f9"`
	F10 bool `json:"f10"`
	F11 bool `json:"f11"`
	F12 bool `json:"f12"`
	F13 bool `json:"f13"`
	F14 bool `json:"f14"`
	F15 bool `json:"f15"`
	F16 bool `json:"f16"`
	F17 bool `json:"f17"`
	F18 bool `json:"f18"`
	F19 bool `json:"f19"`
	F20 bool `json:"f20"`
	F21 bool `json:"f21"`
	F22 bool `json:"f22"`
	F23 bool `json:"f23"`
	F24 bool `json:"f24"`
	F25 bool `json:"f25"`
	F26 bool `json:"f26"`
	F27 bool `json:"f27"`
	F28 bool `json:"f28"`
	// TODO: Reserved
	F29 bool `json:"f29"`
	F30 bool `json:"f30"`
	F31 bool `json:"f31"`
}

type RoutesData struct {
	Routes      map[string]*Route `json:"routes"`
	ActiveRoute string            `json:"active_route"`
}

func (r *Railway) SendPackets() {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.flPacket.Send()
	if r.sendGroupTwo {
		r.fGroupTwoPacket0.Send()
		r.fGroupTwoPacket1.Send()
	}
	if r.sendExpansion0 {
		r.fExpansionPacket0.Send()
	}
	if r.sendExpansion1 {
		r.fExpansionPacket1.Send()
	}
}

// Apply makes any changes to the Railway Module properties
func (r *Railway) Apply() {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.flPacket = nil
	r.fGroupTwoPacket0 = nil
	r.fGroupTwoPacket1 = nil
	r.fExpansionPacket0 = nil
	r.fExpansionPacket1 = nil

	route, ok := r.Routes[r.ActiveRoute]
	if !ok {
		return
	}

	if !r.sendGroupTwo {
		r.sendGroupTwo = route.F5 || route.F6 || route.F7 || route.F8 || route.F9 || route.F10 || route.F11 || route.F12
	}
	if !r.sendExpansion0 {
		r.sendExpansion0 = route.F13 || route.F14 || route.F15 || route.F16 || route.F17 || route.F18 || route.F19 || route.F20
	}
	if !r.sendExpansion1 {
		r.sendExpansion1 = route.F21 || route.F22 || route.F23 || route.F24 || route.F25 || route.F26 || route.F27 || route.F28
	}
	if r.flPacket == nil {
		r.flPacket = packet.NewFunctionGroupOnePacket(r.Driver,
			r.Address, route.F0, route.F1, route.F2, route.F3, route.F4)
	}
	if r.fGroupTwoPacket0 == nil {
		r.fGroupTwoPacket0, r.fGroupTwoPacket1 = packet.NewFunctionGroupTwoPacket(r.Driver,
			r.Address, route.F5, route.F6, route.F7, route.F8, route.F9, route.F10, route.F11, route.F12)
	}
	if r.fExpansionPacket0 == nil {
		r.fExpansionPacket0, r.fExpansionPacket1 = packet.NewFunctionExpansionPacket(r.Driver,
			r.Address, route.F13, route.F14, route.F15, route.F16, route.F17, route.F18, route.F19, route.F20,
			route.F21, route.F22, route.F23, route.F24, route.F25, route.F26, route.F27, route.F28)
	}
}

func (r *Railway) FetchRoutes() error {
	r.mux.Lock()
	defer r.mux.Unlock()
	dataFile, err := r.getDataFile()
	if err != nil {
		return err
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&r.RoutesData)

	if err != nil {
		return err
	}

	defer func(dataFile *os.File) {
		err := dataFile.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}(dataFile)

	return nil
}

func (r *Railway) PersistsRoutes() error {
	r.mux.Lock()
	defer r.mux.Unlock()
	dataFile, err := r.getDataFile()
	if err != nil {
		return err
	}

	dataEncoder := gob.NewEncoder(dataFile)
	err = dataEncoder.Encode(&r.RoutesData)
	if err != nil {
		return err
	}

	defer func(dataFile *os.File) {
		err := dataFile.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}(dataFile)

	return nil
}

func (r *Railway) getDataFile() (*os.File, error) {
	filename := r.getDataFilename()

	return os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0o644)
}

func (r *Railway) getDataFilename() string {
	return fmt.Sprintf("./%s.bin", slug.Make(r.Name+"-railway-module"))
}
