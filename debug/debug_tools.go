package debug

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func Debug(i interface{}) {
	fmt.Print(strings.Repeat("\n", 2))
	fmt.Println(strings.Repeat("#", 10))
	fmt.Printf("%#v\n", i)
	fmt.Println(strings.Repeat("#", 10))
	fmt.Print(strings.Repeat("\n", 2))
}

func DD(i interface{}) {
	Debug(i)
	Die()
}

func SD(i interface{}) {
	spew.Dump(i)
	Die()
}

func Die() {
	os.Exit(0)
}
