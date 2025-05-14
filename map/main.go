package main

import "fmt"

func main() {
	// This is the most concise and readable way to initialize a map with known key-value pairs.
	// It creates a non-nil map and fills it with initial values at the same time.
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"white": "#FFFFFF",
		"black": "#000000",
	}

	printMap(colors)

	// -------------------------------------
	// Other common ways to define maps in Go:
	// -------------------------------------

	// Method 1: Declare the map variable without initialization.
	// This map is nil, meaning it points to no actual memory until assigned.
	// If you try to add entries to it now, it will panic at runtime.
	// Useful when you plan to initialize the map later or under a condition.
	// var colors map[string]string

	// Method 2: Use make() to initialize the map with no entries.
	// This creates a non-nil map, ready to accept key-value assignments.
	// Common when you need a dynamic or empty map at start.
	// colors = make(map[string]string)

	// Method 3: Empty map literal — same as make(), but shorter.
	// Useful when you want an empty map but don't want to call make().
	// colors := map[string]string{}

}

// printMap prints each key-value pair in the given map.
// It uses a for loop with the range keyword to iterate over all entries.
func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("hex of ", key, " is ", value)
	}
}
