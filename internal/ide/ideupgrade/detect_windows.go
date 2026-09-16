//go:build windows

// Copyright (C) 2017-2026 The Rune Authors
// SPDX-License-Identifier: GPL-3.0-or-later

package ideupgrade

import (
	"fmt"
	"os"
	"path/filepath"
)

func detectRunningInstallOS(real string, cfg Config) (detectedInstall, error) {
	dir := filepath.Dir(real)
	return detectedInstall{
		InstallRoot:      dir,
		AppName:          "rune",
		CLIBinaryRelPath: filepath.Base(real),
		ExecutablePath:   real,
	}, fmt.Errorf("%w: in-app upgrade is not supported on Windows yet", ErrUpgradeNotSupported)
}

var _ = os.ErrNotExist
