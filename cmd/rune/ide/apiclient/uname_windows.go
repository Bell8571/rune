//go:build windows

// Copyright (C) 2017-2026 The Rune Authors
// SPDX-License-Identifier: GPL-3.0-or-later

package apiclient

import (
	"os"
	"runtime"
)

func uname() (sysinfo, error) {
	host, _ := os.Hostname()
	return sysinfo{
		Name:    "Windows",
		Node:    host,
		Release: os.Getenv("OS"),
		Version: os.Getenv("PROCESSOR_IDENTIFIER"),
		Machine: runtime.GOARCH,
		OS:      runtime.GOOS,
	}, nil
}
