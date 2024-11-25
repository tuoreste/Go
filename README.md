# Go Conference Ticket Booking Application

This is a simple ticket booking application built using Go (Golang), which allows users to book tickets for a conference and receive confirmation via email. The application demonstrates various Go features, including concurrency with Goroutines, synchronization with WaitGroups, and struct usage for data storage.

## Introduction to Go (Golang)

Go, also known as Golang, is a statically typed, compiled programming language designed by Google. It is known for its simplicity, efficiency, and concurrency support. Go provides built-in features like Goroutines for handling concurrency, making it ideal for scalable applications and networked services.

In this project, you'll see the following key features of Go:

1. **Basic Syntax**: Go syntax is simple and clean, designed for easy readability.
2. **Concurrency**: Using Goroutines and WaitGroups for running functions concurrently.
3. **Structs**: Defining and using structs to represent user data.
4. **Error Handling**: Go encourages explicit error checking, a key aspect of its design.
5. **Packages and Modules**: Structuring a Go project with modules and using external libraries.

## Project Overview

The `Go conference ticket booking application` allows users to:

- Book tickets for a conference
- Receive a confirmation email
- View a list of first names of all booked users

### Features:

- **Concurrency**: Goroutines are used to simulate sending confirmation emails asynchronously.
- **WaitGroup**: Ensures that all Goroutines finish before the program exits.
- **Structs**: `UserData` struct is used to store user booking information.
- **Data Input**: User inputs are collected via the console, and basic validation is performed.

## Installation

To run the application locally, follow these steps:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/tuoreste/Go/book-app.git
   ```
2. Navigate to the project directory:
   ```
   cd book-app
   ```
3. Run the application:
   ```
   go run .
   ```
### Project Structure

```
book-app/
├── sources/
│   └── helper.go
├── main.go
├── go.mod
├── README.md
```
* main.go: Contains the main logic of the conference booking system.
* sources/helper.go: Contains helper functions like ValidateUserInputs.
* go.mod: Go module file to manage dependencies.

### Code Walkthrough

## Main Features

1. User Input: The application prompts the user for their first name, last name, email, and the number of tickets they wish to book.

2. Input Validation: The user inputs are validated for correctness:

* Name should be at least 2 characters long.
* Email should contain the "@" symbol.
* The number of tickets should be positive and not exceed the remaining tickets.

3. Booking and Confirmation: Once valid input is received, the user's booking is recorded, and a Goroutine is spawned to simulate sending a ticket to the user's email.

4. Concurrency with Goroutines: The sendTicket function runs as a Goroutine, which simulates a delay of 10 seconds before "sending" the confirmation ticket.

5. WaitGroup Synchronization: The sync.WaitGroup is used to ensure that the program waits for all Goroutines to finish before exiting.

### Key Concepts in Go

* Goroutines: Lightweight threads managed by the Go runtime, allowing concurrent execution of functions.
* WaitGroups: Synchronization primitive used to wait for a collection of Goroutines to complete.
* Structs: Composite data types that group related fields.
* Slices: Dynamic arrays that grow and shrink in size, used to store user data and bookings.

### Conclusion

This project serves as a practical introduction to Go and demonstrates key features like concurrency, data handling with structs, and modular programming with packages. The simplicity and performance of Go make it a great choice for building efficient, scalable applications.
