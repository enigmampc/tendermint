//go:build !secretcli && linux && !muslc && !darwin
// +build !secretcli,linux,!muslc,!darwin

package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lrandom_api
import "C"
