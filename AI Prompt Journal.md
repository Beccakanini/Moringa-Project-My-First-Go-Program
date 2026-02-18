# Learning Golang 

Prompting Claude

# Prompt 1: Conceptual Understanding of Golang
I'm currently proficient in Python and want to learn Golang. Before diving into code:1. What are the key philosophical differences between Python and Golang?2. What problems was Go designed to solve?3. What mental models should I adjust coming from Python?4. What are common misconceptions developers have about Golang?

# Prompt 2: Learning Key concepts
I want to understand these key concepts: Consts, Structs and Pointers in Golang. Could you break down: 1. How each concept is implemented in Golang 2. How it compares to Python  3. The key syntax and structures I need to understand 4. Common patterns and best practices
Let's not write complex code yet, just focus on structure and concepts.

# My Notes
Claude AI gave simple explanations with examples for each concept. It also explained every concept in comparison to Python.

# Prompt 3: Testing Understanding with exercises
I want to understand Golang programming concepts, with examples. Kindly explain them to me and ask me questions that would test my understanding. Give me questions that I can apply coding. Remember I dont know the syntax

# My Notes
The AI was able to properly explain concepts in comparison with Python Concepts. It gave me a series of 15 Questions which tested my understanding of the 5 Lessons that included: Variables and Types, Functions, Control Structures, Arrays and Slices and Maps. After doing the exercises I submitted my answers and my understanding was graded with detailed explanations where my answers were wrong.


# Prompt 4 : Learning Through Teaching
I'm new to Golang and I've gotten a bit rusty in coding, and I'd like you to be my Golang tutor. Instead of giving direct answers, please help me learn by:
- Asking guiding questions that help me discover solutions
- Breaking down problems into smaller steps
- Giving hints when I'm stuck rather than complete solutions
- Encouraging me to explain my thinking
- Pointing out patterns and connections to previous concepts
- Helping me debug by asking me to examine my assumptions

# My Notes
The prompt was quite clear. The AI was able to encourage my thinking, testing my understanding with different questions and various code comparisons

# 5 Prompt: Guided Implementation
I'm ready to implement my first Golang function. Could you guide me through creating a
'Student' function with name, ID, and grades? Please explain each part of the syntax,
especially the parts that differ from Python's approach to functions and methods.

# My Notes
I had to refine my prompt because it was not quite clear what exactly was to be performed. I refined it by adding some specific purpose such as calculating the average of the grades and printing  student info


# Prompt 6: # Verify Implementation
I've created this Golang Function:
package main

import "fmt"

type Student struct {
	Name   string
	ID     int
	Grades []int
}

func main() {
	student1 := Student{
		Name:   "James",
		ID:     101,
		Grades: []int{78, 74, 90},
	}
	fmt.Printf("Student Name:%s", "ID:%d", "Grades:%d",student1)

}

could you:
1. Verify if I've followed Golang best practices?
2. Explain any improvements I should make?
3. Point out any Python habits that might be showing in my Golang code?

# My Notes
There were similarities in my code with Python. For example:
In Python, you print: print("Name:", name, "ID:", id, "Grades:", grades)
In Go its: fmt.Printf("Name: %s ID: %d Grades: %v", name, id, grades)


# Prompt 7 : Testing with AI (Behavioral Analysis)
I'm learning how to test my-first-go-program, and I want to understand what behaviors I should test:

package main //every Go program starts with this

import "fmt" //used for printing

func main(){  //the main function starts from here
	fmt.Println("What is your name?") //print the question on the cli

	var name string //create a variable to store name
	fmt.Scanln(&name) // Read user input

	fmt.Printf("Hello, %s! Welcome to Go!\n You just run your first Go program",name) //print greeting
}

Rather than generating tests for me, please:
1. Ask me questions about what I think this program does
2. After I answer, help identify any behaviors I missed
3. Ask me what edge cases I think should be tested
4. Help me identify additional edge cases I didn't think of
5. Ask me which test I should write first and why


#  Prompt 8: creating  comprehensive documentation 
Please create a comprehensive README.md file for my project based on the following information:

Project name: Beginner's Guide To Running First Go Program
Description: The program takes user input first name and prints greeting
Key features:
- User Input
- Greeting Message
- Output Greeting

Technologies used: [Golang] Tools:[VSCode]
Installation requirements: [# System Requirements
Go Version 1.20

# Supported Operating Systems:
Windows:7,8,10, 
macOS: 10.13 High Sierra, 10.14 Mojave
Linux: Kernel 2.6.32+

Disk Space: 250MB
RAM: 512 MB+
]

The README should include:
1. Clear project title and description
2. Installation instructions
3. Basic usage examples
4. Features overview
5. Configuration options
6. Troubleshooting section
7. Contributing guidelines
8. License information

Code structure overview:
[Directory: C:\Users\user\Documents\Moringa AI Access Project


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2/18/2026   8:23 AM                my-first-go-app
-a----         2/17/2026  10:46 PM           1525 AI Prompt Journal.md
-a----         2/18/2026   9:02 AM           2749 README.md]
