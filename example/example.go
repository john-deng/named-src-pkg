package example

import "gitlab.vpclub.cn/demo/named-src-pkg/example/subpkg"

func Greeting(name string) string  {
	return "Hello " + name + " from " + subpkg.From()
}
