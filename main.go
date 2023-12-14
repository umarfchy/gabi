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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello wrold"))
	})

	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(products)

	})

	r.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
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

	})

	http.ListenAndServe(":8080", r)
}
