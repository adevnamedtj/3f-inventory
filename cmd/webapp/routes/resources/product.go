package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

const ProductResourceMountPoint = "/products"

type ProductResource struct {
}

// Routes handle router configuration for product
func (rs ProductResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", rs.List)    // GET /products - returns a list of products
	r.Post("/", rs.Create) // POST /products - create a new product
	r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)       // GET /products/{id} - return product by :id
		r.Put("/", rs.Update)    // PUT /products/{id} - update a product by :id
		r.Delete("/", rs.Delete) // DELETE /products/{id} - delete a product by :id
	})

	return r
}

// List : GET /products - returns a list of products
func (rs ProductResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("product list"))
}

// Create : POST /products - create a new product
func (rs ProductResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("product create"))
}

// Get : GET /products/{id} - return product by :id
func (rs ProductResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("product get"))
}

// Update : PUT /products/{id} - update a product by :id
func (rs ProductResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("product update"))
}

// Delete : // DELETE /products/{id} - delete a product by :id
func (rs ProductResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("product delete"))
}
