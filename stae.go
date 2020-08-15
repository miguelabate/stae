package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type Player struct {
	name string
	currentLocation *Room
}

type Room struct {
	name	string
	description	string
	connections map[string]*Room
	contents []interface{}
}

func (p *Player) getCurrentLocation() *Room{
	return p.currentLocation
}

func (p *Player) move(direction string) (*Room, error){
	if newRoom, ok := p.currentLocation.connections[direction]; ok {
		//remove the player fromt his current room
		for i, aContent := range p.currentLocation.contents {
			if IsInstanceOf(&aContent,(*Player)(nil)) {
				p.currentLocation.contents = append(p.currentLocation.contents[:i], p.currentLocation.contents[i+1:]...)
			}
		}
		newRoom.contents = append(newRoom.contents, p) //add to the new room
		p.currentLocation = newRoom //set him in the new room
		return newRoom, nil
	} else {
		log.Printf("Error: playe cant move into that direction: "+direction)
	}
	return (*Room)(nil), errors.New("invalid direction")
}

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}

func main() {
	fmt.Println("Hello, world.")
	kitchen := &Room{"Kitchen", "A kitchen",nil,make([]interface{},1)} //or new (Room)
	//fmt.Println(kitchen)

	connr2 := make(map[string]*Room)
	connr2["n"] = kitchen
	living := &Room{"Living room", "A living",connr2,make([]interface{},1)}

	connr1 := make(map[string]*Room)
	connr1["s"] = living
	kitchen.connections = connr1

	player := Player{"Mike", living}
	living.contents =append(living.contents, player)

	attic := &Room{"Attic", "The dark attic",nil,make([]interface{},0)}
	connr2["u"] = attic //from living to attic

	connr3 := make(map[string]*Room)
	attic.connections = connr3
	connr3["d"] = living

	fmt.Println("Player loc: "+player.getCurrentLocation().name)
	player.move("n")
	fmt.Println("Player loc: "+player.getCurrentLocation().name)
	player.move("s")
	fmt.Println("Player loc: "+player.getCurrentLocation().name)
	player.move("u")
	fmt.Println("Player loc: "+player.getCurrentLocation().name)
	player.move("d")
	fmt.Println("Player loc: "+player.getCurrentLocation().name)

	var command string

	for ok := true; ok; ok = command != "exit" {
		//reading a string
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(">")
		command, _ = reader.ReadString('\n')
		command = strings.Trim(command, "\n")
		player.move(command)
		fmt.Println("Player loc: "+player.getCurrentLocation().name)
	}

}
