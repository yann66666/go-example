// Author: yann
// Date: 2022/5/20
// Desc: ginwrap

package wrap

import (
	"github.com/Hidata-xyz/go-example/pkg/ginwrap"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap/middleware"
)

type EngineWrapOption func(wrap *EngineWrap)

func EngineWrapLoggerOption() EngineWrapOption {
	return func(wrap *EngineWrap) {
		wrap.Use(middleware.Print)
	}
}

func PrintReqParamsOption() EngineWrapOption {
	return func(wrap *EngineWrap) {
		printReq = true
	}
}

func PrintRespParamsOption() EngineWrapOption {
	return func(wrap *EngineWrap) {
		ginwrap.PrintResp = true
	}
}
