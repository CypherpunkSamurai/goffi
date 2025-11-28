package main

import "C"
import (
	"fmt"
)

/**
 * Echo plugin example
 *
 * This plugin prints back any message sent to it.
 */

// main is required for the compiler, even if empty
func main() {}

// PluginRun takes a C-string, processes it, and returns a C-string.
// We use the C calling convention.
//
//export PluginRun
func PluginRun(cInput *C.char) *C.char {
	// 1. Convert C string to Go string
	goInput := C.GoString(cInput)

	// 2. Perform logic (Echo)
	fmt.Printf("[Plugin Echo] Received: %s\n", goInput)
	result := "Echo from DLL: " + goInput

	// 3. Convert Go string back to C string (Allocates memory on heap!)
	return C.CString(result)
}
