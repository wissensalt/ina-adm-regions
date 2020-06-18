package util

import (
	"database/sql"
)

type Pagination struct {
	TotalRecord       int64          `json:"total_record"`
	TotalPageNumber   int64          `json:"total_page_number"`
	CurrentPageNumber int64          `json:"current_page_number"`
	Offset            int64          `json:"offset"`
	Limit             int64          `json:"limit"`
	MetaData          MetaPagination `json:"meta_data"`
}

type MetaPagination struct {
	IsFirst           bool  `json:"is_first"`
	IsLast            bool  `json:"is_last"`
}

type PaginationService interface {
	Info() Pagination
	Next() Pagination
	Prev() Pagination
	First() Pagination
	Last() Pagination
}

// constructor, return the first Pagination info
func (p *Pagination) Init(query string, limit int64, db *sql.DB) {
	p.TotalRecord = getTotalRecord(query, db)
	p.TotalPageNumber = getTotalPageNumber(p.TotalRecord, limit)
	p.CurrentPageNumber = 0
	p.Offset = 0
	p.Limit = limit
	p.MetaData = MetaPagination {
		IsFirst: true,
		IsLast: checkIsLast(p.CurrentPageNumber, p.TotalPageNumber),
	}
}

func (p Pagination) Info() Pagination {
	return constructPage(p)
}

func (p Pagination) Next() Pagination {
	nextPageNumber := p.CurrentPageNumber + 1
	if nextPageNumber > p.TotalPageNumber {
		return constructPage(p)
	}

	p.CurrentPageNumber += 1
	return constructPage(p)
}

func (p Pagination) Prev() Pagination {
	prevPageNumber := p.CurrentPageNumber - 1
	if prevPageNumber <= 0 {
		return constructPage(p)
	}

	p.CurrentPageNumber -= 1
	return constructPage(p)
}

func (p Pagination) First() Pagination {
	p.CurrentPageNumber = 1
	return constructPage(p)
}

func (p Pagination) Last() Pagination {
	p.CurrentPageNumber = p.TotalPageNumber
	return constructPage(p)
}

func constructPage(p Pagination) Pagination  {
	return Pagination{
		TotalRecord:       p.TotalRecord,
		TotalPageNumber:   p.TotalPageNumber,
		CurrentPageNumber: p.CurrentPageNumber,
		Offset:            getOffset(p.Limit, p.CurrentPageNumber),
		Limit:             p.Limit,
		MetaData:		 MetaPagination{
			IsFirst:         checkIsFirst(p.CurrentPageNumber),
			IsLast:          checkIsLast(p.CurrentPageNumber, p.TotalPageNumber),
		},
	}
}

func checkIsFirst(currentPageNumber int64) bool {
	if currentPageNumber == 1 {return true} else {return false}
}

func checkIsLast(currentPageNumber int64, totalPageNumber int64) bool {
	if currentPageNumber == totalPageNumber {return true} else {return false}
}

func getTotalRecord(queryCountTotalRecord string, db *sql.DB) int64 {
	if queryCountTotalRecord == "query" {
		return 50
	}

	rows, err := db.Query(queryCountTotalRecord)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	if rows.Next() {
		var totalNumber int64
		err = rows.Scan(&totalNumber)
		if err != nil {
			panic(err)
		}
		return totalNumber
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return 0
}

func getTotalPageNumber(totalNumber int64, limit int64) int64 {
	return totalNumber/limit
}

func getOffset(limit int64, currentPageNumber int64) int64  {
	offset := (limit * currentPageNumber) - limit
	if offset < 0 {
		return 0
	}

	return offset
}