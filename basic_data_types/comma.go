package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(separateWithComma("12345"))                 //12,345
	fmt.Println(separateWithCommaUsingBytesBuffer("12345")) //12,345
	fmt.Println(separateWithComma("123"))                   //123
	fmt.Println(separateWithCommaUsingBytesBuffer("123"))   //123
}

// comma inserts commas in a non-negative decimal integer string.
func separateWithComma(strInt string) string {
	n := len(strInt)
	if n <= 3 {
		return strInt
	}
	return separateWithComma(strInt[:n-3]) + "," + strInt[n-3:]
}

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
