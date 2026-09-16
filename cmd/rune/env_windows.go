//go:build windows

// Copyright (C) 2017-2026 The Rune Authors
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import "os/exec"

func detachFromTerminal(cmd *exec.Cmd) {}
