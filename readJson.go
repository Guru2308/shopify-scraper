package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FeaturedImage struct {
	Link string `json:"src"`
}

type Variant struct {
	Title         string        `json:"title"`
	Price         string        `json:"price"`
	FeaturedImage FeaturedImage `json:"featured_image"`
}

type Product struct {
	Title   string    `json:"title"`
	Variant []Variant `json:"variants"`
}

type ProductsResponse struct {
	Products []Product `json:"products"`
}

type VariantData struct {
	ImageLink    string
	ProductTitle string
	VariantName  string
	Price        string
}

func ReadJSON() {
	files, err := os.ReadDir("./json/")
	if err != nil {
		fmt.Printf("Error reading directory: %s", err)
	}

	for _, file := range files {
		filePath := fmt.Sprintf("./json/%s", file.Name())

		jsonFile, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error while reading: ", err)
		}

		var products ProductsResponse
		err = json.Unmarshal(jsonFile, &products)

		if err != nil {
			fmt.Printf("Error unmarshalling the JSON file: %s\n", err)
			return
		}

		for _, product := range products.Products {
			fmt.Printf("Product name - %s\n", product.Title)
			for _, variant := range product.Variant {
				fmt.Printf("\tName - %s, Price - %s, Link - %s\n", variant.Title, variant.Price, variant.FeaturedImage.Link)
			}
		}
	}
}

/* func Server() {
	fmt.Println("Hello")
} */
