
# Path Resolver

A simple Go package that resolves the absolute path from a given **current directory** and **relative path**. This package processes paths by simulating a **stack-based resolution** of directories, properly handling components like `..`, `.`, and redundant slashes.

---

## How It Works

### Simple Explanation

The `ResolveAbsolutePath` function works like following a set of navigation instructions on a path, using a stack (like a pile of cards) to keep track of where we are. Here's how it works:

1. **Starting Point**: 
   - We start with an empty stack
   - Every path begins at the root directory `/`

2. **Breaking Down the Path**:
   - We split both the current directory and relative path into pieces
   - For example, `/home/user` becomes `["home", "user"]`

3. **Processing Rules**:
   - For a regular folder name (like "documents"): Add it to our stack
   - For `..`: Remove the last item from our stack (go up one level)
   - For `.` or empty spaces: Skip them (they mean "stay here")
   - For multiple slashes (`///`): Treat them as one slash

4. **Building the Final Path**:
   - Start with a `/`
   - Add each folder name from our stack
   - Put a `/` between each folder name
   - The result is our absolute path

### Visual Example

```
Input:
  Current Dir:  /home/user/docs
  Relative Path: ../../pictures

Step-by-step:
1. Initial stack from current dir: ["home", "user", "docs"]
2. Process "../" → remove "docs"    → ["home", "user"]
3. Process "../" → remove "user"    → ["home"]
4. Process "pictures" → add it      → ["home", "pictures"]
5. Build path: "/" + join with "/" = "/home/pictures"
```


## Features

- **Resolves absolute path** from a given current directory and relative path.
- Correctly handles:
  - `.` (current directory).
  - `..` (parent directory).
  - Redundant slashes (`///`).
  - Root directory (`/`).
- **No external libraries** used—pure Go implementation.

---

## Getting Started

To use the `ResolveAbsolutePath` function, you need to include it in your Go project. Follow these steps to get started:

### Prerequisites

- Go version 1.16 or higher.

### Installation

1. Clone this repository or copy the code into your Go project.
2. Import the `ResolveAbsolutePath` function into your main Go file.

```go
import "path_resolver"
```

---

## Usage

### Example

```go
package main

import (
	"fmt"
	"path_resolver"
)

func main() {
	// Test cases
	currentDir := "/home/user"
	relativePath := "documents"
	absolutePath := path_resolver.ResolveAbsolutePath(currentDir, relativePath)
	fmt.Println("Resolved Absolute Path:", absolutePath)
}
```

### Function Signature

```go
func ResolveAbsolutePath(currentDir string, relativePath string) string
```

### Parameters
- `currentDir`: The **current directory** (absolute path).
- `relativePath`: The **relative path** to resolve from the `currentDir`.

### Returns
- A **string** representing the resolved absolute path.

---

## Example Walkthrough

### Test Case 1: Resolving `/home/user` and `documents`

```go
currentDir := "/home/user"
relativePath := "documents"
result := ResolveAbsolutePath(currentDir, relativePath)
fmt.Println(result) // Output: "/home/user/documents"
```

### Test Case 2: Resolving `/usr/local/bin` and `../../etc/config`

```go
currentDir := "/usr/local/bin"
relativePath := "../../etc/config"
result := ResolveAbsolutePath(currentDir, relativePath)
fmt.Println(result) // Output: "/usr/etc/config"
```

### Test Case 3: Resolving `/home/user/docs` and `../images`

```go
currentDir := "/home/user/docs"
relativePath := "../images"
result := ResolveAbsolutePath(currentDir, relativePath)
fmt.Println(result) // Output: "/home/user/images"
```

---

## Edge Cases

- **Relative Path Points Above Root:**
  - Example: `ResolveAbsolutePath("/home", "../../../etc")` → `/etc`
- **Multiple Redundant Slashes:**
  - Example: `ResolveAbsolutePath("/home//user///docs", "images")` → `/home/user/docs/images`
- **Path Starts with `.` or `..`:**
  - Example: `ResolveAbsolutePath("/home/user", "../../admin")` → `/home/admin`

---

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, feel free to open an issue or submit a pull request.

---

## License

This project is licensed under the MIT License.

---

## Acknowledgments

- Thanks to the Go community for providing great resources and documentation.
