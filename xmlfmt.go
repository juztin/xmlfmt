package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
)

type node struct {
	Attr     []xml.Attr
	XMLName  xml.Name
	Children []node `xml:",any"`
	Text     string `xml:",chardata"`
}

/* Usage:
 * % echo "<xml><test>blah</test></xml>" | go run xmlfmt.go
 */
func main() {
	reader := bufio.NewReader(os.Stdin)
	decoder := xml.NewDecoder(reader)

	n := node{}
	if err := decoder.Decode(&n); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, err := xml.MarshalIndent(n, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}
