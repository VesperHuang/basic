package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":    "mouse",
		"course":  "golang",
		"site":    "web",
		"quality": "good",
	}

	fmt.Println("map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	causeName, ok := m["course"]
	fmt.Println(causeName, ok)
	if causeName, ok = m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key dose not exist")
	}

	fmt.Println("Delete")
	name, ok := m["name"]
	if ok {
		fmt.Println("name:", name)
	} else {
		fmt.Println("key dose not exist")
	}
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
