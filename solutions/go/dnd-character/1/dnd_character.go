package dndcharacter
import (
    "math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	newScore := score - 10
    return int(math.Floor(float64(newScore) / float64(2)))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return rand.Intn(18 - 3 + 1) + 3
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	myChar := &Character{Constitution : Ability()}
    return Character{
        Strength : Ability() ,
        Dexterity : Ability(),
        Constitution: myChar.Constitution,
        Intelligence: Ability(),
        Wisdom: Ability(),
        Charisma: Ability(),
        Hitpoints: 10 + Modifier(myChar.Constitution),
    }
}
