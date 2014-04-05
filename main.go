/*
	BringGo

	Description: A Go CLI client using github.com/darthlukan/libbring to allow users
	to retrieve data provided by Bring's API.  Features include tracking packages by
	tracking number and finding the nearest pickup point based on location criteria
	such as Zip/Postal Code, Latitude and Longitude, and location ID.


	Copyright (C) 2014  Brian C. Tomlinson

	Contact: brian.tomlinson@linux.com

	This program is free software; you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation; either version 2 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License along
	with this program; if not, write to the Free Software Foundation, Inc.,
	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
*/
package main

import (
	"flag"
	"fmt"
	"github.com/darthlukan/libbring"
	"strconv"
)

var (
	version           string = "BringGo version 1.0 \nWritten by Brian Tomlinson <brian.tomlinson@linux.com>\n"
	usageTrack        string = "\nbringgo <tracking number>\n\nExample:\n\tbringgo 123456789\n\n"
	usagePickZip      string = "\nbringgo <postal code>\n\nExample:\n\tbringgo 4006\n\n"
	usagePickLocation string = "\nbringgo <lat> <lon>\n\nExample:\n\tbringgo 60.3952 5.3217\n\n"
	usagePickId       string = "\nbringgo <location ID>\n\nExample:\n\tbringgo 121110\n\n"
	showVersion       *bool  = flag.Bool("version", false, "Output version information and exit.")
	showUsage         *bool  = flag.Bool("usage", false, "Show usage and exit.")
)

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Printf(version)
		return
	}

	if *showUsage {
		fmt.Printf("%s%s%s%s", usageTrack, usagePickZip, usagePickLocation, usagePickId)
		return
	}

	args := flag.Args()

	if len(args) == 1 {
		// TODO: Determine if postal code, ID, or tracking number
	}

	if len(args) == 2 {
		lat, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			panic(err)
		}

		lon, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			panic(err)
		}

		DisplayPickupByLocation(lat, lon)
	}
}

func DisplayPickupByLocation(lat, lon float64) {
	resp, err := libbring.PickByLocation(lat, lon)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The 10 closest pickup points to your location are:\n\n")
	for i := 0; i <= 9; i++ {
		point := resp["pickupPoint"].([]interface{})[i].(map[string]interface{}) // Sorcery!

		fmt.Printf(
			"Name: %v\nAddress: %v\nCity: %v\nPostal Code: %v\nOperating Hours: %v\nDistance: %v\n\n",
			point["name"], point["address"], point["city"],
			point["postalCode"], point["openingHoursEnglish"], point["distanceInKm"])
	}
}
