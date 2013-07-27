package main

import (
	"fmt"
	"os"
)

const (
	pkgName = "github.com/lukesampson/figlet"
	defaultFont = "standard"
)

// Globals affected by command line options
var (
	deutschflag bool
	justification int
	paragraphflag int
	right2left bool
	multibyte int
	cmdinput bool
	smushmode int
)

func printusage() {
	fmt.Println("Usage: figlet [ -cklnoprstvxDELNRSWX ] [ -d fontdirectory ]")
	fmt.Println("              [ -f fontfile ] [ -m smushmode ] [ -w outputwidth ]")
	fmt.Println("              [ -C controlfile ] [ -I infocode ] [ message ]")
}

// gets the font entry for the given character, or the 'missing'
// character if the font doesn't contain this character
func getChar(c rune, f font) []string {
	 l, ok := f.chars[c]
	 if !ok {
	 	l = f.chars[0]
	 }
	 return l
}

//func getLines(f font, msg string) [][]string {

//}

func main() {
	fontsdir, err := findFonts()
	if err != nil {
		fmt.Println(err); os.Exit(1)
	}

	fontname := defaultFont
	fontpath, err := findFont(fontsdir, fontname)
	if err != nil {
		fmt.Println(err); os.Exit(1)
	}
	
	f, err := readFont(fontpath)
	if err != nil {
		fmt.Println(err); os.Exit(1)
	}

	msg := os.Args[1]

	for _, c:= range msg {
		fmt.Println(string(c))
		l := getChar(c, f)

		for _, line := range l {
			fmt.Println(line)
		}	
	}

}