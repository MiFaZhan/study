package words

import (
	"context"
	"star/internal/dao"
)

func (w *Words) Delete(ctx context.Context, uid, id uint) (err error) {
	var (
		cls = dao.Words.Columns()
		orm = dao.Words.Ctx(ctx)
	)
	orm = orm.Where(cls.Id, id)
	if uid > 0 {
		orm = orm.Where(cls.Uid, uid)
	}
	_, err = orm.Delete()
	return
}
