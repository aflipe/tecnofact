package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutescomidas(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	enableCORS(router)

	router.HandleFunc("/AllComidas", func(w http.ResponseWriter, r *http.Request) {
		AllComidas, err := GetAllComidas()
		if err == nil {
			respondWithSuccess(AllComidas, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

}

func setupRoutescomidasbyid(router *mux.Router) {
	enableCORS(router)

	router.HandleFunc("/AllComidasById/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		comidasbyid, err := getComidasById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(comidasbyid, w)
		}
	}).Methods(http.MethodGet)
}

func setupRoutesbebidas(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	enableCORS(router)

	router.HandleFunc("/AllBebidas", func(w http.ResponseWriter, r *http.Request) {
		AllBebidas, err := GetAllBebidas()
		if err == nil {
			respondWithSuccess(AllBebidas, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

}

func setupRoutesbebidasbyid(router *mux.Router) {
	enableCORS(router)

	router.HandleFunc("/AllBebidasById/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		Bebidasbyid, err := getBebidasById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(Bebidasbyid, w)
		}
	}).Methods(http.MethodGet)
}

func setupRoutesbebidastipo(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	enableCORS(router)

	router.HandleFunc("/AllBebidasTipo", func(w http.ResponseWriter, r *http.Request) {
		AllBebidasTipo, err := GetAllBebidastipo()
		if err == nil {
			respondWithSuccess(AllBebidasTipo, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

}

func setupRoutesbebidasTipobyid(router *mux.Router) {
	enableCORS(router)

	router.HandleFunc("/AllBebidasTipoById/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		BebidasTipobyid, err := getBebidastipoById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(BebidasTipobyid, w)
		}
	}).Methods(http.MethodGet)
}

func setupRoutesRestaurante(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	enableCORS(router)

	router.HandleFunc("/AllRestaurante", func(w http.ResponseWriter, r *http.Request) {
		AllRestaurante, err := GetAllRestaurantes()
		if err == nil {
			respondWithSuccess(AllRestaurante, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

}

func setupRestauranteTipobyid(router *mux.Router) {
	enableCORS(router)

	router.HandleFunc("/AllRestauranteTipoById/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		RestauranteTipobyid, err := getRestaurantesById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(RestauranteTipobyid, w)
		}
	}).Methods(http.MethodGet)
}

func setupCreateOrdenar(router *mux.Router) {
	router.HandleFunc("/CreteOrdenar", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var orden ordenar
		err := json.NewDecoder(r.Body).Decode(&orden)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := createOrden(orden)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

// Helper functions for respond with 200 or 500 code
func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
