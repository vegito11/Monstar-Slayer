package main
import (
	"github.com/vegito11/MonsterSlayer/interaction"
	"github.com/vegito11/MonsterSlayer/actions"
)

func main(){
	
	interaction.StartGame()
	vegito := actions.NewPlayer(100, true)
	monster := actions.NewPlayer(100, false)

	for vegito.Health > 0 && monster.Health > 0 {
		actions.ExecuteNextRound(vegito, monster)
	}

	interaction.EndGame(vegito.Health)
}