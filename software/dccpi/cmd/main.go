// Package main implements the DCC protocol for controlling model trains.
// It can support a number of different encoders, which are in charge of
// translating DCC packages into electrical signals. By default, a Raspberry
// Pi driver is provided.
//
// The implementation follows the S-91 Electrical Standard (http://www.nmra.org/sites/default/files/standards/sandrp/pdf/s-9.1_electrical_standards_2006.pdf), the S-92 DCC Communications Standard (http://www.nmra.org/sites/default/files/s-92-2004-07.pdf) and the S-9.2.1 Extended Packet Formats for Digital Command Control standard (http://www.nmra.org/sites/default/files/s-9.2.1_2012_07.pdf).
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexbegoon/go-dcc/software/dccpi/internal/config"
	"github.com/alexbegoon/go-dcc/software/dccpi/internal/controller"
	"github.com/alexbegoon/go-dcc/software/dccpi/internal/driver/dccpi"
	"github.com/alexbegoon/go-dcc/software/dccpi/internal/driver/dummy"
	"github.com/alexbegoon/go-dcc/software/dccpi/internal/logger"
	"github.com/alexbegoon/go-dcc/software/dccpi/internal/server"
	"github.com/stianeikeland/go-rpio/v4"
	"go.uber.org/zap"
)

var (
	configFlag    string
	signalPinFlag uint
	brakePinFlag  uint
)

var DefaultConfigPath = "../conf/initial.json"

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

type repl struct {
	signalCh chan os.Signal
	doneCh   chan struct{}
	ctrl     *controller.Controller
	driver   Driver
	log      *zap.Logger
}

func init() {
	//usr, _ := user.Current()
	//DefaultConfigPath = filepath.Join(usr.HomeDir, ".dccpi")
	flag.Usage = func() {
		fmt.Fprint(os.Stdout, "Usage: dccpi [options]")
		fmt.Fprint(os.Stdout, "Options:")
		flag.PrintDefaults()

	}

	flag.StringVar(&configFlag, "config", DefaultConfigPath,
		"location of a dccpi configuration file")
	flag.UintVar(&signalPinFlag, "signalPin", uint(dccpi.SignalGPIO),
		"GPIO Pin to use for the DCC signal")
	flag.UintVar(&brakePinFlag, "brakePin", uint(dccpi.BrakeGPIO),
		"GPIO Pin to use for the Brake signal (cuts power from tracks")
	flag.Parse()
}

func main() {
	l, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	cfg, err := config.LoadConfig(configFlag)
	if err != nil {
		l.Error("Cannot load configuration. Using empty one.")
		cfg = &config.Config{}
	}

	dccpi.BrakeGPIO = rpio.Pin(brakePinFlag)
	dccpi.SignalGPIO = rpio.Pin(signalPinFlag)

	var dpi Driver
	dpi, err = dccpi.NewDCCPi()
	if err != nil {
		l.Error("DCCPi is not available. Using dummy driver.")
		dpi = &dummy.DCCDummy{
			Log: l,
		}
	}

	ctrl := controller.NewControllerWithConfig(dpi, cfg)
	ctrl.Log = l

	r := &repl{
		signalCh: make(chan os.Signal, 1),
		doneCh:   make(chan struct{}),
		ctrl:     ctrl,
		driver:   dpi,
		log:      l,
	}

	signal.Notify(r.signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGSTOP)

	go func() {
		<-r.signalCh
		r.shutdown()
	}()

	r.ctrl.Start()
	s := server.Serve(r.ctrl)

	defer func(logger *zap.Logger) {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second) // nolint:gomnd
		logger.Error(err.Error())
		err = s.Shutdown(ctx)
		logger.Error("Unable to shutdown server", zap.Error(err))
		_ = logger.Sync()
		cancel()
		os.Exit(0)
	}(l)

	<-r.doneCh
}

func (r *repl) shutdown() {
	r.ctrl.Stop()
	r.log.Info("Tracks were powered off")
	close(r.doneCh)
}
