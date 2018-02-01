package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	bytez, err := ioutil.ReadFile("VERSION")
	if err != nil {
		panic(err)
	}

	currentVersion := strings.TrimSpace(string(bytez))

	parts := strings.Split(currentVersion, ".")
	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	var bump string
	if len(os.Args) > 1 {
		bump = os.Args[1]
	}

	if bump == "major" {
		major++
		minor = 0
		patch = 0
	} else if bump == "minor" {
		minor++
		patch = 0
	} else if bump == "patch" {
		patch++
	}

	fmt.Printf("%d.%d.%d\n", major, minor, patch)

}
