package packet

import (
	"time"
)

// DCC protocol-defined values for reference.
// const (
//	BitOnePartMinDuration  = 55 * time.Microsecond
//	BitOnePartMaxDuration  = 61 * time.Microsecond
//	BitZeroPartMinDuration = 95 * time.Microsecond
//	BitZeroPartMaxDuration = 9900 * time.Microsecond
//	PacketSeparationMin    = 5 * time.Millisecond
//	PacketSeparationMax    = 30 * time.Millisecond
//	PreambleBitsMin        = 14
//)

// Some customizable DCC-related variables.
var (
	BitOnePartDuration  = 55 * time.Microsecond
	BitZeroPartDuration = 100 * time.Microsecond
	Separation          = 15 * time.Millisecond
	PreambleBits        = 16
)

// HeadlightCompatMode controls if one bit in the speed instruction is
// reserved for headlight. This reduces speed steps from 32 to 16 steps.
const HeadlightCompatMode = false

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

// Packet represents the unit of information that can be sent to the DCC
// devices in the system. Packet implements the DCC protocol for converting
// the information into DCC-encoded 1 and 0s.
type Packet struct {
	driver  Driver
	address byte
	data    []byte
	ecc     byte

	// encoded holds an int64 (time.Duration) for each
	// bit in a packet. It is an efficient representation
	// to save extra function calls and IFs when sending
	encoded []time.Duration
}

// NewPacket returns a new generic DCC Packet.
func NewPacket(d Driver, addr byte, data []byte) *Packet {
	ecc := addr
	for _, i := range data {
		ecc ^= i
	}

	return &Packet{
		driver:  d,
		address: addr,
		data:    data,
		ecc:     ecc,
	}
}

// NewBaselinePacket returns a new generic baseline packet.
// Baseline packets are different because they use a 128 address
// space. Therefore the address is forced to start with bit 0.
func NewBaselinePacket(d Driver, addr byte, data []byte) *Packet {
	addr &= 0x7F // 0b01111111 last 7 bits

	return NewPacket(d, addr, data)
}

// NewSpeedAndDirectionPacket returns a new baseline DCC packet with speed and
// direction information.
func NewSpeedAndDirectionPacket(d Driver, addr byte, speed byte, dir byte) *Packet {
	addr &= 0x7F // 0b 0111 1111
	if HeadlightCompatMode {
		// 4 lower bytes
		speed &= 0x0F
	} else {
		// 5 lower bytes
		speed = (speed << 7 >> 3) | (speed >> 1) // nolint:gomnd
	}

	dirB := 0x1 & dir << 5 // nolint:gomnd
	// 0b01DCSSSS
	data := (1 << 6) | dirB | speed // nolint:gomnd

	return &Packet{
		driver:  d,
		address: addr,
		data:    []byte{data},
		ecc:     addr ^ data,
	}
}

// NewFunctionGroupOnePacket returns an advanced DCC packet which allows to
// control FL,F1-F4 functions. FL is usually associated to the headlights.
func NewFunctionGroupOnePacket(d Driver, addr byte, fl, fl1, fl2, fl3, fl4 bool) *Packet {
	var data byte
	var fln, fl1n, fl2n, fl3n, fl4n byte = 0, 0, 0, 0, 0
	if fl {
		fln = 1 << 4 // nolint:gomnd
	}
	if fl1 {
		fl1n = 1
	}
	if fl2 {
		fl2n = 1 << 1 // nolint:gomnd
	}
	if fl3 {
		fl3n = 1 << 2 // nolint:gomnd
	}
	if fl4 {
		fl4n = 1 << 3 // nolint:gomnd
	}

	data = (1 << 7) | fln | fl1n | fl2n | fl3n | fl4n // nolint:gomnd

	return &Packet{
		driver:  d,
		address: addr,
		data:    []byte{data},
		ecc:     addr ^ data,
	}
}

