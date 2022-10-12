package godxcc

import (
	"bufio"
	"log"
	"os"
	"path"
)

type DXCCData struct {
	country string
	waz int
	itu int
	cont string
	lat float64
	lon float64
	utc float64
	prefix string
	entitycode int
}

func locateCty() (*bufio.Reader) {
	basename, err := os.Executable()
	if err != nil {
		log.Fatal("locateCty() basename: %v", err)
	}
	basepath := path.Base(basename)

	fileinfo, err := os.Stat("/usr/share/dxcc/cty.dat")
	if os.IsNotExist(err) {
		fileinfo, err = os.Stat("/usr/local/share/dxcc/cty.dat")
	} else if os.IsNotExist(err) {
		fileinfo, err = os.Stat(basepath + "/cty.dat")
	}
	if err != nil {
		log.Fatal("locateCty() unable to find cty.dat: %v", err)
	}

	filename := fileinfo.Name()
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal("locateCty() unable to open %s: %v", filename, err)
	}
	return bufio.NewReader(fp)
}


