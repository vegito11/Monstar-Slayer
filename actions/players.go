package actions

import (
	"github.com/vegito11/MonsterSlayer/utils"
	"fmt"
)

type Player struct{
	Health int
	round int
	isHero bool
}

func NewPlayer(Health int, isHero bool) *Player{
	pyr := Player{Health, 1, isHero}
	return &pyr
}

func (pyr *Player) reduceHealth(dmg int){
	pyr.Health -= dmg 
	pyr.round++
}

func (pyr *Player) getNextWeapon() int{
	
	if(pyr.isHero){
		if (pyr.round % 3 == 0){
			return utils.GetNumBetween(utils.SPECIAL_ATTACK_MIN_DAMAGE, utils.SPECIAL_ATTACK_MAX_DAMAGE)
		}
		return utils.GetNumBetween(utils.PLAYER_MIN_ATTACK, utils.PLAYER_MAX_ATTACK)
	}else{
		return utils.GetNumBetween(utils.PLAYER_MIN_ATTACK, utils.PLAYER_MAX_ATTACK)
	}
}

func (pyr *Player) restoreHealth(){
	medic := 0
	if(pyr.isHero){
		medic = utils.GetNumBetween(utils.HEALTH_MIN_RESTORE, utils.HEALTH_MAX_RESTORE)
		if (pyr.Health + medic > utils.PLAYER_HEALTH ){
			return
		}
		fmt.Println("\t ➕ Player Health is increased 🤴🧪💪💪 by ", medic)
	}else{
		medic = utils.GetNumBetween(utils.MONSTER_MIN_HEALTH_RESTORE, utils.MONSTER_MAX_HEALTH_RESTORE)
		if (pyr.Health + medic > utils.MONSTER_HEALTH ){
			return
		}
		fmt.Println("\t ➕ Monster Health is increased 💀🧪💪💪 by ", medic)
	}

	pyr.Health += medic
}