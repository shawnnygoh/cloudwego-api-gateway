// Code generated by hertz generator.

package echo

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	echo "github.com/shawnnygoh/cloudwego-api-gateway/biz/model/echo"
)

// EchoMethod .
// @router /echo/echomethod [POST]
func EchoMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req echo.EchoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(echo.EchoResp)

	c.JSON(consts.StatusOK, resp)
}
