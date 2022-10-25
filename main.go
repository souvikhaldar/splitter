package main

import "fmt"

func main() {
	nameToExp := make(map[string]int)
	var numPeeps int
	for {
		var name string
		fmt.Println("Enter name")
		fmt.Scanln(&name)
		var total int
		for {
			var item string
			var exp int
			fmt.Println("Enter item:")
			fmt.Scanln(&item)
			fmt.Println("Enter expenditure:")
			fmt.Scanln(&exp)
			total += exp
			var done string
			fmt.Println("Done with items?(y/n)")
			fmt.Scanln(&done)
			if done == "y" {
				fmt.Printf("Total expenditure by %s is %d\n", name, total)
				nameToExp[name] = total
				total = 0
				break
			}
		}
		numPeeps++

		var itemDone string
		fmt.Println("Done?(y/n)")
		fmt.Scanln(&itemDone)
		if itemDone == "y" {
			perPay := make(map[string]float)
			var totalExp int
			for name, exp := range nameToExp {
				fmt.Printf("%s spent %d\n", name, exp)
				totalExp += exp
			}
			fmt.Println("Total trip expenditure is: ", totalExp)
			for name, exp := range nameToExp {
				fmt.Printf("%s needs to pay %d\n", name, ((totalExp / numPeeps) - exp))
				perPay[name] = (totalExp / numPeeps) - exp
			}
			break
		}
	}
}
