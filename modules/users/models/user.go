package models

type LoginForm struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct{
	ID			int	   			`json:"id"`
	Username 	string 			`json:"username"`
	Email 		string 			`json:"email"`
	Password 	string 			`json:"password"`
	RoleID		int				`json:"role_id,string,omitempty"`
	CompanyID	int				`json:"company_id,string,omitempty"`
	Role		string			`json:"role"`
	Company		string			`json:"company"`
	IsLogin 	bool		 	`json:"is_login"`
}