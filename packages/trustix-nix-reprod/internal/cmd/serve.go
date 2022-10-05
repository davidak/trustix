// Copyright (C) 2022 adisbladis
// Copyright © 2020-2022 The Trustix Authors
//
// SPDX-License-Identifier: GPL-3.0-only

package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/coreos/go-systemd/activation"
	"github.com/nix-community/trustix/packages/trustix-nix-reprod/internal/cron"
	"github.com/nix-community/trustix/packages/trustix-nix-reprod/internal/index"
	"github.com/nix-community/trustix/packages/trustix-nix-reprod/internal/server"
	apiconnect "github.com/nix-community/trustix/packages/trustix-nix-reprod/reprod-api/reprod_apiconnect"
	tclient "github.com/nix-community/trustix/packages/trustix/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var serveListenAddresses []string

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		// config options
		logIndexCronInterval := time.Minute * 10

		err := os.MkdirAll(stateDirectory, 0755)
		if err != nil {
			panic(err)
		}

		db, err := setupDB(stateDirectory)
		if err != nil {
			panic(err)
		}

		client, err := tclient.CreateClient(dialAddress)
		if err != nil {
			panic(err)
		}

		// Start indexing logs
		logIndexCron := cron.NewCronJob(logIndexCronInterval, func() {
			ctx := context.Background()

			err = index.IndexLogs(ctx, db, client)
			if err != nil {
				panic(err)
			}
		})
		defer logIndexCron.Stop()

		apiServer := server.NewAPIServer(db, client)

		errChan := make(chan error)

		createServer := func(lis net.Listener) *http.Server {
			mux := http.NewServeMux()

			mux.Handle(apiconnect.NewReproducibilityAPIHandler(apiServer))

			server := &http.Server{Handler: h2c.NewHandler(mux, &http2.Server{})}

			go func() {
				err := server.Serve(lis)
				if err != nil {
					errChan <- fmt.Errorf("failed to serve: %v", err)
				}
			}()

			return server
		}

		var servers []*http.Server

		// Systemd socket activation
		listeners, err := activation.Listeners()
		if err != nil {
			log.Fatalf("cannot retrieve listeners: %s", err)
		}
		for _, lis := range listeners {
			log.WithFields(log.Fields{
				"address": lis.Addr(),
			}).Info("Using socket activated listener")

			servers = append(servers, createServer(lis))
		}

		for _, addr := range serveListenAddresses {
			u, err := url.Parse(addr)
			if err != nil {
				log.Fatalf("Could not parse url: %v", err)
			}

			family := ""
			host := ""

			switch u.Scheme {
			case "unix":
				family = "unix"
				host = u.Host + u.Path
			case "http":
				family = "tcp"
				host = u.Host
			default:
				log.Fatalf("Socket with scheme '%s' unsupported", u.Scheme)
			}

			lis, err := net.Listen(family, host)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			log.WithFields(log.Fields{
				"address": addr,
			}).Info("Listening to address")

			servers = append(servers, createServer(lis))
		}

		if len(servers) <= 0 {
			log.Fatal("No listeners configured!")
		}

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-quit
			wg := new(sync.WaitGroup)

			log.Info("Received shutdown signal, closing down server gracefully")

			for _, server := range servers {
				server := server
				wg.Add(1)
				go func() {
					defer wg.Done()
					server.Close()
				}()
			}

			wg.Wait()

			log.Info("Done closing down servers")

			close(errChan)
		}()

		for err := range errChan {
			log.Fatal(err)
		}
	},
}

func initServe() {
	serveCommand.Flags().StringSliceVar(&serveListenAddresses, "listen", []string{}, "Listen to address")
}
