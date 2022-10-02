package domain

import (
	"fmt"
)

type HeroAndFireCollision struct {
	*AbstractCollision
}

func NewHeroAndFireCollision(next *ICollision) *HeroAndFireCollision {
	return &HeroAndFireCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *HeroAndFireCollision) match(s1 ISprite, s2 ISprite) bool {
	return s1.getType() == "hero" && s2.getType() == "fire" || s1.getType() == "fire" && s2.getType() == "hero"
}

func (wwc *HeroAndFireCollision) doHandling(s1 ISprite, s2 ISprite) {
	fmt.Println("hero and fire collision")

	var hero *Hero
	if s1.getType() == "hero" && s2.getType() == "fire" {
		hero, _ = s1.(*Hero)
		hero.substractHP(10)
		s2.removeFromWorld()
		if hero.HP <= 0 {
			fmt.Println("hero died")
			s1.removeFromWorld()
		} else {
			s1.updateCoordinate(s2.getCoordinate())
		}

	} else if s1.getType() == "fire" && s2.getType() == "hero" {
		hero, _ = s2.(*Hero)
		hero.substractHP(10)
		s1.removeFromWorld()
		if hero.HP <= 0 {
			hero.removeFromWorld()
			fmt.Println("hero died")
		}
	}
}
