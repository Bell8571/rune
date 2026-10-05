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

package smartware

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestOverlayRunDrafts(t *testing.T) {
	core := CoreDir()
	cli := filepath.Join(core, "src", "cli.ts")
	if _, err := os.Stat(cli); err != nil {
		t.Skip("smartware-core cli missing")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	out, err := OverlayRun(ctx, "jacket.fixture.inspect", "summarize readiness")
	if err != nil {
		t.Fatal(err)
	}
	var rec struct {
		JacketID string `json:"jacketId"`
		Accepted bool   `json:"accepted"`
		Executed bool   `json:"executed"`
	}
	if err := json.Unmarshal(out, &rec); err != nil {
		t.Fatal(err)
	}
	if rec.JacketID != "jacket.fixture.inspect" {
		t.Fatalf("jacket %s", rec.JacketID)
	}
	if rec.Accepted || rec.Executed {
		t.Fatalf("draft must not accept or execute: %+v", rec)
	}
}
