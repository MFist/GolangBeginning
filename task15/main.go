package main

import (
	"GolangBeginning/task15/geometry"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var rect = geometry.Rect{Width: 3, Height: 4}
var circle = geometry.Circle{Radius: 0}

var err error
var radius float64
var width float64
var height float64

func areaHandler(c *gin.Context) {
	if c.Param("radius") != "" {
		radius, err = strconv.ParseFloat(c.Param("radius"), 64)
		circle = geometry.Circle{Radius: radius}
		c.JSON(200, gin.H{"Circle's area": geometry.Geometry.Area(circle)})
	} else {
		width, err = strconv.ParseFloat(c.Param("width"), 64)
		height, err = strconv.ParseFloat(c.Param("height"), 64)
		rect = geometry.Rect{Width: width, Height: height}
		c.JSON(200, gin.H{"Rectangle's area": geometry.Geometry.Area(rect)})
	}
	fmt.Println(err)
}

func perimeterHandler(c *gin.Context) {
	if c.Param("radius") != "" {
		radius, err = strconv.ParseFloat(c.Param("radius"), 64)
		circle = geometry.Circle{Radius: radius}
		c.JSON(200, gin.H{"Circle's perimeter": geometry.Geometry.Perim(circle)})
	} else {
		width, err = strconv.ParseFloat(c.Param("width"), 64)
		height, err = strconv.ParseFloat(c.Param("height"), 64)
		rect = geometry.Rect{Width: width, Height: height}
		c.JSON(200, gin.H{"Rectangle's perimeter": geometry.Geometry.Perim(rect)})
	}
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/rectanglearea/:width/:height", areaHandler)
	r.GET("/rectangleperimeter/:width/:height", perimeterHandler)
	r.GET("/circlearea/:radius", areaHandler)
	r.GET("/circleperimeter/:radius", perimeterHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
