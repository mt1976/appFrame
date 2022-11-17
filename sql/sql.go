package sql

// SqlDateToHTMLDate returns a string version of a SQL Date
func SqlDateToHTMLDate(inDate string) (outDate string) {
	var rtnDate string
	if inDate != "" {
		rtnDate = inDate[0:10]
	}
	return rtnDate
}
