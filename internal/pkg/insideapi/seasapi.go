// Author: yann
// Date: 2022/5/12
// Desc: api

package insideapi

import (
	"fmt"
	"github.com/HiData-xyz/HiData.More/HiExtern/api/seasapi"
	"github.com/HiData-xyz/HiData.More/HiExtern/constants"
)

func GetSeasApi(seasGatewayID, hostGatewayAddr string) *seasapi.SeasApi {
	//return seasapi.GetSeasApi("http://127.0.0.1:8001")
	return seasapi.GetSeasApiByProxy(fmt.Sprintf("http://%s:%d", constants.SeasContainerName, constants.SeasContainerPort), seasGatewayID, hostGatewayAddr)
}
