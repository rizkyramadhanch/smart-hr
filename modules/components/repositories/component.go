package repositories

import(
	"smart-hr/database"
	"smart-hr/library/logger"
	"fmt"
	"smart-hr/modules/components/models"
)

type ComponentRepositories struct{}

type ComponentResult struct {
	ValID 			int		`json:"val_id"`
	EmployeeID  	int 	`json:"employee_id"`
	EmployeeName 	string 	`json:"employee_name"`
	ComponentID  	int 	`json:"component_id"`
	ComponentName 	string  `json:"component_name"`
	Category_id 	int 	`json:"category_id"`
	Category_name 	string 	`json:"category_name"`
	Kind_id 		int 	`json:"kind_id"`
	Kind_name 		string 	`json:"kind_name"`
	Ammount 		int 	`json:"ammount"`
}

func (cr *ComponentRepositories) GetComponent(EmployeeID int) (result []models.Component, err error){
	query := fmt.Sprintf("select b.id, b.code, b.name, a.id, a.component_value from employee_component_values a join master_component b on a.component_id = b.id where a.employee_id = %d" ,EmployeeID)
	arrComponent := []models.Component{}
	component := models.Component{}
	rows, err := database.DB.Query(query) 

	if err != nil {
		logger.Log.Println("[Query Error]", err)
	}

	for rows.Next(){
		errScan := rows.Scan(
			&component.ID,
			&component.Code,
			&component.Name,
			&component.ValID,
			&component.Val,
		)

		if errScan != nil {
			logger.Log.Println("[Error Get Component by Employee]", errScan)
		}

		arrComponent = append(arrComponent, component)
	}
	
	return arrComponent, nil
}

func (cr *ComponentRepositories) GetTotalAllowanceDeduction(EmployeeID int) (result []models.AllowanceDeduction, err error){
	query := fmt.Sprintf("select b.name, sum(a.component_value) from employee_component_values a join kind b on a.kind_id = b.id  where  date_part('year', periode) = 2020 and date_part('month', periode) = 2 group by b.name;")
	rows, err := database.DB.Query(query) 

	arrAllowanceDeduction := []models.AllowanceDeduction{}
	allowanceDeduction := models.AllowanceDeduction{}

	if err != nil {
		logger.Log.Println("[Query Error]", err)
	}

	for rows.Next(){
		errScan := rows.Scan(
			&allowanceDeduction.Name,
			&allowanceDeduction.Total,
		)

		if errScan != nil {
			logger.Log.Println("[Error Get Component by Employee]", errScan)
		}

		arrAllowanceDeduction = append(arrAllowanceDeduction, allowanceDeduction)

	}
	
	return arrAllowanceDeduction, nil
}

//add component salary into employee
func (cr *ComponentRepositories) AddComponentValue(form models.ValueComponent) (result string, err error){
	query := "insert into employee_component_values (employee_id,  component_id, category_id, kind_id, component_value) values ($1, $2, $3, $4, $5)"
	_, err = database.DB.Exec(query, form.EmployeeID, form.ComponentID, form.CategoryID, form.KindID, form.Val) 

	if err != nil {
		logger.Log.Println("[Query Error]", err)
		return "Failed to add new component" , nil
	}

	logger.Log.Println("New component added")
	return "New component for user has been added" , nil
}

//show detail of component value by id val_id
func (cr *ComponentRepositories) GetComponentValue(valID int) (result ComponentResult, err error){
	query := "select ecv.id, ecv.employee_id, e.first_name, ecv.component_id, com.name, ecv.category_id, mc.name, ecv.kind_id, k.name, ecv.component_value from employee_component_values ecv join employees e on ecv.employee_id = e.id join master_category mc on ecv.category_id = mc.id join kind k on ecv.kind_id = k.id join master_component com on ecv.component_id = com.id where ecv.id = $1"
	row := database.DB.QueryRow(query, valID) 

	componentResult := ComponentResult{}

	errScan := row.Scan(
			&componentResult.ValID,
			&componentResult.EmployeeID,
			&componentResult.EmployeeName,
			&componentResult.ComponentID,
			&componentResult.ComponentName,
			&componentResult.Category_id,
			&componentResult.Category_name,
			&componentResult.Kind_id,
			&componentResult.Kind_name,
			&componentResult.Ammount,
		)

	if errScan != nil {
		logger.Log.Println("[Query Error]", errScan)
		return componentResult , nil
	} else {
		logger.Log.Println(componentResult)
		return componentResult , nil
	}
}

//update component_values by val_id
func (cr *ComponentRepositories) EditComponentValue(form models.Ammount) (result string, err error){
	query := "update employee_component_values set component_value = $1 where id = $2"
	_, err = database.DB.Exec(query, form.Ammount, form.ValID) 
	if err != nil {
		logger.Log.Println("Error on update component value")
		return "Error on update component value" , err
	}
	return "Edit component successfully", nil
}