package requesthandler

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	TotalDataInPage int `json:"totalDataInPage"`
	TotalData       int `json:"totalData"`
	TotalPage       int `json:"totalPage"`
	ActivePage      int `json:"activePage"`
}

func InitPagination(r *http.Request) *Pagination {
	pag := Pagination{
		TotalDataInPage: 10,
		TotalData:       0,
		TotalPage:       0,
		ActivePage:      1,
	}

	// Init Pagination
	// Total Data In Page
	if d := r.URL.Query().Get("activePage"); d != "" {
		if c, err := strconv.Atoi(d); err == nil {
			pag.ActivePage = c
		}
	}
	// ActivePage
	if d := r.URL.Query().Get("totalDataInPage"); d != "" {
		if c, err := strconv.Atoi(d); err == nil {
			pag.TotalDataInPage = c
		}
	}

	return &pag
}
