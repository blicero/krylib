// /home/krylon/go/src/github.com/blicero/krylib/id.go
// -*- mode: go; coding: utf-8; -*-
// Created on 25. 10. 2022 by Benjamin Walkenhorst
// (c) 2022 Benjamin Walkenhorst
// Time-stamp: <2026-10-30 10:58:12 krylon>

package krylib

// ID is an identifier that is unique within a given context.
type ID int64

// InvalidID is an ID that does not equal any valid ID.
const InvalidID ID = -1
