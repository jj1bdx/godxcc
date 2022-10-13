// godxcc functions for looking up the DXCC database
// from the callsign.
// Use DXCCGetRecord("CALLSIGN").
// (Note well: the callsign MUST be in the uppercase)
// See loadcty.go for DXCCData definition.

package godxcc

import (
	// "/fmt" // for debug only
	"regexp"
	"strconv"
	"strings"
)

// Get the DXCCData record for a callsign (must be uppercase)
func DXCCGetRecord(call string) DXCCData {
	record := DXCCData{}
	// if the callsign is matched in the fullcall database,
	// use it as is
	dxccdata, matched := tDXCCFullcalls[call]
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
				record = tDXCCPrefixes[s]
			}
		}
	}
	return record
}

// Obtain WPX prefix for a callsign
func getWpxPrefix(call string) string {

	var parta string
	var partb string
	var partc string

	csadditions := []string{"P", "M", "MM", "AM", "A", "QRP", "QRPP", "LGT"}

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
	// Heuristics:
	// If the first part B is a known prefix
	//   then let the main callsign in C be new B
	//     and let prefix B be new A
	// If not:
	// If B is shorter than C,
	//   then let C be new B and let B be new A
	_, existsb := tDXCCPrefixes[partb]
	if existsb || (len(partb) < len(partc)) {
		parta = partb
		partb = partc
		partc = ""
	}

	// fmt.Printf("parta: %s, partb: %s, partc: %s\n", parta, partb, partc)

	// Depending on these values we have to determine the prefix.
	// Following cases are possible:
	//
	// 1.    A and C empty -> only callsign, subcases
	// 1.1    B contains a number -> everything from start to number in B
	// 1.2    B contains no number -> first two letters of B plus 0
	// 2.    A empty and C is not empty, subcases:
	// 2.1    C is only a number -> A with changed number
	// 2.2    C is /P, /M, /MM, /AM, /QRP, etc. -> 1.
	// 2.3    C is something else and will be interpreted as a Prefix
	// 3.    A is not empty, will be taken as prefix, regardless of C
	// 4.    A is not empty, will be taken as prefix, regardless of C

	// A and C are empty from here
	if parta == "" && partc == "" {
		// Case 1
		i := strings.IndexAny(partb, "0123456789")
		if i >= 0 {
			// Case 1.1
			// B contains a number
			// Prefix is all but the last letters
			regcall := regexp.MustCompile(`(.+\d)[A-Z]*`)
			prefixmap := regcall.FindStringSubmatch(partb)
			// fmt.Println("Case 1.1")
			return prefixmap[1]
		} else {
			// Case 1.2
			// B contains no number
			// Pick first two letters + 0
			// fmt.Println("Case 1.2")
			return partb[0:2] + "0"
		}
	}
	// Case 2
	// A is empty and C is not empty from here
	if parta == "" && partc != "" {
		_, err := strconv.Atoi(partc)
		if err == nil {
			// fmt.Printf("num: %d\n", num)
			// Case 2.1
			// C is only a number
			// Regular prefix of B is in prefix1
			regprefix1 := regexp.MustCompile(`(.+\d)[A-Z]*`)
			prefixmap1 := regprefix1.FindStringSubmatch(partb)
			prefix1 := prefixmap1[1]
			// fmt.Printf("prefix1: %s\n", prefix1)
			// Here we need to find out how many digits there are in the
			// prefix, because for example A45XR/0 is A40. If there are 2
			// numbers, the first is not deleted. If course in exotic cases
			// like N66A/7 -> N7 this brings the wrong result of N67, but I
			// think that's rather irrelevant cos such calls rarely appear
			// and if they do, it's very unlikely for them to have a number
			// attached.   You can still edit it by hand anyway..
			regprefix2 := regexp.MustCompile(`^([A-Z]\d)\d$`)
			prefixmap2 := regprefix2.FindStringSubmatch(prefix1)
			// fmt.Printf("prefixmap2: %v\n", prefixmap2)
			if len(prefixmap2) == 2 {
				// For example:
				// prefix1 = "A45", partc = "0"	-> prefix = "A40"
				// fmt.Println("Case 2.1 a")
				return prefixmap2[1] + partc
			} else {
				// Otherwise cut all numbers
				// Prefix without number in prefix3
				// and add attached number
				prefix2 := strings.TrimRight(prefix1, "0123456789")
				// fmt.Println("Case 2.1 b")
				return prefix2 + partc
			}
		} else {
			// Case 2.2
			// If C is in csaddition, See Case 1.1
			foundincs := false
			for i := range csadditions {
				// For known modifiers, see Case 1.1
				if partc == csadditions[i] {
					foundincs = true
					break
				}
			}
			if foundincs {
				// Same as Case 1.1
				regprefix4 := regexp.MustCompile(`(.+\d)[A-Z]*`)
				prefixmap4 := regprefix4.FindStringSubmatch(partb)
				// fmt.Println("Case 2.2 a")
				return prefixmap4[1]
			} else {
				// if two or more numbers in partc: ignore
				i1 := strings.IndexAny(partc, "0123456789")
				if i1 >= 0 {
					i2 := strings.IndexAny(partc[i1+1:], "0123456789")
					if i2 >= 0 {
						// 2 or more digits
						// See Case 1.1
						regprefix5 := regexp.MustCompile(`(.*[A-Z])\d+`)
						prefixmap5 := regprefix5.FindStringSubmatch(partb)
						// fmt.Println("Case 2.2 b")
						return prefixmap5[1]
					} else {
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
				} else {
					// Same as Case 1.1
					regprefix6 := regexp.MustCompile(`(.+\d)[A-Z]*`)
					prefixmap6 := regprefix6.FindStringSubmatch(partb)
					// fmt.Println("Case 2.3 c")
					return prefixmap6[1]

				}
			}

		}
	}

	// Case 3
	if parta != "" {
		// Case 3: A is not empty
		i := strings.IndexAny(parta, "0123456789")
		if i >= 0 {
			// if ends in number: good prefix
			// fmt.Println("Case 3 a")
			return parta
		} else {
			// fmt.Println("Case 3 b")
			return parta + "0"
		}
	}

	// Return empty string for unparsable prefix
	return ""
}
