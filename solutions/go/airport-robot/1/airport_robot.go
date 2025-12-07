package airportrobot
import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(na string) string
}

type greeterMain struct{
    Language string
    greeting string
}

func (gG greeterMain) LanguageName() string{
    return gG.Language
}

func (gG greeterMain) Greet(na string) string{
    return fmt.Sprintf("%s %s!", gG.greeting, na)
}

func SayHello(name string, greeter Greeter) string{
    return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct {}
func (iI Italian) LanguageName() string{
    return "Italian"
}

func (iI Italian) Greet(na string) string{
    return fmt.Sprintf("%s %s!", "Ciao", na)
}

type Portuguese struct {}
func (Po Portuguese) LanguageName() string{
    return "Portuguese"
}

func (Po Portuguese) Greet(na string) string{
    return fmt.Sprintf("%s %s!", "Olá", na)
}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
