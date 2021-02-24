// /Users/krylon/go/src/krylib/msgbuf_test.go
// -*- mode: go; coding: utf-8; -*-
// Created on 02. 07. 2016 by Benjamin Walkenhorst
// (c) 2016 Benjamin Walkenhorst
// Time-stamp: <2019-09-13 22:00:24 krylon>

package krylib

import (
	"fmt"
	"testing"
	"time"
)

var buffer *MessageBuffer

const (
	TestCnt   = 5
	TestDelay = time.Millisecond * 25
)

func TestCreateMessageBuffer(t *testing.T) {
	if buffer = CreateMessageBuffer(); buffer == nil {
		t.Fatal("Error creating MessageBuffer: CreateMessageBuffer returned nil!")
	} else if !buffer.Running() {
		t.Fatal("MessageBuffer is not running!")
	}
} // func TestCreateMessageBuffer(t *testing.T)

func TestAddMessages(t *testing.T) {
	for i := 0; i < TestCnt; i++ {
		m := fmt.Sprintf("Test Message #%d/%d", i, TestCnt)
		buffer.AddMessage(m)
	}

	time.Sleep(time.Millisecond * 25)

	var messages []Message

	if buffer.Empty() {
		t.Fatalf("Error adding messages to MessageBuffer: No messages in buffer after adding %d messages.",
			TestCnt)
	} else if messages = buffer.GetAllMessages(); messages == nil {
		t.Fatal("Failed to get messages from buffer!")
	} else if len(messages) != TestCnt {
		t.Fatalf("Got unpexpected number of messages from buffer: %d (expected %d)",
			len(messages), TestCnt)
	}
} // func TestAddMessages(t *testing.T)

func TestGetMessages(t *testing.T) {
	var msg *Message

	go func() {
		var m string

		for x := 0; x < TestCnt; x++ {
			m = fmt.Sprintf("Msg #%d/%d -- %s",
				x, TestCnt, time.Now().Format("02. 01. 2006, 15:04:05"))
			buffer.AddMessage(m)
			time.Sleep(TestDelay)
		}
	}()

	time.Sleep(TestDelay + (time.Millisecond * 2))

	if msg = buffer.GetOneMessage(); msg == nil {
		t.Fatal("Error getting one message from Buffer after we just added one to it (I think!)!")
	}

}
