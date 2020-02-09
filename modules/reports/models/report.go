package models

type ReportBPJS struct {
	ID 			int 	`json:"trx_id"`
	EmployeeID  int     `json:"employee_id"`
	BpjsID 		int 	`json:"bpjs_id"`
	Year 		int 	`json:"year"`
	Month 		int 	`json:"month"`
	Sallary 	int 	`json:"sallary"`
	Premi 		int 	`json:"premi"`
	NIK 		string 	`json:"nik"`
	FirstName 	string 	`json:"first_name"`
	LastName 	string 	`json:"last_name"`
	OfficeName 	string 	`json:"office_name"`
	Age 		int 	`json:"age"`
}


//transaction_id	bpjs_trx_id	"year"	"month"	sallary	premi	nik	first_name	last_name	office_name	age