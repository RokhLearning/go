// https://www.youtube.com/watch?v=Rg1w27A0DrU&t=3s
package main

import "fmt"

// MyInterface defines a contract with two methods that any type must implement.
type MyInterface interface {
	hello() string        // Method that returns a greeting message
	by(name string) error // Method that does something with a name (e.g., prints it)
}

// MyStruct is a concrete type that implements MyInterface.
type MyStruct struct{}

//  This line is a **compile-time assertion**.
// It ensures that *MyStruct implements MyInterface.
// If it doesn’t, the compiler will throw an error.
// This is a good way to be sure that your implementation is correct.
var _ MyInterface = (*MyStruct)(nil)

func main() {
	// Declare a variable of the interface type.
	var greeter MyInterface

	//  Method 1: Assign a pointer of MyStruct to the interface variable.
	// This works because *MyStruct implements all methods required by MyInterface.
	// This is the recommended and clean way to use interfaces in Go.
	greeter = &MyStruct{}

	// Call the methods via the interface variable.
	fmt.Println(greeter.hello()) // Outputs: hello babe
	greeter.by("king of kings")  // Outputs: by king of kings

	//  Method 2: Use the concrete struct directly.
	// Here we dereference the pointer (&MyStruct{}), which gives us a value of type MyStruct.
	// Even though the methods are defined on *MyStruct, Go allows us to call them on MyStruct too.
	// The compiler will automatically take the address of the value to match the pointer receiver.
	structInstance := *&MyStruct{}

	// These still call the pointer receiver methods.
	// Go automatically rewrites: structInstance.hello() -> (&structInstance).hello()
	structInstance.hello()
	structInstance.by("king")
}

// hello implements MyInterface. It returns a greeting message.
// Defined on a pointer receiver (*MyStruct) to allow interface satisfaction.
func (s *MyStruct) hello() string {
	return "hello babe"
}

// by implements MyInterface. It prints a goodbye message and returns nil.
// Also defined on a pointer receiver.
func (s *MyStruct) by(name string) error {
	fmt.Printf("by %+v\n", name)
	return nil
}
