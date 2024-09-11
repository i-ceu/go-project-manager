package enums

const (
	ONE_TEN          = "1-10"
	TEN_TWENTYFIVE   = "10-25"
	TWENTYFIVE_FIFTY = "25-50"
	ABOVE_FIFTY      = "above_50"
)

func GetAllSizes() []string {
	return []string{ONE_TEN, TEN_TWENTYFIVE, TWENTYFIVE_FIFTY, ABOVE_FIFTY}
}
