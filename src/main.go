package main

import (
	"strconv"

	"github.com/astaxie/beego"
)

func main() {
	/*
		This would match routes like the following:
		/sum/3/5
		/product/6/23
	*/
	beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
	beego.Run()
}

type mainController struct {
	beego.Controller
}

func (c *mainController) Get() {
	// Obtain the values of the route parameters defined in the route above
	operation := c.Ctx.Input.Param(":operation")
	num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
	num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))
	c.Data["operation"] = operation
	c.Data["num1"] = num1
	c.Data["num2"] = num2
	c.TplName = "result.html"

	// Perform the calculation depending on the 'operation' route parameter
	switch operation {
	case "sum":
		c.Data["result"] = Sum(num1, num2)
	case "product":
		c.Data["result"] = Product(num1, num2)
	default:
		c.TplName = "invalid-route.html"
	}

}

func Sum(num1 int, num2 int) int {
	return num1 + num2
}

func Product(num1 int, num2 int) int {
	return num1 * num2
}
