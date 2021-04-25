package bird

import "fmt"
import "github.com/benice7029/fish"

func Bark() string {
	fmt.Println("Bird chu chu!")
	fish.Action()
	return "Bird chu chu!"
}
