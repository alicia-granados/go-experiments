package data

import "fmt"

type Duration float32 // in hour

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

func (c Course) String() string {
	return fmt.Sprintf("---%v---%v", c.Name, c.Instructor.FirstName)
}
