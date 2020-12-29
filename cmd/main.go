package main

import (
	"flag"
	"fmt"
	"generateTool/cmd/gormgen"
	"log"
	"os"
	"strings"
)

type config struct {
	input             string
	inputController   string
	imports           []gormgen.ImportPkg
	importsController []gormgen.ImportPkg
	structs           []string
	queryPath         string
}

var (
	cnf          config
	logName      string
	transformErr bool
)

func parseFlags() {

	var input, structs, imports, queryPath, importsController, inputController string
	flag.StringVar(&structs, "structs", "", "[Required] The name of schema structs to generate structs for, comma seperated")
	flag.StringVar(&input, "input", "", "[Required] The name of the input file dir")
	flag.StringVar(&inputController, "inputController", "", "[Required] The name of the inputController file dir")
	flag.StringVar(&imports, "imports", "", "[Required] The name of the import  to import package")
	flag.StringVar(&importsController, "importsController", "", "[Required] The name of the importsController  to import package")
	flag.StringVar(&queryPath, "queryPath", "", "[Required] The name of the structs file dir")
	flag.StringVar(&logName, "logName", "", "[Option] The name of log db error")

	flag.BoolVar(&transformErr, "transformErr", false, "[Option] The name of transform db err")
	flag.Parse()
	//fmt.Println("imports:", imports)
	if input == "" || structs == "" || len(imports) == 0 || queryPath == "" {
		fmt.Println("----------------")
		flag.Usage()
		os.Exit(1)
	}

	cnf = config{
		input:           input,
		inputController: inputController,
		structs:         strings.Split(structs, ","),
		queryPath:       queryPath,
	}
	s := strings.Split(imports, ",")
	for _, v := range s {
		cnf.imports = append(cnf.imports, gormgen.ImportPkg{
			Pkg: v,
		})
	}
	inmpotsCon := strings.Split(importsController, ",")
	for _, inm := range inmpotsCon {
		cnf.importsController = append(cnf.importsController, gormgen.ImportPkg{
			Pkg: inm,
		})
	}
}

func main() {
	parseFlags()
	fmt.Println("cnf:", cnf)

	p := gormgen.NewParser(cnf.queryPath)
	fmt.Println("p:", p)
	gen := gormgen.NewGenerator(cnf.input).SetImportPkg(cnf.imports).SetLogName(logName)
	fmt.Println("gen:", gen)
	if transformErr {
		gen = gen.TransformError()
	}
	if err := gen.ParserASTPath(p, cnf.structs, cnf.input).Generate().Format().Flush(); err != nil {
		log.Fatalln(err)
	}
	//生成controller
	fmt.Println("controller-------------")
	genController := gormgen.NewGenerator(cnf.inputController).SetImportPkg(cnf.importsController).SetLogName(logName)
	if transformErr {
		genController = genController.TransformError()
	}
	if err := genController.ParserASTPath(p, cnf.structs, cnf.inputController).GenerateController().Format().FlushController(); err != nil {
		log.Fatalln(err)
	}

}
