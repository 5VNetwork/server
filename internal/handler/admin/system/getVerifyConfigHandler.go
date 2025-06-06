package system

import (
	"github.com/gin-gonic/gin"
	"github.com/perfect-panel/server/internal/logic/admin/system"
	"github.com/perfect-panel/server/internal/svc"
	"github.com/perfect-panel/server/pkg/result"
)

// Get verify config
func GetVerifyConfigHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := system.NewGetVerifyConfigLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetVerifyConfig()
		result.HttpResult(c, resp, err)
	}
}
