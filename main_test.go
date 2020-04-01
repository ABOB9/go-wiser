package gowiser

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func Test0(t *testing.T) {
	str := "<a class=\"test\"/>"
	tk := html.NewTokenizer(strings.NewReader(str))

	fmt.Println(tk.Next())
	fmt.Println(tk.Next())
	fmt.Println(tk.Next())
}
