package datasource

import (
	"go-commons/passport/model"
)

var bases = map[int64]baseModel.UserInfo{
	1:{
		ID		:1000001,
		Name	:"Jumpman",
		Year	:1990,
	},
	2:{
		ID		:1000002,
		Name	:"Jumpman",
		Year	:1990,
	},
}