// NewFunctionGroupTwoPacket returns two advanced DCC packet which allows to
// control F5-F12 functions.
func NewFunctionGroupTwoPacket(d Driver, addr byte, f5, f6, f7, f8, f9, f10, f11, f12 bool) (*Packet, *Packet) {
	var data0, data1 byte
	var f5n, f6n, f7n, f8n, f9n, f10n, f11n, f12n byte = 0, 0, 0, 0, 0, 0, 0, 0

	if f5 {
		f5n = 1
	}

	if f6 {
		f6n = 1 << 1 // nolint:gomnd
	}

	if f7 {
		f7n = 1 << 2 // nolint:gomnd
	}

	if f8 {
		f8n = 1 << 3 // nolint:gomnd
	}

	if f9 {
		f9n = 1
	}

	if f10 {
		f10n = 1 << 1 // nolint:gomnd
	}

	if f11 {
		f11n = 1 << 2 // nolint:gomnd
	}

	if f12 {
		f12n = 1 << 3 // nolint:gomnd
	}

	data0 = (1 << 7) | (1 << 5) | (1 << 4) | f5n | f6n | f7n | f8n // nolint:gomnd
	data1 = (1 << 7) | (1 << 5) | f9n | f10n | f11n | f12n         // nolint:gomnd

	return &Packet{
			driver:  d,
			address: addr,
			data:    []byte{data0},
			ecc:     addr ^ data0,
		},
		&Packet{
			driver:  d,
			address: addr,
			data:    []byte{data1},
			ecc:     addr ^ data1,
		}
}

// NewFunctionExpansionPacket returns two advanced 2-byte DCC packet which allows to
// control F13-F28 functions.
func NewFunctionExpansionPacket(d Driver, addr byte, f13, f14, f15, f16, f17, f18,
	f19, f20, f21, f22, f23, f24, f25, f26, f27, f28 bool,
) (*Packet, *Packet) {
	var dataA0, dataA1, dataB0, dataB1 byte
	var f13n, f14n, f15n, f16n, f17n, f18n,
		f19n, f20n, f21n, f22n, f23n,
		f24n, f25n, f26n, f27n, f28n byte = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0

	if f13 {
		f13n = 1
	}

	if f14 {
		f14n = 1 << 1 // nolint:gomnd
	}

	if f15 {
		f15n = 1 << 2 // nolint:gomnd
	}

	if f16 {
		f16n = 1 << 3 // nolint:gomnd
	}

	if f17 {
		f17n = 1 << 4 // nolint:gomnd
	}

	if f18 {
		f18n = 1 << 5 // nolint:gomnd
	}

	if f19 {
		f19n = 1 << 6 // nolint:gomnd
	}

	if f20 {
		f20n = 1 << 7 // nolint:gomnd
	}

	if f21 {
		f21n = 1
	}

	if f22 {
		f22n = 1 << 1 // nolint:gomnd
	}

	if f23 {
		f23n = 1 << 2 // nolint:gomnd
	}

	if f24 {
		f24n = 1 << 3 // nolint:gomnd
	}

	if f25 {
		f25n = 1 << 4 // nolint:gomnd
	}

	if f26 {
		f26n = 1 << 5 // nolint:gomnd
	}

	if f27 {
		f27n = 1 << 6 // nolint:gomnd
	}

	if f28 {
		f28n = 1 << 7 // nolint:gomnd
	}

	dataA0 = 0b11011110
	dataA1 = f13n | f14n | f15n | f16n | f17n | f18n | f19n | f20n

	dataB0 = 0b11011111
	dataB1 = f21n | f22n | f23n | f24n | f25n | f26n | f27n | f28n

	return &Packet{
			driver:  d,
			address: addr,
			data:    []byte{dataA0, dataA1},
			ecc:     addr ^ dataA0 ^ dataA1,
		},
		&Packet{
			driver:  d,
			address: addr,
			data:    []byte{dataB0, dataB1},
			ecc:     addr ^ dataB0 ^ dataB1,
		}
}

// NewBroadcastResetPacket returns a new broadcast baseline DCC packet which
// makes the decoders erase their volatile memory and return to power up
// state. This stops all locomotives at non-zero speed.
func NewBroadcastResetPacket(d Driver) *Packet {
	return &Packet{
		driver:  d,
		address: 0,
		data:    []byte{0},
		ecc:     0,
	}
}

