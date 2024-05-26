package controllers

import "rustdesk-api-server/utils/common"

type IndexController struct {
	BaseController
}

func (ctl *IndexController) Index() {
	ctl.JSON(common.JsonResult{
		Code:  1,
		Msg:   "Email:chen1350962574@Gmail.com | 1350962574@QQ.com",
		Data:  nil,
		Count: 0,
	})
}
