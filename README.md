# Path Resolver : DSA Evaluation Task 1 

This repository contains the solution code and explanation for DSA evaluation question #1.
---

## How It Works

### Simple Explanation

The `ResolveAbsolutePath` function works by using a stack to keep track of where we are. Here's how it works:

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


## Getting Started

To get started with this project, follow these steps:

1. Clone the repository
```bash
git clone https://github.com/ukhanseecs/dsa-evaluation-task-1.git
cd dsa-evaluation-task-1
```

2. Run the test cases
```bash
go run .
```
or
```bash
go run main.go resolver.go
```

The program will execute all test cases and display the results.
