package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/DarthPestilane/easytcp"
	"github.com/DarthPestilane/easytcp/internal/examples/fixture"
	"github.com/DarthPestilane/easytcp/internal/examples/tcp/custom_packet/common"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	easytcp.SetLogger(log)

	s := easytcp.NewServer(&easytcp.ServerOption{
		// specify codec and packer
		Codec:  &easytcp.JsonCodec{},
		Packer: &common.CustomPacker{},
	})

	s.AddRoute("json01-req", handler, fixture.RecoverMiddleware(log), logMiddleware)

	go func() {
		if err := s.Run(fixture.ServerAddr); err != nil {
			log.Errorf("serve err: %s", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	if err := s.Stop(); err != nil {
		log.Errorf("server stopped err: %s", err)
	}
}

func handler(ctx easytcp.Context) { _ = "STUB: not implemented"; return }

func logMiddleware(next easytcp.HandlerFunc) easytcp.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(easytcp.HandlerFunc)
}
