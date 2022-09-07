package appmodel

import (
	"bridge/bridge/libs/store"
	"bridge/bridge/pbs/apppb"

	"github.com/pkg/errors"
)

type App struct {
	Id      int64
	Name    string
	Runtime RuntimeType
	Image   string
}

func (a *App) ToPB() *apppb.App {
	return &apppb.App{
		Name:    a.Name,
		Image:   a.Image,
		Runtime: a.Runtime.ToPBEnum(),
	}
}

func AddApp(app *apppb.App) (*App, error) {
	sqlStr := "insert into apps(name, runtime, image) values (?,?,?)"
	ret, err := store.MySQL.Exec(sqlStr, app.Name, app.Runtime.String(), app.Image)
	if err != nil {
		return nil, errors.WithMessagef(err, "insert failed")
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		return nil, errors.WithMessagef(err, "get LastInsertId failed")
	}

	return &App{
		Id:      theID,
		Name:    app.Name,
		Runtime: RuntimeType(app.Runtime.String()),
		Image:   app.Image,
	}, nil
}
