package models

import (
	"math"
)

type PostPageResult struct {
	List      []PostJson `json:"list"`
	Total     int        `json:"total"`    //文章总数
	Page      int        `json:"page"`     //页数
	PageSize  int        `json:"pageSize"` //每一页的博文数
	TotalPage int        //总页数
}

func Pagination(postJsons *[]PostJson, page int, pageSize int, totalPost int) PostPageResult {
	totalPage := int(math.Floor(float64(totalPost / pageSize)))
	if (totalPost % pageSize) != 0 {
		totalPage++
	}
	result := PostPageResult{
		List:      *postJsons,
		Total:     totalPost,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
	}
	if page < 1 {
		result.Page = 1
	}
	if page > result.TotalPage {
		result.Page = result.TotalPage
	}
	return result
}
