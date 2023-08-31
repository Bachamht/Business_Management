package model

import (
	"fmt"
)

type Client_Info struct {
	ClientName string `json:"c_name"`
	Company    string `json:"c_company"`
	Phone      string `json:"c_phone"`
	Detail     string `json:"detail"`
	Progess    string `json:"progess,omitempty"`
	Finished   bool   `json:"finished,omitempty"`
}

type SaleMan_Info struct {
	ID             int
	Name           string
	Company        string
	PhoneNumber    string
	ConfictContent string
}
type Conflict_Client struct {
	Name        string
	Company     string
	PhoneNumber string
}

type Conflict_userView struct {
	Name_b         string
	Company_b      string
	Phone_b        string
	ConfictContent string
}

// 查看业务类型
type ViewBusiness struct {
	CName    string `json:"c_name"`
	CCompany string `json:"c_company"`
	CPhone   string `json:"c_phone"`
	Detail   string `json:"detail"`
	Progress string `json:"progress"`
}

// 更新信息
type UpdateBusiness struct {
	CName    string `json:"c_name"`
	CCompany string `json:"c_company"`
	CPhone   string `json:"c_phone"`
	Detail   string `json:"detail"`
	Progress string `json:"progress"`
	Finished bool   `json:"finished"`
}

// 新增客户信息
func SubmitClientInfo(clientInfo *Client_Info, pNumber string) error {
	var saleMan *SaleMan_Info
	saleMan = GetPersonnalInfo(pNumber)

	conflictingClients := checkConflicts(clientInfo, pNumber)
	err := insertConflictInfo(saleMan, conflictingClients)
	if err != nil {
		fmt.Println("Error inserting conflict data into database:", err)
	}

	insertQuery := "INSERT INTO t_client_info (salesman_name, salesman_number, salesman_company, client_name, client_number, client_company, detail, progress, finished) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = DB.Exec(insertQuery, saleMan.Name, saleMan.PhoneNumber, saleMan.Company, clientInfo.ClientName, clientInfo.Phone, clientInfo.Company, clientInfo.Detail, "", 0)
	if err != nil {
		fmt.Println("Error inserting data into database:", err)
		return err
	}
	return nil
}

// 查看历史业务
func GetUnfinished(pNumber string) ([]ViewBusiness, error) {

	phoneNumber, err1 := DecodeBase64(pNumber)
	if err1 != nil {
		fmt.Println("decode error:", err)
	}

	var unfinished []ViewBusiness
	query := "SELECT client_name, client_company, client_number, detail, progress FROM t_client_info WHERE finished = false and salesman_number = ?"

	rows, err := DB.Query(query, phoneNumber)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var business ViewBusiness
		if err := rows.Scan(&business.CName, &business.CCompany, &business.CPhone, &business.Detail, &business.Progress); err != nil {
			return nil, err
		}
		unfinished = append(unfinished, business)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return unfinished, nil
}

// 获取个人信息
func GetPersonnalInfo(phoneNumber string) *SaleMan_Info {
	phoneNumber, err1 := DecodeBase64(phoneNumber)
	if err1 != nil {
		fmt.Println("decode error:", err)

	}
	query := "SELECT name, company FROM t_user_info WHERE phone_number = ?"
	row := DB.QueryRow(query, phoneNumber)

	var salesman SaleMan_Info
	salesman.PhoneNumber = phoneNumber
	err := row.Scan(
		&salesman.Name,
		&salesman.Company,
	)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	return &salesman
}

// 更新进度
func UpdateProgress(businesses []UpdateBusiness, pNumber string) error {

	phoneNumber, err1 := DecodeBase64(pNumber)
	if err1 != nil {
		fmt.Println("decode error:", err)
	}

	for _, business := range businesses {
		cNumber := business.CPhone
		updateQuery := "UPDATE t_client_info SET progress = ?, finished = ? WHERE client_number = ? and salesman_number = ?"
		_, err = DB.Exec(updateQuery, business.Progress, business.Finished, cNumber, phoneNumber)

		if err != nil {
			return err
		}
	}
	return nil
}

// 查看历史业务
func ViewHistory(pNumber string) ([]ViewBusiness, error) {

	phoneNumber, err1 := DecodeBase64(pNumber)
	if err1 != nil {
		fmt.Println("decode error:", err)
	}

	var finished []ViewBusiness
	query := "SELECT client_name, client_company, client_number, detail, progress FROM t_client_info WHERE  finished = true and salesman_number = ?"

	rows, err := DB.Query(query, phoneNumber)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var business ViewBusiness
		if err := rows.Scan(&business.CName, &business.CCompany, &business.CPhone, &business.Detail, &business.Progress); err != nil {
			return nil, err
		}
		finished = append(finished, business)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return finished, nil

}

