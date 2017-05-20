package arangoapi

import "gopkg.in/kataras/iris.v6"



type TimeElapsed struct {
	Startime int64 `json:"startime"`
	Endtime uint64 `json:"endtime"`
	Duration int64 `json:"duration"`
}

func Read(ctx *iris.Context) {

	var timer TimeElapsed




	ctx.JSON(iris.StatusOK, timer)
}
