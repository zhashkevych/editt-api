package publication

import "errors"

var (
	ErrNoPublication      = errors.New("no publication found")
	ErrWordsLimitExceeded = errors.New("words count exceeds limit")
)

