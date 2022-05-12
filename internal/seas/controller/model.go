// Author: yann
// Date: 2022/5/12
// Desc: controller

package controller

import (
	"github.com/HiData-xyz/HiData.More/HiExtern/api/smodel"
	"github.com/HiData-xyz/HiData.More/HiExtern/constants"
	"github.com/Hidata-xyz/go-example/internal/pkg/auth"
	"github.com/Hidata-xyz/go-example/internal/pkg/code"
	"github.com/Hidata-xyz/go-example/internal/seas/errors"
	"github.com/Hidata-xyz/go-example/internal/seas/service"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap/wrap"
	"github.com/gin-gonic/gin"
)

type seasModel struct {
	service                          *service.SeasModel
	internationalWatersSeasGatewayID string
}

func NewSeasModel(internationalWatersSeasGatewayID string) *seasModel {
	return &seasModel{
		service:                          service.NewSeasModel(""),
		internationalWatersSeasGatewayID: internationalWatersSeasGatewayID,
	}
}

type SaveModelReq struct {
	smodel.SeasModel
}

func (s *SaveModelReq) DecodeType() wrap.DecodeType {
	return wrap.DecodeTypeJson
}

func (s *seasModel) SaveModel(ctx *gin.Context, params wrap.IBase) *ginwrap.Response {
	req := *(params.(*SaveModelReq))
	base, err := auth.MustGetBase(ctx)
	if err != nil {
		return auth.BaseErrResponse(err)
	}

	//检查和初始化请求参数
	switch req.SeasType {
	case constants.SeasTypeInlandSea:
		//内海不能创建功能模型
		if req.Type == constants.ModelTypeFunc && req.ID == "" {
			return code.NewSeasErrResponse(code.CannotCreateFuncModel, nil)
		}
		//如果是创建聚合模型, 赋值组织信息
		if req.ID == "" {
			req.SeasGatewayID = base.GID
			req.GroupID = base.GID
			req.UID = base.UID
		}
	case constants.SeasTypeInternationalWaters:
		if req.ID == "" {
			req.SeasGatewayID = s.internationalWatersSeasGatewayID
			req.GroupID = base.GID
			req.UID = base.UID
		}
	default:
		return code.NewSeasErrResponse(code.SeasBadRequest, errors.ErrSeasTypeCannotNull)
	}
	id, err := s.service.SaveModel(ctx, &req.SeasModel)
	if err != nil {
		return code.NewSeasErrResponse(code.SaveModelFail, err)
	}

	return ginwrap.Response200(id)
}
