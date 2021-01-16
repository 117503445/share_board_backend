package service

import (
	"fmt"
	"shareboard/model"
	"shareboard/util"
)

// BoardGet 获取指定的 Page
// 如果不能合法创建,返回 nil
func PageGet(boardID string, pageNumber int) *model.Page {
	//util.Log().Debug("boardid %v\n", boardID)
	//util.Log().Debug("pageNumber %v\n", pageNumber)

	var page model.Page

	model.DB.Where(&model.Page{
		PageNumber: pageNumber,
		BoardID:    boardID,
		Data:       "",
	}).First(&page)

	//if page == nil {
	//
	//}

	if page.ID == 0 {
		page.Data =  "{\"version\":\"2.1.0\",\"objects\":[]}"
		page.PageNumber = pageNumber
		page.BoardID = boardID
		model.DB.Create(&page)
		util.Log().Debug("create page %v\n", page.ID)

	}

	fmt.Println(page)

	return &page
}
