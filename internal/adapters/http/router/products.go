package router

import (
	"fmt"
	"net/http"
	"poke-store/internal/application/products"

	"github.com/gorilla/mux"
)

const (
	CreateProductRoute = "/v1/products"
)

type ProductRouter struct { //adds router to route
	router  *mux.Router
	service products.Service //expecting the initialization of product router, product service
}

func NewProductRouter(router *mux.Router) ProductRouter {
	service := products.NewService()
	return ProductRouter{router: router, service: service}
}

func (pr *ProductRouter) trialing(response http.ResponseWriter, request *http.Request) {
	// en el router van: 1- request validation, 2- mapeo de status code, 3- mapeo de respuestas(servicios)
	// 1. request body, 2. meterlo, 3. validarlo, 4. meterlo a la capa de servicio (llamar a la capa de servicio)
	// single of responsibility
	fmt.Printf("%v", request.Body)
	pr.service.Save()
}

func (pr *ProductRouter) ConfigureRoutes() {
	pr.router.HandleFunc(CreateProductRoute, pr.trialing).Methods("POST")
}
