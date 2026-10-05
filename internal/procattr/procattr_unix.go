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

// Package procattr builds syscall.SysProcAttr values that stay compilable
// on platforms whose process model does not have POSIX session fields.
package procattr

import "syscall"

// ProcessGroup starts the child as the leader of a new process group.
func ProcessGroup() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{Setpgid: true}
}

// Session makes the child a session leader with a controlling terminal.
func Session() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{Setsid: true, Setctty: true}
}

// SessionFlags applies the requested session and controlling-tty flags.
func SessionFlags(setsid, setctty bool) *syscall.SysProcAttr {
	return &syscall.SysProcAttr{Setsid: setsid, Setctty: setctty}
}

// LeadsGroup reports whether attr asks the child to head a new process group.
func LeadsGroup(attr *syscall.SysProcAttr) bool {
	return attr != nil && attr.Setpgid && attr.Pgid == 0
}
