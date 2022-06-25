package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(separateWithComma("12345"))                 //12,345
	fmt.Println(separateWithCommaUsingBytesBuffer("12345")) //12,345
	fmt.Println(separateWithComma("123"))                   //123
	fmt.Println(separateWithCommaUsingBytesBuffer("123"))   //123

	fmt.Println(separateSignedFloatingNumbersWithComma("-12345.678")) //-12,345.678
}

// comma inserts commas in a non-negative decimal integer string.
func separateWithComma(strInt string) string {
	n := len(strInt)
	if n <= 3 {
		return strInt
	}
	return separateWithComma(strInt[:n-3]) + "," + strInt[n-3:]
}

//Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.
func separateWithCommaUsingBytesBuffer(strInt string) string {
	n := len(strInt)
	if n <= 3 {
		return strInt
	}

	buf := new(bytes.Buffer)
	//write first number group
	i := n % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(strInt[:i])

	// rest part
	for j := i + 3; j < n; {
		buf.WriteString("," + strInt[i:j])
		i, j = j, j+3
	}
	buf.WriteString("," + strInt[i:])
	return buf.String()
}

// Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers and an optional sign.
func separateSignedFloatingNumbersWithComma(strFloat string) string {
	buf := new(bytes.Buffer)

	// find decimal point of floating number
	n := strings.Index(strFloat, ".")
	if n == -1 {
		n = len(strFloat)
	}

	// check if the number has a sign
	if strings.HasPrefix(strFloat, "+") || strings.HasPrefix(strFloat, "-") {
		buf.WriteByte(strFloat[0])
		strFloat = strFloat[1:]
		n--
	}

	i := n % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(strFloat[:i])

	for j, c := range strFloat[i:n] {
		if j%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(c)
	}
	buf.WriteString(strFloat[n:])

	return buf.String()
}
