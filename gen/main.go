package main

import (
	"context"
	"fmt"

	"github.com/hinss/api/gen/biz"
	"github.com/hinss/api/gen/conf"
	"github.com/hinss/api/gen/dal"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	biz.Query(context.Background())
}
