package helper

import (
	"errors"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/app/dto"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorWithMsg(ctx *gin.Context, msg string, err error) {
	res := dto.Response{
		Message: msg,
	}
	switch {
	case errors.Is(err, constant.ErrTypeInvalidParams):
		res.Code = constant.CodeErrBadRequest
	default:
		res.Code = constant.CodeErrInternalServer
	}
	ctx.JSON(http.StatusOK, res)
	ctx.Abort()
}

func SuccessWithData(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := dto.Response{
		Code: constant.CodeSuccess,
		Data: data,
	}
	ctx.JSON(http.StatusOK, res)
	ctx.Abort()
}

func SuccessWithOutData(ctx *gin.Context) {
	res := dto.Response{
		Code:    constant.CodeSuccess,
		Message: "success",
	}
	ctx.JSON(http.StatusOK, res)
	ctx.Abort()
}

func SuccessWithMsg(ctx *gin.Context, msg string) {
	res := dto.Response{
		Code:    constant.CodeSuccess,
		Message: msg,
	}
	ctx.JSON(http.StatusOK, res)
	ctx.Abort()
}
