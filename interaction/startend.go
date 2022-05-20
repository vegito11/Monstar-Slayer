package interaction
import (
	"fmt"
	"github.com/vegito11/MonsterSlayer/utils"
)

func StartGame(){
	fmt.Println("\n 🤴 ⚔️ 💀 Monster Slayer 🤴 ⚔️ 💀")
	fmt.Println(" 🔥 Game Started ... 🔥 ")
	fmt.Println(" 🌠 Good Luck 🍀🌠 ")
	fmt.Println(" Player Health 🤴 : ", utils.PLAYER_HEALTH)
	fmt.Println(" Monster Health 💀 : ", utils.MONSTER_HEALTH)
	fmt.Println("=================================================")
}

func EndGame(Health int){

	fmt.Println("\n\n 🤴 ⚔️ 💀 Monster Slayer 🤴 ⚔️ 💀")
	win_emojis := "🤺🎯🩸🩸☠️ \n 🎆🤴🪄🎆🧨"
	lose_emojis := "💀🪓🩸🩸🤴🥵😵‍💫"
	
	if(Health > 0){
		fmt.Println("Congratulation 🪄 !! You have Defeated Monster 😎 ... ")
		for _, icon := range(win_emojis){
			fmt.Printf("%c", icon)
		}
	}else{
		fmt.Println("Ooops 😵‍💫 !! You have been Defeated by Monster 😵‍💫 ... ")
		for _, icon := range(lose_emojis){
			fmt.Printf("%c", icon)
		}
	}
	fmt.Println("\n🏁🏁🏁🏁 Game Ended ... 🏁🏁🏁🏁 ")
}