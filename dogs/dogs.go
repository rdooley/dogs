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

func NewDog(name, owner, details string) (Dog, error) {
	currentDogs := LoadDogs()
	var ID int
	// Cheap auto inc ID
	for _, dog := range currentDogs {
		ID = dog.ID
	}
	ID++
	newDog := Dog{
		ID:      ID,
		Name:    name,
		Owner:   owner,
		Details: details,
	}
	log.Printf(name)
	log.Printf(owner)
	log.Printf(details)
	currentDogs = append(currentDogs, newDog)
	err := WriteDogs(currentDogs)
	if err != nil {
		return newDog, err
	}
	return newDog, nil
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
func WriteDogs(dogs []Dog) error {
	json, err := json.MarshalIndent(dogs, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(DogFile, json, 0644)
	if err != nil {
		return err
	}
	return nil
}
