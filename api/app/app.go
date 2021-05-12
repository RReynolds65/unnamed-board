package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/api/app/config"
	"github.com/api/app/handler"
	"github.com/api/app/model"
	"github.com/api/gorilla/mux"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB *gorm.DB
}

func (a *App) Initialize(config *config.Config){
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
	config.DB.Username,
	config.DB.Password,
	config.DB.Name,
	config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil{
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()


	func (a *App) setRouters(){
		a.Get("/Posts", a.GetAllPosts)
	}
	
}