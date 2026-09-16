//go:build windows

// Copyright (C) 2017-2026 The Rune Authors
// SPDX-License-Identifier: GPL-3.0-or-later

package vte

import (
	"context"

	"github.com/unstablebuild/rune-go-sdk/api/workspaceapi"
)

type ptyGather struct {
	ready chan []byte
	err   error
}

func newPtyGather(context.Context, workspaceapi.File) (*ptyGather, bool) {
	return nil, false
}

func (g *ptyGather) release([]byte) {}
