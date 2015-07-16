package environment

import (
	"gopkg.in/redis.v3"
)

type Env struct {
	CL *redis.Client
}