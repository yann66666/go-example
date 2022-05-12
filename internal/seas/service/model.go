// Author: yann
// Date: 2022/5/12
// Desc: service

package service

import (
	"github.com/HiData-xyz/HiData.More/HiExtern/api"
	"github.com/HiData-xyz/HiData.More/HiExtern/api/smodel"
	"github.com/HiData-xyz/HiData.More/HiExtern/lib/logging"
	"github.com/Hidata-xyz/go-example/internal/pkg/insideapi"
	"github.com/gin-gonic/gin"
)

type SeasModel struct {
	hostGatewayAddr string
	*logging.ZapEventLogger
}

func NewSeasModel(hostGatewayAddr string) *SeasModel {
	return &SeasModel{hostGatewayAddr: hostGatewayAddr, ZapEventLogger: logging.Logger("Seas-Model")}
}

func (s *SeasModel) SaveModel(ctx *gin.Context, model *smodel.SeasModel) (string, error) {
	//保存模型
	resp, err := insideapi.GetSeasApi(model.SeasGatewayID, s.hostGatewayAddr).SaveModel(model)
	if err = api.CheckApiError(resp, err); err != nil {
		s.Errorf("更新或者创建模型失败: %s", err.Error())
		return "", err
	}

	return resp.Data.ID, nil
}
