//go:generate ./vendor/github.com/goadesign/goa/goagen/goagen main -d github.com/potatogopher/goa-uuid-ex/design -o ./controllers
//go:generate ./vendor/github.com/goadesign/goa/goagen/goagen app -d github.com/potatogopher/goa-uuid-ex/design --notest
//go:generate ./vendor/github.com/goadesign/goa/goagen/goagen client -d github.com/potatogopher/goa-uuid-ex/design --notool
//go:generate ./vendor/github.com/goadesign/goa/goagen/goagen swagger -d github.com/potatogopher/goa-uuid-ex/design -o ./public
//go:generate ./vendor/github.com/goadesign/goa/goagen/goagen --design=github.com/potatogopher/goa-uuid-ex/design gen --pkg-path=github.com/goadesign/gorma

package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/potatogopher/goa-uuid-ex/config"
	"github.com/potatogopher/goa-uuid-ex/controllers"
	"github.com/potatogopher/goa-uuid-ex/helpers"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open("postgres", config.Database.String())
	if err != nil {
		panic(err)
	}
	gormDB.DB().SetMaxOpenConns(50)

	service := helpers.MyNewService()

	controllers.SetupDatabase(gormDB, true)
	controllers.Mount(service)

	// Start service
	if err := service.ListenAndServe(":8000"); err != nil {
		service.LogError("startup", "err", err)
	}
}
