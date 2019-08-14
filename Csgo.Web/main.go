package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Root struct {
	Func []Funcionario `json:""`
}

// A Response struct to map the Entire Response
type Funcionario struct {
	ID           int     `json:idFuncionario`
	Name         string  `json:"name"`
	CPF          string  `json:"cpf"`
	DataCadastro string  `json:"dataCad"`
	Cargo        string  `json:"cargo"`
	UF           string  `json:"ufNasc"`
	Salario      float64 `json:"salario"`
	Status       string  `json:"status"`
}

func main() {
	response, err := http.Get("https://ledacards20190524012526.azurewebsites.net/api/Home/BuscaCPF/80478555490")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	//fmt.Println(responseData)
	if err != nil {
		log.Fatal(err)
		fmt.Println("ERRO")
	}

	var responseObject Root
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject)
	//fmt.Println(len(responseObject.Name))

	//	for i := 0; i < len(responseObject.Name); i++ {
	//	fmt.Println(responseObject.Name)
	//	}

}
