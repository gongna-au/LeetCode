package main

import (
	"fmt"
	"time"
)

type MyError struct {
	t     time.Time
	reson string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.t, e.reson)
}

func main() {
	MyError{
		time.Now(),
		"unexpext error",
	}

}
