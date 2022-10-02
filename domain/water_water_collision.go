package domain

import (
	"fmt"
)

type WaterAndWaterCollision struct {
	*AbstractCollision
}

func NewWaterAndWaterCollision(next *ICollision) *WaterAndWaterCollision {
	return &WaterAndWaterCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *WaterAndWaterCollision) match(s1 ISprite, s2 ISprite) bool {
	return s1.getType() == "water" && s2.getType() == "water"
}

func (wwc *WaterAndWaterCollision) doHandling(s1 ISprite, s2 ISprite) {
	fmt.Printf("water and water collision, cannot move x1 to x2\n")
}
