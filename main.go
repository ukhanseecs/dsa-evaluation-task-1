package main

import (
	"fmt"
)

func main() {
	// Test cases resembling real-world scenarios
	testCases := []struct {
		currentDir   string
		relativePath string
		expected     string
	}{
		// Basic navigation
		{"/home/user", "documents", "/home/user/documents"},


		// Moving up directories
		{"/home/user/docs", "../images", "/home/user/images"},

		// Edge cases for root "/"
		{"/", "..", "/"},

		// Multiple consecutive slashes
		{"/home//user///docs/", "./images//", "/home/user/docs/images"},

		// Paths with ./ (current directory)
		{"/home/user", "./downloads", "/home/user/downloads"},

		// Complex nested paths
		{"/home/user/projects/go", "../../../usr/local/bin", "/usr/local/bin"},
		{"/home/user/projects/go", "../../../../../etc/passwd", "/etc/passwd"},
	}

	// Run test cases
	for _, test := range testCases {
		result := ResolveAbsolutePath(test.currentDir, test.relativePath)
		status := "✓"
		if result != test.expected {
			status = "✗"
		}
		fmt.Printf("[%s] Current: %q + Relative: %q\n    Got:      %q\n    Expected: %q\n\n",
			status, test.currentDir, test.relativePath, result, test.expected)
	}
}
