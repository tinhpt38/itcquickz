package main

import (
	"changeme/app/model"
	"changeme/app/module/auth"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Login(idnum string) (res model.LoginResponse) {
	ls := auth.LoginService{}

	user, err := ls.Login(idnum)
	if err != nil {
		res = model.LoginResponse{
			Message: err.Error(),
			Status:  false,
		}
		return 
	}
	res = model.LoginResponse{
		Message: "Đăng nhập thành công",
		Status:  true,
		User: &user,
	}
	return 
}
