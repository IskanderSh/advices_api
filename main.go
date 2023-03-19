package main

import (
	"practice7/files"
	"fmt"
)

func main() {
	word := "life"
	id := 11
	ran_advice := api.RandomAdvice()
	advice_by_id := api.AdviceById(id)
	count_of_advices := api.CountOfQuery(word)
	searched_advices := api.SearchAdvice(word)

	fmt.Println("Random advice is:", ran_advice.Slip.Advice)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Id of random advice is:", ran_advice.Slip.Id)
	fmt.Println("-----------------------------------------------------")
	fmt.Println(fmt.Sprintf("Advice by id %v:", id), advice_by_id.Slip.Advice)
	fmt.Println("-----------------------------------------------------")
	fmt.Println(fmt.Sprintf("Count of advices with word %v:", word), count_of_advices.TotalResults)
	fmt.Println("-----------------------------------------------------")
	fmt.Println(fmt.Sprintf("All this %v advices \n", count_of_advices.TotalResults), searched_advices)
}