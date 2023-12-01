package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	day := os.Args[1]

	dayTmpl, _ := template.ParseFiles("templates/day.tmpl")
	testTmpl, _ := template.ParseFiles("templates/test.tmpl")

	dayFile, _ := os.Create(fmt.Sprintf("calendar/day%s.go", day))
	testFile, _ := os.Create(fmt.Sprintf("calendar/day%s_test.go", day))
	os.Create(fmt.Sprintf("input/Day%s.txt", day))

	dayTmpl.Execute(dayFile, map[string]interface{}{ "Day": day })
	testTmpl.Execute(testFile, map[string]interface{}{ "Day": day })
}
