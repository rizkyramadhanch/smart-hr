package models

type Employee struct {
	ID 				int 	`json:"id"`
	EmployeeID 		string 	`json:"employee_id"`
	NIK				string	`json:"nik"`
	BpjsID			string	`json:"bpjs_id"`
	JoinDate 		string 	`json:"join_date"`
	FirstName 		string 	`json:"firstname"`
	LastName 		string 	`json:"lastname"`
	PlaceOfBirth 	string 	`json:"place_of_birth"`
	Birthday 		string 	`json:"birthday"`
	Address 		string 	`json:"address"`
	Age  			int 	
	Sallary  		int 	
	DepartmentID    int     
	LevelID    		int     
	CompanyID    	int     
}