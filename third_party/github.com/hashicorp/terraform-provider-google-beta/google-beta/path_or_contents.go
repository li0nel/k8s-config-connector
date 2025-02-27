package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

// If the argument is a path, pathOrContents loads it and returns the contents,
// otherwise the argument is assumed to be the desired contents and is simply
// returned.
//
// The boolean second return value can be called `wasPath` - it indicates if a
// path was detected and a file loaded.
func pathOrContents(poc string) (string, bool, error) {
	return verify.PathOrContents(poc)
}
