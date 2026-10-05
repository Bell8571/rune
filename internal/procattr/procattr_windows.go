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

//go:build windows

// Package procattr builds syscall.SysProcAttr values that stay compilable
// on platforms whose process model does not have POSIX session fields.
package procattr

import "syscall"

// ProcessGroup is a no-op attribute on Windows. POSIX process groups are
// not part of syscall.SysProcAttr there.
func ProcessGroup() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{}
}

// Session is a no-op attribute on Windows. Setsid and Setctty do not exist
// on syscall.SysProcAttr there.
func Session() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{}
}

// SessionFlags is a no-op attribute on Windows.
func SessionFlags(_, _ bool) *syscall.SysProcAttr {
	return &syscall.SysProcAttr{}
}

// LeadsGroup is always false on Windows: SysProcAttr cannot express a
// POSIX process-group leader.
func LeadsGroup(*syscall.SysProcAttr) bool {
	return false
}
