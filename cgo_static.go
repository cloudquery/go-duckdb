//go:build !(duckdb_use_lib || duckdb_from_source) && (darwin || (linux && amd64) || (windows && amd64))

package duckdb

/*
#cgo LDFLAGS: -lduckdb
#cgo darwin,amd64 LDFLAGS: -lc++ -L${SRCDIR}/deps/darwin_amd64
#cgo darwin,arm64 LDFLAGS: -lc++ -L${SRCDIR}/deps/darwin_arm64
#cgo linux,amd64 LDFLAGS: -lstdc++ -lm -ldl -L${SRCDIR}/deps/linux_amd64
#cgo windows,amd64 LDFLAGS: -lstdc++ -lm -ldl -L${SRCDIR}/deps/windows_amd64
#include <duckdb.h>
*/
import "C"
