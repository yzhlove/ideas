package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {

	a := app.New()
	w := a.NewWindow("test windows")

	w.SetContent(layout())

	w.Resize(fyne.NewSize(1200, 650))
	w.ShowAndRun()

}

func layout() fyne.CanvasObject {

	left := widget.NewTextGridFromString(csv1)
	left.ShowLineNumbers = true
	//left.SetRowStyle(0,
	//	&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}})
	left.SetStyleRange(2, 12, 2, 12,
		&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}})
	left.SetRowStyle(2,
		&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 64, B: 192, A: 128}})

	left.SetStyleRange(3, 24, 3, 24,
		&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}})
	left.SetRowStyle(3,
		&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 64, B: 192, A: 128}})

	left.SetRow(0, widget.TextGridRow{
		Cells: left.Rows[0].Cells,
		Style: &widget.CustomTextGridStyle{
			FGColor: color.White,
			BGColor: &color.NRGBA{R: 64, G: 64, B: 192, A: 128}},
	},
	)

	right := widget.NewTextGridFromString(csv2)
	right.ShowLineNumbers = true
	right.SetStyleRange(2, 12, 2, 12,
		&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}})
	right.SetRowStyle(2,
		&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 64, B: 192, A: 128}})

	return container.NewGridWithColumns(2, container.NewScroll(left), container.NewScroll(right))
}

var (
	csv1 = `ID	OnceTicket	OnceDiamond	TenTimesTicket	TenTimesDiamond	LimitCount	ATypeProb	BTypeProb	PRDCount	PRDProb	ATypeFirstCount	ATypeCount	GuaranteePkgId	TargetCount	TargetPkgId	ATypePkg	BTypePkg	CTypePkg	ATypeUpProb	BTypeUpProb	StorageId	ExtraPkgId	GuranteeSR	StartTime	EndTime
1	null	300	null	3000	999	200	1500	51	50	20	100	10009	3	10008	10001	10002	10003	5000	0	1	10099	10006		
2	null	300	null	3000	999	200	1500	51	50	20	100	10009	3	10008	10001	10002	10003	5000	0	2	10099	10006		
3	null	300	null	3000	999	200	1500	51	50	20	100	10009	3	10008	10001	10002	10003	5000	0	3	10099	10006		
4	null	300	null	3000	999	200	1500	51	50	20	100	10009	3	10008	10001	10002	10003	5000	0	4	10099	10006		
`
	csv2 = `ID	OnceTicket	OnceDiamond	TenTimesTicket	TenTimesDiamond	LimitCount	ATypeProb	BTypeProb	PRDCount	PRDProb	ATypeFirstCount	ATypeCount	GuaranteePkgId	TargetCount	TargetPkgId	ATypePkg	BTypePkg	CTypePkg	ATypeUpProb	BTypeUpProb	StorageId	ExtraPkgId	GuranteeSR	StartTime	EndTime
1	null	300	null	3000	999	200	1500	51	50	20	100	10009	3	10008	10001	10002	10003	5000	0	1	10099	10006		
2	null	600	null	3000	999	200	1500	51	50	20	100	10009	3	10008	10001	10002	10003	5000	0	2	10099	10006		
3	null	300	null	8000	999	200	1500	51	50	20	100	10009	3	10008	10001	10002	10003	5000	0	3	10099	10006		
4	null	300	null	3000	999	200	1700	51	50	20	100	10009	3	10008	10001	10002	10003	5000	0	4	10099	10006		
`
)
