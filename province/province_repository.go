package province

import "database/sql"

type Province struct {
	Id   int64          `json:"id"`
	Name sql.NullString `json:"name"`
}

func FindAllWithPagination()  {
	
}