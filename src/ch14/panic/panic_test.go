package panic

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from", err)
		}
	}()

	fmt.Println("Start")
	//os.Exit(-1)
	panic(errors.New("Something Wrong!"))
}
