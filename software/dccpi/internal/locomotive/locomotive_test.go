package locomotive

// import (
//	"github.com/alexbegoon/go-dcc/internal/driver/dummy"
//	"github.com/alexbegoon/go-dcc/internal/packet"
//	"testing"
//)
//
// func TestApply(t *testing.T) {
//	l := &Locomotive{
//		Name:        "loco",
//		Address:     3,
//		speedPacket: packet.NewBroadcastIdlePacket(&dummy.DCCDummy{}),
//	}
//	l.Apply()
//	if l.speedPacket != nil {
//		t.Error("should have cleared the speed packet")
//	}
//}
//
// func TestString(t *testing.T) {
//	l := &Locomotive{
//		Name:      "loco",
//		Address:   4,
//		Direction: Forward,
//		Speed:     4,
//		Fl:        true,
//		F1:        true,
//		F2:        true,
//		F3:        true,
//		F4:        true,
//	}
//	l.String()
//}
