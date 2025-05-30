package impl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"sky-take-out/pojo/entity"
	"sky-take-out/pojo/vo"
	"sky-take-out/resources/mapperParams"
	"sky-take-out/server/service"
	"strings"
	"time"
)

type ReportServiceImpl struct{}

func (r *ReportServiceImpl) GetTurnoverStatistics(begin time.Time, end time.Time) (vo.TurnoverReportVO, error) {
	dateList := ""
	for t := begin; !t.After(end); t = t.Add(time.Hour * 24) {
		dateList = dateList + t.Format("2006-01-02") + ","
	}
	dateList = strings.TrimSuffix(dateList, ",")

	m := map[interface{}]interface{}{
		"begin":  begin,
		"end":    end,
		"status": entity.COMPLETED,
	}

	turnoverList, err := mapperParams.OrderMapper.SumByMap(m)

	return vo.TurnoverReportVO{
		DateList:     dateList,
		TurnoverList: strings.Join(turnoverList, ","),
	}, err
}

func (r *ReportServiceImpl) GetSalesTop10(begin time.Time, end time.Time) (vo.SalesTop10ReportVO, error) {
	salesTop10, err := mapperParams.OrderMapper.GetSalesTop(begin, end)
	if err != nil {
		return vo.SalesTop10ReportVO{}, err
	}
	names := ""
	numbers := ""
	for _, goods := range salesTop10 {
		names = fmt.Sprintf("%s%s,", names, goods.Name)
		numbers = fmt.Sprintf("%s%d,", numbers, goods.Number)
	}
	names = strings.TrimSuffix(names, ",")
	numbers = strings.TrimSuffix(numbers, ",")
	return vo.SalesTop10ReportVO{
		NameList:   names,
		NumberList: numbers,
	}, err
}

func (r *ReportServiceImpl) GetUserStatistics(begin time.Time, end time.Time) (userVO vo.UserReportVO, err error) {
	dateList := ""
	newUserList := ""
	totalUserList := ""
	for t := begin; !t.After(end); t = t.Add(24 * time.Hour) {
		newUser, err := getUserCount(t, t.Add(24*time.Hour))
		if err != nil {
			return userVO, err
		}
		totalUser, err := getUserCount(*new(time.Time), t.Add(24*time.Hour))
		if err != nil {
			return userVO, err
		}
		dateList = dateList + t.Format("2006-01-02") + ","
		newUserList = newUserList + fmt.Sprintf("%d,", newUser)
		totalUserList = totalUserList + fmt.Sprintf("%d,", totalUser)
	}
	dateList = strings.TrimSuffix(dateList, ",")
	newUserList = strings.TrimSuffix(newUserList, ",")
	totalUserList = strings.TrimSuffix(totalUserList, ",")
	return vo.UserReportVO{
		DateList:      dateList,
		TotalUserList: newUserList,
		NewUserList:   totalUserList,
	}, err
}

func (r *ReportServiceImpl) GetOrderStatistics(begin time.Time, end time.Time) (reportVO vo.OrderReportVO, err error) {
	dateList := ""
	orderCountList := ""
	validOrderCountList := ""
	var totalOrderCount int64
	var validOrderCount int64
	for t := begin; !t.After(end); t = t.Add(24 * time.Hour) {
		dateList = dateList + t.Format("2006-01-02") + ","

		//总订单
		orderCount, err := getOrderCount(t, t.Add(24*time.Hour), -1)
		if err != nil {
			return reportVO, err
		}
		totalOrderCount += orderCount
		orderCountList = fmt.Sprintf("%s%d,", orderCountList, orderCount)

		//有效订单
		validCount, err := getOrderCount(t, t.Add(24*time.Hour), entity.COMPLETED)
		if err != nil {
			return reportVO, err
		}
		validOrderCount += validCount
		validOrderCountList = fmt.Sprintf("%s%d,", validOrderCountList, validCount)

	}
	dateList = strings.TrimSuffix(dateList, ",")
	orderCountList = strings.TrimSuffix(orderCountList, ",")
	validOrderCountList = strings.TrimSuffix(validOrderCountList, ",")
	var orderCompleteRate float64
	if totalOrderCount != 0 {
		orderCompleteRate = float64(validOrderCount) / float64(totalOrderCount)
	} else {
		orderCompleteRate = 0
	}
	return vo.OrderReportVO{
		DateList:            dateList,
		OrderCountList:      orderCountList,
		ValidOrderCountList: validOrderCountList,
		TotalOrderCount:     totalOrderCount,
		ValidOrderCount:     validOrderCount,
		OrderCompletionRate: orderCompleteRate,
	}, err
}

