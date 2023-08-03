package mock

import (
	"fmt"
)

// TitleInfo represents information of an indivual title.
type TitleInfo struct {
	Title string // The title
}

var TitleInfoMap map[string]TitleInfo

func init() {
	fmt.Println("mock init")
	TitleInfoMap = make(map[string]TitleInfo)
	TitleInfoMap["Mr"] = TitleInfo{Title: "Mr"}
	TitleInfoMap["Mrs"] = TitleInfo{Title: "Mrs"}
	TitleInfoMap["Miss"] = TitleInfo{Title: "Miss"}
	TitleInfoMap["Ms"] = TitleInfo{Title: "Ms"}
	TitleInfoMap["Dr"] = TitleInfo{Title: "Dr"}
	TitleInfoMap["Prof"] = TitleInfo{Title: "Prof"}
}

func GetList() []string {
	rtn := []string{}
	for k := range TitleInfoMap {
		rtn = append(rtn, k)
	}
	return rtn
}

func IsValidTitle(in string) bool {
	_, ok := TitleInfoMap[in]
	return ok
}