// 检测冲突
func checkConflicts(clientInfo *Client_Info, pNumber string) *[]SaleMan_Info {

	var conflictingClients []SaleMan_Info

	//numberCheck := map[int]int
	nameCheck := map[int]int{}

	companyCheck := map[int]int{}

	pnumber, err1 := DecodeBase64(pNumber)
	if err1 != nil {
		fmt.Println("decode error:", err)

	}

	// 检查号码冲突
	queryNumber := "SELECT id, salesman_name, salesman_number, salesman_company, client_name, client_number, client_company  FROM t_client_info WHERE client_number = ? and salesman_number != ?"
	rows, err := DB.Query(queryNumber, clientInfo.Phone, pnumber)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var existingClient SaleMan_Info
		var conflictClient Conflict_Client
		existingClient.ConfictContent += "客户号码冲突 "
		err := rows.Scan(&existingClient.ID, &existingClient.Name, &existingClient.PhoneNumber, &existingClient.Company, &conflictClient.Name, &conflictClient.PhoneNumber, &conflictClient.Company)
		if err != nil {
			fmt.Println("error:", err)
		}

		if conflictClient.Name == clientInfo.ClientName {
			existingClient.ConfictContent += " 客户名字冲突 "
			nameCheck[existingClient.ID] = 1
		}

		if conflictClient.Company == clientInfo.Company {
			existingClient.ConfictContent += " 客户公司信息冲突 "
			companyCheck[existingClient.ID] = 1
		}

		conflictingClients = append(conflictingClients, existingClient)
	}

	// 检查名字冲突
	queryName := "SELECT id, salesman_name, salesman_number, salesman_company, client_name, client_number, client_company  FROM t_client_info WHERE client_name = ? and salesman_number != ?"
	rows, err = DB.Query(queryName, clientInfo.ClientName, pnumber)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var existingClient SaleMan_Info
		var conflictClient Conflict_Client
		err := rows.Scan(&existingClient.ID, &existingClient.Name, &existingClient.PhoneNumber, &existingClient.Company, &conflictClient.Name, &conflictClient.PhoneNumber, &conflictClient.Company)
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println("message2:", nameCheck[existingClient.ID])

		if nameCheck[existingClient.ID] != 1 {

			if conflictClient.Company == clientInfo.Company {
				existingClient.ConfictContent += " 客户公司信息冲突 "
				companyCheck[existingClient.ID] = 1
			}

			existingClient.ConfictContent += " 客户名字冲突 "
			conflictingClients = append(conflictingClients, existingClient)
		}
	}

	// 检查公司信息冲突
	queryCompany := "SELECT id,  salesman_name, salesman_number, salesman_company  FROM t_client_info WHERE client_company = ? and salesman_number != ?"
	rows, err = DB.Query(queryCompany, clientInfo.Company, pnumber)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var existingClient SaleMan_Info
		err := rows.Scan(&existingClient.ID, &existingClient.Name, &existingClient.PhoneNumber, &existingClient.Company)
		if err != nil {
			fmt.Println("error:", err)
		}

		if companyCheck[existingClient.ID] != 1 {
			existingClient.ConfictContent += " 客户公司信息冲突 "
			conflictingClients = append(conflictingClients, existingClient)
		}
	}

	return &conflictingClients
}

// 插入冲突消息
func insertConflictInfo(saleMan *SaleMan_Info, conflictingClients *[]SaleMan_Info) error {

	for _, conflictingClient := range *conflictingClients {
		insertQuery := "INSERT INTO t_conflict_info (name, phone, company, conflict_content, name_b, phone_b, company_b) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err := DB.Exec(insertQuery, saleMan.Name, saleMan.PhoneNumber, saleMan.Company, conflictingClient.ConfictContent, conflictingClient.Name, conflictingClient.PhoneNumber, conflictingClient.Company)
		if err != nil {
			return err
		}

	}

	return nil
}

// 查看冲突记录
func ViewConflict(pNumber string) ([]Conflict_userView, error) {
	phoneNumber, err1 := DecodeBase64(pNumber)
	if err1 != nil {
		fmt.Println("decode error:", err)
	}

	var ConflictInfo []Conflict_userView
	query := "SELECT name_b, phone_b, company_b, conflict_content FROM t_conflict_info WHERE phone = ? "

	rows, err := DB.Query(query, phoneNumber)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var conflict Conflict_userView
		if err := rows.Scan(&conflict.Name_b, &conflict.Phone_b, &conflict.Company_b, &conflict.ConfictContent); err != nil {
			return nil, err
		}
		ConflictInfo = append(ConflictInfo, conflict)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ConflictInfo, nil

}
