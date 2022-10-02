package domain

import (
	"fmt"
)

type FireAndFireCollision struct {
	*AbstractCollision
}

func NewFireAndFireCollision(next *ICollision) *FireAndFireCollision {
	return &FireAndFireCollision{
		AbstractCollision: NewAbstractCollision(next),
	}
}

func (wwc *FireAndFireCollision) match(s1 ISprite, s2 ISprite) bool {
	return s1.getType() == "fire" && s2.getType() == "fire"
}

func (wwc *FireAndFireCollision) doHandling(s1 ISprite, s2 ISprite) {
	fmt.Printf("fire and fire collision, cannot move x1 to x2\n")
}
