package cleanurl_test

import (
	"github.com/kozmos/clean-url"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanURL(t *testing.T) {
	assert.Equal(t, "foobar.com", cleanurl.Clean("https://foobar.com"))
	assert.Equal(t, "foobar.com", cleanurl.Clean("http://foobar.com/"))
	assert.Equal(t, "foobar.com", cleanurl.Clean("http://foobar.com#"))
	assert.Equal(t, "foobar.com", cleanurl.Clean("http://www.foobar.com?"))
	assert.Equal(t, "foobar.com", cleanurl.Clean("http://www.foobar.com/?"))
	assert.Equal(t, "foobar.com", cleanurl.Clean("http://www.foobar.com/?#"))
	assert.Equal(t, "foobar.com/yo/lo?span=eggs&hey=there", cleanurl.Clean("http://foobar.com/yo/lo?span=eggs&hey=there#top"))
	assert.Equal(t, "foobar.com/yo/lo", cleanurl.Clean("http://foobar.com/yo/lo?utm_source=hi"))
	assert.Equal(t, "yolo.com/blog?reffoo=bar&span=eggs", cleanurl.Clean("http://www.yolo.com/blog/?span=eggs&utm_content=bufferc60dd&utm_medium=social&utm_source=twitter.com&utm_campaign=buffer&ref=asd&reffoo=bar&empty"))
}
