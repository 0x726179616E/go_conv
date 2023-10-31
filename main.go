package main

import (
	"conv/conv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// storage for value to be converted
	var num float64

	// declare flags
	var from string // from units
	var t bool      // convert temperature
	var w bool      // convert weight
	var d bool      // convert distance
	var v bool      // convert volume

	// define flags
	flag.BoolVar(&t, "t", false, "convert temperature")
	flag.BoolVar(&w, "w", false, "convert weight")
	flag.BoolVar(&d, "d", false, "convert distance")
	flag.BoolVar(&v, "v", false, "convert volume")
	flag.StringVar(&from, "from", "", "valid units to convert from:\nTemperature:\n\t\"c\"\n\t\"f\"\n\t\"k\"\nWeight:\n\t\"kg\"\n\t\"lbs\"\nDistance:\n\t\"m\"\n\t\"ft\"\nVolume:\n\t\"l\"\n\t\"gal\"")

	// parse flags
	flag.Parse()

	numArgs := len(os.Args)

	// handle cases specified by number of commandline arguments
	switch numArgs {
	case 2:
		// load "unit to convert from" from stdin
		var unit string
		fmt.Fprintf(os.Stdout, "Enter a unit to convert from: ")
		_, err := fmt.Scan(&unit)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading unit to convert from: %s", err)
			os.Exit(1)
		}
		from = unit

		// load "value to convert" from stdin
		var val float64
		fmt.Fprintf(os.Stdout, "Enter a value to be converted: ")
		_, err = fmt.Scan(&val)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading value to be converted: %s", err)
			os.Exit(1)
		}
		num = val

	case 3:
		// load input value from stdin
		var val float64
		fmt.Fprintf(os.Stdout, "Enter your value to be converted: ")
		_, err := fmt.Scan(&val)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading value to be converted: %s", err)
			os.Exit(1)
		}
		num = val
	case 4:
		// load input value from commandline argument
		val, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading value to convert: %s", err)
			os.Exit(1)
		}
		num = val
	default:
		fmt.Fprintf(os.Stdout, "Usage: ./main <Measurement Type> <Value>\n")
		os.Exit(1)
	}

	// handle cases specified by type of measurement
	switch {
	// handle temperature conversions
	case t:
		fmt.Fprintf(os.Stdout, "converting temperature:\n")
		switch from {
		// converting from celsius
		case "c":
			temp := conv.Celsius(num)
			fmt.Fprintf(os.Stdout, "%v = %v = %v\n", temp, conv.CToF(temp), conv.CToK(temp))
		// converting from fahrenheit
		case "f":
			temp := conv.Fahrenheit(num)
			fmt.Fprintf(os.Stdout, "%v = %v = %v\n", temp, conv.FToC(temp), conv.FToK(temp))
		// converting from kelvin
		case "k":
			temp := conv.Kelvin(num)
			fmt.Fprintf(os.Stdout, "%v = %v = %v\n", temp, conv.KToC(temp), conv.KToF(temp))
		// handle invalid "from" unit for temperature conversion
		default:
			fmt.Fprintf(os.Stderr, "Error: the valid temperature units to convert from are: \"c\" \"f\" and \"k\"\n")
			os.Exit(1)
		}
	// handle weight conversions
	case w:
		fmt.Fprintf(os.Stdout, "converting weight:\n")
		switch from {
		// converting from kilograms to pounds
		case "kg":
			weight := conv.Kilograms(num)
			fmt.Fprintf(os.Stdout, "%v = %v\n", weight, conv.KgToLbs(weight))
		// converting from pounds to kilograms
		case "lbs":
			weight := conv.Pounds(num)
			fmt.Fprintf(os.Stdout, "%v = %v\n", weight, conv.LbsToKg(weight))
		// handle invalid "from" unit for weight conversion
		default:
			fmt.Fprintf(os.Stderr, "Error: the valid weight units to convert from are: \"kg\" and \"lbs\"\n")
			os.Exit(1)
		}
	// handle distance conversions
	case d:
		fmt.Fprintf(os.Stdout, "converting distance:\n")
		switch from {
		case "m":
			dist := conv.Meters(num)
			fmt.Fprintf(os.Stdout, "%v = %v\n", dist, conv.MToFt(dist))
		case "ft":
			dist := conv.Feet(num)
			fmt.Fprintf(os.Stdout, "%v = %v\n", dist, conv.FtToM(dist))
		// handle invalid "from" unit for distance conversion
		default:
			fmt.Fprintf(os.Stderr, "Error: the valid distance units to convert from are: \"m\" and \"ft\"\n")
			os.Exit(1)
		}
	// handle volume conversions
	case v:
		fmt.Fprintf(os.Stdout, "converting volume:\n")
		switch from {
		// converting from liters to gallons
		case "l":
			vol := conv.Liters(num)
			fmt.Fprintf(os.Stdout, "%v = %v\n", vol, conv.LToGal(vol))
		// converting from gallons to liters
		case "gal":
			vol := conv.Gallons(num)
			fmt.Fprintf(os.Stdout, "%v = %v\n", vol, conv.GalToL(vol))
		// handle invalid "from" unit for volume conversion
		default:
			fmt.Fprintf(os.Stderr, "Error: the valid volume units to convert from are: \"l\" and \"gal\"\n")
			os.Exit(1)
		}
	// handle case when invalid flag for measurement type is passed
	default:
		fmt.Fprintf(os.Stdout, "Error: No flag provided.\nUsage of ./main:\n -t\tconvert temperature\n -d\tconvert distance\n -w\tconvert weight\n -v\tconvert volume\n")
	}
}
