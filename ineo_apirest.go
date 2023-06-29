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

var (
	APIKEY_FAKE = "cd66a4f9-8a9b-4a8c-a02a-ff2d7d1c3e3c"
	APIKEY1     = "c6aa27d6-fc9b-4223-8665-1de556cd7a09"
	APIKEY2     = "b4e77488-5330-2a27-2d07-4ce64d0af88d"

	//URL_BASE+":"+PORT+URL_ADDRESS1+URL_ADDRESS2
	//URL_BASE="http://51.178.19.210"
	URL_BASE     = "http://127.0.0.1" //localhost server
	URL_ADDRESS1 = "/api/v1/"
	URL_ADDRESS2 = "gateways"
	PORT_FAKE    = "8080" //Fake PORT localhost
	PORT1        = "51882"
	PORT2        = "1882"
)

func main() {
	// URL de la API
	//url := "http://51.178.19.210:1882/api/v1/pirs/70b3d56371d3e040/"
	//url := "http://51.178.19.210:1882/api/v1/pirs/"
	url := URL_BASE + ":" + PORT_FAKE + URL_ADDRESS1 + URL_ADDRESS2

	// créer une nouvelle requête GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Échec de la demande GET:", err)
		return
	}

	// Configurer les headers requis
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("APIKey", APIKEY_FAKE) //APIKEY

	fmt.Println("URL :", url)
	fmt.Println("Content-Type :", req.Header.Get("Content-Type")) //DEBUG Content-Type
	fmt.Println("APIKey :", req.Header.Get("APIKey"))             //DEBUG APIKey

	// Faire la requête HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Échec de la demande:", err)
		return
	}
	defer resp.Body.Close()

	// lire la réponse sous forme de chaîne de texte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur de lecture de la réponse:", err)
		return
	}
	fmt.Println(string(body))
}
