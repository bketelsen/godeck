package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/bketelsen/godeck/version"
	"github.com/bketelsen/libgo/events"
	"github.com/karalabe/hid"
)

var lastState []byte
var startTimes map[int]time.Time

func EventHandler(e events.Event) {
	switch t := e.(type) {
	case KeyEvent:
		dispatchEvent(e.(KeyEvent))
	default:
		// we don't care
		fmt.Printf("%v, %T", t, t)
	}

}

func dispatchEvent(ke KeyEvent) {
	fmt.Println("Key:", ke.Key, "State:", ke.State, "Press", ke.Press)
	long := "-long"
	var script string
	if ke.Press == PressLong {
		script = fmt.Sprintf("/home/bketelsen/src/github.com/bketelsen/godeck/rules/%d%s.sh", int(ke.Key), long)
	} else {
		script = fmt.Sprintf("/home/bketelsen/src/github.com/bketelsen/godeck/rules/%d.sh", int(ke.Key))
	}
	fmt.Println(script)
	cmd := exec.Command("/bin/sh", script)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func eventFromBytes(bb []byte) KeyEvent {

	var pressed int
	var down bool
	var diff time.Duration
	for x, b := range bb[1:] {
		if lastState[x] != b {
			if b == 1 {
				down = true
				startTimes[x] = time.Now()
				pressed = x
			} else {
				down = false
				diff = time.Since(startTimes[x])
				pressed = x
			}
			break
		}
	}

	for j, s := range bb[1:] {
		lastState[j] = s
	}

	var state State
	if down {
		state = KeyDown
	} else {
		state = KeyUp
	}

	var press Press
	if diff > 0 {
		if diff > 500*time.Millisecond {
			press = PressLong
		}
	}
	return KeyEvent{
		Key:      Key(pressed),
		State:    state,
		Duration: diff,
		Press:    press,
	}
}

func main() {
	fmt.Println("Version:", version.Version)
	fmt.Println("HID Supported:", hid.Supported())
	devices := hid.Enumerate(0x0fd9, 0x0060)

	fmt.Println("Elgato Device Info:")
	for _, d := range devices {
		fmt.Println(d.Path)
		fmt.Println(d.VendorID)
		fmt.Println(d.ProductID)
		fmt.Println(d.Release)
		fmt.Println(d.UsagePage)
		fmt.Println(d.Usage)
		fmt.Println(d.Interface)
	}
	if len(devices) > 0 {
		di := devices[0]
		fmt.Println(di)

		d, err := di.Open()
		if err != nil {
			fmt.Println("Unable to open device:", err)
			os.Exit(1)
		}

		lastState = make([]byte, 17)
		startTimes = make(map[int]time.Time)
		quit := make(chan bool)
		kbs := &events.Subscriber{
			Handler: EventHandler,
		}
		events.Subscribe(kbs)
		go func() {
			b := make([]byte, 17)
			for {
				read, err := d.Read(b)
				if err != nil {
					fmt.Println("Error Reading from device:", err)
					quit <- true
				}
				if read == 0 {
					continue
				}
				events.Publish(eventFromBytes(b))
			}

		}()
		<-quit
	}
}
