package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Image string `json:"image"`
}

func GetItems() []string {
	// URL de la API a la que deseas hacer la solicitud
	url := "https://fakestoreapi.com/products?limit=5"

	// Realizar la solicitud GET a la API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error al hacer la solicitud: %s", err)
		return nil
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error al leer la respuesta: %s", err)
		return nil
	}

	// Deserializar el JSON en una estructura de datos Go
	var products []Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		fmt.Printf("Error al deserializar JSON: %s", err)
		return nil
	}

	// Crear un slice de strings para almacenar los datos formateados (Title, Price, e Image)
	var productStrings []string
	for _, product := range products {
		// Formatear los datos como una cadena y agregarlos al slice
		productString := fmt.Sprintf("%s \nPrice: %.2f \nImage: %s", product.Title, product.Price, product.Image)
		productStrings = append(productStrings, productString)
	}

	// Devolver el slice de strings
	return productStrings
}
