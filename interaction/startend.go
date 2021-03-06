package interaction
import (
	"fmt"
	"github.com/vegito11/MonsterSlayer/utils"
)

func StartGame(){
	fmt.Println("\n π€΄ βοΈ π Monster Slayer π€΄ βοΈ π")
	fmt.Println(" π₯ Game Started ... π₯ ")
	fmt.Println(" π  Good Luck ππ  ")
	fmt.Println(" Player Health π€΄ : ", utils.PLAYER_HEALTH)
	fmt.Println(" Monster Health π : ", utils.MONSTER_HEALTH)
	fmt.Println("=================================================")
}

func EndGame(Health int){

	fmt.Println("\n\n π€΄ βοΈ π Monster Slayer π€΄ βοΈ π")
	win_emojis := "π€Ίπ―π©Έπ©Έβ οΈ \n ππ€΄πͺππ§¨"
	lose_emojis := "ππͺπ©Έπ©Έπ€΄π₯΅π΅βπ«"
	
	if(Health > 0){
		fmt.Println("Congratulation πͺ !! You have Defeated Monster π ... ")
		for _, icon := range(win_emojis){
			fmt.Printf("%c", icon)
		}
	}else{
		fmt.Println("Ooops π΅βπ« !! You have been Defeated by Monster π΅βπ« ... ")
		for _, icon := range(lose_emojis){
			fmt.Printf("%c", icon)
		}
	}
	fmt.Println("\nππππ Game Ended ... ππππ ")
}