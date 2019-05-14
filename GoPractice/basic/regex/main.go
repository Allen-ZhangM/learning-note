package main

import (
	"fmt"
	"regexp"
)

const text = "My email is 250081124@qq.com"

func main() {

	compile, _ := regexp.Compile(`.+@.+\..+`)
	s := compile.FindString(text)
	fmt.Println(s)

}
