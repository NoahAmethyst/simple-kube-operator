package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
)

// gracefulShutdown waits for termination syscalls and doing clean up operations after received it
func gracefulShutdown(_ context.Context, _ *grpc.Server) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGSTOP, syscall.SIGKILL, syscall.SIGHUP)
	go func() {
		sig := <-signalChannel
		defer close(signalChannel)
		log.Info().Msgf("receive signal:%+v,graceful shutdown", sig)

		log.Info().Msgf("graceful shutdown done")

	}()
}
