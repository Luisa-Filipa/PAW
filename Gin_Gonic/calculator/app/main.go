package main

import (
	"math"
	"strconv"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/add/:a/:b", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": add(c.Param("a"), c.Param("b")),
		})
	})

	router.GET("/sub/:a/:b", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": sub(c.Param("a"), c.Param("b")),
		})
	})

	router.GET("/mult/:a/:b", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": mult(c.Param("a"), c.Param("b")),
		})
	})

	router.GET("/div/:a/:b", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": div(c.Param("a"), c.Param("b")),
		})
	})

	router.GET("/mod/:a/:b", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": mod(c.Param("a"), c.Param("b")),
		})
	})

	router.GET("/pow/:a/:b", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": pow(c.Param("a"), c.Param("b")),
		})
	})

	router.Run(":8080")
}


func add(a ,b string) string {
	val1,err1 := strconv.Atoi(a)
	val2,err2 := strconv.Atoi(b)

	if err1 != nil || err2 !=nil {
		return "Introduza apenas números"
	}

	return strconv.Itoa(val1+val2)
}

func sub(a ,b string) string {
	val1,err1 := strconv.Atoi(a)
	val2,err2 := strconv.Atoi(b)

	if err1 != nil || err2 !=nil {
		return "Introduza apenas números"
	}

	return strconv.Itoa(val1-val2)
}

func mult(a ,b string) string {
	val1,err1 := strconv.Atoi(a)
	val2,err2 := strconv.Atoi(b)

	if err1 != nil || err2 !=nil {
		return "Introduza apenas números"
	}

	return strconv.Itoa(val1*val2)
}

func div(a ,b string) string {
	val1,err1 := strconv.ParseFloat(a,64)
	val2,err2 := strconv.ParseFloat(b,64)

	if err1 != nil || err2 !=nil {
		return "Introduza apenas números"
	}

	res:=val1/val2

	return strconv.FormatFloat(res, 'g',6, 64)
}

func mod(a ,b string) string  {
	val1,err1 := strconv.ParseFloat(a,64)
	val2,err2 := strconv.ParseFloat(b,64)

	if err1 != nil || err2 !=nil {
		return "Introduza apenas números"
	}

	res:=math.Mod(val1,val2)

	return strconv.FormatFloat(res, 'g',6, 64)
}

func pow(a ,b string) string  {
	val1,err1 := strconv.ParseFloat(a,64)
	val2,err2 := strconv.ParseFloat(b,64)

	if err1 != nil || err2 !=nil {
		return "Introduza apenas números"
	}

	res:=math.Pow(val1,val2)

	return strconv.FormatFloat(res, 'g',6, 64)
}