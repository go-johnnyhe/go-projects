package versionPack

import "runtime"

func Version() string {
	return runtime.Version()
}
