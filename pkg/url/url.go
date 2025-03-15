package url

import (
	"errors"
	"net/url"
	"path"
)

func GetBasePath(inputUrl string) (string, error) {
	u, err := url.Parse(inputUrl)
	if err != nil {
		return "", err
	}
	if len(u.Host) == 0 {
		return "", errors.New("host is required")
	}
	return path.Base(u.Path), nil
}
