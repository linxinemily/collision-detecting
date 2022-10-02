package domain

import (
	"fmt"
)

type HeroAndHeroCollision struct {
	*AbstractCollision
}

func NewHeroAndHeroCollision(next *ICollision) *HeroAndHeroCollision {
	return &HeroAndHeroCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *HeroAndHeroCollision) match(s1 ISprite, s2 ISprite) bool {
	return s1.getType() == "hero" && s2.getType() == "hero"
}

func (wwc *HeroAndHeroCollision) doHandling(s1 ISprite, s2 ISprite) {
	fmt.Println("hero and hero collision, can not move x1 to x2")
}
