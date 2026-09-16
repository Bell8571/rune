//go:build !windows

// Copyright (C) 2017-2026 The Rune Authors
// SPDX-License-Identifier: GPL-3.0-or-later

package syntax

import "github.com/ebitengine/purego"

const (
	dlNow    = purego.RTLD_NOW
	dlGlobal = purego.RTLD_GLOBAL
)

func sysDlopen(path string, flags int) (uintptr, error) {
	return purego.Dlopen(path, flags)
}

func sysDlsym(lib uintptr, name string) (uintptr, error) {
	return purego.Dlsym(lib, name)
}

func sysDlclose(lib uintptr) error {
	return purego.Dlclose(lib)
}
