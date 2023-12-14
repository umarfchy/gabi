package main

import (
	"encoding/json"

	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lucsky/cuid"
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func generateProduct(noOfItems int) []Product {
	_products := []Product{}

	for i := 0; i < noOfItems; i++ {
		_product := Product{
			ID:          cuid.New(),
			Name:        gofakeit.ProductName(),
			Description: gofakeit.ProductDescription(),
		}

		_products = append(_products, _product)
	}

	return _products
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	products := generateProduct(10)

	getHelloWorld := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello wrold"))
	}

	getAllProducts := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}

	getProductById := func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

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

		newProductId := cuid.New()
		newProduct := Product{
			ID:          newProductId,
			Name:        newProductInfo.Name,
			Description: newProductInfo.Description,
		}

		products = append(products, newProduct)

		w.Write([]byte("hey, the product is added. Cheers!!!"))

		http.Error(w, "Product not found", http.StatusNotFound)
	}

	updateProductById := func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

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
		id := chi.URLParam(r, "id")

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
		w.Write([]byte("deleted product with id " + id))
	}

	r.Get("/", getHelloWorld)
	r.Get("/products", getAllProducts)
	r.Get("/products/{id}", getProductById)
	r.Post("/products/", createProductById)
	r.Put("/products/{id}", updateProductById)
	r.Delete("/products/{id}", deleteProductById)
	http.ListenAndServe(":8080", r)
}
