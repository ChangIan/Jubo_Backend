package controllerV1

import (
	"log"

	"github.com/gin-gonic/gin"

	modelBase "changian.com/jubo/model/base"
	modelCommon "changian.com/jubo/model/common"

	serviceV1 "changian.com/jubo/service/v1"
)

// Controller
// Method - Get
// 查詢患者
func GetPatient(context *gin.Context) {

	result := new(modelBase.Response)

	// 沒有 JWT 先給 OperatorID = 1
	operatorClaims := modelCommon.OperatorClaims{
		OperatorID: 1,
	}

	patients, err := serviceV1.QueryPatient(&operatorClaims)

	if err != nil {
		log.Fatal(err)

		result.ErrorCode = 1
		result.ErrorMessage = "查詢失敗，請稍候執行"
	} else {
		result.ErrorCode = 0
		result.ErrorMessage = "查詢成功"
		result.Data = patients
	}

	context.JSON(200, result)
}