func (r *ReportServiceImpl) ExportBusinessData(w gin.ResponseWriter) error {

	var WorkSpaceService service.WorkSpaceService = new(WorkSpaceServiceImpl)

	dateBegin := time.Now().AddDate(0, 0, -30) // 30天前到昨天
	dateEnd := time.Now().AddDate(0, 0, -1)

	// 获取汇总数据
	businessDataVO, err := WorkSpaceService.GetBusinessData(
		time.Date(dateBegin.Year(), dateBegin.Month(), dateBegin.Day(), 0, 0, 0, 0, time.Local),
		time.Date(dateEnd.Year(), dateEnd.Month(), dateEnd.Day()+1, 0, 0, 0, 0, time.Local),
	)
	if err != nil {
		return fmt.Errorf("获取业务数据失败: %v", err)
	}

	// 打开Excel模板文件
	f, err := excelize.OpenFile("resources/template/运营数据报表模板.xlsx")
	if err != nil {
		return fmt.Errorf("打开Excel模板失败: %v", err)
	}
	defer func() {
		// 确保文件被关闭
		if err := f.Close(); err != nil {
			fmt.Printf("关闭Excel文件时出错: %v\n", err)
		}
	}()

	// 获取Sheet1
	sheet := "Sheet1"

	// 设置标题行数据
	axis := fmt.Sprintf("B2")
	f.SetCellStr(sheet, axis, fmt.Sprintf("时间：%s至%s",
		dateBegin.Format("2006-01-02"),
		dateEnd.Format("2006-01-02")))

	// 设置汇总数据
	f.SetCellFloat(sheet, "C4", businessDataVO.Turnover, 2, 64)
	f.SetCellFloat(sheet, "E4", businessDataVO.OrderCompletionRate, 4, 64)
	f.SetCellInt(sheet, "G4", businessDataVO.NewUsers)

	f.SetCellInt(sheet, "C5", businessDataVO.ValidOrderCount)
	f.SetCellFloat(sheet, "E5", businessDataVO.UnitPrice, 2, 64)

	// 设置每日数据
	for i := 0; i < 30; i++ {
		currentDate := dateBegin.AddDate(0, 0, i)

		businessData, err := WorkSpaceService.GetBusinessData(
			time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.Local),
			time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day()+1, 0, 0, 0, 0, time.Local),
		)
		if err != nil {
			return fmt.Errorf("获取%s业务数据失败: %v", currentDate.Format("2006-01-02"), err)
		}

		row := 8 + i
		f.SetCellStr(sheet, fmt.Sprintf("B%d", row), currentDate.Format("2006-01-02"))
		f.SetCellFloat(sheet, fmt.Sprintf("C%d", row), businessData.Turnover, 2, 64)
		f.SetCellInt(sheet, fmt.Sprintf("D%d", row), businessData.ValidOrderCount)
		f.SetCellFloat(sheet, fmt.Sprintf("E%d", row), businessData.OrderCompletionRate, 4, 64)
		f.SetCellFloat(sheet, fmt.Sprintf("F%d", row), businessData.UnitPrice, 2, 64)
		f.SetCellInt(sheet, fmt.Sprintf("G%d", row), businessData.NewUsers)
	}

	// 写入到输出流
	if err := f.Write(w); err != nil {
		return fmt.Errorf("写入Excel数据失败: %v", err)
	}

	return nil
}

func getOrderCount(begin time.Time, end time.Time, status int) (int64, error) {
	m := map[interface{}]interface{}{
		"begin":  begin,
		"end":    end,
		"status": status,
	}
	return mapperParams.OrderMapper.CountByMap(m)
}

func getUserCount(begin time.Time, end time.Time) (int64, error) {
	m := map[interface{}]interface{}{
		"begin": begin,
		"end":   end,
	}
	return mapperParams.UserMapper.CountByMap(m)
}
