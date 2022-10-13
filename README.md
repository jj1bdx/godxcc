# godxcc: DXCC CTY.DAT database for Go

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

* CTY.DAT license is under cty.dat.copyright.txt
* cty.dat reference: Big CTY by Jim Reisert, AD1C <https://www.country-files.com/big-cty/>

