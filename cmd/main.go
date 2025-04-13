package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/configs"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/dataBase"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/handlers"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	err := configs.Environment()
	if err != nil {
		log.Fatalf("Error environment func: %v", err)
	}

	db := dataBase.ConnectToDB()
	defer db.Close()

	router := mux.NewRouter()

	authMiddleware := middleware.NewAuthMiddleware()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/dummyLogin") {
				next.ServeHTTP(w, r)
				return
			}
			authMiddleware(next).ServeHTTP(w, r)
		})
	})

	RegisterHandlers(router, db)

	port := configs.APIPort
	if port == "" {
		port = "8080"
	}
	port = ":" + port

	log.Println("application running on port" + port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}
}

func RegisterHandlers(router *mux.Router, db *pgxpool.Pool) {
	router.HandleFunc("/dummyLogin", handlers.DummyLoginHandler).Methods(http.MethodPost)
	// router.HandleFunc("/register", handlers.RegisterHandler).Methods(http.MethodPost)
	// router.HandleFunc("/login", handlers.LoginHandler).Methods(http.MethodPost)

	router.HandleFunc("/pvz", handlers.CreatePVZHandler).Methods(http.MethodPost)
	router.HandleFunc("/pvz", handlers.GetPVZListHandler).Methods(http.MethodGet)

	router.HandleFunc("/receptions", handlers.CreateReceptionHandler).Methods(http.MethodPost)
	router.HandleFunc("/products", handlers.AddProductHandler).Methods(http.MethodPost)
	router.HandleFunc("/pvz/{pvzId}/delete_last_product", handlers.DeleteLastProductHandler).Methods(http.MethodPost)
	router.HandleFunc("/pvz/{pvzId}/close_last_reception", handlers.CloseLastReceptionHandler).Methods(http.MethodPost)
}
