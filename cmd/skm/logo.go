package main

import (
	"github.com/fatih/color"
)

var (
	// Version is the default version of SKM
	Version = "0.0.1"
	logo    = `
  ──────▄▀▄─────▄▀▄
  ─────▄█░░▀▀▀▀▀░░█▄
  ─▄▄──█░░░░░░░░░░░█──▄▄
  █▄▄█─█░░▀░░┬░░▀░░█─█▄▄█	

███████╗██╗  ██╗███╗   ███╗
██╔════╝██║ ██╔╝████╗ ████║
███████╗█████╔╝ ██╔████╔██║
╚════██║██╔═██╗ ██║╚██╔╝██║
███████║██║  ██╗██║ ╚═╝ ██║
╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝
                           
							   
Made with ❤️ by Anda Toshiki at t.me/toshikidev
SKM %s-Bulk SSH key management from terminal at ease


`
)

func displayLogo() {
	color.Cyan(logo, Version)
}
