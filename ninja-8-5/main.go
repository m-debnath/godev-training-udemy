package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (n ByLast) Len() int           { return len(n) }
func (n ByLast) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByLast) Less(i, j int) bool { return n[i].Last < n[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("Original slice")
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, val := range v.Sayings {
			fmt.Println("\t", val)
		}
	}

	// sort by age
	fmt.Println("\nAfter sort by age:")
	sort.Sort(ByAge(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, val := range v.Sayings {
			fmt.Println("\t", val)
		}
	}

	// sort by last name
	fmt.Println("\nAfter sort by last name:")
	sort.Sort(ByLast(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, val := range v.Sayings {
			fmt.Println("\t", val)
		}
	}

	// sorting Sayings for each user
	for _, v := range users {
		sort.Strings(v.Sayings)
	}
	fmt.Println("\nAfter sorting all sayings:")
	sort.Sort(ByLast(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, val := range v.Sayings {
			fmt.Println("\t", val)
		}
	}
}
