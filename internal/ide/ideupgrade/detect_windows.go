//go:build windows

// Copyright (C) 2017-2026 The Rune Authors
// SPDX-License-Identifier: GPL-3.0-or-later

package ideupgrade

func detectRunningInstallOS(real string, cfg Config) (detectedInstall, error) {
	return detectedInstall{ExecutablePath: real}, &ErrUpgradeNotSupported{
		Path:   real,
		Reason: errNoAppBundle,
	}
}
