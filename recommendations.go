package main

/*
recommendations.go

This file includes data and concepts from Chapter 2 of Programming Collective Intelligence.
It emulates what the author is referring to as the file "recommendations.py" including the
dataset, which he asks to be included in the file.
*/
import (
	"encoding/json"
	"fmt"
	"math"
)

/*
We need a dictionary of movie critics and their ratings of a small set of movies
format:
    critics = {"Lisa Rose": {"Lady in the Water": 2.5,
                             "Snakes on a Plane": 3.5,
                             "Just My Luck": 3.0,
                             "Superman Returns": 3.5,
                             "You, Me and Dupree": 2.5,
                             "The Night Listener": 3.0},
               "Gene Seymour": {"Lady in the Water": 3.0,
                                "Snakes on a Plane": 3.5,
                                "Just My Luck": 1.5,
                                "Superman Returns": 5.0,
                                "The Night Listener": 3.0,
                                "You, Me and Dupree": 3.5},
                "Michael Phillips": {"Lady in the Water": 2.5,
                                     "Snakes on a Plane": 3.0,
                                     "Superman Returns": 3.5,
                                     "The Night Listener": 4.0},
                "Claudia Puig": {"Snakes on a Plane": 3.5,
                                "Just My Luck": 3.0,
                                "The Night Listener": 4.5,
                                "Superman Returns": 4.0,
                                "You, Me and Dupree": 2.5},
                "Mick LaSalle": {"Lady in the Water": 3.0,
                                 "Snakes on a Plane": 4.0,
                                 "Just My Luck": 2.0,
                                 "Superman Returns": 3.0,
                                 "The Night Listener": 3.0,
                                 "You, Me and Dupree": 2.0},
                "Jack Matthews": {"Lady in the Water": 3.0,
                                  "Snakes on a Plane": 4.0,
                                  "The Night Listener": 3.0,
                                  "Superman Returns": 5.0,
                                  "You, Me and Dupree": 3.5},
                "Toby": {"Snakes on a Plane":4.5,
                         "You, Me and Dupree":1.0,
                         "Superman Returns":4.0}}
*/

// simDistance computes a distance-based similarity score for person 1 and person 2
// This is based on Euclidean distance, and accepts a string:value as the data input
func simDistance(data map[string]interface{}, person1 string, person2 string) float64 {
	// fmt.Println("Running simDistance")
	// Get list of shared items
	var shared = map[string]int{}
	sumOfSquares := 0.0
	// for item in prefs
	d1 := data[person1].(map[string]interface{})
	d2 := data[person2].(map[string]interface{})
	for i := range d1 {
		for j := range d2 {
			if j == i {
				if a1, ok := d1[i].(float64); !ok {
					panic("person1 invalid type")
				} else if a2, ok := d2[i].(float64); !ok {
					panic("person2 invalid type")
				} else {
					shared[i] = 1
					tempResult := (a2 - a1)
					// fmt.Printf("\nData for person2: %v\n", tempResult)
					// fmt.Printf("\n Delta is this: %v\n", tempResult)
					sumOfSquares = sumOfSquares + math.Pow(tempResult, 2)
				}
			}
		}
	}
	if len(shared) == 0 {
		return 0
	}
	return 1 / (1 + sumOfSquares)
}

func main() {
	payload := []byte(`{"Lisa Rose": {"Lady in the Water": 2.5,
                             "Snakes on a Plane": 3.5,
                             "Just My Luck": 3.0,
                             "Superman Returns": 3.5,
                             "You, Me and Dupree": 2.5,
                             "The Night Listener": 3.0},
               "Gene Seymour": {"Lady in the Water": 3.0,
                                "Snakes on a Plane": 3.5,
                                "Just My Luck": 1.5,
                                "Superman Returns": 5.0,
                                "The Night Listener": 3.0,
                                "You, Me and Dupree": 3.5},
                "Michael Phillips": {"Lady in the Water": 2.5,
                                     "Snakes on a Plane": 3.0,
                                     "Superman Returns": 3.5,
                                     "The Night Listener": 4.0},
                "Claudia Puig": {"Snakes on a Plane": 3.5,
                                "Just My Luck": 3.0,
                                "The Night Listener": 4.5,
                                "Superman Returns": 4.0,
                                "You, Me and Dupree": 2.5},
                "Mick LaSalle": {"Lady in the Water": 3.0,
                                 "Snakes on a Plane": 4.0,
                                 "Just My Luck": 2.0,
                                 "Superman Returns": 3.0,
                                 "The Night Listener": 3.0,
                                 "You, Me and Dupree": 2.0},
                "Jack Matthews": {"Lady in the Water": 3.0,
                                  "Snakes on a Plane": 4.0,
                                  "The Night Listener": 3.0,
                                  "Superman Returns": 5.0,
                                  "You, Me and Dupree": 3.5},
                "Toby": {"Snakes on a Plane":4.5,
                         "You, Me and Dupree":1.0,
                         "Superman Returns":4.0}
                }`)
	var critics map[string]interface{}
	err := json.Unmarshal(payload, &critics)
	if err != nil {
		panic(err)
	}
	// Here's where we print examples like found on page 9
	fmt.Println(critics["Lisa Rose"])
	fmt.Println(critics["Toby"])
	fmt.Println(critics["Toby"].(map[string]interface{})["Snakes on a Plane"])

	fmt.Printf("\n\nsimDistance between Lisa Rose and Gene Seymour: \n%v\n", simDistance(critics, "Lisa Rose", "Gene Seymour"))
	// Is there a way to set it up to lookup with critics["Toby"]["Snakes on a Plane"] ?
	// I cannot figure out how to make the type/struct to fit properly to accomplish taht
}
