package pkg

import (
	"fmt"
	"syscall"
	"unsafe"
)

/**
* This File Helps run .dll or .so plugins
* in the cwd ./plugins folder.
* It is used by the manager application to run plugins dynamically.
 */

// ExecutePlugin loads a DLL and calls 'PluginRun'
func ExecutePlugin(dllPath string, message string) {
	fmt.Println("--- Loading:", dllPath, "---")

	// 1. Load the DLL
	// On Windows, we use LoadLibrary via syscall
	dll, err := syscall.LoadDLL(dllPath)
	if err != nil {
		fmt.Printf("Failed to load DLL: %v\n", err)
		return
	}
	// Defer releasing the DLL handle
	defer dll.Release()

	// 2. Find the exported C function 'PluginRun'
	proc, err := dll.FindProc("PluginRun")
	if err != nil {
		fmt.Printf("Function 'PluginRun' not found in DLL\n")
		return
	}

	// 3. Prepare Input: Convert Go string to C-String (Byte Slice + Null Terminator)
	// We handle this manually because we aren't using cgo in the manager, just syscall
	cStrBytes := append([]byte(message), 0)
	cStrPtr := unsafe.Pointer(&cStrBytes[0])

	// 4. Call the function
	// Call returns: r1 (return value), r2 (unused), lastErr
	r1, _, _ := proc.Call(uintptr(cStrPtr))

	// 5. Process Output: Convert resulting pointer back to Go String
	if r1 == 0 {
		fmt.Println("Plugin returned null")
		return
	}

	resultStr := ptrToString(r1)
	fmt.Printf("Host received: %s\n", resultStr)

	// Note: In a production app, you should export a 'FreeString' function
	// in the plugin and call it here to free r1 memory to avoid leaks.
}

// ptrToString converts a raw C memory pointer (char*) to a Go string
func ptrToString(ptr uintptr) string {
	// Iterate memory until we hit a 0 byte (null terminator)
	var bytes []byte
	for {
		// Read one byte at the pointer + offset
		val := *(*byte)(unsafe.Pointer(ptr))
		if val == 0 {
			break
		}
		bytes = append(bytes, val)
		ptr++
	}
	return string(bytes)
}
