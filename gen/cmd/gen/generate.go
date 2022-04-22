package main

import (
	flag "github.com/spf13/pflag"
	"gorm.io/gen"
	"github.com/hinss/api/gen/conf"
	"github.com/hinss/api/gen/dal"
)

var needPrepare bool

func init() {
	flag.BoolVarP(&needPrepare, "prepare", "p", false, "need prepare step to execute create table sql!")
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	flag.Parse()

	if needPrepare {
		prepare(dal.DB) // prepare table for generate
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../../goodserver/v1", // 输出的路径
		OutFile: "xxxxx.go", //输出的文件名
		ModelPkgPath: "../../../goodserver/v1",
	})

	g.UseDB(dal.DB)

	// generate all table from database
	//g.ApplyBasic(g.GenerateAllTable()...)
	g.GenerateModelAs("goods", "Goods2")

	g.Execute()
}
