package domain

type ICollision struct {
	Collision
}

func NewICollision(c Collision) *ICollision {
	return &ICollision{
		Collision: c,
	}
}

func (ac *ICollision) handle(s1 ISprite, s2 ISprite) {
	if ac.match(s1, s2) {
		ac.doHandling(s1, s2)
	} else if ac.getNext() != nil {
		ac.getNext().handle(s1, s2)
	}
}

type Collision interface {
	match(s1 ISprite, s2 ISprite) bool
	doHandling(s1 ISprite, s2 ISprite)
	getNext() *ICollision
}

type AbstractCollision struct {
	next *ICollision
}

func NewAbstractCollision(n *ICollision) *AbstractCollision {
	return &AbstractCollision{
		next: n,
	}
}

func (ac *AbstractCollision) getNext() *ICollision {
	return ac.next
}
