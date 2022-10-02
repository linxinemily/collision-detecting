package domain

import (
	"fmt"
)

type WaterAndFireCollision struct {
	*AbstractCollision
}

func NewWaterAndFireCollision(next *ICollision) *WaterAndFireCollision {
	return &WaterAndFireCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *WaterAndFireCollision) match(s1 ISprite, s2 ISprite) bool {
	return s1.getType() == "water" && s2.getType() == "fire" || s1.getType() == "fire" && s2.getType() == "water"
}

func (wwc *WaterAndFireCollision) doHandling(s1 ISprite, s2 ISprite) {
	fmt.Println("water and fire collision")
	s1.removeFromWorld()
	s2.removeFromWorld()
}
