package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/other"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/user"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	User    user.RouterGroup
	Other   other.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
