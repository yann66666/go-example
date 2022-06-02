// Author: yann
// Date: 2022/5/11
// Desc: smodel

package seas

import (
	"github.com/Hidata-xyz/go-example/internal/seas/controller"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap/wrap"
)

func InitRouter(group wrap.IGroupWrap) {
	api := controller.NewSeasModel("ZZZ")
	group.POST("model", &controller.SaveModelReq{}, api.SaveModel)
}
