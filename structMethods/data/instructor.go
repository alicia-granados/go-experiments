package data

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func NewInstructor(name, lastname string) Instructor {
	return Instructor{FirstName: name, LastName: lastname}
}
