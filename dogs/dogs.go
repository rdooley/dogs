package dogs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Dog struct {
	ID      int
	Name    string
	Owner   string
	Details string
}

var DogFile = "dogs.json"

func NewDog(name, owner, details string) Dog {
	currentDogs := LoadDogs()
	var ID int
	// Cheap auto inc ID
	for _, dog := range currentDogs {
		if dog.ID > ID {
			ID = dog.ID
		}
	}
	ID++
	newDog := Dog{
		ID:      ID,
		Name:    name,
		Owner:   owner,
		Details: details,
	}
	currentDogs = append(currentDogs, newDog)
	WriteDogs(currentDogs)
	return newDog
}

// Load all the dogs
func LoadDogs() (dogs []Dog) {
	content, err := ioutil.ReadFile(DogFile)
	if err != nil {
		log.Print("dog.json does not exist")
		return dogs
	}
	err = json.Unmarshal(content, &dogs)
	if err != nil {
		log.Print("Error reading dogs.json")
		return dogs
	}
	return dogs
}

// Load a specific dog, bool that is true if it is found
func LoadDog(ID int) (Dog, bool) {
	dogs := LoadDogs()
	for _, dog := range dogs {
		if ID == dog.ID {
			return dog, true
		}
	}
	return Dog{}, false
}

// Write Dogs to our json file
func WriteDogs(dogs []Dog) {
	json, _ := json.MarshalIndent(dogs, "", "\t")
	_ = ioutil.WriteFile(DogFile, json, 0644)
}

// Update a dog
func UpdateDog(dog Dog, name, owner, details string) Dog {
	if name != "" {
		dog.Name = name
	}
	if owner != "" {
		dog.Owner = owner
	}
	if details != "" {
		dog.Details = details
	}
	dogs := LoadDogs()
	var index int
	for i, d := range dogs {
		if d.ID == dog.ID {
			index = i
		}
	}
	dogs = append(dogs[:index], dogs[index+1:]...)
	dogs = append(dogs, dog)

	WriteDogs(dogs)
	return dog
}
