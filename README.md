# <img src="logo.png" alt="AdiLang Logo" width="50"/> AdiLang

[![Go Version](https://img.shields.io/badge/go-1.21%2B-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A minimalist toy programming language built with ❤️ in Go. Modern syntax meets playful experimentation.

---

## 🚀 Features

<div align="center">

| **Feature**               | **Syntax Example**                          |
|---------------------------|---------------------------------------------|
| **🛠️ Variables**          | `var(name = "AdiLang")`                     |
| **📜 Literals**            | `42` (number), `"hello"` (string)           |
| **💬 Comments**            | `// Single-line`<br>`% Multi-line %`        |
| **🖨️ Print Statements**   | `out->"Hello World!"`                       |
| **🌀 Loops**               | `fordude i in range(5) { ... }`             |
| **🤔 conditional**               | `ifdude condition { ... }`             |
| **🎈 Block Level Design**               | `ifdude condition { ... }`             |

</div>

---

## 📚 Syntax Guide

### Variables & Literals
```adilang
// String variable
var(greeting = "Hello, AdiLang!")

// Number variable
var(answer = 42)
var(newAnswer = answer)

// Print variables
out->greeting
out->newAnswer // output -> 42

// Control flow
// Simple for loop
fordude i in range(3) {
    out->"aditya"
}

%
Output:
aditya
aditya 
aditya
%
```

## Installation Guide 

### Clone repository
```
git clone https://github.com/your-username/adilang.git
cd adilang
```
### Build interpreter
```
go build -o adilang
```
### Run sample program
```
./adilang hello.adi
```

