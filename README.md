# Inefficient Error Handling in Go

This repository demonstrates a common issue in Go: verbose and repetitive error handling using `errors.New`.  The `bug.go` file shows a simple function that handles division by zero, but the error handling is somewhat cumbersome. The `bugSolution.go` file provides a more concise and efficient approach using custom error types and error wrapping.

## Problem

The primary problem is that `errors.New` creates a new error object for each error condition.  For complex applications, this can lead to a significant amount of boilerplate code and make error management difficult.

## Solution

The improved solution uses custom error types with methods to provide more context and easier error handling.  Error wrapping is also employed to preserve the original error information when wrapping a new error.