package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

func RandomAdvice() Advice {
	// Структура для осуществления HTTP-запросов
	client := &http.Client{}
	// создание нового запроса
	req, err := http.NewRequest("GET", "https://api.adviceslip.com/advice", nil)
	if err != nil {
		log.Fatal(err)
	}
	// выполнение запроса с записью в переменную resp
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// отложенное закрытие запроса
	defer resp.Body.Close()
	// преобразование тела запроса в набор байт
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// объявление переменной для записи полученных данных
	var data Advice
	// запись данных в структуру
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func AdviceById(id int) Advice {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.adviceslip.com/advice/%v", id), nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
		log.Fatal(err)
	}
	var data Advice
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func SearchAdvice(query string) []string {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.adviceslip.com/advice/search/%v", query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}	
	resp, err := client.Do(req)
	if err!=nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Searchadvice
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]string, 0)
	for _, slip := range data.Slips {
		result = append(result, slip.Advice + "\n")
	}
	return result
}

func CountOfQuery(query string) Searchadvice {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.adviceslip.com/advice/search/%v", query)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	
	defer response.Body.Close()
	bodyByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Searchadvice
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}