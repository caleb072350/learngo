package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	// The type of a string and a raw string literal
	// is the same. They both are strings.
	// So, they both can be used as a string value.
	var s string
	s = "How are you?"
	fmt.Println(s)

	s = `How are you?`
	fmt.Println(s)

	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)

	s = `
<html>
	<body>"Hello"</body>
</html>`
	fmt.Println(s)

	//win path
	fmt.Println("c:\\my\\dir\\file")
	fmt.Println(`c:\my\dir\file`)

	name, last := "carl", "max"
	fmt.Println(name + " " + last)

	name += " edward"
	fmt.Println(name + " " + last)

	// you can combine raw string and string literals
	fmt.Println(`hello` + `,` + ` ` + `how` + ` ` + `are` + ` ` + "today?")

	eq := "1 + 2 = "
	sum := 1 + 2

	// you need to convert it using strconv.Itoa
	// Itoa = Interger to ASCII
	fmt.Println(eq + strconv.Itoa(sum))

	eq = strconv.FormatBool(true) + " " + strconv.FormatBool(false)
	fmt.Println(eq)

	// len of a string
	// len(s)
	fmt.Println(len(last))

	//unicode 编码长度
	name = "晨辉"
	fmt.Println(len(name))
	fmt.Println(utf8.RuneCountInString(name))

	msg := os.Args[1]

	l := len(msg)

	ss := msg + strings.Repeat("!", l)

	ss = strings.ToUpper(ss)

	fmt.Println(ss)

}
