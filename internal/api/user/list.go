package user

import (
	"context"
	"log"

	"github.com/xm-chentl/eddydoor/internal/contract"
	"github.com/xm-chentl/eddydoor/internal/model/global"
	"github.com/xm-chentl/eddydoor/internal/model/views"
	"github.com/xm-chentl/eddydoor/internal/response"
	"github.com/xm-chentl/goresource"
)

type ListAPI struct {
	contract.IAdminSession
	MySqlDB goresource.IResource `inject:"mysql"`

	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (s ListAPI) Call(ctx context.Context) (res interface{}, err error) {
	db := s.MySqlDB.Db(ctx)
	query := db.Query()
	total, err := query.Count(&global.User{})
	if err != nil {
		log.Println(err)
		err = response.ErrDataException
		return
	}

	entries := make([]global.User, 0)
	if err = query.Page(s.Page).PageSize(s.PageSize).ToArray(&entries); err != nil {
		// log.Println(string(debug.Stack()))
		log.Println(err)
		err = response.ErrDataException
		return
	}

	respEntries := make([]views.ResponseUser, len(entries))
	for i := 0; i < len(entries); i++ {
		respEntries[i] = entries[i].ToResponse()
	}

	res = views.DataGrid{
		Total: total,
		Rows:  respEntries, // todo: 可根据前端需要过滤下载
	}

	return
}
