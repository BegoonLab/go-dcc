package controller

import (
	"testing"
	"time"

	"github.com/alexbegoon/go-dcc/internal/config"
	"github.com/alexbegoon/go-dcc/internal/driver/dummy"
	"github.com/alexbegoon/go-dcc/internal/locomotive"
	"github.com/alexbegoon/go-dcc/internal/packet"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewController(t *testing.T) {
	cfg, err := config.LoadConfig("../../test/config.json")
	assert.NoError(t, err)
	c := NewControllerWithConfig(&dummy.DCCDummy{Log: zap.NewExample()}, cfg)
	c.Stop()
}

func TestAddLoco(t *testing.T) {
	c := NewController(&dummy.DCCDummy{Log: zap.NewExample()})
	assert.NotEmpty(t, c)
	c.AddLoco(&locomotive.Locomotive{})
}

func TestRmLoco(t *testing.T) {
	c := NewController(&dummy.DCCDummy{Log: zap.NewExample()})
	c.AddLoco(&locomotive.Locomotive{Name: "abc"})
	c.RmLoco(&locomotive.Locomotive{Name: "abc"})

	if _, ok := c.GetLoco("abc"); ok {
		t.Error("loco should have been deleted")
	}
}

func TestLocos(t *testing.T) {
	c := NewController(&dummy.DCCDummy{Log: zap.NewExample()})
	c.AddLoco(&locomotive.Locomotive{Name: "abc"})
	l := c.Locos()
	if len(l) != 1 || l[0].Name != "abc" {
		t.Error("Locos() does not work")
	}
}

func TestCommand(t *testing.T) {
	d := &dummy.DCCDummy{Log: zap.NewExample()}
	c := NewController(d)
	assert.NotEmpty(t, c)
	p := packet.NewBroadcastIdlePacket(d)
	c.Command(p)
	c.Start()
	time.Sleep(250 * time.Millisecond) // nolint:forbidigo
	c.Stop()
}

func TestStart(t *testing.T) {
	c := NewController(&dummy.DCCDummy{Log: zap.NewExample()})
	assert.NotEmpty(t, c)
	c.Start()
	time.Sleep(1 * time.Second) // nolint:forbidigo
	c.AddLoco(&locomotive.Locomotive{Name: "abc", Address: 10})
	time.Sleep(1 * time.Second) // nolint:forbidigo
	c.Stop()
}
