package main

import (
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	products := []Product{
		{
			ID:          1,
			Name:        "Bamboo Fiber Reusable Coffee Cups",
			Description: "Sustainable and durable coffee cups made from bamboo fibers, ideal for everyday use.",
		},
		{
			ID:          2,
			Name:        "Stainless Steel Travel Mugs",
			Description: "Insulated travel mugs made from stainless steel, perfect for keeping beverages hot or cold while on the go.",
		},
		{
			ID:          3,
			Name:        "Glass Tea Infuser Bottles",
			Description: "Eco-friendly glass bottles with built-in tea infusers, great for brewing loose leaf tea or infused water.",
		},
		{
			ID:          4,
			Name:        "Ceramic Reusable Espresso Cups",
			Description: "Stylish and sustainable espresso cups made from ceramic, suitable for home or office use.",
		},
		{
			ID:          5,
			Name:        "Collapsible Silicone Coffee Cups",
			Description: "Portable and collapsible cups made from silicone, convenient for travel and reducing single-use cup waste.",
		},
		{
			ID:          6,
			Name:        "Cornstarch Biodegradable Coffee Cups",
			Description: "Eco-friendly coffee cups made from cornstarch, biodegradable and compostable.",
		},
	}

	getHelloWorld := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello wrold"))
	}

	getAllProducts := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}

	getProductById := func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)

		if err != nil {
			http.Error(w, "Invalid prouduct ID", http.StatusBadRequest)
		}

		for _, product := range products {

			if product.ID == id {
				json.NewEncoder(w).Encode(product)
				return
			}

		}

		http.Error(w, "Product not found", http.StatusNotFound)
	}

	createProductById := func(w http.ResponseWriter, r *http.Request) {

		var newProductInfo struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		}

		if err := json.NewDecoder(r.Body).Decode(&newProductInfo); err != nil {
			http.Error(w, "Unable to parse body.", http.StatusBadRequest)
			return
		}

		newProductIndex := products[len(products)-1].ID + 1

		newProduct := Product{
			ID:          newProductIndex,
			Name:        newProductInfo.Name,
			Description: newProductInfo.Description,
		}

		products = append(products, newProduct)

		w.Write([]byte("hey, the product is added. Cheers!!!"))

		http.Error(w, "Product not found", http.StatusNotFound)
	}

	updateProductById := func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)

		if err != nil {
			http.Error(w, "Invalid prouduct ID", http.StatusBadRequest)
		}

		var updatedProductInfo struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		}

		if err := json.NewDecoder(r.Body).Decode(&updatedProductInfo); err != nil {
			http.Error(w, "Unable to parse body.", http.StatusBadRequest)
			return
		}

		index := -1
		for currentIdx, product := range products {

			if product.ID == id {
				index = currentIdx
				break
			}

		}

		if index < 0 {
			http.Error(w, "Unable to find the product.", http.StatusBadRequest)
		}

		newProduct := Product{
			ID:          id,
			Name:        updatedProductInfo.Name,
			Description: updatedProductInfo.Description,
		}

		_products := append(products[:index], newProduct)
		products = append(_products, products[index+1:]...)

		w.Write([]byte("hey, the product is updated. Cheers!!!"))

		http.Error(w, "Product not found", http.StatusNotFound)
	}

	deleteProductById := func(w http.ResponseWriter, r *http.Request) {
		idParams := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParams)

		if err != nil {
			http.Error(w, "Invalid prouduct ID", http.StatusBadRequest)
		}

		index := -1

		for currentIdx, product := range products {
			if product.ID == id {
				index = currentIdx
				break
			}
		}

		if index < 0 {
			http.Error(w, "product not found", http.StatusNotFound)
		}

		products = append(products[:index], products[index+1:]...)
		w.Write([]byte("deleted product with id " + idParams))
	}

	r.Get("/", getHelloWorld)
	r.Get("/products", getAllProducts)
	r.Get("/products/{id}", getProductById)
	r.Post("/products/", createProductById)
	r.Put("/products/{id}", updateProductById)
	r.Delete("/products/{id}", deleteProductById)
	http.ListenAndServe(":8080", r)
}
