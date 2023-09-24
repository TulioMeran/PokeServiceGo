package main

import (
	"TulioMeran/example/services"
	"fmt"
)

func main() {

	apiService := services.PokeServiceImpl{}

	getAllChannel := make(chan int)
	getPokemonByNameChannel := make(chan int)

	go func() {
		data, err := apiService.GetAll()

		if err != nil {
			fmt.Errorf(err.Error())
			return
		}

		fmt.Println(data)
		getAllChannel <- 1
	}()

	go func() {
		data, err := apiService.GetPokemonByName("pikachu")

		if err != nil {
			fmt.Errorf(err.Error())
			return
		}

		fmt.Println(data)
		getPokemonByNameChannel <- 2
	}()

	<-getAllChannel
	<-getPokemonByNameChannel
	fmt.Println("finished")

}
