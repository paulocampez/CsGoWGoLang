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
	Funcionario []Funcionario `json:"root"`
}

// A Response struct to map the Entire Response
type Funcionario struct {
	ID           int     `json:idFuncionario`
	Name         string  `json:"nome"`
	CPF          string  `json:"cpf"`
	DataCadastro string  `json:"dataCad"`
	Cargo        string  `json:"cargo"`
	UF           string  `json:"ufNasc"`
	Salario      float64 `json:"salario"`
	Status       string  `json:"status"`
}

func main() {
	response, err := http.Get("https://ledacards20190524012526.azurewebsites.net/api/Home/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	//fmt.Println(ioutil.ReadAll(response.Body))
	responseData, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(responseData))
	if err != nil {
		log.Fatal(err)
		fmt.Println("ERRO")
	}

	var responseObject []Funcionario
	json.Unmarshal(responseData, &responseObject)

	//fmt.Println(string(responseObject.Name))
	//fmt.Println(len(responseObject.Name))
	//fmt.Println(responseObject[0].Name)
	for i := 0; i < len(responseObject); i++ {
		fmt.Println(responseObject[i].Name)
	}

}
