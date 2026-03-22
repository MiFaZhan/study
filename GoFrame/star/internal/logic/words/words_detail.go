package words

import (
	"context"
	"star/internal/dao"
	"star/internal/model/entity"
)

func (w *Words) Detail(ctx context.Context, uid, id uint) (word *entity.Words, err error) {
	var (
		cls = dao.Words.Columns()
		orm = dao.Words.Ctx(ctx)
	)
	orm = orm.Where(cls.Id, id)
	if uid > 0 {
		orm = orm.Where(cls.Uid, uid)
	}
	err = orm.Scan(&word)
	return
}
