// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1

package types

type ConvertRequest struct {
	LongUrl string `json:"longUrl" validate:"required"`
}

type ConvertResponse struct {
	ShortUrl string `json:"shortUrl"`
}

type ShowRequest struct {
	ShortUrl string `path:"shortUrl" validate:"required"`
}

type ShowResponse struct {
	LongUrl string `json:"longUrl"`
}
