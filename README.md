# godxcc: DXCC CTY.DAT database for Go

## NOTE WELL

* Run `godxcc.LoadCty()` for initialization
* `godxcc.DXCCGetRecord()` argument callsign must be fully capitalized and verified

### cty.dat file

cty.dat file is no longer embedded in the Go code and must be externally provided.
The filename must be all *lowercased*.

The file search path is as follows, respectively:

* /usr/share/dxcc/cty.dat
* /usr/local/share/dxcc/cty.dat
* (directory where the executable file resides)/cty.dat

See [the Big CTY page by Jim Reisert, AD1C](https://www.country-files.com/big-cty/)
for the cty.dat file.

## usage

```go
	// Run the loader before using API
	godxcc.LoadCty()

	// type godxcc.DXCCData
	var dxccdata godxcc.DXCCData
	// Caution: DXCCGetRecord argument callsign
	// must be fully capitalized and verified
	dxccdata := godxcc.DXCCGetRecord("JJ1BDX")
```

## Demo code

* dxcc: lookup command of the DXCCGetRecord
* testparse: testing DXCCGetRecord output

## LICENSE

MIT

