package cir

type CIRDetails struct {

	CaseNumber string `json:"caseNumber"`
	CustomerName string `json:"customerName"`
	Subject string `json:"subject"`
	Description string `json:"description"`
	IssueDate string `json:"issueDate"`
	SiteName string `json:"siteName"`
	UpdateDate string `json:"updatedate"`
	ProductCompany string `json:"productCompany"`
	CustomerWantDate string `json:"customerWantDate"`
	EscalationLevel string `json:"escalationLevel"`
	TypeOfIssue string `json:"typeOfIssue"`
	Status string `json:"status"`
	Aging string `json:"aging"`
	ResolutionCommittedDate string `json:"resolutionCommittedDate"`

}
