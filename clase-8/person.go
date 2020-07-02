package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func createPerson(url, token string, person *Person) GeneralResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)
	if err != nil {
		log.Fatalf("Error en marshal de persona: %v", err)
	}

	resp := httpClient(http.MethodPost, url, token, data)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error en lectura del body en login: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Se esperaba código 201, se obtuvo: %d, respuesta: %s", resp.StatusCode, string(body))
	}

	dataResponse := GeneralResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error en el unmarshal del body en login: %v", err)
	}

	return dataResponse
}
