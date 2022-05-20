package actions
import (
	"fmt"
	"github.com/vegito11/MonsterSlayer/interaction"
)

func ExecuteNextRound(prince *Player , monster *Player){
	round := prince.round
	fmt.Printf("\n\n ⚔️ ⚔️ ⚔️ ⚔️ ⚔️  Round %v Started  ⚔️ ⚔️ ⚔️ ⚔️ ⚔️ \n", round)

	specialRound := round % 3 == 0 
	for{
		ch := interaction.GetChoices(specialRound)

		if (ch == 1){/* Heal Player */
			prince.restoreHealth()
			break
		}else if(ch == 2){ /* Attack Monster */
			pyr_weapon := prince.getNextWeapon()
			fmt.Printf("\n\t ➕ Monster Dealt 💀🌡️  %v damage", pyr_weapon)
			monster.reduceHealth(pyr_weapon)
			break
		}else{
			fmt.Println(" !!! Invalid Choice , Try Again")
		}

	}

	mst_weapon := monster.getNextWeapon()
	fmt.Printf("\n\t ➕ Player Dealt 🤴🩸 %v damage \n", mst_weapon)
	prince.reduceHealth(mst_weapon)
	
	if(monster.Health < 8 && monster.Health > 0){
		monster.restoreHealth()
	}

	fmt.Printf("\n\t Prince Health 🤴🧪 : %v \n", prince.Health)
	fmt.Printf("\t Monster Health 💀🧪 : %v \n", monster.Health)
	fmt.Printf("\n\n 🪓🪓🪓🪓🪓🪓 Round %v Ended  🪓🪓🪓🪓🪓🪓 ", round)
}
 