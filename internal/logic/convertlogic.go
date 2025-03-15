package logic

import (
	"context"
	"errors"
	"net/url"

	"go-shorturl/internal/svc"
	"go-shorturl/internal/types"
	"go-shorturl/model"
	"go-shorturl/pkg/connect"
	"go-shorturl/pkg/encode"
	"go-shorturl/pkg/sequence"
	urlTool "go-shorturl/pkg/url"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	ErrorInternal = errors.New("internal error")
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	seqence sequence.Sequence
	encoder *encode.Base62
	blackset map[string]struct{}
	domainName string
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	seqConfig := svcCtx.Config.Sequence
	encConfig := svcCtx.Config.Encode
	
	blackset := make(map[string]struct{}, len(encConfig.BlackList))
	for _, item := range encConfig.BlackList {
		blackset[item] = struct{}{}
	}

	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		seqence: sequence.NewMySQLSequence(seqConfig.DSN, seqConfig.Table, seqConfig.Field, seqConfig.Value),
		encoder: encode.NewBase62(encConfig.Table, encConfig.BlackList),
		blackset: blackset,
		domainName: svcCtx.Config.DomainName,
	}
}

// Convert convert a long url into a short one
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// todo: add your logic here and delete this line
	// 1. verify input long url (if exist a corresponded short url)
	// 1.1 check if a null url
	// 1.2 check if the url is reachable
	if !connect.Get(req.LongUrl) {
		return nil, errors.New("invalid url")
	}
	// 1.3 check if the short one exists
	hash := encode.Sum([]byte(req.LongUrl))
	_, err = l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, hash)
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, errors.New("input url has already converted")
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, ErrorInternal
	}

	// 1.4 check if input is a short url (avoid loop convert)
	base, err := urlTool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("url.GetBasePath failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, ErrorInternal
	}
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, base)

	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, errors.New("input url has already converted")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, ErrorInternal
	}

	var shortUrl string
	for {
		// 2. get number
		seq, err := l.seqence.Next()
		if err != nil { 
			logx.Errorw("Sequence.Next failed", logx.LogField{Key: "err", Value: err.Error()})
			return nil, ErrorInternal
		}
		
		// 3. convert number into shortUrl url
		shortUrl = l.encoder.Encode(seq)
		if _, ok := l.blackset[shortUrl]; !ok {
			break
		}
	}
	// 4. store the map between the long and the short url
	shortUrl, err = url.JoinPath(l.domainName, shortUrl)
	if err != nil {
		logx.Errorw("url.JoinPath failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, ErrorInternal
	}
	_, err = l.svcCtx.ShortUrlModel.Insert(
		l.ctx,
		&model.ShortUrlMap{
			Lurl: req.LongUrl,
			Md5: hash,
			Surl: shortUrl,
		},
	)
	if err != nil {
		logx.Errorw("ShortUrlModel.Insert failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, ErrorInternal
	}
	// 5. response
	return &types.ConvertResponse{
		ShortUrl: shortUrl,
	}, nil
}
