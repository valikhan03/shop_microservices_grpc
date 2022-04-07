package repository

import (
	"fmt"
	"strings"
)

func BuildQuery(title string, category string, minprice int64, maxprice int64) string {
	var query = "SELECT * FROM products WHERE "
	base_len := len(query)
	if len(strings.TrimSpace(title)) != 0{
		query_params := strings.Replace(title, " ", " & ", -1)
		query += fmt.Sprintf("to_tsvector(products.title) @@ to_tsquery('%s')", query_params)	
		//after tests add || to_tsvector(products.description) @@ to_tsquery(params...)
	}

	if len(strings.TrimSpace(category)) != 0{
		if len(query)>base_len{
			query += " AND "
		}
		query += fmt.Sprintf("category=%s", category)
	}

	if(maxprice > 0){
		if len(query)>base_len{
			query += " AND "
		}
		query += fmt.Sprintf("price>%d AND price<%d", minprice, maxprice)
	}

	return query
}

