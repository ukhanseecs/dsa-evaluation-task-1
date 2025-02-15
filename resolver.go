package main

import "strings"

// ResolveAbsolutePath returns the absolute path from a given current directory and relative path
func ResolveAbsolutePath(currentDir string, relativePath string) string {
	// Stack to hold the path components
	var dirStack []string

	// Helper function to process a path and modify the stack
	processPath := func(path string, dirStack *[]string) {
		// Split the path by "/"
		pathSegments := strings.Split(path, "/")

		// Process each segment in the path
		for _, segment := range pathSegments {
			if segment == "" || segment == "." {
				// Ignore empty segments (// or .)
				continue
			} else if segment == ".." {
				// Go up one level by popping from the stack
				if len(*dirStack) > 0 {
					*dirStack = (*dirStack)[:len(*dirStack)-1]
				}
			} else {
				// Push the segment onto the stack
				*dirStack = append(*dirStack, segment)
			}
		}
	}

	// Process the current directory, if not root
	if currentDir != "/" {
		processPath(currentDir, &dirStack)
	}

	// Process the relative path
	processPath(relativePath, &dirStack)

	// Manually construct the absolute path
	// Start with "/"
	absolutePath := "/"

	// Manually build the path using a loop
	for i := 0; i < len(dirStack); i++ {
		// Add each directory name to the absolute path with a "/"
		absolutePath += dirStack[i]

		// Only add "/" if it's not the last element
		if i < len(dirStack)-1 {
			absolutePath += "/"
		}
	}

	// Return the final absolute path
	return absolutePath
}
