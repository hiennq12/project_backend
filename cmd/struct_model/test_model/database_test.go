package testmodel_test

type TestRequest struct {
	Id   int64
	Name string
}

type TestResponse struct {
	LastInsertId int64
	RowEffect    int64
}
