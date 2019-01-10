package rest

import (
	"net/http"
	"tantan-simplify/errors"

	"github.com/gin-gonic/gin"
)

// basic response structure
type BaseResponse struct {
	StatusCode  int32  `json:"status"`
	EnglishText string `json:"desc"`
}

// error code
const (
	// operation succeed
	Success int32 = 0

	// ----------------------------- //
	// params error （2560 = 0x0A00）//
	// ---------------------------- //

	// -------------------------------- //
	// illegal request （6656 = 0x1A00）//
	// ------------------------------- //
	Operated     = 6656
	UnknownState = 6657
	UserNotFound = 6658

	// --------------------------------------- //
	// internal service error（10752 = 0x2A00）//
	// -------------------------------------- //

	// unknown internal error
	UnknownInternalError = 10752
)

// error code -> HTTP status code
var restStatusCodeToHTTPStatusCode = map[int32]int{
	Success:              http.StatusOK,
	Operated:             http.StatusBadRequest,
	UnknownState:         http.StatusBadRequest,
	UserNotFound:         http.StatusBadRequest,
	UnknownInternalError: http.StatusInternalServerError,
}

var errorToRestStatusCode = map[error]int32{
	errors.ErrOperated:     Operated,
	errors.ErrUnknownState: UnknownState,
	errors.ErrUserNotFound: UserNotFound,
	errors.ErrUnknown:      UnknownInternalError,
}

// REST 错误返回
// 只有 err != nil 的时候才返回错误的描述性文本信息（返回成功的文本描述没有太大价值）
// 如果 err 不在 errorToRestStatusCode 表中，默认为 errors.ErrUnknown 错误
func responseWithError(c *gin.Context, err error) {
	if err != nil {
		status, ok := errorToRestStatusCode[err]
		if !ok {
			err = errors.ErrUnknown
		}
		c.JSON(restStatusCodeToHTTPStatusCode[status], gin.H{
			"status": errorToRestStatusCode[err],
			"desc":   err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(restStatusCodeToHTTPStatusCode[Success], gin.H{
		"status": Success,
	})
	c.Abort()
}

// REST 返回 items 数据，items 是一个列表数据
// 调用这个接口时，默认 status = Success。只有这样，调用这个接口才有价值
func responseWithItems(c *gin.Context, items interface{}) {
	if items != nil {
		c.JSON(restStatusCodeToHTTPStatusCode[Success], items)
		c.Abort()
		return
	}
}
