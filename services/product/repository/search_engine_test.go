package repository

import(
	"testing"
)


func TestBuildQuery(t *testing.T){
	//case 1
	q1 := BuildQuery("", "", 0, 1500)
	if q1 != "SELECT * FROM products WHERE price>0 AND price<1500"{
		t.Fatalf("\nError query is %s\n", q1)
	}

	//case 2




	//case 3
}