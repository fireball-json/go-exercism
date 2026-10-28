package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct {
}
type Portuguese struct {
}

var _ Greeter = Italian{}
var _ Greeter = Portuguese{}

func (c Italian) LanguageName() string {
	return "Italian"
}
func (c Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao {name}!", name)
}
func (c Portuguese) LanguageName() string {
	return "Portuguese"
}
func (c Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá {%v}!", name)
}
