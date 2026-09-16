//go:build !windows

// Copyright (C) 2017-2026 The Rune Authors
// SPDX-License-Identifier: GPL-3.0-or-later

package apiclient

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/unix"
)

func uname() (sysinfo, error) {
	var data unix.Utsname
	if err := unix.Uname(&data); err != nil {
		return sysinfo{}, fmt.Errorf("uname: %v", err)
	}
	return sysinfo{
		Name:    utsnameToString(data.Sysname),
		Node:    utsnameToString(data.Nodename),
		Release: utsnameToString(data.Release),
		Version: utsnameToString(data.Version),
		Machine: utsnameToString(data.Machine),
		OS:      runtime.GOOS,
	}, nil
}
