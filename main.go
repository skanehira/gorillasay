package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
)

const gorilla = `
%s
          \
           \
            \
            ,.-" "-.,
           /   ===   \
          /  =======  \
       __|  (o)   (0)  |__
      / _|    .---.    |_ \
     | /.----/ O O \----.\ |
      \/     |     |     \/
      |                   |
      |                   |
      |                   |
      _\   -.,_____,.-   /_
  ,.-"  "-.,_________,.-"  "-.,
 /          |       |          \
|           l.     .l           |
|            |     |            |
l.           |     |           .l
 |           l.   .l           | \,
 l.           |   |           .l   \,
  |           |   |           |      \,
  l.          |   |          .l        |
   |          |   |          |         |
   |          |---|          |         |
   |          |   |          |         |
   /"-.,__,.-"\   /"-.,__,.-"\"-.,_,.-"\
  |            \ /            |         |
  |             |             |         |
   \__|__|__|__/ \__|__|__|__/ \_|__|__/
`

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	var length int

	for _, w := range args {
		wl := runewidth.StringWidth(w)
		if wl > length {
			length = wl
		}
	}
	length = length / 2

	str := strings.Builder{}

	for i := 0; i < length; i++ {
		str.WriteString("🍌")
	}

	str.WriteString("\n")
	str.WriteString(strings.Join(args, "\n"))
	str.WriteString("\n")

	for i := 0; i < length; i++ {
		str.WriteString("🍌")
	}

	fmt.Printf(gorilla, str.String())
}
