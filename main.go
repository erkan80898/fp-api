package main

import (
	lib "flx/lib"

	"github.com/kr/pretty"
)

func main() {

	result := lib.GetListingVariantBySku([]string{"F_M_BB_AST"})
	pretty.Print(result)

}
