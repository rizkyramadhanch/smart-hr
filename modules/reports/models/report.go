package models


// import EmplModels "smart-hr/modules/employees/models"


type ReportBPJS struct {
	ID         int    `json:"trx_id"`
	EmployeeID int    `json:"employee_id"`
	BpjsID     int    `json:"bpjs_id"`
	Year       int    `json:"year"`
	Month      int    `json:"month"`
	Sallary    int    `json:"sallary"`
	Premi      int    `json:"premi"`
	NIK        string `json:"nik"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	OfficeName string `json:"office_name"`
	Age        int    `json:"age"`
}

type Sallary struct {
	Periode        string  					`json:"periode"`
	Code           string  					`json:"code"`
	Ammount        float64 					`json:"ammount"`
	Component      string  					`json:"component"`
	Category       string  					`json:"category"`
	Kind           string  					`json:"kind"`
	TotalAllowance float64 					`json:"total_allowance"`
	TotalDeduction float64 					`json:"total_deduction"`
	TakeHomePay    float64 					`json:"take_home_pay"`
}



type Val struct {
	ID 		int `json:"id"`
	Ammount int `json:"ammount"`
}

//transaction_id	bpjs_trx_id	"year"	"month"	sallary	premi	nik	first_name	last_name	office_name	age
