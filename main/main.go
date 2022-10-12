package main

import (
	"fmt"
	"github.com/jj1bdx/godxcc"
)

func printinfo(call string) {
	fmt.Printf("call: %s, DXCCData: %v\n", call, godxcc.DXCCGetRecord(call))
}

func main() {
	godxcc.LoadCty()
	printinfo("jj1bdx")
	printinfo("8J1RL")
	printinfo("KL7/JJ1BDX")
	printinfo("W1AW/KH0")
	printinfo("3D2AG")
}
