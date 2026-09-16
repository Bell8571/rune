//go:build !windows

// Copyright (C) 2017-2026 The Rune Authors
// SPDX-License-Identifier: GPL-3.0-or-later

package debug

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func StartPProfOnSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR1)
	go CapturePanicReport(func() {
		<-ch
		signal.Stop(ch)
		if _, err := StartPProfHTTP("127.0.0.1:0"); err != nil {
			log.Errorf("StartPProfOnSignal: %v", err)
		}
	})
}
