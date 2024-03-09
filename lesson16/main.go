package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	transportCosts := map[string]float64{}

	fmt.Println(transportCosts)

	transportCosts["bus"] = 2
	transportCosts["miniBus"] = 2.5
	transportCosts["taxi"] = 10
	transportCosts["123 fdff"] = 5.6
	transportCosts["bicycle"] = 0

	fmt.Println(transportCosts)

	delete(transportCosts, "bus")

	fmt.Println(transportCosts)

	delete(transportCosts, "bus")

	s, _ := json.MarshalIndent(transportCosts, "", "\t")
	fmt.Println(string(s))

	fmt.Println("Price for mini-bus:", transportCosts["miniBus"])
	fmt.Println("Price for taxi:", transportCosts["taxi"])
	fmt.Println("Price for bus:", transportCosts["bus"])
	fmt.Println("Price for bus:", transportCosts["afefsrgrsgdrsgrs"])

	fmt.Println("\nBeautifully printed:")
	for key, value := range transportCosts {
		fmt.Println("Price for", key, "is:", value, "somoni")
		value = 34
		key = "abc"
	}

	fmt.Println(transportCosts)

	if transportCosts["bicycle"] == 0 {
		fmt.Println("Такого ключа нет")
	}

	bicycle, exists := transportCosts["bicycle"]
	taxi, b := transportCosts["taxi"]
	airplane, xyz := transportCosts["airplane"]

	fmt.Println(bicycle, exists)
	fmt.Println(taxi, b)
	fmt.Println(airplane, xyz)

	transportList := []string{
		"bus",
		"skateboard",
		"ship",
		"airplane",
		"bicycle",
		"miniBus",
		"elephant",
		"taxi",
	}

	fmt.Println(transportList)

	for _, transport := range transportList {
		fmt.Println(transport)
	}

	fmt.Println("\nTable of costs:")
	for _, transport := range transportList {
		//cost, ok := transportCosts[transport]
		//if ok {
		//	fmt.Println("Price for", transport, "is:", cost, "somoni")
		//}

		if cost, ok := transportCosts[transport]; ok {
			fmt.Println("Price for", transport, "is:", cost, "somoni")
		}
	}

	fmt.Println(transportCosts["spaceShip"])

	transportCosts["spaceShip"] -= 10000

	fmt.Println(transportCosts["spaceShip"])

	fmt.Printf("%T\t%v\n", transportCosts, transportCosts)

	fmt.Println("\nExperimenting with values types:")
	people := map[string][]string{}

	fmt.Printf("%T\t%v\n", people, people)

	people["teachers"] = []string{
		"Siyovush",
		"Farhod",
		"Akai Khurshed",
	}

	fmt.Println(people)

	people["students"] = []string{
		"Faridun",
		"Muhammad",
		"Ubaidullo",
		"Shokirjon",
	}

	fmt.Println(people)

	people["managers"] = []string{
		"Firuza",
		"Umedakhon",
	}

	fmt.Println(people)

	s, _ = json.MarshalIndent(people, "", "\t")
	fmt.Println(string(s))

	experiment := map[string]*[]*float64{}
	fmt.Printf("%T\t%v\n", experiment, experiment)

	x := 4.5
	experiment["first"] = &[]*float64{
		&x,
	}

	fmt.Println(experiment, *(*experiment["first"])[0])

	experiment2 := map[int]map[string][]bool{}
	fmt.Printf("%T\t%v\n", experiment2, experiment2)

	experiment2[100] = map[string][]bool{
		"First": []bool{
			false,
			false,
			true,
		},
		"Second": {
			true,
			true,
			true,
			true,
			true,
		},
	}

	experiment2[1000] = map[string][]bool{
		"Winter": []bool{
			true,
		},
		"Summer": {
			false,
			false,
		},
		"Spring": []bool{
			true,
			false,
			true,
		},
		"Fall": {
			false,
			false,
			true,
			false,
		},
	}

	fmt.Println(experiment2)

	s, _ = json.MarshalIndent(experiment2, "", "\t")
	fmt.Println(string(s))

}
