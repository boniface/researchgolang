package redisapi

import "gopkg.in/kataras/iris.v6"
import "time"
import "strconv"

type TimeElapsed struct {
	Startime int64 `json:"startime"`
	Endtime  int64 `json:"endtime"`
	Duration int64 `json:"duration"`
}

func Read(ctx *iris.Context) {

	var timer TimeElapsed

	timer.Startime = time.Now().UnixNano()

	ReadRecord()

	timer.Endtime = time.Now().UnixNano()
	timer.Duration = timer.Endtime - timer.Startime

	ctx.JSON(iris.StatusOK, timer)
}

func ReadAll(ctx *iris.Context) {

	var timer TimeElapsed

	timer.Startime = time.Now().UnixNano()

	ReadAllRecord()

	timer.Endtime = time.Now().UnixNano()
	timer.Duration = timer.Endtime - timer.Startime

	ctx.JSON(iris.StatusOK, timer)
}

func Create(ctx *iris.Context) {
	var timer TimeElapsed

	timer.Startime = time.Now().UnixNano()

	CreateRecord()

	timer.Endtime = time.Now().UnixNano()
	timer.Duration = timer.Endtime - timer.Startime

	ctx.JSON(iris.StatusOK, timer)
}

func CreateInNumbers(ctx *iris.Context) {

	number := ctx.Param("number")
	num, err := strconv.ParseInt(number, 10, 64)

	if err != nil{
		println(" Not a Valid String ", number)
	}

	var timer TimeElapsed

	timer.Startime = time.Now().UnixNano()

	var start int64 = 1

	for i := start; i <= num; i++{
		CreateRecord()
	}

	timer.Endtime = time.Now().UnixNano()
	timer.Duration = timer.Endtime - timer.Startime

	ctx.JSON(iris.StatusOK, timer)
}

func Delete(ctx *iris.Context) {

	var timer TimeElapsed

	timer.Startime = time.Now().UnixNano()

	DeleteRecord()

	timer.Endtime = time.Now().UnixNano()
	timer.Duration = timer.Endtime - timer.Startime

	ctx.JSON(iris.StatusOK, timer)
}