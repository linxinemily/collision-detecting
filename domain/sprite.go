package domain

import (
	"fmt"
)

type ISprite interface {
	getCoordinate() int
	setCoordinate(coordinate int)
	getType() string
	setWorld(world *World)
	removeFromWorld()
	updateCoordinate(coordinate int)
}

func NewSprite(_type string) *Sprite {
	return &Sprite{
		_type: _type,
	}
}

type Sprite struct {
	coordinate int
	_type      string
	world      *World
}

func (s *Sprite) getCoordinate() int {
	return s.coordinate
}

func (s *Sprite) setCoordinate(coordinate int) {
	s.coordinate = coordinate
}

func (s *Sprite) getType() string {
	return s._type
}

func (s *Sprite) setWorld(world *World) {
	s.world = world
}

func (s *Sprite) removeFromWorld() {
	s.world.sprites[s.coordinate] = nil
	fmt.Println("after remove")
	PrintItems(s.world.sprites)
}

func (s *Sprite) updateCoordinate(coordinate int) {
	if s.world.getSpriteInPosition(coordinate) != nil {
		panic("new coordinate is not nil")
	}
	s.world.Move(s.coordinate, coordinate)
	s.coordinate = coordinate

	fmt.Println("after update")
	PrintItems(s.world.sprites)
}

/** function for debugging **/
func (s Sprite) String() string {
	switch s.getType() {
	case "hero":
		return "🦸"
	case "water":
		return "💧"
	case "fire":
		return "🔥"
	default:
		return "unknown"
	}
}

/** function for debugging **/
func PrintItems(items []ISprite) {
	fmt.Println(items)
	for index, val := range items {
		if val == nil {
			fmt.Printf("[%v] nil ", index)
		} else {
			fmt.Printf("[%v] %s ", index, val)
			if hero, ok := val.(*Hero); ok {
				fmt.Printf("HP: %d", hero.HP)
			}
		}
	}
	fmt.Println()
}
