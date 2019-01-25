# GODECK

## Installation

Put the file `99-hid.rules` in `/etc/udev/rules.d` to allow your HID devices to be world readable.

> This is a security choice you need to make for yourself.  Don't get mad at me if someone starts streaming your keypresses over the internets.  You've been warned.

You may need a bunch of libusb type packages installed.  I'll figure that out on a clean install.

## Events
In the `/rules` directory are 30 scripts.  The X.sh script fires for key down, and key up (I think), and X-long.sh fires for key presses (after release of the key) of 500ms or longer.

Still working out firing too many events, may change this.

## Hardcoded
The rules dir is hardcoded to my GOPATH/project directory.  Fork it for now.

## License 
MIT

