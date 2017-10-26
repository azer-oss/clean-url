# clean-url

A tiny library for cleaning up the pollution around URLs, so we can figure out if two different URLs might actually point to same page.

Examples:

```go
Clean("https://foobar.com/?")
// => foobar.com

Clean("foobar.com/?utm_source=span&utm_content=eggs")
// => foobar.com

Clean("http://foobar.com/?yo=lo#heythere")
// => foobar.com/?yo=lo
```

# Usage

```go
import (
  "github.com/kozmos/clean-url"
)

func main () {
  cleanurl.Clean("foobar.com/?utm_source=span&utm_content=eggs")
  // => foobar.com
}
```

# Kozmos' Usage

This library is vital for [Kozmos](https://getkozmos.com), as it has to avoid polluted URLs to provide more
reliable and accurate information.
