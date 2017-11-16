// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// encapsulate webhook as a windwos service
package main

import (
	"log"

	"github.com/kardianos/service"
)

var logger service.Logger

type webhook struct{}

func (w *webhook) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go w.run()
	return nil
}
func (w *webhook) run() {
	main_as_run()
}
func (w *webhook) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "webhook",
		DisplayName: "webhook as a windows service",
		Description: "webhook as a windows service",
	}

	wh := &webhook{}
	s, err := service.New(wh, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
