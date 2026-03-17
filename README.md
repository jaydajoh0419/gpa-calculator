# GPA Calculator

A command-line GPA calculator written in Go for university students on a 4.0 scale.

## What it does

- Takes in your courses, credit hours, and letter grades
- Calculates your weighted GPA on a 4.0 scale
- Supports all letter grades including +/− variants (A+, A, A-, B+, B, etc.)

## Tech Stack

- Go
- `fmt` package for CLI input/output

## Setup

### Prerequisites

- [Go installed](https://go.dev/dl/) (version 1.18+)

### Run it

1. Clone the repo
```bash
   git clone https://github.com/jaydajoh0419/gpa-calculator.git
   cd gpa-calculator
```

2. Run the program
```bash
   go run main.go
```

3. Follow the prompts — enter each course name, credits, and grade when asked.

## Example
```
How many courses? 3

Course 1 name: CSE486
Credits: 3
Grade: A

Course 2 name: MTH411
Credits: 3
Grade: B+

Course 3 name: LIN467
Credits: 3
Grade: A-

Your GPA: 3.57
```

## What I learned

- Building interactive CLI tools in Go
- Working with maps and error handling
- Calculating weighted averages programmatically
