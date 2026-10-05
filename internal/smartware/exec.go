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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// CoreDir is the on-disk smartware-core checkout. Override with SMARTWARE_CORE.
func CoreDir() string {
	if v := strings.TrimSpace(os.Getenv("SMARTWARE_CORE")); v != "" {
		return v
	}
	return `C:\src\smartware-core`
}

// OverlayRun drafts a job against an installed jacket. Does not execute the jacket.
func OverlayRun(ctx context.Context, jacketID, job string) (json.RawMessage, error) {
	return runCLI(ctx, "overlay", "run", "--jacket", jacketID, "--job", job)
}

// OverlayAccept seals a draft run. Does not execute the jacket.
func OverlayAccept(ctx context.Context, runID string) (json.RawMessage, error) {
	return runCLI(ctx, "overlay", "accept", runID)
}

func runCLI(ctx context.Context, args ...string) (json.RawMessage, error) {
	core := CoreDir()
	cli := filepath.Join(core, "src", "cli.ts")
	cmdArgs := append([]string{"--import", "tsx", cli}, args...)
	cmd := exec.CommandContext(ctx, "node", cmdArgs...)
	cmd.Dir = core
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("smartware cli: %w: %s", err, strings.TrimSpace(stderr.String()))
	}
	out := bytes.TrimSpace(stdout.Bytes())
	if !json.Valid(out) {
		return nil, fmt.Errorf("smartware cli: output is not json")
	}
	return json.RawMessage(out), nil
}
