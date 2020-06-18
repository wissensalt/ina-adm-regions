package example

import (
	"database/sql"
	"fmt"
	"github.com/wissensalt/ina-adm-regions/constant"
	"github.com/wissensalt/ina-adm-regions/util"
)

func Run()  {
	pagination := new(util.Pagination)
	pagination.Init("query", 5, new(sql.DB))

	var servicePagination util.PaginationService = pagination

	fmt.Printf("Init Data : "+constant.PrintAsJson+constant.NewLine+constant.NewLine, pagination)

	for i:=int64(0) ; i<pagination.TotalPageNumber; i++ {
		nextPage := servicePagination.Next()
		fmt.Printf(constant.PrintAsJson+constant.NewLine, nextPage)
		servicePagination = nextPage
	}

	fmt.Printf(constant.NewLine+constant.NewLine)

	for i:=pagination.TotalPageNumber ; i>=0; i-- {
		prevPage := servicePagination.Prev()
		fmt.Printf(constant.PrintAsJson+constant.NewLine, prevPage)
		servicePagination = prevPage
	}

	fmt.Printf("First >>> "+constant.PrintAsJson+constant.NewLine,servicePagination.First())
	fmt.Printf("Last >>> "+constant.PrintAsJson+constant.NewLine,servicePagination.Last())
}
