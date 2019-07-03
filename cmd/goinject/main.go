package main

import (
	"fmt"
	"github.com/facebookgo/inject"
)

type DBEngine struct {
	Name string
}

type UserDB struct {
	Db *DBEngine `inject:""`
}

type UserService struct {
	Db *UserDB `inject:""`
}

type App struct {
	Name string
	User *UserService `inject:""`
}

func (a *App) Create() string {
	return "create app, in db name:" + a.User.Db.Db.Name + " app name :" + a.Name
}

type Object struct {
	App *App
}

func Init() *Object {
	db := DBEngine{Name: "db1"}
	var g inject.Graph
	app := App{Name: "go-app"}

	_ = g.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: &db},
	)
	_ = g.Populate()
	return &Object{
		App: &app,
	}

}

func main() {
	obj := Init()
	fmt.Println(obj.App.Create())
}
