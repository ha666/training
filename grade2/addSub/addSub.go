package addSub

import (
	"errors"
	"fmt"
	"github.com/ha666/golibs"
	"github.com/ha666/golibs/util/grand"
	"github.com/signintech/gopdf"
	"os/user"
	"runtime"

	"github.com/ha666/logs"
)

var (
	fontPath string
	fontName string
)

func Do() {
	logs.Info("hello Add Sub")
	initFont()
	genFormulas()
}

func initFont() {

	//region 根据当前系统选择字体
	os_name := runtime.GOOS
	switch os_name {
	case "linux":
		{
			fontName = "Microsoft YaHei"
			fontPath = "/usr/share/fonts/MSYH.TTF"
		}
	case "windows":
		{
			fontName = "微软雅黑"
			fontPath = "C:/Windows/Fonts/msyh.ttf"
		}
	case "darwin":
		{
			fontName = "微软雅黑"
			user, err := user.Current()
			if nil != err {
				logs.Error("查询当前目录出错:%s", err.Error())
				return
			}
			fontPath = user.HomeDir + "/Library/Fonts/MSYH.TTF"
		}
	default:
		{
			logs.Emergency("【initFont】未知系统:%s", os_name)
		}
	}
	//endregion

	//region 验证字体信息
	if golibs.Length(fontName) <= 0 {
		logs.Emergency("【initFont】字体名称未找到")
	}
	if golibs.Length(fontPath) <= 0 {
		logs.Emergency("【initFont】字体路径未找到")
	}
	//endregion

}

func genFormulas() {

	var (
		leftMargin = 10.0
		topMargin  = 45.0
		rowHeigh   = 30
		rows       = 25
	)

	//region 添加页面
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}})
	pdf.AddPage()
	err := pdf.AddTTFFont(fontName, fontPath)
	if err != nil {
		logs.Error("添加字体出错:%s", err.Error())
		return
	}
	//endregion

	//region 画表格
	{
		pdf.SetLineWidth(1)
		pdf.SetLineType("solid")
		for i := 0; i < rows+1; i++ {
			pdf.Line(leftMargin, topMargin+float64(i*rowHeigh), 585, topMargin+float64(i*rowHeigh))
		}
		pdf.Line(leftMargin, topMargin, leftMargin, topMargin+float64(rows*rowHeigh))
		pdf.Line(585/2, topMargin, 585/2, topMargin+float64(rows*rowHeigh))
		pdf.Line(585, topMargin, 585, topMargin+float64(rows*rowHeigh))
		err = pdf.SetFont(fontName, "", 14)
		if err != nil {
			logs.Error("设置字体出错:%s", err.Error())
			return
		}
	}
	//endregion

	//region 写标题
	{
		text01 := "二年级加减法训练"
		w01, _ := pdf.MeasureTextWidth(text01)
		left := (595.28 - w01) / 2.0
		pdf.SetX(left)
		pdf.SetY(15)
		err = pdf.Cell(nil, text01)
		if err != nil {
			logs.Error("写文本出错:%s", err.Error())
			return
		}
	}
	//endregion

	//region 写内容
	{
		for i := 0; i < 2; i++ {
			for j := 0; j < rows; j++ {
				formula, err := genFormula()
				if err != nil {
					logs.Error("生成算式没有成功:%s", err.Error())
					return
				}
				pdf.SetX(float64(i)*585/2 + leftMargin*2)
				pdf.SetY(topMargin + float64(j*rowHeigh) + 10)
				err = pdf.Cell(nil, formula)
				if err != nil {
					logs.Error("写文本出错:%s", err.Error())
					return
				}
			}
		}
	}
	//endregion

	//region 写入pdf
	{
		err = pdf.WritePdf("hello.pdf")
		if err != nil {
			logs.Error("生成pdf出错:%s", err.Error())
			return
		}
		logs.Info("生成成功")
	}
	//endregion

}

func genFormula() (string, error) {

	for i := 0; i < 10; i++ {

		//region 先随机出算法
		isAdd := true
		a, b := 0, 0
		s := grand.Rand(100, 999)
		isAdd = s%2 == 0
		//endregion

		//region 随机出数字
		a = grand.Rand(1, 99)
		b = grand.Rand(1, 99)
		//endregion

		//region 验证结果
		if a == b {
			continue
		}
		//endregion

		//region 生成算式
		if isAdd {
			return fmt.Sprintf("%d+%d=", a, b), nil
		} else {
			if b > a {
				return fmt.Sprintf("%d-%d=", b, a), nil
			} else {
				return fmt.Sprintf("%d-%d=", a, b), nil
			}
		}
		break
		//endregion

	}

	return "", errors.New("生成失败")

}
