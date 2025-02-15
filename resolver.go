package main

import (
	"strings"
)

// ResolveAbsolutePath returns the absolute path from a given current directory and relative path
func ResolveAbsolutePath(currentDir, relativePath string) string {
	// If relativePath is absolute, ignore currentDir
	if strings.HasPrefix(relativePath, "/") {
		currentDir = ""
	}

	// Convert currentDir into a stack
	var dirStack []string
	if currentDir != "" {
		dirStack = strings.Split(strings.TrimPrefix(currentDir, "/"), "/")
	}

	// Process relativePath
	for _, segment := range strings.Split(relativePath, "/") {
		switch segment {
		case "", ".": // Ignore empty or current directory references
			continue
		case "..": // Move up one level, or ignore if already at root
			if len(dirStack) > 0 {
				dirStack = dirStack[:len(dirStack)-1]
			} else if currentDir == "" { // If we're already at root, stay at root
				// Do nothing, we remain at the root
			}
		default:
			dirStack = append(dirStack, segment)
		}
	}

	// Ensure we don't retain leftover parts of currentDir if .. escapes completely
	if len(dirStack) == 0 {
		return "/" // Return root if stack is empty
	}
	return "/" + strings.Join(dirStack, "/")
}
