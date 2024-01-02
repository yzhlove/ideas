package main

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {

	df := diffmatchpatch.New()
	chars1, chars2, lines := df.DiffLinesToChars(csv1, csv2)
	diffs := df.DiffMain(chars1, chars2, false)
	dfs := df.DiffCharsToLines(diffs, lines)

	for _, d := range dfs {
		fmt.Printf("type:%s content:%q\n", d.Type, d.Text)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println(df.DiffPrettyText(dfs))
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
