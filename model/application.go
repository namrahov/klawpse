package model

type PageableApplicationDto struct {
	List           *[]Application `json:"list"`
	HasNextPage    bool           `json:"hasNextPage"`
	LastPageNumber int            `json:"lastPageNumber"`
	TotalCount     int            `json:"totalCount"`
}

type Application struct {
	tableName struct{} `sql:"application" pg:",discard_unknown_columns"`

	Id              int64        `sql:"id"  json:"id"`
	RequestId       int64        `sql:"request_id" json:"requestId"`
	CheckedId       int64        `sql:"checked_id" json:"checkedId"`
	Person          string       `sql:"person" json:"person"`
	CustomerType    CustomerType `sql:"customer_type" json:"customerType"`
	CustomerName    string       `sql:"customer_name" json:"customerName"`
	FilePath        string       `sql:"file_path" json:"filePath"`
	CourtName       string       `sql:"court_name" json:"courtName"`
	JudgeName       string       `sql:"judge_name" json:"judgeName"`
	DecisionNumber  string       `sql:"decision_number" json:"decisionNumber"`
	DecisionDate    string       `sql:"decision_date" json:"decisionDate"`
	IsChecked       bool         `sql:"is_checked" json:"isChecked"`
	Status          Status       `sql:"status" json:"status"`
	StatusHistoryId int64        `sql:"status_history_id" json:"statusHistoryId"`
	Comments        []Comment    `json:"comments"`
	Documents       []Document   `json:"documents"`
	CreatedAt       string       `sql:"created_at" json:"createdAt"`
	BankDetails     []BankDetail `sql:"-" json:"bankDetails"`
}

type CustomerType string

const (
	Person   CustomerType = "PERSON"
	Taxpayer              = "TAXPAYER"
)

type Status string

const (
	Received   Status = "RECEIVED"
	Inprogress        = "IN_PROGRESS"
	Sent              = "SENT"
	Hold              = "HOLD"
)

type Priority string

const (
	Standard Priority = "STANDARD"
	High              = "HIGH"
)

type BankDetail struct {
	CustomerNo   string `json:"customerNo"`
	CustomerName string `json:"customerName"`
	RequestDate  string `json:"requestDate"`
}
