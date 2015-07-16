/**
CONTROLLERS TEST FILE
**/

package controllers

import (
	"testing"
	"time"

	"github.com/cyarie/linksecret-generator/linkgen/controllers"
	"github.com/cyarie/linksecret-generator/linkgen/environment"
	"github.com/cyarie/linksecret-generator/linkgen/generator"
	"gopkg.in/redis.v3"
)

func Setup() *environment.Env {
	// Setup our redis connection
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Don't need to actually assign this here, just set CL to the Redis client we made up above
	env := &environment.Env{
		CL: client,
	}

	return env
}

func genHash() string {
	const time_form = "2015-07-13 17:28:27.043022906"
	test_time, _ := time.Parse(time_form, "2015-07-13 17:28:27.043022906")
	test_struct := generator.LinkGen{
		"http://google.com",
		"cyarie@gmail.com",
		test_time,
	}

	gen_hash := generator.GenerateLink(test_struct)

	return gen_hash
}

func TestRedisInsert(t *testing.T) {
	env := Setup()
	rc := env.CL
	// Let's start by generating a hash
	const time_form = "2015-07-13 17:28:27.043022906"
	test_time, _ := time.Parse(time_form, "2015-07-13 17:28:27.043022906")
	test_struct := generator.LinkGen{
		"http://google.com",
		"cyarie@gmail.com",
		test_time,
	}

	gen_hash := generator.GenerateLink(test_struct)

	err := controllers.RedisInsertLink(test_struct, gen_hash, env)
	if err != nil {
		t.Error(err)
	}

	// Test out both the link we expect the hash to redirect to, and the username it is attached to
	rl, _ := rc.HGet(gen_hash, "link").Result()

	if rl != "http://google.com" {
		t.Errorf("Expected http://google.com, got %s", rl)
	}

	ru, _ := rc.HGet(gen_hash, "username").Result()

	if ru != "cyarie@gmail.com" {
		t.Errorf("Expected cyarie@gmail.com, got %s", ru)
	}

	// Let's clean up Redis instance
	rc.Del(gen_hash)
}
