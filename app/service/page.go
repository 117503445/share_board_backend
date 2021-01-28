package service

import (
	"github.com/gogf/gf/frame/g"
	"share_board/app/dao"
	"share_board/app/model"
)

func GetPageByBoardIDPageNumber(boardID string, pageNumber int) *model.Page {
	var page model.Page
	_ = dao.Page.Where("board_id = ?", boardID).Where("page_number = ?", pageNumber).Struct(&page)
	//if err != nil {
	//	g.Log().Line().Debug(reflect.TypeOf(err))
	//	g.Log().Line().Error(err)
	//}
	if page.Id == 0 {
		page.Data = "{\"version\":\"2.1.0\",\"objects\":[]}"
		page.PageNumber = pageNumber
		page.BoardId = boardID
		if _, err := dao.Page.Insert(page); err != nil {
			g.Log().Line().Error(err)
		}
	}
	return &page
}
