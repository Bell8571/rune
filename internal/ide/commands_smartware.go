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

package ide

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"strings"

	"unstable.build/rune/internal/smartware"
)

func init() {
	exCommands["smartwareoverlay"] = commandAll{
		man: textapiManual(
			"Draft a Smartware overlay run against an installed jacket. Does not execute the jacket. Requires SMARTWARE_CORE.",
			"<jacket-id> <job...>",
		),
		handler: (*ex).smartwareOverlay,
	}
}

func textapiManual(summary, synopsis string) interface{} {
	return struct {
		Summary  string
		Synopsis string
	}{Summary: summary, Synopsis: synopsis}
}
