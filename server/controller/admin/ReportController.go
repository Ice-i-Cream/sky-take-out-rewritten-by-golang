package admin

import (
	"github.com/gin-gonic/gin"
	"log"
	"sky-take-out/resources/functionParams"
	"sky-take-out/resources/serviceParams"
	"time"
)

type ReportController struct{}

func (r *ReportController) TurnoverStatistics(ctx *gin.Context) {

	exec := func(ctx *gin.Context) (data interface{}, err error) {
		log.Println("营业额数据统计")
		cstZone, _ := time.LoadLocation("Asia/Shanghai")
		begin, err := time.Parse("2006-01-02", ctx.Query("begin"))
		begin = begin.In(cstZone).Add(time.Hour * (-8))
		end, err := time.Parse("2006-01-02", ctx.Query("end"))
		end = end.In(cstZone).Add(time.Hour * (-8))
		return serviceParams.ReportService.GetTurnoverStatistics(begin, end)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (r *ReportController) UserStatistics(ctx *gin.Context) {

	exec := func(ctx *gin.Context) (data interface{}, err error) {
		log.Println("营业额数据统计")
		cstZone, _ := time.LoadLocation("Asia/Shanghai")
		begin, err := time.Parse("2006-01-02", ctx.Query("begin"))
		begin = begin.In(cstZone).Add(time.Hour * (-8))
		end, err := time.Parse("2006-01-02", ctx.Query("end"))
		end = end.In(cstZone).Add(time.Hour * (-8))
		return serviceParams.ReportService.GetUserStatistics(begin, end)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (r *ReportController) OrderStatistics(ctx *gin.Context) {

	exec := func(ctx *gin.Context) (data interface{}, err error) {
		cstZone, _ := time.LoadLocation("Asia/Shanghai")
		begin, err := time.Parse("2006-01-02", ctx.Query("begin"))
		begin = begin.In(cstZone).Add(time.Hour * (-8))
		end, err := time.Parse("2006-01-02", ctx.Query("end"))
		end = end.In(cstZone).Add(time.Hour * (-8))
		return serviceParams.ReportService.GetOrderStatistics(begin, end)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (r *ReportController) Top10(ctx *gin.Context) {

	exec := func(ctx *gin.Context) (data interface{}, err error) {
		cstZone, _ := time.LoadLocation("Asia/Shanghai")
		begin, err := time.Parse("2006-01-02", ctx.Query("begin"))
		begin = begin.In(cstZone).Add(time.Hour * (-8))
		end, err := time.Parse("2006-01-02", ctx.Query("end"))
		end = end.In(cstZone).Add(time.Hour * (-8))
		return serviceParams.ReportService.GetSalesTop10(begin, end)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}

func (r *ReportController) Export(ctx *gin.Context) {

	exec := func(ctx *gin.Context) (data interface{}, err error) {

		return nil, serviceParams.ReportService.ExportBusinessData(ctx.Writer)
	}
	data, err := exec(ctx)
	functionParams.PostProcess(ctx, err, data)
}
