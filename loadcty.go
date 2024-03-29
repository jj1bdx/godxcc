// Loader of BIG CTY file cty.dat for godxcc.
// See <https://www.country-files.com/big-cty/>
// for the details of cty.dat format and updates.
// NOTE: godxcc uses an external file of cty.dat

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

var ctyFile []byte

// Full DXCC record data for a given callsign/prefix
type DXCCData struct {
	Waecountry string
	Waz        int
	Ituz       int
	Cont       string
	Lat        float64
	Lon        float64
	Utc        float64
	Waeprefix  string
	Dxccprefix string
	Entitycode int
}

// Tables/maps for DXCC prefixes and full callsigns
// parsed and loaded from cty.dat
// This provides a reference to DXCCData
// with Waz (CQZ) and Ituz (ITUZ) modification
type DXCCRef struct {
	Waeprefix string
	Waz       int
	Ituz      int
}

var tDXCCBase = make(map[string]DXCCData, 400)
var tDXCCPrefixes = make(map[string]DXCCRef, 10000)
var tDXCCFullcalls = make(map[string]DXCCRef, 20000)

// Get complete DXCCFullcalls for a callsign
func GetDXCCFullcalls(callsign string) (DXCCData, bool) {
	dxccref, refexists := tDXCCFullcalls[callsign]
	if refexists {
		dxccdata, dataexists := tDXCCBase[dxccref.Waeprefix]
		if dataexists {
			dxccdata.Waz = dxccref.Waz
			dxccdata.Ituz = dxccref.Ituz
			return dxccdata, true
		}
	}
	return DXCCData{}, false
}

// Get complete DXCCPrefixes for a prefix
func GetDXCCPrefixes(prefix string) (DXCCData, bool) {
	dxccref, refexists := tDXCCPrefixes[prefix]
	if refexists {
		dxccdata, dataexists := tDXCCBase[dxccref.Waeprefix]
		if dataexists {
			dxccdata.Waz = dxccref.Waz
			dxccdata.Ituz = dxccref.Ituz
			return dxccdata, true
		}
	}
	return DXCCData{}, false
}

// Locate cty.dat and open the file.
// Returns the *bufio.Reader for the file.
//
// Search path:
//
//	/usr/share/dxcc, /usr/local/share/dxcc,
//	and the path where the program resides.
func locateCty() *bufio.Reader {
	basename, err := os.Executable()
	if err != nil {
		log.Fatalf("locateCty() basename: %v", err)
	}
	basedir := path.Dir(basename)

	var filename string
	filename = "/usr/share/dxcc/cty.dat"
	_, err = os.Stat(filename)
	if !os.IsNotExist(err) {
	} else {
		// fmt.Printf("locateCty(): %s does not exist\n", filename)
		filename = "/usr/local/share/dxcc/cty.dat"
		_, err = os.Stat(filename)
		if !os.IsNotExist(err) {
		} else {
			// fmt.Printf("locateCty(): %s does not exist\n", filename)
			filename = basedir + "/cty.dat"
			_, err = os.Stat(filename)
			if !os.IsNotExist(err) {
			} else {
				log.Fatalf("locateCty() unable to find cty.dat: %v", err)
			}
		}
	}

	fp, err := os.Open(filename)
	if err != nil {
		log.Fatalf("locateCty() unable to open %s: %v", filename, err)
	}
	return bufio.NewReader(fp)
}

