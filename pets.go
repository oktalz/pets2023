package pets

import "math/rand"

var prefix = [...]string{"big", "small", "tiny", "angry"}
var pet = [...]string{"cat", "dog", "fish", "elephant"}

func PetName() string {
	return prefix[rand.Intn(len(prefix))] + "-" + pet[rand.Intn(len(pet))]
}
