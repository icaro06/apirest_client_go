/*
APIRest for ineo  A.Villanueva

OK get gateways with APIKey
curl -X GET -H "Content-Type: application/json" -H "APIKey: cd66a4f9-8a9b-4a8c-a02a-ff2d7d1c3e3c" 127.0.0.1:8080/api/v1/gateways

L’URL de base sera : http://{{server_IP}}:51882/api/v1/...
Système en production (sur le réseau LPA)
L’URL de base sera : http://192.168.1.52:1882/api/v1/...
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL de la API
	//url := "http://51.178.19.210:8080/#/organizations/1/applications"
	//url := "http://51.178.19.210:1882/api/v1/pirs/70b3d56371d3e040/"
	//url := "http://51.178.19.210:1882/api/v1/pirs/"
	url := "http://127.0.0.1:8080/api/v1/gateways"

	// Crear una nueva solicitud GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Échec de la demande GET:", err)
		return
	}

	// Configurar los headers requeridos
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("APIKey", "cd66a4f9-8a9b-4a8c-a02a-ff2d7d1c3e3c") //Fake APIKEY
	/*
		req.Header.Set("APIKey Niveau 1", "c6aa27d6-fc9b-4223-8665-1de556cd7a09")
		req.Header.Set("APIKey Niveau 2", "b4e77488-5330-2a27-2d07-4ce64d0af88d")

		fmt.Println("Content-Type :", req.Header.Get("Content-Type"))
		fmt.Println("APIKey Niveau 1 :", req.Header.Get("APIKey Niveau 1"))
		fmt.Println("APIKey Niveau 2 :", req.Header.Get("APIKey Niveau 2"))
	*/
	// Realizar la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Échec de la demande:", err)
		return
	}
	defer resp.Body.Close()

	// Leer y procesar la respuesta
	// ...

	// Ejemplo de lectura de la respuesta como cadena de texto
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}
	fmt.Println(string(body))
}
