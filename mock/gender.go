package mock

// GenderInfo represents information of an individuals gender.

type GenderInfo struct {
	Name        string
	Description string
}

var GenderInfoMap map[string]GenderInfo

func init() {
	GenderInfoMap = make(map[string]GenderInfo)
	GenderInfoMap["Cis"] = GenderInfo{Name: "Cisgender", Description: "Cisgender is a term to describe someone whose gender matches what they were assigned at birth."}
	GenderInfoMap["Trans"] = GenderInfo{Name: "Transgender", Description: "Someone whose gender identity is different from the one they were assigned at birth might identify as transgender."}
	GenderInfoMap["Two-Spirit"] = GenderInfo{Name: "Two Sprit", Description: "A modern English term that an Indigenous person might identify as that comes from the traditional knowledge of Indigenous peoples in Canada/Turtle Island/North America. It can mean a person who walks between genders; one who carries the gifts of both males and female."}
	GenderInfoMap["Non-Binary"] = GenderInfo{Name: "Non Binary", Description: "Someone who does not identify as a man or a woman, or solely as one of those two genders. "}
	GenderInfoMap["Gender-Fluid"] = GenderInfo{Name: "Gender Fluid", Description: "Gender fluid may refer to a gender which varies over time. Someone who identifies as gender fluid may fluctuate between genders or express multiple genders at the same time. "}
	GenderInfoMap["Gender-Neutral"] = GenderInfo{Name: "Gender Neutral", Description: "Someone who feels they are neither male or female may identify as gender neutral. They may also identify as agender, genderless, non-binary and/or ungendered. "}
}

func GetGenderList() []string {
	rtn := []string{}
	for k := range GenderInfoMap {
		rtn = append(rtn, k)
	}
	return rtn
}

func IsValidGender(in string) bool {
	_, ok := GenderInfoMap[in]
	return ok
}
