package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "Puke",
		"course": "golang",
	}

	// m2 == empty map
	m2 := make(map[string]int)

	// m3 == nil
	var m3 map[string]int

	fmt.Println(m, m2, m3)

	fmt.Println("\n Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)

	if failName, ok := m["cccourse"]; ok {
		fmt.Println(failName, ok)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("\n Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
