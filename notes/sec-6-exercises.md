# Exercises

## Exercise 1 - Add a new static page to your app using the new pattern.

## Exercise 2 - Experiment with errors

Explore the `errors` package and documentation to see if you can learn how
to determine if a wrapper is a speciffic type. Eg:

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    err := B()
    // TODO Determine if the  `err` variable is an `ErrNotFound`
}

var ErrNotFound = errors.New("not found")

func A() error {
	return ErrNotFound
}

func B() error { 
	err := A()
  if err != nil {
  	return fmt.Errorf("b: %w", err)
  }
  return nil
}
```

*Hint: Look at `error.Is`*

As a followup, read about `errors.As` and see if you can use it of think of cases where it might be useful.