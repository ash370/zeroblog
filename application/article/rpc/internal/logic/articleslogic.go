package logic

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"strconv"
	"time"

	"zeroblog/application/article/rpc/internal/code"
	"zeroblog/application/article/rpc/internal/model"
	"zeroblog/application/article/rpc/internal/svc"
	"zeroblog/application/article/rpc/internal/types"
	"zeroblog/application/article/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
)

const (
	prefixArticles = "biz#articles#%d#%d"
	articlesExpire = 3600 * 24 * 2
)

type ArticlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticlesLogic {
	return &ArticlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticlesLogic) Articles(in *pb.ArticlesRequest) (*pb.ArticlesResponse, error) {
	if in.SortType != types.SortLikeNum && in.SortType != types.SortPublishTime {
		return nil, code.SortTypeInvalid
	}
	if in.AuthorId <= 0 {
		return nil, code.UserIdInvalid
	}
	if in.PageSize == 0 {
		in.PageSize = types.DefaultPageSize
	}
	if in.Cursor == 0 {
		//说明是第一次请求
		if in.SortType == types.SortLikeNum {
			in.Cursor = types.DefaultSortLikeCursor
		} else {
			in.Cursor = time.Now().Unix()
		}
	}

	var (
		sortType        string
		sortLikeNum     int64
		sortPublishTime string
	)
	if in.SortType == types.SortLikeNum {
		sortType = "like_num"
		sortLikeNum = in.Cursor //请求的点赞数,查询DB时条件是小于当前
	} else {
		sortType = "publish_time"
		sortPublishTime = time.Unix(in.Cursor, 0).Format("2006-01-02 15:04:05") //请求的发布时间
	}

	var (
		err            error
		isCache, isEnd bool
		lastId, cursor int64
		curPage        []*pb.ArticleInfo
		articles       []*model.Article //从数据库返回的
	)

	//先从缓存获取文章id
	articleIds, _ := l.cacheArticles(l.ctx, in.AuthorId, in.Cursor, in.PageSize, in.SortType) //忽略错误，尽量获取数据

	if len(articleIds) > 0 {
		isCache = true //缓存命中

		if articleIds[len(articleIds)-1] == -1 { //已经加载到最后一页了
			isEnd = true
		}

		//根据业务缓存返回的文章id,查询自动缓存的行记录
		articles, err := l.articleByIds(l.ctx, articleIds)
		if err != nil {
			return nil, err
		}

		//根据指定排序类型进行排序
		var cmpFunc func(a, b *model.Article) int
		switch sortType {
		case "like_num":
			cmpFunc = func(a, b *model.Article) int {
				return cmp.Compare(b.LikeNum, a.LikeNum)
			}
		case "publish_time":
			cmpFunc = func(a, b *model.Article) int {
				return cmp.Compare(b.PublishTime.Unix(), a.PublishTime.Unix())
			}
		}
		slices.SortFunc(articles, cmpFunc)

		//构建返回页面
		for _, article := range articles {
			curPage = append(curPage, &pb.ArticleInfo{
				Id:          int64(article.Id),
				Title:       article.Title,
				Content:     article.Content,
				LikeNum:     article.LikeNum,
				AuthorId:    int64(article.AuthorId),
				CommentNum:  article.CommentNum,
				PublishTime: article.PublishTime.Unix(),
			})
		}
	} else {
		//直接回表
		//TODO：防止缓存击穿
		v, err, _ := l.svcCtx.SingleFlightGroup.Do(fmt.Sprintf("ArticlesByAuthorId:%d:%d", in.AuthorId, in.SortType), func() (interface{}, error) {
			return l.svcCtx.ArticleModel.ArticlesByAuthorId(l.ctx, uint64(in.AuthorId), 0, sortLikeNum, sortPublishTime, sortType, types.DefaultLimit)
		})
		if err != nil {
			logx.Errorf("ArticlesByAuthorId authorId: %d sortType: %s error: %v", in.AuthorId, sortType, err)
			return nil, err
		}
		if v == nil {
			return &pb.ArticlesResponse{}, nil
		}
		articles = v.([]*model.Article)
		if len(articles) == 0 {
			return &pb.ArticlesResponse{}, nil
		}

		var firstPageArticles []*model.Article
		if len(articles) > int(in.PageSize) { //数据库返回了大于请求的页面大小
			firstPageArticles = articles[:int(in.PageSize)]
		} else {
			firstPageArticles = articles
			isEnd = true
		}

		for _, article := range firstPageArticles {
			curPage = append(curPage, &pb.ArticleInfo{
				Id:          int64(article.Id),
				Title:       article.Title,
				Content:     article.Content,
				AuthorId:    int64(article.AuthorId),
				LikeNum:     article.LikeNum,
				CommentNum:  article.CommentNum,
				PublishTime: article.PublishTime.Unix(),
			})
		}
	}

	//去重,如果数据变动,可能会重复查出
	if len(curPage) > 0 {
		pageLast := curPage[len(curPage)-1]
		lastId = pageLast.Id
		if in.SortType == types.SortPublishTime {
			cursor = pageLast.PublishTime
		} else {
			cursor = pageLast.LikeNum
		}
		if cursor < 0 {
			cursor = 0
		}
		for k, article := range curPage {
			if in.SortType == types.SortPublishTime {
				if article.PublishTime == in.Cursor && article.Id == in.ArticleId {
					//包含了上一次的最后一条记录,去掉
					curPage = curPage[k:]
				}
			} else {
				if article.LikeNum == in.Cursor && article.Id == in.ArticleId {
					curPage = curPage[k:]
				}
			}
		}
	}

	res := &pb.ArticlesResponse{
		IsEnd:     isEnd,
		Cursor:    cursor,
		ArticleId: lastId,
		Articles:  curPage,
	}

	//缓存未命中时异步写缓存
	if !isCache {
		threading.GoSafe(func() {
			if len(articles) < types.DefaultLimit && len(articles) > 0 {
				//加结束标识符
				articles = append(articles, &model.Article{Id: -1})
			}
			err = l.addCacheArticles(context.Background(), articles, in.AuthorId, in.SortType)
			if err != nil {
				logx.Errorf("addCacheArticles error: %v", err)
			}
		})
	}

	return res, nil
}

