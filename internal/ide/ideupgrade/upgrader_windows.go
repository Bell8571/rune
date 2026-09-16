//go:build windows

// Copyright (C) 2017-2026 The Rune Authors
// SPDX-License-Identifier: GPL-3.0-or-later

package ideupgrade

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
)

func defaultInstallRoot() string {
	return filepath.Join(os.Getenv("LOCALAPPDATA"), "Rune")
}

func defaultAppName() string          { return "rune" }
func defaultCLIBinaryRelPath() string { return "rune.exe" }

type windowsPlatformOps struct{ httpClient *http.Client }

func newPlatformOps(client *http.Client) platformOps {
	return windowsPlatformOps{httpClient: client}
}

func (w windowsPlatformOps) Download(ctx context.Context, url, dest string, progress func(n, total int64)) error {
	return downloadOver(ctx, w.httpClient, url, dest, progress)
}
func (windowsPlatformOps) VerifySHA256(path, want string) error { return verifySHA256(path, want) }
func (windowsPlatformOps) MountDMG(context.Context, string) (string, func() error, error) {
	return "", nil, ErrUnsupported
}
func (windowsPlatformOps) AssessGatekeeper(context.Context, string) error { return ErrUnsupported }
func (windowsPlatformOps) VerifyCodesign(context.Context, string) error   { return ErrUnsupported }
func (windowsPlatformOps) Ditto(ctx context.Context, src, dst string) error {
	return copyDir(src, dst)
}
func (windowsPlatformOps) ExtractTarGz(context.Context, string, string, func(int64, int64)) error {
	return ErrUnsupported
}
func (windowsPlatformOps) Symlink(target, linkPath string) error {
	return os.Symlink(target, linkPath)
}
func (windowsPlatformOps) RenameAtomic(src, dst string) error { return os.Rename(src, dst) }
func (windowsPlatformOps) RemoveAll(path string) error        { return os.RemoveAll(path) }
func (windowsPlatformOps) FreeSpace(string) (uint64, error)   { return 1 << 40, nil }
