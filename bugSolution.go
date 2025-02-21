type DivisionByZeroError struct {
    n int
}

func (e *DivisionByZeroError) Error() string {
    return fmt.Sprintf("division by zero: %d", e.n)
}

func calculate(a, b int) (int, error) {
    if b == 0 {
        return 0, &DivisionByZeroError{b}
    }
    return a / b, nil
}

func main() {
    result, err := calculate(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
} 
