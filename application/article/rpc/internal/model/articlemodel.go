package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
		ArticlesByAuthorId(ctx context.Context, authorId uint64, status int, likeNum int64, pubTime, sortType string, limit int) ([]*Article, error)
	}

	customArticleModel struct {
		*defaultArticleModel
	}
)

// NewArticleModel returns a model for the database table.
func NewArticleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ArticleModel {
	return &customArticleModel{
		defaultArticleModel: newArticleModel(conn, c, opts...),
	}
}

func (m *customArticleModel) ArticlesByAuthorId(ctx context.Context, authorId uint64, status int, likeNum int64, pubTime, sortType string, limit int) ([]*Article, error) {
	var (
		err      error
		sql      string
		anyType  any
		articles []*Article
	)

	if sortType == "like_num" {
		anyType = likeNum
		sql = fmt.Sprintf("select "+articleRows+" from "+m.table+" where author_id = ? and status = ? and like_num < ? order by %s desc limit ?", sortType)
	} else {
		anyType = pubTime
		sql = fmt.Sprintf("select "+articleRows+" from "+m.table+" where author_id = ? and status = ? and publish_time < ? order by %s desc limit ?", sortType)
	}
	err = m.QueryRowsNoCacheCtx(ctx, &articles, sql, authorId, status, anyType, limit)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
