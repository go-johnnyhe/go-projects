package myVersion

import "runtime"

func Version() string {
	return runtime.Version()
}
