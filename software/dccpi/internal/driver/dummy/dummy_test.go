package dummy

import (
	"testing"
	"time"
)

func TestGuessBuffer(t *testing.T) {
	d := DCCDummy{}
	d.TracksOn()
	d.Low()
	d.High()
	d.Low()
	d.High()
	if GuessBuffer.String() != "11" {
		t.Error("it should guess 1")
	}
	d.TracksOff()
	d.TracksOn()
	d.Low()
	time.Sleep(5000 * time.Microsecond) // nolint:forbidigo
	d.High()
	time.Sleep(5000 * time.Microsecond) // nolint:forbidigo
	d.Low()
	time.Sleep(5000 * time.Microsecond) // nolint:forbidigo
	d.High()
	d.Low()
	time.Sleep(time.Second) // nolint:forbidigo
	d.High()

	if GuessBuffer.String() != "00\n" {
		t.Error("it should guess 0")
	}
}
