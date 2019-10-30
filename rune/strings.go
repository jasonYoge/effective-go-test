package main

import (
	"fmt"
	s "strings"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello, World")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower(upper))
	f("%s\n", s.Title(upper))

	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))

	f("%v\n", s.HasPrefix("Mihalis", "Mi"))
	f("%v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))

	f("%v\n", s.Index("Mihalis", "ha"))

	f("%v\n", s.Count("Mihalis", "i"))

	f("%s\n", s.Repeat("ab", 5))

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s", s.TrimLeft(" \nThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	f("%v\n", s.Compare("Mihalis", "MIHALIS"))
	f("%v\n", s.Compare("MIHALIS", "MIHalis"))

	f("%v\n", s.Fields("Thisis\na\tstring!"))
	f("%s\n", s.Split("abcd efg", ""))

	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))

	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))
}
