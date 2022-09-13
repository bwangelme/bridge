package appmodel

import (
	"bridge/bridge/libs/store"
	"bridge/bridge/pbs/apppb"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type App struct {
	Id      int64
	Name    string
	Runtime RuntimeType
	Image   string
}

func (a *App) String() string {
	return fmt.Sprintf("%v %v %v %v", a.Id, a.Name, a.Runtime, a.Image)
}

func (a *App) ToPB() *apppb.App {
	return &apppb.App{
		Name:    a.Name,
		Image:   a.Image,
		Runtime: a.Runtime.ToPBEnum(),
	}
}

func AppsToPB(apps []*App) []*apppb.App {
	var res []*apppb.App
	for _, app := range apps {
		res = append(res, app.ToPB())
	}
	return res
}

func Add(app *apppb.App) (*App, error) {
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

func List(start, limit int64) ([]int64, error) {
	var ids []int64

	sqlStr := "select id from apps limit ?, ?"
	err := store.MySQL.Select(&ids, sqlStr, start, limit)

	if err != nil {
		return nil, errors.Wrapf(err, "exec sql %v failed", sqlStr)
	}
	return ids, nil
}

func GetMulti(ids []int64) ([]*App, error) {
	if len(ids) > 500 {
		return nil, errors.New("get too much once")
	}

	var apps []App
	// sqlx.In 将一个 ? 转换成 len(ids) 个 ?
	// Select 中再将参数挨个传递进去
	// 防止 sql 注入的工作还是 mysql 做的
	query, args, err := sqlx.In("select * from apps where id in (?)", ids)
	query = store.MySQL.Rebind(query)
	err = store.MySQL.Select(&apps, query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "select failed sql %v", query)
	}

	var res []*App
	for _, a := range apps {
		res = append(res, &a)
	}
	return res, nil
}
