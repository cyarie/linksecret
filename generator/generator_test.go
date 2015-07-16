package generator_test

import (
	"testing"
	"time"

	"github.com/cyarie/linksecret/application/generator"
)

func TestHashGenerator(t *testing.T) {
	const time_form = "2015-07-13 17:28:27.043022906"
	test_time, _ := time.Parse(time_form, "2015-07-13 17:28:27.043022906")
	test_struct := generator.LinkGen{
		"http://google.com",
		"cyarie@gmail.com",
		test_time,
	}

	test_hash := "cf85562809680b42959c6421872119e8"
	gen_hash := generator.GenerateLink(test_struct)

	if test_hash != gen_hash {
		t.Errorf("Expected a hash of %s, got %s", test_hash, gen_hash)
	}
}
