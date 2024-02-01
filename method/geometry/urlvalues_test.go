package geometry

import (
	"fmt"
	"net/url"
	"testing"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	// fmt.Println("M: %#v", m)
}

func TestUrlValues(t *testing.T) {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"      (first value)
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // panic: assignment to entry in nil map
}
