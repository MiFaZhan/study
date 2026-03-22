package words

import (
	"context"
	"star/api/words/v1"
	"star/internal/dao"
	"star/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

type UpdateInput struct {
	Uid                uint
	Word               string
	Definition         string
	ExampleSentence    string
	ChineseTranslation string
	Pronunciation      string
	ProficiencyLevel   v1.ProficiencyLevel
}

func (w *Words) Update(ctx context.Context, id uint, in UpdateInput) error {
	var cls = dao.Words.Columns()

	count, err := dao.Words.Ctx(ctx).
		Where(cls.Uid, in.Uid).
		Where(cls.Word, in.Word).
		WhereNot(cls.Id, id).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("单词已存在")
	}

	_, err = dao.Words.Ctx(ctx).Data(do.Words{
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Where(cls.Id, id).Where(cls.Uid, in.Uid).Update()
	if err != nil {
		return err
	}
	return nil
}
