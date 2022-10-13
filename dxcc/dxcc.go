// dxcc: search callsigns with godxcc library
// usage: dxcc <callsign>...

package main

import (
	"fmt"
	"github.com/jj1bdx/godxcc"
	"os"
	"regexp"
	"strings"
)

func main() {
	argc := len(os.Args)
	Usage := func() {
		execname := os.Args[0]
		fmt.Fprintln(os.Stderr,
			"dxcc: search callsigns with godxcc library")
		fmt.Fprintf(os.Stderr,
			"Usage: %s [callsign]... \n\n", execname)
		fmt.Fprintf(os.Stderr,
			"dxcc modified by Kenji Rikitake, JJ1BDX.\n"+
				"(c) 2022 Kenji Rikitake, JJ1BDX.\n"+
				"\n"+
				"Originally from:\n"+
				"dxcc in perl at <https://github.com/jj1bdx/dj1yfk-dxcc>\n"+
				"using the godxcc library at <https://github.com/jj1bdx/godxcc>\n"+
				"\n")
	}

	godxcc.LoadCty()

	if argc == 1 {
		Usage()
		return
	}

	for i := 1; i < argc; i++ {

		entry := os.Args[i]
		call := strings.ToUpper(entry)

		// Pick up the prefix part of the testcall
		regcheck := regexp.MustCompile(`^([A-Z0-9\/]+)$`)
		checkmap := regcheck.FindStringSubmatch(call)
		var callsign string
		if len(checkmap) == 0 {
			fmt.Printf("Invalid Call:   %s\n\n", call)
			continue
		} else {
			callsign = checkmap[1]
		}

		dxccdata := godxcc.DXCCGetRecord(callsign)

		fmt.Printf("Callsign:       %s\n", callsign)
		fmt.Printf("Main Prefix:    %s\n", dxccdata.Dxccprefix)
		fmt.Printf("WAE Prefix:     %s\n", dxccdata.Waeprefix)
		fmt.Printf("Country Name:   %s\n", dxccdata.Waecountry)
		fmt.Printf("WAZ Zone:       %d\n", dxccdata.Waz)
		fmt.Printf("ITU Zone:       %d\n", dxccdata.Ituz)
		fmt.Printf("Continent:      %s\n", dxccdata.Cont)
		fmt.Printf("Latitude:       %f\n", dxccdata.Lat)
		fmt.Printf("Longitude:      %f\n", dxccdata.Lon)
		fmt.Printf("UTC shift:      %f\n", dxccdata.Utc)
		fmt.Printf("Entity Code:    %d\n", dxccdata.Entitycode)
		fmt.Printf("\n")
	}

	return

}
