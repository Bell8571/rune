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

`go.mod` pins the fork commits that contain the Windows fixes. Those
commits are on the upstream default branches and are not tagged yet:

- `github.com/hajimehoshi/ebiten/v2` -> `unstablebuild/ebiten` `v2.7.6-0.20260916232601-6f1d59169775` ([unstablebuild/ebiten#1](https://github.com/unstablebuild/ebiten/pull/1)), one commit after `v2.7.5-ub.32`
- `github.com/unstablebuild/tcell/v3` -> `v3.6.6-0.20260919122710-90d3af1d7f12` ([unstablebuild/tcell#6](https://github.com/unstablebuild/tcell/pull/6)), three commits after `v3.6.5`

`v2.7.5-ub.27` and `tcell` `v3.6.5` do not contain those fixes, so a
Windows GUI build fails in `ui_glfw.go` (`glfw.InitHint`) and
`console_win.go` (`cScreen`).

Tree-sitter is cgo, so the GUI binary needs a C compiler (`CGO_ENABLED=1`)
even though ebiten itself is built with `-tags=ebitensinglethread`.
`github.com/unstablebuild/rune-go-sdk@v0.2.0` also reads
`syscall.SysProcAttr.Setsid` and `Setctty`, which do not exist on
Windows and are not fixed in any published SDK tag. Apply
`scripts/rune-go-sdk-windows.patch` to that module before building.

```bash
go mod download github.com/unstablebuild/rune-go-sdk@v0.2.0
mod="$(go env GOMODCACHE)/github.com/unstablebuild/rune-go-sdk@v0.2.0"
patch -d "$mod" -p1 < scripts/rune-go-sdk-windows.patch
CGO_ENABLED=1 go build -tags=ebitensinglethread -o rune.exe ./cmd/rune
```

Transparent themes (`gui.window_opacity`, `gui.window_blur_radius`)
activate Mica on Windows 11 and Acrylic on Windows 10.
