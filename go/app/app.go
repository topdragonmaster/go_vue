package app

import (
	"fmt"
	"log"
	"net/http"
	"sniper/door/app/middleware"

	"sniper/door/app/handler"
	"sniper/door/app/model"
	"sniper/door/config"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {

	userController := handler.NewBaseHandler(a.DB)
	// Routing for handling the projects
	a.Post("/login", userController.Login)
	// a.Post("/getusers", middleware.AuthMiddleware(userController.GetUsers))

	AuthRouter := a.Router.Methods(http.MethodPost).Subrouter()
	AuthRouter.HandleFunc("/getusers", userController.GetUsers).Methods("Post")
	AuthRouter.Use(middleware.AuthMiddleware)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

// func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		handler(a.DB, w, r)
// 	}
// }
