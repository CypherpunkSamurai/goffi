package main

/**
 * This is the Main Entrypoint for Manager
 * Usage:
 *   go run manager/main.go [command] --flags
 *
 * Commands:
 *   list					List all available plugins
 *   info <plugin>			Show detailed information about a plugin
 *   run <plugin> <args>	Run a specific plugin
 *   help					Show help information
 *
 * Flags:
 *   --verbose
 *       Enable verbose output
 *   --version
 *       Show version information
 */

import (
	"fmt"
	"log"
	"manager/pkg" // Assumes your module name is 'manager'
	"path/filepath"
)

func main() {
	// 1. Define where to look for plugins
	// Going up one level from manager to find the 'plugins' folder
	pluginDir, _ := filepath.Abs("../plugins")

	fmt.Printf("Scanning for plugins in: %s\n", pluginDir)

	// 2. Find all .dll files
	dlls, err := pkg.FindPlugins(pluginDir)
	if err != nil {
		log.Fatal(err)
	}

	if len(dlls) == 0 {
		fmt.Println("No plugins found. Did you run 'go build -buildmode=c-shared' in the echo folder?")
		return
	}

	// 3. Loop and execute
	for _, dll := range dlls {
		pkg.ExecutePlugin(dll, "Hello from Manager!")
	}
}
