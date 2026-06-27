// /home/krylon/go/src/krylib/errors.go
// -*- mode: go; coding: utf-8; -*-
// Created on 13. 09. 2019 by Benjamin Walkenhorst
// (c) 2019 Benjamin Walkenhorst
// Time-stamp: <2026-06-27 11:52:59 krylon>

package krylib

import "errors"

// ErrNotImplemented indicates that some functionality is not implemented, yet.
var ErrNotImplemented error = errors.New("not implemented") // nolint: golint

// ErrInvalidValue indicates that some value is invalid.
var ErrInvalidValue error = errors.New("invalid value")

// ErrNil indicates a parameter is nil that is not supposed to be.
var ErrNil error = errors.New("argument is nil")

// Local Variables:  //
// compile-command: "go generate && go vet && go build -v -p 16 && gometalinter && go test -v" //
// End: //
