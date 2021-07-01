package url

import (
	"net/url"
	"path"
)

// JoinURLAndPaths joins the given URL with an array of paths. url must not be nil.
func JoinURLAndPaths(url *url.URL, paths ...string) *url.URL {
	urlCopy := *url
	for _, p := range paths {
		urlCopy.Path = path.Join(urlCopy.Path, p)
	}
	return &urlCopy
}
