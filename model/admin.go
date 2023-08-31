package model

type ApplicantInfo struct {
	Name      string `json:"name"`
	Company   string `json:"company"`
	Phone     string `json:"phone"`
	Permitted int    `json:"permitted"`
}

// 业务类型
type HistoryBusiness struct {
	Name     string `json:"name"`
	Company  string `json:"company"`
	Phone    string `json:"phone"`
	CName    string `json:"c_name"`
	CCompany string `json:"c_company"`
	CPhone   string `json:"c_phone"`
	Detail   string `json:"detail"`
	Progress string `json:"progress"`
	Finished bool   `json:"finished"`
}

// 业务员信息
type SaleMan struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Phone   string `json:"phone"`
}

type ConflictAdmin struct {
	Name_a          string
	Company_a       string
	Phone_a         string
	Name_b          string
	Company_b       string
	Phone_b         string
	ConflictContent string
}

// 查看业务员申请
func GetApplicantInfo() ([]ApplicantInfo, error) {

	rows, err := DB.Query("SELECT name, company, phone_number, permitted FROM t_user_info where permitted = 0")
	if err != nil {
		return nil, err
	}

	Applicant := []ApplicantInfo{}

	for rows.Next() {
		var applicant ApplicantInfo
		if err := rows.Scan(&applicant.Name, &applicant.Company, &applicant.Phone, &applicant.Permitted); err != nil {
			return nil, err
		}
		Applicant = append(Applicant, applicant)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Applicant, nil
}

// 通过审核
func UpdatePermitted(applicants []ApplicantInfo) error {

	for _, applicant := range applicants {
		_, err := DB.Exec("UPDATE t_user_info SET permitted = ? WHERE phone_number = ?",
			applicant.Permitted, applicant.Phone)
		if err != nil {
			return err
		}
	}
	return nil
}

// 查看历史业务
func ViewHistoryAdmin() ([]HistoryBusiness, error) {

	var History []HistoryBusiness
	query := "SELECT salesman_name, salesman_number, salesman_company, client_name, client_company, client_number, detail, progress, finished FROM t_client_info"
	rows, err := DB.Query(query) // 使用解析后的手机号
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var business HistoryBusiness
		err := rows.Scan(&business.Name, &business.Company, &business.Phone, &business.CName, &business.CCompany, &business.CPhone, &business.Detail, &business.Progress, &business.Finished)
		if err != nil {
			return nil, err
		}
		History = append(History, business)
	}
	return History, nil
}

// 查看业务员列表
func InfoSaleMan() ([]SaleMan, error) {
	var SaleManAdmin []SaleMan
	query := "SELECT name, company, phone_number FROM t_user_info where permitted = 1"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var sMan SaleMan
		err := rows.Scan(&sMan.Name, &sMan.Company, &sMan.Phone)
		if err != nil {
			return nil, err
		}
		SaleManAdmin = append(SaleManAdmin, sMan)
	}
	return SaleManAdmin, nil
}

// 查看冲突记录
func ViewConflictAdmin() ([]ConflictAdmin, error) {

	var ConflictAd []ConflictAdmin
	query := "SELECT name_b, phone_b, company_b, name, phone, company, conflict_content FROM t_conflict_info"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		var ConflictA ConflictAdmin
		err := rows.Scan(&ConflictA.Name_b, &ConflictA.Phone_b, &ConflictA.Company_b, &ConflictA.Name_a, &ConflictA.Phone_a, &ConflictA.Company_a, &ConflictA.ConflictContent)
		if err != nil {
			return nil, err
		}
		ConflictAd = append(ConflictAd, ConflictA)
	}
	return ConflictAd, nil
}