// Read cty.dat and store parsed data
// into tDXCCBase, tDXCCPrefixes and tDXCCFullcalls
func LoadCty() {

	var lastdxccref DXCCRef
	reader := locateCty()

	for line, err := reader.ReadBytes('\n'); line != nil || err != nil; line, err = reader.ReadBytes('\n') {
		if err != nil {
			if err != io.EOF {
				log.Fatalf("LoadCty(): %v", err)
			}
			break // when io.EOF break the loop!
		}
		if line[0] != ' ' {
			// New DXCC data starts
			var dxccdata DXCCData
			var err error
			linemap := strings.Split(string(line), ":")
			dxccdata.Waecountry = strings.TrimSpace(linemap[0])
			dxccdata.Waz, err = strconv.Atoi(strings.TrimSpace(linemap[1]))
			if err != nil {
				log.Fatalf("LoadCty() dxccdata.Waz: %v", err)
			}
			dxccdata.Ituz, err = strconv.Atoi(strings.TrimSpace(linemap[2]))
			if err != nil {
				log.Fatalf("LoadCty() dxccdata.Ituz: %v", err)
			}
			dxccdata.Cont = strings.TrimSpace(linemap[3])
			dxccdata.Lat, err = strconv.ParseFloat(strings.TrimSpace(linemap[4]), 64)
			if err != nil {
				log.Fatalf("LoadCty() dxccdata.Lat: %v", err)
			}
			dxccdata.Lon, err = strconv.ParseFloat(strings.TrimSpace(linemap[5]), 64)
			if err != nil {
				log.Fatalf("LoadCty() dxccdata.Lon: %v", err)
			}
			dxccdata.Utc, err = strconv.ParseFloat(strings.TrimSpace(linemap[6]), 64)
			if err != nil {
				log.Fatalf("LoadCty() dxccdata.Utc: %v", err)
			}
			dxccdata.Waeprefix = strings.TrimSpace(linemap[7])
			dxccprefix, exists := tWAEToDXCC[dxccdata.Waeprefix]
			if !exists {
				dxccdata.Dxccprefix = dxccdata.Waeprefix
			} else {
				dxccdata.Dxccprefix = dxccprefix
			}
			entitycode, exists := ctyToEntitycode[dxccdata.Dxccprefix]
			if !exists {
				dxccdata.Entitycode = 0
			} else {
				dxccdata.Entitycode = entitycode
			}
			tDXCCBase[dxccdata.Waeprefix] = dxccdata
			lastdxccref = DXCCRef{dxccdata.Waeprefix, dxccdata.Waz, dxccdata.Ituz}
		} else {
			// prefix line for the previous DXCC
			// Remove ending semicolon
			linetrimmed := strings.TrimRight(strings.TrimSpace(string(line)), ";,")
			words := strings.Split(linetrimmed, ",")
			for i := range words {
				// Initialize dxccref here
				dxccref := lastdxccref
				word := words[i]
				// CQ Zone in ()
				regwaz := regexp.MustCompile(`\((\d+)\)`)
				wazstr := regwaz.FindStringSubmatch(word)
				if len(wazstr) == 2 && wazstr[1] != "" {
					// Trim parentheses
					wazval, err := strconv.Atoi(wazstr[1])
					if err != nil {
						log.Fatalf("LoadCty() wazval: %v", err)
					}
					dxccref.Waz = wazval
				}
				// ITU Zone in ()
				regituz := regexp.MustCompile(`\[(\d+)\]`)
				ituzstr := regituz.FindStringSubmatch(word)
				if len(ituzstr) == 2 && ituzstr[1] != "" {
					// Trim square brackets
					ituzval, err := strconv.Atoi(ituzstr[1])
					if err != nil {
						log.Fatalf("LoadCty() ituzval: %v", err)
					}
					dxccref.Ituz = ituzval
				}
				// Check fullcall (begins with "=") or not
				pos := strings.IndexAny(word, "([<{~")
				var callstr string
				if pos >= 0 {
					callstr = word[:pos]
				} else {
					callstr = word
				}
				if callstr[0:1] == "=" {
					// Fullcall
					fullcall := callstr[1:]
					tDXCCFullcalls[fullcall] = dxccref
					// testdata, _ := GetDXCCFullcalls(fullcall)
					// fmt.Printf("tDXCCFullcalls[%s] = %v\n", fullcall, testdata)
				} else {
					// Normal prefix
					tDXCCPrefixes[callstr] = dxccref
					// testdata, _ := GetDXCCPrefixes(callstr)
					// fmt.Printf("tDXCCPrefixes[%s] = %v\n", callstr, testdata)
				}
			}
			// fmt.Printf("\n")
		}
	}
}
