package generator

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

type LinkGen struct {
	Url  string
	User string
	Ts   time.Time
}

func GenerateLink(lg LinkGen) string {
	// Concatenate the LinkGen fields together to provide a string to hash
	catStr := fmt.Sprintf("%s|%s|%t", lg.Url, lg.User, lg.Ts)

	data := []byte(catStr)

	hash := md5.New()
	hash.Write(data)

	return hex.EncodeToString(hash.Sum(nil))
}
