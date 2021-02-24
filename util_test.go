// /Users/krylon/go/src/krylib/util_test.go
// -*- coding: utf-8; mode: go; -*-
// Created on 22. 08. 2015 by Benjamin Walkenhorst
// (c) 2015 Benjamin Walkenhorst
// Time-stamp: <2019-09-13 22:15:16 krylon>

package krylib

import (
	"errors"
	"fmt"
	"net/url"
	"testing"
)

func TestFibonacci(t *testing.T) {
	expectedValues := []int64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

	for idx, val := range expectedValues {
		f := Fibonacci(idx)
		if f != val {
			t.Fatalf("Error in Fibonacci series: element number %d is supposed to be %d, but we got %d",
				idx+1, val, f)
		}
	}
}

func parserDummy(s string) (u *url.URL, e error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%v\n", err)
			u = nil
			//e = errors.New(err)
			switch t := err.(type) {
			case error:
				e = t
			case string:
				e = errors.New(t)
			default:
				e = errors.New("invalid type recovered from panic")
			}
		}
	}()

	var addr *url.URL = ParseURL(s)
	return addr, nil
} // func parser_dummy(s string) (u *url.URL, e error)

func TestParseURL(t *testing.T) {
	var mustParse = []string{
		"http://www.google.de/",
		"https://blog.fefe.de/",
		"https://en.wikipedia.org/wiki/Go",
		"https://www.golang.org/",
		"https://www.heise.de/",
		"https://finn.krylon.net:8081/pkg/net/url/",
	}

	var addr *url.URL
	var err error

	for _, valid := range mustParse {
		if addr, err = parserDummy(valid); err != nil {
			t.Errorf("Error parsing %s: %s",
				valid, err.Error())
		} else if addr == nil {
			t.Errorf("No URL was actually returned for string %s!", valid)
		} else if addr.String() != valid {
			t.Errorf("Parsed URL and raw string are not equal: %s <-> %s",
				valid, addr.String())
		}
	}
} // func TestParseURL(t *testing.T)
