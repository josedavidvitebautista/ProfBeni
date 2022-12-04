package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getBenito(numb int) {
	channel := make(chan string, 10)
	resp, erro := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(numb))
	if erro != nil {
		panic(erro)
	}
	defer resp.Body.Close()
	fmt.Println("Status: ", resp.Status)

	lector := bufio.NewScanner(resp.Body)

	for i := 0; lector.Scan(); i++ {
		channel <- lector.Text()
	}

	if erro := lector.Err(); erro != nil {
		panic(erro)
	}

	beni := <-channel
	fmt.Println(beni)

}

func getBenito2() {
	channel := make(chan string, 10000)

	url := "https://api.covidtracking.com/v1/us/daily.json"
	req, _ := http.NewRequest("GET", url, nil)
	resp, erro := http.Get("https://api.covidtracking.com/v1/us/daily.json")
	if erro != nil {
		panic(erro)
	}
	defer resp.Body.Close()

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	lector, _ := ioutil.ReadAll(res.Body)

	channel <- string(lector)

	beni := <-channel
	fmt.Println(beni)

}

func main() {
	fmt.Println("Hola Mundo :v")
	for i := 1; i <= 10; i++ {
		go getBenito(i)
	}
	go getBenito2()

	var input string
	fmt.Scanln(&input)

}
