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
		OutPath: "/Users/zhanghaoxuan/go/src/api/goodserver/query/v1", // query输出的路径
		//OutFile: "banner_query.go", //输出的文件名
		ModelPkgPath: "/Users/zhanghaoxuan/go/src/api/goodserver/model/v1",
		Mode: gen.WithDefaultQuery,
	})

	g.UseDB(dal.DB)

	// generate all table from database
	g.ApplyBasic(g.GenerateModelAs("banner", "Banner"))

	g.Execute()
}
