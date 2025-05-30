package service

import (
	"github.com/gin-gonic/gin"
	"sky-take-out/pojo/vo"
	"time"
)

type ReportService interface {
	GetTurnoverStatistics(begin time.Time, end time.Time) (vo.TurnoverReportVO, error)
	GetSalesTop10(begin time.Time, end time.Time) (vo.SalesTop10ReportVO, error)
	GetUserStatistics(begin time.Time, end time.Time) (vo.UserReportVO, error)
	GetOrderStatistics(begin time.Time, end time.Time) (vo.OrderReportVO, error)
	ExportBusinessData(writer gin.ResponseWriter) error
}
