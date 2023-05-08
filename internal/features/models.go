package features

import "errors"

var ErrMediaNotFound = errors.New("media not found")
var ErrInvalidMediaType = errors.New("invalid media type")

type Rating struct {
	Rating float32 `json:"rating"`
	Count  int     `json:"count"`
}
