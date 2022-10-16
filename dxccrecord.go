// godxcc functions for looking up the DXCC database
// from the callsign.
// Use DXCCGetRecord("CALLSIGN").
// (Note well: the callsign MUST be in the uppercase)
// See loadcty.go for DXCCData definition.

package godxcc

import (
	// "fmt" // for debug only
	"regexp"
	"strconv"
	"strings"
)

// Get the DXCCData record for a callsign (must be uppercase)
func DXCCGetRecord(call string) DXCCData {
	record := DXCCData{}
	// if the callsign is matched in the fullcall database,
	// use it as is
	dxccdata, matched := GetDXCCFullcalls(call)
	if matched {
		return dxccdata
	}
	// Check WPX prefix if callsign contains a stroke/slash
	var testcall string
	stroke := strings.IndexAny(call, "/")
	if stroke >= 0 {
		testcall = getWpxPrefix(call) + "AA"
	} else {
		testcall = call
	}
	// Pick up the prefix part of the testcall
	regprefix := regexp.MustCompile(`^([A-Z0-9\/]+)`)
	prefixmap := regprefix.FindStringSubmatch(testcall)
	testprefix := prefixmap[1]
	// Use the longest match result for the prefix
	matchlen := 0
	for s := range tDXCCPrefixes {
		if strings.HasPrefix(testprefix, s) {
			lens := len(s)
			if matchlen <= lens {
				matchlen = lens
				newrecord, matched := GetDXCCPrefixes(s)
				if matched {
					record = newrecord
				}
			}
		}
	}
	return record
}

// Callsign additions which should not be recognized as prefixes
var csadditions = map[string]bool{

	"90KK":  true, // ?
	"A":     true, // ?
	"AE":    true, // FCC Rules Part 97.119(f)(3)
	"AG":    true, // FCC Rules Part 97.119(f)(2)
	"AM":    true, // aeronautical mobile
	"B":     true, // ?
	"IEJ50": true, // JAs, please don't add this
	"J":     true, // ?
	"KT":    true, // FCC Rules Part 97.119(f)(1)
	"L":     true, // ?
	"LGT":   true, // Lighthouse? Don't add this
	"M":     true, // mobile
	"MM":    true, // marine mobile
	"OKA50": true, // JAs, please don't add this
	"OKA60": true, // JAs, please don't add this
	"P":     true, // portable
	"QRP":   true, // Don't add this, please
	"QRPP":  true, // Don't add this, please
	"REN":   true, // Regional identifier?
	"SO200": true, // JAs, please don't add this
}

// Obtain WPX prefix for a callsign
func getWpxPrefix(call string) string {

	var parta string
	var partb string
	var partc string

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

	// If C is in csaddition, remove it
	// Example: KL7/JJ1BDX/M -> KL7/JJ1BDX
	// Example: JJ1BDX/AM -> JJ1BDX
	_, existscs := csadditions[partc]
	if existscs {
		partc = ""
	}

	// If the non-empty parts are:
	//  A/B -> then A is the prefix, leave them as they are
	//  A/B/C -> then A is the prefix, leave them as they are
	//  B -> then B is the callsign, leave them as they are

	// So if A is empty and C is NOT empty:
	if parta == "" && partc != "" {
		//  B/C -> use the following heuristics
		// How to distinguish KL7/JJ1BDX correctly
		//  when KL7 is B and JJ1BDX is C?
		// Heuristics:
		// If the first part B is a known prefix
		// (withOUT the longest prefix match)
		//   then let the main callsign in C be new B
		//     and let prefix B be new A
		// If not:
		//   if B is shorter than C,
		//     then let C be new B and let B be new A
		_, existsb := tDXCCPrefixes[partb]
		if existsb || (len(partb) < len(partc)) {
			parta = partb
			partb = partc
			partc = ""
		}
	}

	// fmt.Printf("parta: %s, partb: %s, partc: %s\n", parta, partb, partc)

	// Using A/B/C (where B is the main callsign),
	// we need to process as follows:
	// 1.    If A is not empty: A is the prefix (C is ignored)
	// 2.    If A and C are empty:
	// 2.1    B contains a number ->
	//          Get prefix part of B
	// 2.2    B contains no number ->
	//          first two letters of B plus "0"
	// 3.    If A is empty and C is not empty:
	// 3.1    C is only one-digit number ->
	//          prefix part of B replacing the last digit with C
	// 3.2    C is two or more digits: ignore C, use 2
	// 3.3    For other Cs: C is the prefix

	// 1.    If A exists: A is the prefix (C is ignored)
	if parta != "" {
		i := strings.LastIndexAny(parta, "0123456789")
		if i >= 1 {
			// if the string length is 2 or more and
			// the string ends in number: good prefix
			// No prefix will be shorter than length 2
			// fmt.Println("Case 1.a")
			return parta
		} else {
			// fmt.Println("Case 1.b")
			return parta + "0"
		}
	}

	// A is empty here

	// obtain prefix part of B (main callsign)
	var prefixofb string
	i := strings.IndexAny(partb, "0123456789")
	if i >= 0 {
		// Case 2.1
		// B contains a number
		// Prefix is all but the last letters
		regcall := regexp.MustCompile(`(.+\d)[A-Z]*`)
		prefixmap := regcall.FindStringSubmatch(partb)
		// fmt.Println("Case 2.1")
		prefixofb = prefixmap[1]
	} else {
		// Case 2.2
		// B contains no number
		// Pick first two letters of B + "0"
		// fmt.Println("Case 2.2")
		prefixofb = partb[0:2] + "0"
	}

	// 2.    If A and C are empty:
	if partc == "" {
		return prefixofb
	}

	// A is empty and C is not empty here

	// 3.    If A is empty and C is not empty:

	// 3.1    C is only one-digit number ->
	//        prefix part of B replacing the last digit with C
	num, err := strconv.Atoi(partc)
	if err == nil {
		// fmt.Printf("num: %d\n", num)
		// C contains only number
		if num < 10 {
			// fmt.Println("Case 3.1")
			// C is a one-digit number
			// Example:
			//  A35ABC/0: A35 -> A30
			lb := len(prefixofb)
			if lb >= 2 {
				return prefixofb[:lb-1] + partc
			} else {
				// 1-letter prefix:
				// illegal prefix anyway
				return ""
			}

		} else {
			// 3.2    C is two or more digits: ignore C, use 2
			// fmt.Println("Case 3.2")
			return prefixofb
		}
	} else {
		// 3.3    For other Cs: C is the prefix
		// fmt.Println("Case 3.3")
		// C must be a prefix!
		l := len(partc)
		// if B ends in a digit, it will be a good prefix
		if strings.ContainsAny(partc[l-1:l], "0123456789") {
			// fmt.Println("Case 2.3 a")
			return partc
		} else {
			// Add Zero at the end
			// fmt.Println("Case 2.3 b")
			return partc[:l] + "0"
		}
	}

	// Return empty string for unparsable prefix
	// This code is supposed to be unreachable
	return ""
}
