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
	"image"
	"sync"
	"sync/atomic"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/unstablebuild/rune-go-sdk/term"
	"github.com/unstablebuild/rune-go-sdk/tui"
	"github.com/unstablebuild/tcell/v3"
	"unstable.build/rune/internal/cell"
	"unstable.build/rune/internal/term/gui/font"
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

// GUI implements a graphical TUI runtime as an alternative runtime to what
// the tui packages provides.
type GUI struct {
	ctx               context.Context
	cancelCtx         func()
	mu                sync.Locker
	fontManager       *font.Manager
	updateChan        chan term.Event
	handler           tui.Handler
	writer            *cell.BufferWriter
	mouse             *mouse
	input             *input
	drag              *dragPoller
	theme             string
	bgOpacity         float64
	fgOpacity         float64
	defaultWidth      int
	defaultHeight     int
	explicitSize      bool
	startPositionX    int
	startPositionY    int
	printFPS          bool
	bgBlurRadius      int
	enableTransparent bool
	enableLigatures   bool
	forceFullRepaint  bool
	renderOffset      image.Point
	cursorAttributes  term.Attributes
	defaultAttr       term.Attributes
	renderer          *renderer

	originalColorValues map[tcell.Color]int32
	originalValuesColor map[int32]tcell.Color
	initialTheme        string
	colorThemes         map[string]Theme

	cursor struct {
		pos   term.Coordinates
		style term.CursorStyle
		show  bool
	}

	pendingEvents []term.Event
	needsDraw     bool
	needsRender   bool
	width         int
	height        int
	lastPositionX int
	lastPositionY int
	iteration     int64
	deviceScale   float64

	links linkScanner

	// echoLikely arms the once-per-tick echo wait. It is learned, not
	// configured: an interrupt pending at tick entry right after a
	// single-key tick means the focused handler echoes asynchronously
	// (a terminal); a timed-out wait disarms it, so handlers that
	// update synchronously (the editor) never pay the wait.
	echoLikely  bool
	prevTickKey bool

	interruptPending atomic.Bool
	// processWindowClosed turns a pending window close request into
	// events for the handler. WithCloseRequestEvent installs it; it
	// defaults to a no-op, leaving ebiten's default behavior in place.
	processWindowClosed func() []term.Event
	closingHandled      bool
	closeOnce           sync.Once
}
