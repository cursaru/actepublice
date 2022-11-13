package structures

type News struct {
	Headline string
	Body     string
}

type SmsStructure struct {
	Id         int
	TemplateId int
	SmsId      int
	Agent      string
}

type GeneralStructure struct {
	SmsList []SmsStructure
}

var SmsEmptyList = SmsStructure{
	Id:         0,
	TemplateId: 0,
	SmsId:      0,
	Agent:      "",
}

type EmptyInt struct {
	Nope int
}

var EmptyInteger = EmptyInt{
	Nope: 0,
}

type SmsReportRow struct {
	SMSID         int
	TemplateId    int
	TemplateName  string
	SMSText       string
	MediatelAgent string
	PhoneNumber   string
	Processedat   string
	Sentat        string
	SMSStatus     string
}

type SmsReport struct {
	Report []SmsReportRow
}
