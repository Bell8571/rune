// Copyright (C) 2017-2026 The Rune Authors
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or (at
// your option) any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

//go:build unix

package main

import (
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func ignoredSignal(sig os.Signal) bool {
	// Raised when socket urgent data is ready. It is not a shutdown.
	return sig == syscall.SIGURG
}

func dupStderr(f *os.File) error {
	// syscall.Dup2 isn't defined on linux/arm64 (the kernel only exposes
	// Dup3 there); golang.org/x/sys/unix papers over the difference.
	return unix.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
}
