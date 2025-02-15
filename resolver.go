package main

import (
	"strings"
)

// ResolveAbsolutePath returns the absolute path from a given current directory and relative path
func ResolveAbsolutePath(currentDir, relativePath string) string {
	// Stack to store the final absolute path components
	var dirStack []string

	// If relativePath is absolute, ignore currentDir completely
	if strings.HasPrefix(relativePath, "/") {
		dirStack = []string{} // Reset the stack
	} else {
		// Convert currentDir into a stack
		dirStack = strings.Split(strings.TrimPrefix(currentDir, "/"), "/")
	}

	// Process relativePath
	for _, segment := range strings.Split(relativePath, "/") {
		switch segment {
		case "", ".": // Ignore empty or current directory references
			continue
		case "..": // Move up one level, but only if we are not at root
			if len(dirStack) > 0 {
				dirStack = dirStack[:len(dirStack)-1]
			}
		default:
			dirStack = append(dirStack, segment)
		}
	}

	// Construct the final absolute path
	if len(dirStack) == 0 {
		return "/" // Return root if the stack is empty
	}
	return "/" + strings.Join(dirStack, "/")
}

