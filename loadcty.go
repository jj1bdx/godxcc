package godxcc

import (
	"bufio"
	"bytes"
	_ "embed"
	// "fmt" // for debug only
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

//go:embed cty.dat
var ctyFile []byte

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

// Locate cty.dat and open the file.
// Returns the *bufio.Reader for the file.
// Search path: /usr/share/dxcc, /usr/local/share/dxcc,
//              and the path where the program resides.

func locateCty() *bufio.Reader {
	basename, err := os.Executable()
	if err != nil {
		log.Fatalf("locateCty() basename: %v", err)
	}
	basepath := path.Base(basename)

	var filename string
	filename = "/usr/share/dxcc/cty.dat"
	_, err = os.Stat(filename)
	if !os.IsNotExist(err) {
	} else {
		filename = "/usr/local/share/dxcc/cty.dat"
		_, err = os.Stat(filename)
		if !os.IsNotExist(err) {
		} else {
			filename = basepath + "/cty.dat"
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

// Read cty.dat and

func LoadCty() {

	var lastdxccdata DXCCData

	// reader := locateCty()
	reader := bufio.NewReader(bytes.NewReader(ctyFile))

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
			dxccdata.waecountry = strings.TrimSpace(linemap[0])
			dxccdata.waz, err = strconv.Atoi(strings.TrimSpace(linemap[1]))
			if err != nil {
				log.Fatalf("LoadCty() dxccdata.waz: %v", err)
			}
			dxccdata.ituz, err = strconv.Atoi(strings.TrimSpace(linemap[2]))
			if err != nil {
				log.Fatalf("LoadCty() dxccdata.ituz: %v", err)
			}
			dxccdata.cont = strings.TrimSpace(linemap[3])
			dxccdata.lat, err = strconv.ParseFloat(strings.TrimSpace(linemap[4]), 64)
			if err != nil {
				log.Fatalf("LoadCty() dxccdata.lat: %v", err)
			}
			dxccdata.lon, err = strconv.ParseFloat(strings.TrimSpace(linemap[5]), 64)
			if err != nil {
				log.Fatalf("LoadCty() dxccdata.lon: %v", err)
			}
			dxccdata.utc, err = strconv.ParseFloat(strings.TrimSpace(linemap[6]), 64)
			if err != nil {
				log.Fatalf("LoadCty() dxccdata.utc: %v", err)
			}
			dxccdata.waeprefix = strings.TrimSpace(linemap[7])
			dxccprefix, exists := WAEToDXCC[dxccdata.waeprefix]
			if !exists {
				dxccdata.dxccprefix = dxccdata.waeprefix
			} else {
				dxccdata.dxccprefix = dxccprefix
			}
			entitycode, exists := ctyToEntitycode[dxccdata.dxccprefix]
			if !exists {
				dxccdata.entitycode = 0
			} else {
				dxccdata.entitycode = entitycode
			}
			lastdxccdata = dxccdata
		} else {
			// prefix line for the previous DXCC
			// Remove ending semicolon
			linetrimmed := strings.TrimRight(strings.TrimSpace(string(line)), ";,")
			words := strings.Split(linetrimmed, ",")
			for i := range words {
				word := words[i]
				// Use saved data, dxccdata may be modified
				dxccdata := lastdxccdata
				// CQ Zone in ()
				regwaz := regexp.MustCompile(`\((\d+)\)`)
				wazstr := regwaz.FindStringSubmatch(word)
				if len(wazstr) == 2 && wazstr[1] != "" {
					// Trim parentheses
					wazval, err := strconv.Atoi(wazstr[1])
					if err != nil {
						log.Fatalf("LoadCty() wazval: %v", err)
					}
					dxccdata.waz = wazval
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
					dxccdata.ituz = ituzval
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
					DXCCFullcalls[fullcall] = dxccdata
					// fmt.Printf("DXCCFullcalls[%s] = %v\n", fullcall, dxccdata)
				} else {
					// Normal prefix
					DXCCPrefixes[callstr] = dxccdata
					// fmt.Printf("DXCCPrefixes[%s] = %v\n", callstr, dxccdata)
				}
			}
			// fmt.Printf("\n")
		}
	}
}
