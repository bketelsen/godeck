package main

import (
	"strconv"
	"time"
)

type State int

const (
	KeyDown = 0
	KeyUp   = 1
)

type Press int

const (
	PressShort = 0
	PressLong  = 1
)

type Key int

const (
	Key0 Key = iota
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	Key10
	Key11
	Key12
	Key13
	Key14
)

type KeyEvent struct {
	Key       Key
	State     State
	Press     Press
	Duration  time.Duration
	StartTime time.Time
}

func (t KeyEvent) Name() string {
	return "KeyPress:" + strconv.Itoa(int(t.Key))
}
func (t KeyEvent) Created() time.Time {
	return t.StartTime
}

func NewKeyEvent(key Key, state State) *KeyEvent {
	ts := &KeyEvent{
		Key:       key,
		State:     state,
		StartTime: time.Now(),
	}
	return ts

}
