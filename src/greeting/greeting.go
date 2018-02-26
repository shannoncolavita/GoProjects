package greeting

import "fmt"

type Salutation struct {
	Name string
	Greeting string
}

type Printer func(string) ()

func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {
	for _, s := range salutation {
		message, alt := CreateMessage(s.Name, s.Greeting)

		if isFormal {

			do(GetPrefix(s.Name) + message)
		} else {
			do(alt)
		}
	}

}

func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string {
		"Bob" : "Mr ",
		"Joe" : "Dr ",
		"Amy" : "Dr ",
		"Mary" : "Mrs ",
	}

	prefixMap["Joe"] = "Jr "

	delete(prefixMap, "Mary")
	return prefixMap[name]
}

func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	case int: fmt.Println("int")
	case string: fmt.Println("string")
	case Salutation: fmt.Println("salutation")
	default: fmt.Println("unknown")
	}

}
func CreateMessage(name string, greeting string) (message string, alt string) {
	message = greeting + " " + name
	alt = "HEY!" + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

