package formula

import (
	"github.com/titus12/formula/internal/cache"
	"github.com/titus12/formula/opt"
)

//Register custom function which implement opt.Function
func Register(f *opt.Function) error {
	return cache.Register(f)
}

//RegisterGlobalParameter register global parameter which will be used in all the runtime
//func RegisterGlobalParameter(name string, value interface{}) {
//
//}
