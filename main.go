package main

import (
	"flag"
	"fmt"

	"github.com/jjournet/alaconf/pkg/alacritty"
)

var (
	color    = flag.String("color", "", "color theme to use")
	getcolor = flag.Bool("getcolor", false, "get current color theme")
)

func main() {
	flag.Parse()

	myconf := alacritty.NewConfig("")
	if *getcolor {
		fmt.Println(alacritty.GetColorTheme(&myconf))
	} else if *color != "" {
		alacritty.ChangeColorTheme(&myconf, *color)
		alacritty.SaveConfig(&myconf)
	}
}
