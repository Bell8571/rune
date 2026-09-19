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

package debug

import (
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// StartPProfHTTP serves pprof on addr and returns the bound address.
func StartPProfHTTP(addr string) (string, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return "", fmt.Errorf("pprof listen %q: %w", addr, err)
	}
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	bound := ln.Addr().String()
	log.Infof("pprof server listening on http://%s/debug/pprof/", bound)
	go CapturePanicReport(func() {
		if err := http.Serve(ln, nil); err != http.ErrServerClosed {
			log.Errorf("pprof serve: %v", err)
		}
	})
	return bound, nil
}
