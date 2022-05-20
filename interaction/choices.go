package interaction
import (
	"fmt"
)

func GetChoices(special bool) int{
	fmt.Println("\n\t Select Any of the Choice:")
	fmt.Println("\t 1) Heal 🧪")
	if (special){
		fmt.Print("\t 2) Attack Monster with Special Weapon ♨️♨️♨️ ")
	}else{
		fmt.Print("\t 2) Attack Monster ♨️ ")
	}
	fmt.Print("\t\t Your Choice: ")

	var ch int
	fmt.Scanln(&ch)
	fmt.Println("")
	return ch
}
