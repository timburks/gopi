/*
	Go Language Raspberry Pi Interface
	(c) Copyright David Thorpe 2016
	All Rights Reserved

	For Licensing and Usage information, please see LICENSE.md

	gpio.go is a command-line utility to manipulate the GPIO ports
*/
package main

import (
	"flag"
	"fmt"
	"github.com/djthorpe/gopi/rpi"
	"github.com/djthorpe/gopi/rpi/gpio"
	"os"
	"path"
)

////////////////////////////////////////////////////////////////////////////

func main() {

	pi := rpi.New()
	defer pi.Terminate()

	////////////////////////////////////////////////////////////////////////////

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <command>\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	////////////////////////////////////////////////////////////////////////////

	//args := flag.Args()

	////////////////////////////////////////////////////////////////////////////
	// obtain information about the Pi and then create the GPIO interface

	var (
		model *rpi.Model
		g     *gpio.Pins
		err   error
	)

	if model, err = pi.GetModel(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(-1)
	}

	if g, err = gpio.New(model); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(-1)
	}
	defer g.Terminate()

	////////////////////////////////////////////////////////////////////////////
	// Enumerate through the pins of the GPIO connector, starting at 1

	for i := uint(1); i <= g.NumberOfPins; i++ {
		pin, err := g.GetPin(gpio.PhysicalPin(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(-1)
		}
		fmt.Printf("Pin: %v => %+v\n", i, pin)
	}
}