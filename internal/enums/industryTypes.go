package enums

const (
	IT            = "IT"
	MARKETING     = "marketing"
	INSURANCE     = "insurance"
	MANUFACTURING = "manufacturing"
	QA            = "quality assurance"
	SOFWARE_DEV   = "software development"
)

func GetAllIndustries() []string {
	return []string{IT, MARKETING, INSURANCE, MANUFACTURING, QA, SOFWARE_DEV}
}
