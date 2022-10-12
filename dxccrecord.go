package godxcc

import (
	"bufio"
	// "fmt" // for debug only
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// Obtain WPX prefix for a callsign
func getWpxPrefix (call string) string {

	var parta string	
	var partb string	
	var partc string	
	var prefix string

	lidadditions := []string{"QRP", "QRPP", "LGT"}
	csadditions := []string{"P", "M", "MM", "AM", "A"}

	// First check if the call is in the proper format, A/B/C where A and C
        // are optional (prefix of guest country and P, MM, AM etc) and B is the
        // callsign. 
	// Only letters, figures and "/" is accepted, no further check if the
        // callsign "makes sense".

	// Possible formats:
        //  JJ1BDX: B = JJ1BDX
        //  JJ1BDX/KL7: B = JJ1BDX / C = KL7
        //  KL7/JJ1BDX/P : two slashes, easy to determine: 
	//    A = KL7 / B = JJ1BDX / C = P

	callparts := strings.SplitN(call, "/", 3)
	switch len(callparts) {
	case 1:
		parta = ""
		partb = callparts[0]
		partc = ""
	case 2: 
		parta = ""
		partb = callparts[0]
		partc = callparts[1]
	case 3: 
		parta = callparts[0]
		partb = callparts[1]
		partc = callparts[2]
	}

	// Then how to distinguish KL7/JJ1BDX correctly?
        // If the first part is a known prefix, then let the part (in B) be A
        // and let the main callsign (in C) part be B
        // if not: HEURISTIC: if the first part length is smaller than second part,
        // the first part is highly likely to be a prefix

	_, exists:= DXCCPrefixes[partb]
	if exists || (len(partb) < len(partc)) {
		parta = partb
		partb = partc
		partc = ""
	}

	// Remove liddish callsign additions like /QRP and /LGT.	
	if partc != "" {
		for i := range lidadditions {
			if partc == lidadditions[i] {
				partc = ""
			}
		}
	}

	// Depending on these values we have to determine the prefix.
        // Following cases are possible:
        // 
        // 1.    A and C nil -> only callsign, subcases
        // 1.1   B contains a number -> everything from start to number
        // 1.2   B contains no number -> first two letters plus 0
        // 2.    A nil, subcases:
        // 2.1    C is only a number -> A with changed number
        // 2.2    C is /P,/M,/MM,/AM -> 1.
        // 2.3    C is something else and will be interpreted as a Prefix
        // 3.    A is not nil, will be taken as prefix, regardless of C

	if parta == "" && partc == "" {
		// Case 1
		i := strings.IndexAny(partb, "0123456789")
        	if i >= 0 {
			// Case 1.1
			// B contains a number
			// Prefix is all but the last letters
			regcall := regexp.MustCompile(`(.+\d)[A-Z]*`)
                        prefixmap := regcall.FindStringSubmatch(partb)
                	if len(prefixmap) == 2 && prefixmap[1] != "" {
				prefix = prefixmap[1]
			}
		} else {
			// Case 1.2
			// B contains no number
			// Pick first two letters + 0
			prefix = partb[0:2] + "0"
		}
	} else if parta == "" && partc != "" {
		// Case 2
		// CALL/X
		
	}
}



/*
type DXCCData struct {
	waecountry string
	waz        int
	ituz       int
	cont       string
	lat        float64
	lon        float64
	utc        float64
	waeprefix  string
	dxccprefix string
	entitycode int
}

var DXCCPrefixes = map[string]DXCCData{}
var DXCCFullcalls = map[string]DXCCData{}
*/

