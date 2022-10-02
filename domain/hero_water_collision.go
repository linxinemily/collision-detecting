package domain

import (
	"fmt"
)

type HeroAndWaterCollision struct {
	*AbstractCollision
}

func NewHeroAndWaterCollision(next *ICollision) *HeroAndWaterCollision {
	return &HeroAndWaterCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *HeroAndWaterCollision) match(s1 ISprite, s2 ISprite) bool {
	return s1.getType() == "hero" && s2.getType() == "water" || s1.getType() == "water" && s2.getType() == "hero"
}

func (wwc *HeroAndWaterCollision) doHandling(s1 ISprite, s2 ISprite) {
	fmt.Println("hero and water collision")

	var hero *Hero
	if s1.getType() == "hero" && s2.getType() == "water" {
		hero, _ = s1.(*Hero)
		hero.addHP(10)
		s2.removeFromWorld()
		hero.updateCoordinate(s2.getCoordinate())
	} else if s1.getType() == "water" && s2.getType() == "hero" {
		hero, _ = s2.(*Hero)
		hero.addHP(10)
		s1.removeFromWorld() // just remove water
	}
}
