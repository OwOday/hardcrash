package main

import (
	"github.com/OwOday/hardcrash/src/hardcrash"
)

func main() {
	var hc = hardcrash.New()
	//var ender = hardcrash.Ender{} //need a way to bind ender functions to ender.html through webui
	//var picoemp = goglitch.PicoEMP{}
	//gog.Registertool(&ender, "ender")
	//gog.Registerglitcher(&picoemp, "picoemp")
	hc.Start()
}
