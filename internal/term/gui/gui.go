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

package gui

import (
	"context"
	"errors"
	"fmt"
	"image"
	"sync"
	"sync/atomic"
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	log "github.com/sirupsen/logrus"
	"github.com/unstablebuild/blue/iterator"
	"github.com/unstablebuild/blue/logging"
	"github.com/unstablebuild/rune-go-sdk/term"
	"github.com/unstablebuild/rune-go-sdk/tui"
	"github.com/unstablebuild/tcell/v3"
	"unstable.build/rune/internal/cell"
	"unstable.build/rune/internal/term/gui/drawrect"
	"unstable.build/rune/internal/term/gui/font"
	"unstable.build/rune/internal/term/gui/wingfx"
)

var (
	_ ebiten.Game = (*GUI)(nil)

	// ErrHandlerExited is returned by GUI.Run to indicate that
	// the root tui.Handler exited.
	ErrHandlerExited = errors.New("tui handler exited")
)

const (
	defaultWidth, defaultHeight = 800, 600
	// echoPollInterval is the sleep slice while awaiting a
	// post-keystroke interrupt.
	echoPollInterval = 50 * time.Microsecond
)

// echoWaitBudget bounds the once-per-tick wait for the focused
// handler's asynchronous post-keystroke update (a pty echo), letting it
// render in the keystroke's own frame instead of the next one. It must
// stay well under a frame period: with vsync the present time is
// unchanged as long as Update plus Draw still fit the frame. It is a
// variable so tests can widen it for deterministic timing margins.
var echoWaitBudget = 2 * time.Millisecond
