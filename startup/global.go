package global

import (
	"database/sql"

	modelCommon "changian.com/jubo/model/common"
)

// 全域變數 | Pointer
var (
	Config     *modelCommon.Config
	Postgresql *sql.DB
)
