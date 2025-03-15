package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	DomainName string

	ShortUrlDB struct {
		DSN string
	}

	Sequence struct {
		DSN string
		Table string
		Field string
		Value string
	}

	Encode struct {
		Table string
		BlackList []string
	}
}
