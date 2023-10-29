package main

import (
	"conv/conv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// declare flags
	var from string // from units
	var to string   // to units
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
	flag.StringVar(&to, "to", "", "valid units to convert from:\nTemperature:\n\t\"c\"\n\t\"f\"\n\t\"k\"\nWeight:\n\t\"kg\"\n\t\"lbs\"\nDistance:\n\t\"m\"\n\t\"ft\"\nVolume:\n\t\"l\"\n\t\"gal\"")

	// parse flags
	flag.Parse()

	// store input value to convert
	val, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading value to convert: %s", err)
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
			temp := conv.Celsius(val)
			switch to {
			// to fahrenheit
			case "f":
				fmt.Fprintf(os.Stdout, "%v = %v\n", temp, conv.CToF(temp))
			// to kelvin
			case "k":
				fmt.Fprintf(os.Stdout, "%v = %v\n", temp, conv.CToK(temp))
			// handle invalid "to" unit for celsius conversion
			default:
				fmt.Fprintf(os.Stderr, "valid temperature units to convert to:\n\t\"f\"\n\t\"k\"\n")
				os.Exit(1)
			}
		// converting from fahrenheit
		case "f":
			temp := conv.Fahrenheit(val)
			switch to {
			// to celsius
			case "c":
				fmt.Fprintf(os.Stdout, "%v = %v\n", temp, conv.FToC(temp))
			// to kelvin
			case "k":
				fmt.Fprintf(os.Stdout, "%v = %v\n", temp, conv.FToK(temp))
			// handle invalid "to" unit for fahrenheit conversion
			default:
				fmt.Fprintf(os.Stderr, "valid temperature units to convert to:\n\t\"c\"\n\t\"k\"\n")
				os.Exit(1)
			}
		// converting from kelvin
		case "k":
			temp := conv.Kelvin(val)
			switch to {
			// to celsius
			case "c":
				fmt.Fprintf(os.Stdout, "%v = %v\n", temp, conv.KToC(temp))
			// to fahrenheit
			case "f":
				fmt.Fprintf(os.Stdout, "%v = %v\n", temp, conv.KToF(temp))
			// handle invalid "to" unit for kelvin conversion
			default:
				fmt.Fprintf(os.Stderr, "valid temperature units to convert to:\n\t\"c\"\n\t\"f\"\n")
				os.Exit(1)
			}
		default:
			fmt.Fprintf(os.Stderr, "valid temperature units to convert from:\n\t\"c\"\n\t\"f\"\n\t\"\"k\"\n")
			os.Exit(1)
		}
	// handle weight conversions
	case w:
		fmt.Fprintf(os.Stdout, "converting weight:\n")
		switch from {
		// converting from kilograms
		case "kg":
			weight := conv.Kilograms(val)
			if to == "lbs" {
				fmt.Fprintf(os.Stdout, "%v = %v\n", weight, conv.KgToLbs(weight))
			} else {
				fmt.Fprintf(os.Stderr, "valid weight units to convert to:\n\t\"lbs\"\n")
				os.Exit(1)
			}
		// converting from pounds
		case "lbs":
			weight := conv.Pounds(val)
			if to == "lbs" {
				fmt.Fprintf(os.Stdout, "%v = %v\n", weight, conv.LbsToKg(weight))
			} else {
				fmt.Fprintf(os.Stderr, "valid weight units to convert to:\n\t\"kg\"\n")
				os.Exit(1)
			}
		default:
			fmt.Fprintf(os.Stderr, "valid weight units to convert from:\n\t\"kg\"\n\t\"lbs\"\n")
			os.Exit(1)
		}
	// handle distance conversions
	case d:
		fmt.Fprintf(os.Stdout, "converting distance:\n")
		switch from {
		case "m":
			dist := conv.Meters(val)
			if to == "ft" {
				fmt.Fprintf(os.Stdout, "%v = %v\n", dist, conv.MToFt(dist))
			} else {
				fmt.Fprintf(os.Stderr, "valid distance units to convert to:\n\t\"ft\"\n")
				os.Exit(1)
			}
		case "ft":
			dist := conv.Feet(val)
			if to == "m" {
				fmt.Fprintf(os.Stdout, "%v = %v\n", dist, conv.FtToM(dist))
			} else {
				fmt.Fprintf(os.Stderr, "valid distance units to convert to:\n\t\"m\"\n")
				os.Exit(1)
			}
		default:
			fmt.Fprintf(os.Stderr, "valid distance units to convert from:\n\t\"m\"\n\t\"ft\"\n")
			os.Exit(1)
		}
	// handle volume conversions
	case v:
		fmt.Fprintf(os.Stdout, "converting volume:\n")
		switch from {
		case "l":
			vol := conv.Liters(val)
			if to == "gal" {
				fmt.Fprintf(os.Stdout, "%v = %v\n", vol, conv.LToGal(vol))
			} else {
				fmt.Fprintf(os.Stderr, "valid volume units to convert to:\n\t\"gal\"\n")
				os.Exit(1)
			}
		case "gal":
			vol := conv.Gallons(val)
			if to == "l" {
				fmt.Fprintf(os.Stdout, "%v = %v\n", vol, conv.GalToL(vol))
			} else {
				fmt.Fprintf(os.Stderr, "valid volume units to convert to:\n\t\"l\"\n")
				os.Exit(1)
			}
		default:
			fmt.Fprintf(os.Stderr, "valid volume units to convert from:\n\t\"l\"\n\t\"gal\"\n")
			os.Exit(1)
		}
	//
	default:
		fmt.Fprintf(os.Stdout, "No flag provided.\nUsage of ./main:\n -t\tconvert temperature\n -d\tconvert distance\n -w\tconvert weight\n -v\tconvert volume\n")
	}
}
