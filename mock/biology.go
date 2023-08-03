package mock

// BiologyInfo represents information of an indivual biology.
type BiologyInfo struct {
	Name        string
	Description string
}

var BiologyInfoMap map[string]BiologyInfo

func init() {
	BiologyInfoMap = make(map[string]BiologyInfo)
	BiologyInfoMap["M"] = BiologyInfo{Name: "Male", Description: ""}
	BiologyInfoMap["F"] = BiologyInfo{Name: "Female", Description: ""}
	BiologyInfoMap["I"] = BiologyInfo{Name: "Intersex", Description: ""}
	BiologyInfoMap["O"] = BiologyInfo{Name: "Other", Description: ""}
}

func GetBiologyList() []string {
	rtn := []string{}
	for k := range BiologyInfoMap {
		rtn = append(rtn, k)
	}
	return rtn
}

func GetBiologyInfo(biology string) BiologyInfo {
	return BiologyInfoMap[biology]
}

func IsValidBiology(biology string) bool {
	_, ok := BiologyInfoMap[biology]
	return ok
}
