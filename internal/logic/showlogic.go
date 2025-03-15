package logic

import (
	"context"

	"go-shorturl/internal/svc"
	"go-shorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// todo: add your logic here and delete this line
	row, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, req.ShortUrl)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, err
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, ErrorInternal
	}

	return &types.ShowResponse{
		LongUrl: row.Lurl,
	}, nil
}
