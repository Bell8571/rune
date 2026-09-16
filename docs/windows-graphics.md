# Windows graphics

Rune's GPU path on Windows is Direct3D 11 (via the
`github.com/unstablebuild/ebiten` fork) plus Desktop Window Manager for
the frame.

| OS | GPU API | compositor extras |
| --- | --- | --- |
| Linux | OpenGL | X11/Wayland window class |
| macOS | Metal | Cocoa glass bar, blur |
| Windows | Direct3D 11 | per-monitor DPI v2, dark title bar, rounded corners, Mica/Acrylic |

`internal/term/gui/wingfx` owns the compositor extras. It is a no-op on
non-Windows builds.

## Build

```bash
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -tags=ebitensinglethread -o rune.exe ./cmd/rune
```

Ebiten on Windows does not need CGO.

## Remaining upstream work

Two Unstable Build forks still need Windows-shaped fixes before
`go build` is green on `GOOS=windows`:

1. `github.com/unstablebuild/ebiten/v2` — `internal/ui/ui_glfw.go` calls
   `glfw.InitHint` (Cocoa menubar) on every OS. `InitHint` is only
   defined in the Unix GLFW binding. Guard that call with `runtime.GOOS == "darwin"`.
2. `github.com/unstablebuild/tcell/v3` — `console_win.go` is out of date
   with the fork's `screenImpl` / `NewEventKey` signatures.

Those live outside this repository. Once they land, the Rune-side
Windows graphics path in this branch is the rest of the story: fonts,
emoji, DWM backdrop, presets, and the compile stubs for pty gather /
upgrade / uname.

Transparent themes (`gui.window_opacity`, `gui.window_blur_radius`)
activate Mica on Windows 11 and Acrylic on Windows 10.