// NewBroadcastIdlePacket returns a new broadcast baseline DCC packet
// on which decoders perform no action.
func NewBroadcastIdlePacket(d Driver) *Packet {
	return &Packet{
		driver:  d,
		address: 0xFF, // nolint:gomnd
		data:    []byte{0},
		ecc:     0xFF, // nolint:gomnd
	}
}

// NewBroadcastStopPacket returns a new broadcast baseline DCC packet which
// tells the decoders to stop all locomotives. If softStop is false, an
// emergency stop will happen by cutting power off the engine.
func NewBroadcastStopPacket(d Driver, dir byte, softStop bool, ignoreDir bool) *Packet {
	var speed byte
	if !softStop {
		speed = 1
	}

	if ignoreDir {
		speed |= 1 << 4 // nolint:gomnd
	}

	dirB := 0x1 & dir // nolint:gomnd

	data := (1 << 6) | (dirB << 5) | speed // nolint:gomnd

	return &Packet{
		driver:  d,
		address: 0x0, // nolint:gomnd
		data:    []byte{data},
		ecc:     0x0 ^ data, // nolint:gomnd
	}
}

// delayPoll causes a active delay for the specified time
// by actively polling the clock. Unfortunately, for latencies
// under 100us, it is not possible to sleep reliably with
// syscall.Nanosleep().
func delayPoll(now time.Time, d time.Duration) {
	for {
		if time.Since(now) > d {
			return
		}
	}
}

// PacketPause performs a pause by sleeping
// during the Separation time.
func (p *Packet) PacketPause() {
	// Not really needed
	p.driver.Low()
	time.Sleep(Separation) // nolint:forbidigo
	p.driver.High()
}

// Send encodes and sends a packet using the Driver associated to it.
func (p *Packet) Send() {
	if p.driver == nil {
		panic("No driver set")
	}
	if p.encoded == nil {
		p.build()
	}

	// The way p.encoded is we reduce function calls and ifs
	for _, b := range p.encoded {
		p.driver.Low()
		now := time.Now()
		delayPoll(now, b)
		p.driver.High()
		now = time.Now()
		delayPoll(now, b)
	}
}

// Length returns the length of the DCC-encoded representation
// of a packet.
func (p *Packet) Length() int {
	l := 0
	l += PreambleBits // Preamble
	l++               // Packet start
	l += 8            // Address byte
	for i := 0; i < len(p.data); i++ {
		l++    // Data start
		l += 8 // Data byte
	}
	l++    // ECC start
	l += 8 // ECC byte
	l++    // Packet end

	return l
}

// By prebuilding packages we ensure more consistent Send() times.
func (p *Packet) build() {
	enc := make([]time.Duration, 0, p.Length())

	unpackByte := func(b byte) []time.Duration {
		bs := make([]time.Duration, 8) // nolint:gomnd
		for i := uint8(0); i < 8; i++ {
			bit := (b >> (7 - i)) & 0x1 // nolint:gomnd
			if bit == 0 {
				bs[i] = BitZeroPartDuration
			} else {
				bs[i] = BitOnePartDuration
			}
		}

		return bs
	}

	// Preamble
	for i := 0; i < PreambleBits; i++ {
		enc = append(enc, BitOnePartDuration)
	}

	// Packet start bit
	enc = append(enc, BitZeroPartDuration)

	// Address
	enc = append(enc, unpackByte(p.address)...)

	// Data
	for _, d := range p.data {
		enc = append(enc, BitZeroPartDuration) // Data start
		enc = append(enc, unpackByte(d)...)    // Data
	}

	// ECC
	enc = append(enc, BitZeroPartDuration) // ECC start
	enc = append(enc, unpackByte(p.ecc)...)

	// Packet end
	enc = append(enc, BitOnePartDuration)

	p.encoded = enc
}

func (p *Packet) String() string {
	if p.encoded == nil {
		p.build()
	}
	var str string
	for _, b := range p.encoded {
		switch {
		case b == BitZeroPartDuration:
			str += "0"
		case b == BitOnePartDuration:
			str += "1"
		default:
			panic("bad encoding")
		}
	}

	return str
}
