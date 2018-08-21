package example

import "example.io/named-src-pkg/example/subpkg"

func Greeting(name string) string  {
	return "Hello " + name + " from" + subpkg.From()
}
