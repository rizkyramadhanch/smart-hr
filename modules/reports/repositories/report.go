package repositories

import (
	"fmt"
	"smart-hr/database"
	"smart-hr/library/logger"
	ComModel "smart-hr/modules/components/models"
	"smart-hr/modules/components/repositories"
	ReportModel "smart-hr/modules/reports/models"
)

type ReportRepositories struct{}

type SallaryReport struct {
	ID                 int
	NIK                string
	FirstName          string
	Lastname           string
	JoinDate           string
	Department         string
	Overtime           int
	Component          []ComModel.Component
	AllowanceDeduction []ComModel.AllowanceDeduction
}

func (rr *ReportRepositories) ShowBPJSReport() (result []ReportModel.ReportBPJS, err error) {
	query := `select 
		b.id as bpjs_trx_id,
		e.id as employee_id,
		e.bpjs_id,
		b."year",
		b."month",
		e.sallary,
		b.premi,
		e.nik,
		e.first_name,
		e.last_name,
		c.name,
		e.age
			from bpjs_transaction b 
		join employees e on b.bpjs_id = e.bpjs_id
		join department d on d.id = e.department_id
		join level l on l.id = e.level_id
		join companies c on c.id = e.company_id`

	rows, err := database.DB.Query(query)
	if err != nil {
		logger.Log.Println(err)
		return result, err
	}

	reportBPJS := ReportModel.ReportBPJS{}

	for rows.Next() {
		errScan := rows.Scan(
			&reportBPJS.ID,
			&reportBPJS.EmployeeID,
			&reportBPJS.BpjsID,
			&reportBPJS.Year,
			&reportBPJS.Month,
			&reportBPJS.Sallary,
			&reportBPJS.Premi,
			&reportBPJS.NIK,
			&reportBPJS.FirstName,
			&reportBPJS.LastName,
			&reportBPJS.OfficeName,
			&reportBPJS.Age,
		)

		if errScan != nil {
			logger.Log.Println(errScan)
			return result, err
		}

		result = append(result, reportBPJS)
	}

	return result, nil
}

func (rr *ReportRepositories) ShowBPJSReportByID(userID int) (result []ReportModel.ReportBPJS, err error) {

	query := fmt.Sprintf("select b.id as bpjs_trx_id, e.id as employee_id, e.bpjs_id,")
	query = fmt.Sprintf("%s b.year, b.month, e.sallary, b.premi, e.nik, e.first_name, e.last_name,", query)
	query = fmt.Sprintf("%s c.name, e.age from bpjs_transaction b  ", query)
	query = fmt.Sprintf("%s join employees e on b.bpjs_id = e.bpjs_id  ", query)
	query = fmt.Sprintf("%s join department d on d.id = e.department_id  ", query)
	query = fmt.Sprintf("%s join level l on l.id = e.level_id  ", query)
	query = fmt.Sprintf("%s join companies c on c.id = e.company_id  ", query)
	query = fmt.Sprintf("%s where e.id = $1 order by year, month", query)

	rows, err := database.DB.Query(query, userID)
	if err != nil {
		logger.Log.Println(err)
		return result, err
	}

	reportBPJS := ReportModel.ReportBPJS{}

	for rows.Next() {
		errScan := rows.Scan(
			&reportBPJS.ID,
			&reportBPJS.EmployeeID,
			&reportBPJS.BpjsID,
			&reportBPJS.Year,
			&reportBPJS.Month,
			&reportBPJS.Sallary,
			&reportBPJS.Premi,
			&reportBPJS.NIK,
			&reportBPJS.FirstName,
			&reportBPJS.LastName,
			&reportBPJS.OfficeName,
			&reportBPJS.Age,
		)

		if errScan != nil {
			logger.Log.Println(errScan)
			return result, err
		}

		result = append(result, reportBPJS)
	}

	return result, nil
}

//SallaryReport All Employee
func (rr *ReportRepositories) SallaryReportAll() (result []SallaryReport, err error) {
	query := fmt.Sprintf("select id, nik, first_name, last_name from employees")

	rows, err := database.DB.Query(query)
	if err != nil {
		logger.Log.Println(err)
		return result, err
	}

	employee := SallaryReport{}
	arrEmployee := []SallaryReport{}

	for rows.Next() {
		errScan := rows.Scan(
			&employee.ID,
			&employee.NIK,
			&employee.FirstName,
			&employee.Lastname,
		)

		if errScan != nil {
			logger.Log.Println(errScan)
			return result, err
		}
		//masih dalam looping
		repoCom := repositories.ComponentRepositories{}
		arrComponents, errCom := repoCom.GetComponent(employee.ID)

		if errCom != nil {
			logger.Log.Println("[Error Repo Com]")
		}

		AD, errAD := repoCom.GetTotalAllowanceDeduction(employee.ID)

		if errAD != nil {
			logger.Log.Println("[Error Repo Com]")
		}

		employee.Component = arrComponents
		employee.AllowanceDeduction = AD

		arrEmployee = append(arrEmployee, employee)
	}

	return arrEmployee, nil
}

func (rr *ReportRepositories) SallaryReport(employeeID int) (result SallaryReport, err error) {
	query := fmt.Sprintf("select id, nik, first_name, last_name from employees where id = %d", employeeID)

	row  := database.DB.QueryRow(query)

	employee := SallaryReport{}

		errScan := row.Scan(
			&employee.ID,
			&employee.NIK,
			&employee.FirstName,
			&employee.Lastname,
		)

		if errScan != nil {
			logger.Log.Println(errScan)
			return result, err
		}
		//masih dalam looping
		repoCom := repositories.ComponentRepositories{}
		arrComponents, errCom := repoCom.GetComponent(employee.ID)

		if errCom != nil {
			logger.Log.Println("[Error Repo Com]")
		}

		AD, errAD := repoCom.GetTotalAllowanceDeduction(employee.ID)

		if errAD != nil {
			logger.Log.Println("[Error Repo Com]")
		}

		employee.Component = arrComponents
		employee.AllowanceDeduction = AD

	return employee, nil
}

