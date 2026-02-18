# 🚀 Beginner's Guide to Running Your First Go Program

A simple, beginner-friendly introduction to Go programming that demonstrates how to take user input and display a personalized greeting message.

---

## 📖 Table of Contents

- [Overview](#overview)
- [Why Go?](#why-go)
- [System Requirements](#system-requirements)
- [Installation](#installation)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Code Explanation](#code-explanation)
- [Features](#features)
- [Troubleshooting](#troubleshooting)
- [Learning Resources](#learning-resources)
- [Contributing](#contributing)
- [License](#license)

---

## 🎯 Overview

### What is This Project?

This is a beginner-friendly toolkit designed to help anyone take their first steps in Go programming. The program demonstrates fundamental Go concepts through a simple interactive application that:

1. Prompts the user for their name
2. Reads the input from the terminal
3. Displays a personalized greeting message

### What You'll Learn

By following this guide, you will:
- ✅ Install and configure Go on your system
- ✅ Set up your first Go project
- ✅ Understand basic Go syntax and structure
- ✅ Work with user input and output
- ✅ Run and build Go programs
- ✅ Debug common beginner issues

### Target Audience

This guide is perfect for:
- Complete programming beginners
- Developers from other languages (Python, JavaScript, Java) wanting to learn Go
- Students learning Go for coursework
- Anyone curious about Go's simplicity and power

---

## 💡 Why Go?

Go (Golang) is an open-source, statically-typed, compiled programming language created by Google. Here's why it's worth learning:

### Key Benefits

| Feature | Description |
|---------|-------------|
| **Simple & Clean** | Minimal syntax, easy to read and write |
| **Fast Compilation** | Compiles to native machine code in seconds |
| **Built-in Concurrency** | Goroutines make parallel programming easy |
| **Static Typing** | Catches errors at compile-time, not runtime |
| **Standard Library** | Rich built-in packages for common tasks |
| **Single Binary** | Deploy anywhere - no dependencies needed |

### Real-World Usage

Go powers critical infrastructure at major companies:

- **Google**: Kubernetes, Docker, Cloud services
- **Uber**: High-performance microservices
- **Netflix**: Server-side applications
- **Dropbox**: Backend infrastructure
- **Twitch**: Chat and video streaming services
- **American Express**: Payment processing systems

**Learn more**: [How Companies Use Golang](https://renaldid.medium.com/how-companies-use-golang-7-real-world-examples-you-need-to-know-f9a93d86ca25)

---

## 💻 System Requirements

### Go Version

- **Minimum**: Go 1.20 or higher
- **Recommended**: Go 1.21+ (latest stable version)

### Supported Operating Systems

| OS | Minimum Version | Notes |
|----|----------------|-------|
| **Windows** | 7, 8, 10, 11 | 64-bit recommended |
| **macOS** | 10.13 High Sierra, 10.14 Mojave or later | Intel or Apple Silicon |
| **Linux** | Kernel 2.6.32+ | Most modern distributions supported |

### Hardware Requirements

- **Disk Space**: 250 MB minimum for Go installation
- **RAM**: 512 MB minimum (1 GB recommended)
- **Processor**: Any modern CPU (32-bit or 64-bit)

### Required Tools

- **Text Editor**: VS Code (recommended), Notepad++, Sublime Text, or any code editor
- **Terminal**: Command Prompt (Windows), Terminal (macOS/Linux), or PowerShell

### Optional but Recommended

- **VS Code Extensions**:
  - [Go extension by Google](https://marketplace.visualstudio.com/items?itemName=golang.go)
  - Provides syntax highlighting, IntelliSense, and debugging
- **Git**: For version control and cloning repositories

---

## 🔧 Installation

### Step 1: Install Go

#### Windows

1. Visit the [official Go downloads page](https://go.dev/dl/)
2. Download the `.msi` installer for Windows (e.g., `go1.21.x.windows-amd64.msi`)
3. Run the installer
4. Follow the installation wizard (use default settings)
5. **Important**: The installer will automatically add Go to your PATH

#### macOS

**Option 1: Using the installer**
1. Visit [go.dev/dl](https://go.dev/dl/)
2. Download the `.pkg` file for macOS
3. Open and run the installer
4. Follow the installation steps

**Option 2: Using Homebrew**
```bash
brew install go
```

#### Linux

**Option 1: Using package manager (Ubuntu/Debian)**
```bash
sudo apt update
sudo apt install golang-go
```

**Option 2: Manual installation**
```bash
wget https://go.dev/dl/go1.21.x.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.x.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### Step 2: Verify Installation

Open a **new** terminal window and run:
```bash
go version
```

**Expected output**:
```
go version go1.21.x windows/amd64
```

**If you see this**, Go is installed correctly! ✅

**If you get an error**, see the [Troubleshooting](#troubleshooting) section.

---

### Step 3: Set Up Your First Project

#### 3.1 Create Project Folder

Open your terminal and run:

**On Windows:**
```bash
cd Desktop
mkdir my-first-go-app
cd my-first-go-app
```

**On macOS/Linux:**
```bash
cd ~/Desktop
mkdir my-first-go-app
cd my-first-go-app
```

**What this does**: Creates a folder called `my-first-go-app` on your Desktop and navigates into it.

---

#### 3.2 Initialize Go Module

Inside the `my-first-go-app` folder, run:
```bash
go mod init my-first-go-app
```

**Expected output**:
```
go: creating new go.mod: module my-first-go-app
```

**What this does**: Creates a `go.mod` file that tells Go "this is a Go project".

---

#### 3.3 Create Your Code File

**On Windows:**
```bash
type nul > main.go
```

**On macOS/Linux:**
```bash
touch main.go
```

**What this does**: Creates an empty file called `main.go` where you'll write your code.

---

#### 3.4 Write Your Code

Open `main.go` in VS Code or any text editor:
```bash
code main.go
```

Copy and paste this code into `main.go`:
```go
package main // Every Go program starts with this

import "fmt" // Used for printing and reading input

func main() { // The main function - program starts here
	// Print question to terminal
	fmt.Println("What is your name?")

	// Create variable to store user input
	var name string

	// Read user input from terminal
	fmt.Scanln(&name)

	// Print personalized greeting
	fmt.Printf("Hello, %s! Welcome to Go!\nYou just ran your first Go program!\n", name)
}
```

**Save the file** (Ctrl+S or Cmd+S)

---

#### 3.5 Run Your Program

In the terminal (make sure you're in the `my-first-go-app` folder):
```bash
go run main.go
```

**You should see**:
```
What is your name?
```

**Type your name and press Enter**:
```
What is your name?
Becca
Hello, Becca! Welcome to Go!
You just ran your first Go program!
```

🎉 **Congratulations!** You just ran your first Go program!

---

## 📁 Project Structure

After setup, your project folder should look like this:
```
my-first-go-app/
├── go.mod              # Go module file (defines project)
├── main.go             # Your Go code (the actual program)
└── (optional files)
    ├── README.md       # Documentation
    └── .gitignore      # Git ignore file (if using version control)
```

### File Descriptions

| File | Purpose | Created By |
|------|---------|-----------|
| `go.mod` | Defines the Go module and tracks dependencies | `go mod init` command |
| `main.go` | Contains your program code | You (manually created) |

---

## 🎮 Usage

### Running the Program

**Method 1: Quick Run (Development)**
```bash
go run main.go
```
- Compiles and runs immediately
- No executable file created
- Good for testing during development

**Method 2: Build Executable (Production)**
```bash
# Build the executable
go build -o greeting

# Run it
./greeting           # macOS/Linux
greeting.exe         # Windows
```
- Creates a standalone executable file
- Can distribute to others
- Faster if running multiple times

### Example Session
```bash
$ go run main.go
What is your name?
Alice
Hello, Alice! Welcome to Go!
You just ran your first Go program!

$ go run main.go
What is your name?
Bob
Hello, Bob! Welcome to Go!
You just ran your first Go program!
```

### Advanced Usage

**Run with custom output name**:
```bash
go build -o myprogram
./myprogram
```

**Cross-compile for different OS**:
```bash
# Build for Windows from macOS/Linux
GOOS=windows GOARCH=amd64 go build -o greeting.exe

# Build for Linux from Windows
set GOOS=linux
set GOARCH=amd64
go build -o greeting
```

---

## ✨ Features

### Core Features

| Feature | Description | Code Location |
|---------|-------------|---------------|
| **User Input** | Captures user's name from terminal | `fmt.Scanln(&name)` |
| **Input Validation** | Stores input in a string variable | `var name string` |
| **Formatted Output** | Displays personalized greeting message | `fmt.Printf()` |
| **Error-Free** | Simple, beginner-friendly code with no complex error handling | Entire `main.go` |

### What This Demonstrates

**Go Fundamentals**:
- ✅ Package declaration (`package main`)
- ✅ Importing packages (`import "fmt"`)
- ✅ Main function (`func main()`)
- ✅ Variable declaration (`var name string`)
- ✅ Standard input (`fmt.Scanln()`)
- ✅ Standard output (`fmt.Println()`, `fmt.Printf()`)
- ✅ String formatting (`%s` placeholder)
- ✅ Pointers (basic usage with `&name`)

---

## 📚 Code Explanation

Let's break down the code line by line:

### Complete Code
```go
package main

import "fmt"

func main() {
	fmt.Println("What is your name?")
	
	var name string
	fmt.Scanln(&name)
	
	fmt.Printf("Hello, %s! Welcome to Go!\nYou just ran your first Go program!\n", name)
}
```

### Line-by-Line Breakdown

#### Line 1: Package Declaration
```go
package main
```
- **Purpose**: Declares this file belongs to the `main` package
- **Why `main`?**: Go requires an executable program to have `package main`
- **Rule**: Every Go file starts with a package declaration

---

#### Line 3: Import Statement
```go
import "fmt"
```
- **Purpose**: Imports the `fmt` (format) package from Go's standard library
- **What is `fmt`?**: Provides functions for formatted I/O (input/output)
- **What we use**: `Println`, `Scanln`, `Printf`

---

#### Line 5: Main Function
```go
func main() {
```
- **Purpose**: Defines the entry point of the program
- **When it runs**: Go automatically calls `main()` when you run the program
- **Rule**: Every executable Go program must have exactly one `main()` function

---

#### Line 6: Display Prompt
```go
fmt.Println("What is your name?")
```
- **Purpose**: Prints text to the terminal with a newline at the end
- **`Println`**: "Print Line" - automatically adds `\n` (newline) at the end
- **Output**: `What is your name?` followed by a new line

---

#### Line 8: Variable Declaration
```go
var name string
```
- **Purpose**: Creates a variable to store the user's input
- **Syntax**: `var <name> <type>`
- **Type**: `string` means it can hold text
- **Initial value**: Empty string `""` (Go's zero value for strings)

---

#### Line 9: Read Input
```go
fmt.Scanln(&name)
```
- **Purpose**: Waits for user to type and press Enter, then stores input
- **`Scanln`**: "Scan Line" - reads until newline character
- **`&name`**: The `&` symbol gets the memory address of `name`
  - **Why?**: `Scanln` needs the address to modify the variable
  - **Pointer**: This is a basic use of pointers in Go

---

#### Line 11: Display Greeting
```go
fmt.Printf("Hello, %s! Welcome to Go!\nYou just ran your first Go program!\n", name)
```
- **Purpose**: Prints a formatted message with the user's name
- **`Printf`**: "Print Formatted" - uses placeholders for variables
- **`%s`**: String placeholder - replaced with value of `name`
- **`\n`**: Newline character - moves to next line
- **Example**: If `name = "Alice"`, output is:
```
  Hello, Alice! Welcome to Go!
  You just ran your first Go program!
```

---

### Common Printf Placeholders

| Placeholder | Type | Example |
|-------------|------|---------|
| `%s` | String | `"Hello"` |
| `%d` | Integer | `42` |
| `%f` | Float | `3.14` |
| `%t` | Boolean | `true` |
| `%v` | Any value | Auto-detects type |

---

## ⚙️ Configuration Options

This simple program has no configuration files, but here are ways to customize it:

### Customizing the Greeting Message

**Current code**:
```go
fmt.Printf("Hello, %s! Welcome to Go!\nYou just ran your first Go program!\n", name)
```

**Customization examples**:

**1. Change the greeting style**:
```go
// Formal
fmt.Printf("Greetings, %s. Welcome to the world of Go programming.\n", name)

// Casual
fmt.Printf("Hey %s! 👋 You're now a Go developer!\n", name)

// Enthusiastic
fmt.Printf("🎉 WOW! %s, you just crushed your first Go program! 🚀\n", name)
```

**2. Add more personalization**:
```go
fmt.Println("What is your name?")
var name string
fmt.Scanln(&name)

fmt.Println("How old are you?")
var age int
fmt.Scanln(&age)

fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
fmt.Printf("In 10 years, you'll be %d!\n", age+10)
```

**3. Add colors (advanced)**:
```go
// ANSI color codes
const (
    ColorReset  = "\033[0m"
    ColorGreen  = "\033[32m"
    ColorYellow = "\033[33m"
)

fmt.Printf("%sHello, %s!%s Welcome to Go!\n", ColorGreen, name, ColorReset)
```

### Building with Different Options

**Standard build**:
```bash
go build
```

**Custom executable name**:
```bash
go build -o greeting
```

**Optimized build (smaller size)**:
```bash
go build -ldflags="-s -w" -o greeting
```

**Cross-platform builds**:
```bash
# For Windows
GOOS=windows GOARCH=amd64 go build -o greeting.exe

# For macOS
GOOS=darwin GOARCH=amd64 go build -o greeting

# For Linux
GOOS=linux GOARCH=amd64 go build -o greeting
```

---

## 🔧 Troubleshooting

### Common Issues & Solutions

---

#### Issue 1: "go: command not found" or "'go' is not recognized"

**Symptom**: When you type `go version`, you see an error

**Cause**: Go is not installed or not in your system PATH

**Solutions**:

**Windows**:
1. Verify installation:
   - Check if `C:\Program Files\Go\bin` exists
   - If not, reinstall Go from [go.dev/dl](https://go.dev/dl)

2. Fix PATH:
   - Press `Win + X` → System → Advanced system settings
   - Click "Environment Variables"
   - Under "System variables", find "Path"
   - Click "Edit" → "New"
   - Add: `C:\Program Files\Go\bin`
   - Click OK, restart terminal

**macOS/Linux**:
```bash
# Check if Go is installed
which go

# If not found, check installation path
ls /usr/local/go/bin/go

# Add to PATH (add to ~/.bashrc or ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin

# Reload terminal
source ~/.bashrc  # or source ~/.zshrc
```

---

#### Issue 2: "go: cannot find main module"

**Symptom**: Error when running `go run main.go`

**Cause**: You're not in the project directory or `go.mod` doesn't exist

**Solution**:
```bash
# Check current directory
pwd  # macOS/Linux
cd   # Windows

# Navigate to project folder
cd Desktop/my-first-go-app

# If go.mod missing, create it
go mod init my-first-go-app

# Then run
go run main.go
```

---

#### Issue 3: "main.go: no such file or directory"

**Symptom**: Error says it can't find `main.go`

**Cause**: File doesn't exist or wrong directory

**Solution**:
```bash
# Check if file exists
ls        # macOS/Linux
dir       # Windows

# If main.go not listed, create it
touch main.go           # macOS/Linux
type nul > main.go      # Windows

# Then add code and run
go run main.go
```

---

#### Issue 4: "syntax error: unexpected semicolon or newline before {"

**Symptom**: Compilation error about braces

**Cause**: Opening brace `{` on wrong line (Go requires it on same line)

**Wrong**:
```go
func main()
{  // ❌ ERROR: brace on new line
    fmt.Println("Hello")
}
```

**Correct**:
```go
func main() {  // ✅ CORRECT: brace on same line
    fmt.Println("Hello")
}
```

**Solution**: Use `gofmt` to auto-format:
```bash
gofmt -w main.go
```

---

#### Issue 5: "imported and not used: fmt"

**Symptom**: Compilation error about unused import

**Cause**: You imported `fmt` but didn't use any of its functions

**Solution**:
- Either remove the import if you don't need it
- Or make sure you're using `fmt.Println()`, `fmt.Scanln()`, etc.

**Example of the problem**:
```go
package main

import "fmt"  // ❌ ERROR: imported but not used

func main() {
    // Empty function - not using fmt
}
```

**Fix**:
```go
package main

import "fmt"  // ✅ OK: now we use it

func main() {
    fmt.Println("Hello!")  // Using fmt
}
```

---

#### Issue 6: Program Runs but No Output Appears

**Symptom**: `go run main.go` executes but nothing prints

**Possible causes**:

**Cause 1: Terminal buffering (Windows)**
```bash
# Instead of go run, try:
go build -o myprogram.exe
myprogram.exe
```

**Cause 2: Output going to wrong place**
```bash
# Redirect output explicitly
go run main.go 2>&1
```

**Cause 3: Code issue**
```go
// Make sure you have:
fmt.Println("What is your name?")  // Not just println

// And import fmt:
import "fmt"
```

---

#### Issue 7: "cannot use name (variable of type string) as type *string"

**Symptom**: Error with `fmt.Scanln`

**Cause**: Forgot the `&` before variable name

**Wrong**:
```go
fmt.Scanln(name)  // ❌ Missing &
```

**Correct**:
```go
fmt.Scanln(&name)  // ✅ Correct
```

---

#### Issue 8: Input Not Reading Correctly (Spaces)

**Symptom**: Only first word is captured (e.g., "John Doe" becomes "John")

**Cause**: `Scanln` stops at first space

**Current code**:
```go
fmt.Scanln(&name)  // Only reads "John" from "John Doe"
```

**Solution for full line**:
```go
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

reader := bufio.NewReader(os.Stdin)
name, _ := reader.ReadString('\n')
name = strings.TrimSpace(name)  // Remove newline
```

**For this beginner project**, the simple `Scanln` is fine for single-word names.

---

#### Issue 9: Permission Denied When Building

**Symptom**: "permission denied" when running `go build`

**Cause**: Previous executable file is running or locked

**Solution**:
```bash
# Windows: Close any terminal running the program, then:
del myprogram.exe
go build

# macOS/Linux:
rm ./myprogram
go build
```

---

#### Issue 10: "module declares its path as X but was required as Y"

**Symptom**: Module path mismatch error

**Cause**: `go.mod` module name doesn't match what you're trying to import

**Solution**:
```bash
# Delete go.mod and recreate
rm go.mod
go mod init my-first-go-app

# Or edit go.mod directly to fix module name
```

---

### Still Having Issues?

If you're still stuck:

1. **Check Go version**:
```bash
   go version
```
   Make sure it's 1.20 or higher

2. **Verify file contents**:
```bash
   # View file
   cat main.go      # macOS/Linux
   type main.go     # Windows
```

3. **Try minimal example**:
```go
   package main
   import "fmt"
   func main() {
       fmt.Println("Test")
   }
```

4. **Search error message**:
   - Copy exact error
   - Search on [Stack Overflow](https://stackoverflow.com/questions/tagged/go)
   - Check [Go Forum](https://forum.golangbridge.org/)

5. **Ask for help**:
   - [r/golang on Reddit](https://www.reddit.com/r/golang/)
   - [Go Discord](https://discord.gg/golang)
   - Include: OS, Go version, exact error, your code

---

## 📖 Learning Resources

### Official Documentation

- **[A Tour of Go](https://go.dev/tour/)** - Interactive tutorial (highly recommended!)
- **[Go Documentation](https://go.dev/doc/)** - Official docs
- **[Effective Go](https://go.dev/doc/effective_go)** - Best practices guide
- **[Go by Example](https://gobyexample.com/)** - Code examples for common tasks
- **[Go Playground](https://go.dev/play/)** - Run Go code in browser

### Video Tutorials

- **[Go Tutorial for Beginners - freeCodeCamp](https://www.youtube.com/watch?v=YS4e4q9oBaU)** - 7-hour comprehensive course
- **[Learn Go Programming - Tech With Tim](https://www.youtube.com/watch?v=un6ZyFkqFKo)** - Beginner-friendly series
- **[Go Programming - Traversy Media](https://www.youtube.com/watch?v=SqrbIlUwR0U)** - Crash course

### Books (Free)

- **[The Go Programming Language](https://www.gopl.io/)** - By Brian Kernighan
- **[Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/)** - By Jon Bodner
- **[Go 101](https://go101.org/)** - Free online book

### Interactive Practice

- **[Exercism - Go Track](https://exercism.org/tracks/go)** - Coding exercises with mentorship
- **[LeetCode - Go](https://leetcode.com/)** - Algorithm challenges
- **[Codewars - Go](https://www.codewars.com/)** - Gamified coding challenges

### Community

- **[r/golang](https://www.reddit.com/r/golang/)** - Reddit community
- **[Go Forum](https://forum.golangbridge.org/)** - Official forum
- **[Gophers Slack](https://gophers.slack.com/)** - Active Slack community
- **[Stack Overflow - Go Tag](https://stackoverflow.com/questions/tagged/go)** - Q&A

### Articles & Blogs

- **[Go Blog](https://go.dev/blog/)** - Official Go team blog
- **[How Companies Use Golang](https://renaldid.medium.com/how-companies-use-golang-7-real-world-examples-you-need-to-know-f9a93d86ca25)** - Real-world examples
- **[Why Go is Worth Learning](https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65)** - Motivation

---

## 🤝 Contributing

This is a learning project, but contributions are welcome!

### How to Contribute

1. **Fork the repository**
2. **Create a feature branch**:
```bash
   git checkout -b feature/improvement
```
3. **Make your changes**
4. **Test your changes**:
```bash
   go run main.go
   go build
```
5. **Commit your changes**:
```bash
   git commit -m "Add: description of changes"
```
6. **Push to your fork**:
```bash
   git push origin feature/improvement
```
7. **Open a Pull Request**

### Contribution Ideas

- Add more examples (with age, favorite color, etc.)
- Improve error handling
- Add input validation
- Create alternative versions (quiz, calculator, etc.)
- Improve documentation
- Add screenshots
- Translate README to other languages
- Fix typos or clarify confusing sections

### Code Style

- Follow standard Go formatting (`gofmt`)
- Add comments for beginners
- Keep it simple - this is for learning!
- Test all code before submitting

---

## 📄 License

This project is created for educational purposes as part of the Moringa School AI Capstone Project.

### License Type

**MIT License** - Free to use, modify, and distribute
```
MIT License

Copyright (c) 2026 [Your Name]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

### What This Means

You are free to:
- ✅ Use this code for learning
- ✅ Modify it for your own projects
- ✅ Share it with others
- ✅ Use it in commercial projects

---

## 🙏 Acknowledgments

- **Go Team at Google** - For creating an amazing language
- **Moringa School** - For the AI Capstone Project opportunity
- **Claude AI** - For assisting in the learning process and documentation
- **Go Community** - For excellent learning resources

---

## 📞 Support & Contact

### Need Help?

- **Issues**: Open an issue on GitHub
- **Questions**: Use GitHub Discussions
- **Bugs**: Report via GitHub Issues with:
  - Your OS and Go version
  - Exact error message
  - Steps to reproduce

### Author

**[Your Name]**
- GitHub: [@yourusername](https://github.com/yourusername)
- Email: your.email@example.com
- Project Link: [https://github.com/yourusername/my-first-go-app](https://github.com/yourusername/my-first-go-app)

---

## 🎓 Next Steps

After completing this tutorial, try:

1. **Modify the program**:
   - Ask for age and calculate years to 100
   - Ask multiple questions
   - Add emoji to output

2. **Build something new**:
   - Simple calculator
   - Temperature converter
   - Guessing game

3. **Learn more concepts**:
   - Functions with parameters
   - Structs and methods
   - Error handling
   - File I/O

4. **Take the next step**:
   - [A Tour of Go](https://go.dev/tour/) - Interactive tutorial
   - [Go by Example](https://gobyexample.com/) - More examples
   - Build a CLI tool or simple web server

---

## 🎉 Congratulations!

You've completed your first Go program! This is just the beginning of your Go journey.

**What you've accomplished**:
- ✅ Installed and configured Go
- ✅ Created your first Go project
- ✅ Understood basic Go syntax
- ✅ Worked with variables and user input
- ✅ Compiled and ran a Go program

**Keep coding, keep learning, and welcome to the Go community!** 🚀

---

*Last updated: February 18, 2026*
*Version: 1.0.0*