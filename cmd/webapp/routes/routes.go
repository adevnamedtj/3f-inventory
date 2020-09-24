package routes

import (
	"github.com/ckalagara/3f-inventory/cmd/webapp/routes/resources"
	"github.com/go-chi/chi"
)

// Mount handles mounting for all resources
func Mount(r *chi.Mux) {

	// mounting product routes
	r.Mount(resources.ProductResourceMountPoint, resources.ProductResource{}.Routes())
}
