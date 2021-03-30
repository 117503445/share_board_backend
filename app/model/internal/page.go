// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// Page is the golang structure for table page.
type Page struct {
    Id         uint64      `orm:"id,primary"  json:"id"`         //   
    CreatedAt  *gtime.Time `orm:"created_at"  json:"createdAt"`  //   
    UpdatedAt  *gtime.Time `orm:"updated_at"  json:"updatedAt"`  //   
    DeletedAt  *gtime.Time `orm:"deleted_at"  json:"deletedAt"`  //   
    PageNumber int         `orm:"page_number" json:"pageNumber"` //   
    BoardId    string      `orm:"board_id"    json:"boardId"`    //   
    Data       string      `orm:"data"        json:"data"`       //   
}