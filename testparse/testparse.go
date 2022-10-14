// callsign parsing test program

package main

import (
	"bufio"
	"fmt"
	"github.com/jj1bdx/godxcc"
	"io"
	"os"
	// "github.com/pkg/profile" // for profiling only
	"strings"
)

func printinfo(call string) {
	fmt.Printf("call: %s, DXCCData: %v\n", call, godxcc.DXCCGetRecord(call))
}

func main() {
	// defer profile.Start(profile.ProfilePath("."), profile.MemProfile).Stop()

	godxcc.LoadCty()

	fp, _ := os.Open("callsigns.txt")
	reader := bufio.NewReader(fp)

	end := false
	for !end {
		call, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				end = true
				break
			} else {
				fmt.Printf("err: %v\n", err)
				return
			}
		}
		printinfo(strings.TrimSpace(call))
	}
}
