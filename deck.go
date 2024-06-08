package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// basically just think of this as a class, remember a class has several methods and property
// we can then create an instance of this class and access it properties and methods.

//this is a receiver on a function
//reciever sets up methods on variables that we create

// every variable of type deck now gets access to the print method
// d is the instance of the variable we are working with, or an actual copy of the variable
// we call it d and not cards because by convention we match it with the first or two letter of the actual type
// not compulsory though
// (d deck) is called a reciever, reciever sets up methods on variables that we create
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//we replaced i and j with _ because in golang, we get an error when we dont use a declared variable
	// and since we dont need to read i and j values in our loop, can should replace it with _
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {

	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	newString := strings.Split(string(bs), ",")
	return deck(newString)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// func twoSums(nums []int, target int) []int {
// 	placeholder  := []int{}
// 	len := len(placeholder);
// 	for i := 0; i < len; i++ {
// 		for j := i + 1; j < len; j ++ {
// 			if nums[i] + nums[j] == target {
// 			placeholder = 	append(placeholder, i);
// 			placeholder =	 append(placeholder, j);
// 			}
// 		}
// 	}
// 	return placeholder;
// }
