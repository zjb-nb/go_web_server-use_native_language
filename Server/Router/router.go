package router

import (
	"fmt"
)

func Home(ctx *MyContext) {
	fmt.Fprintf(ctx.W, "%s\n", "hello word")
}
