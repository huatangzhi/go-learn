package poly

import (
	"fmt"
	"testing"
)
type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

func (g *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(\"Hello World!\")"
}

func WriteFirstProgram(p Programmer)  {
	fmt.Printf("%T %v\n",p,p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	WriteFirstProgram(goProg)
	WriteFirstProgram(javaProg)
}
