package controllers

import (
	"log"

	"github.com/cyarie/linksecret/application/environment"
	"github.com/cyarie/linksecret/application/generator"
)

func RedisInsertLink(lg generator.LinkGen, hash string, env *environment.Env) error {
	var err error

	rc := env.CL

	res := rc.HMSet(hash, "link", lg.Url, "username", lg.User)
	if res.Err() != nil {
		log.Println(res.Err())
		return res.Err()
	}

	log.Println(res.Result())

	return err
}