// 根据id查文章详情
func (l *ArticlesLogic) articleByIds(ctx context.Context, ids []int64) ([]*model.Article, error) {
	articles, err := mr.MapReduce[int64, *model.Article, []*model.Article](func(source chan<- int64) {
		for _, id := range ids {
			if id == -1 {
				continue
			}
			source <- id
		}
	}, func(id int64, writer mr.Writer[*model.Article], cancel func(error)) {
		article, err := l.svcCtx.ArticleModel.FindOne(ctx, id)
		if err != nil {
			cancel(err)
		}
		writer.Write(article)
	}, func(pipe <-chan *model.Article, writer mr.Writer[[]*model.Article], cancel func(error)) {
		var articles []*model.Article
		for article := range pipe {
			articles = append(articles, article)
		}
		writer.Write(articles)
	})
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func articlesKey(uid int64, sortType int32) string {
	return fmt.Sprintf(prefixArticles, uid, sortType)
}
func (l *ArticlesLogic) cacheArticles(ctx context.Context, uid, cursor, ps int64, sortType int32) ([]int64, error) {
	key := articlesKey(uid, sortType)
	b, err := l.svcCtx.BizRedis.ExistsCtx(ctx, key)

	if err != nil {
		logx.Errorf("ExistsCtx key: %s err: %v", key, err)
	}
	if b {
		err = l.svcCtx.BizRedis.ExpireCtx(ctx, key, articlesExpire) //对于hotcache增加过期时间
		if err != nil {
			logx.Errorf("ExpireCtx key: %s err: %v", key, err)
		}
	}

	pairs, err := l.svcCtx.BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, key, 0, cursor, 0, int(ps))
	//cursor是最后,如果按点赞数排序,点赞多的在后面,从cursor往下查点赞是递减
	if err != nil {
		logx.Errorf("ZrevrangebyscoreWithScoresAndLimitCtx key: %s err: %v", key, err)
		return nil, err
	}

	var ids []int64
	for _, pair := range pairs {
		id, err := strconv.ParseInt(pair.Key, 10, 64)
		if err != nil {
			logx.Errorf("strconv.ParseInt key:%s err: %v", pair.Key, err)
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (l *ArticlesLogic) addCacheArticles(ctx context.Context, articles []*model.Article, authorId int64, sortType int32) error {
	if len(articles) == 0 {
		return nil
	}

	key := articlesKey(authorId, sortType)

	for _, article := range articles {
		var score int64
		if sortType == types.SortLikeNum {
			score = article.LikeNum
		} else if sortType == types.SortPublishTime && article.Id != -1 {
			score = article.PublishTime.Local().Unix()
		}
		if score < 0 {
			score = 0
		}

		_, err := l.svcCtx.BizRedis.ZaddCtx(ctx, key, score, strconv.Itoa(int(article.Id)))
		if err != nil {
			return err
		}
	}

	return l.svcCtx.BizRedis.ExpireCtx(ctx, key, articlesExpire)
}
