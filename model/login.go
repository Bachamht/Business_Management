package model

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"
)

type User_info struct {
	ProfilePhoto string `json:"profile_photo,omitempty"`
	Name         string `json:"name"`
	Company      string `json:"company"`
	PhoneNumber  string `json:"phone"`
	Password     string `json:"password"`
	Permitted    int    `json:"permitted,omitempty"`
}

type SessionInfo struct {
	PhoneNumber    string
	ExpirationTime time.Time
}

var Sessions = map[string]SessionInfo{}

// 设置session有效期五分钟
const sessionDuration = 5 * time.Minute

// 检查手机号重复
func CheckNumber(phoneNumber string) *User_info {

	pNumber, err1 := DecodeBase64(phoneNumber)
	if err1 != nil {
		fmt.Println("decode error:", err)

	}
	query := "SELECT * FROM t_user_info WHERE phone_number = ?"
	row := DB.QueryRow(query, pNumber)

	var user User_info
	err := row.Scan(
		&user.Name,
		&user.PhoneNumber,
		&user.Password,
		&user.Company,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil // 没有找到重复
		}
		fmt.Println("Error retrieving user:", err)
		return nil
	}

	return &user
}

// 新注册用户插入数据库
func InsertUser(user User_info) error {

	pNumber, err1 := DecodeBase64(user.PhoneNumber)
	if err1 != nil {
		fmt.Println("decode error:", err)
	}

	pword, err1 := DecodeBase64(user.Password)
	if err1 != nil {
		fmt.Println("decode error:", err)
	}

	query := "INSERT INTO t_user_info (profile_photo, name, phone_number, password, company, permitted) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := DB.Exec(query, user.ProfilePhoto, user.Name, pNumber, pword, user.Company, user.Permitted)
	if err != nil {
		return err
	}
	return nil
}

// 获得对应账户的密码
func GetPasswrd(account string) (string, error, bool) {
	pNumber, err1 := DecodeBase64(account)
	if err1 != nil {
		fmt.Println("decode error:", err)

	}

	query := "SELECT password FROM t_admin_info WHERE account = ?"
	row := DB.QueryRow(query, pNumber)

	var passwordAdmin string
	err := row.Scan(&passwordAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			query := "SELECT password FROM t_user_info WHERE phone_number = ?"
			row := DB.QueryRow(query, pNumber)
			var password string
			err := row.Scan(&password)
			if err != nil {
				if err == sql.ErrNoRows {
					return "", nil, false
				}
				return "", err, false
			}
			return password, nil, false
		}
		return "", err, false
	}
	return passwordAdmin, nil, true
}

// 判定审核是否通过
func IsPermitted(phoneNumber string) bool {
	pNumber, err1 := DecodeBase64(phoneNumber)
	if err1 != nil {
		fmt.Println("decode error:", err)

	}
	queryAdmin := "SELECT password FROM t_admin_info WHERE account = ?"
	rowAdmin := DB.QueryRow(queryAdmin, pNumber)
	var password string
	errAdmin := rowAdmin.Scan(&password)
	if errAdmin != nil {
		fmt.Println("error:", err)
	}
	if password != "" {
		return true
	}
	query := "SELECT permitted FROM t_user_info WHERE phone_number = ?"
	row := DB.QueryRow(query, pNumber)
	var permitted int
	err := row.Scan(&permitted)
	if err != nil {
		fmt.Println("error:", err)
	}
	if permitted == 0 {
		return false
	}
	return true
}

// base64解码
func DecodeBase64(encodedStr string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

// 生成session
func GenerateSession() string {
	tokenBytes := make([]byte, 16)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		panic("Error generating session token")
	}
	return fmt.Sprintf("%x", tokenBytes)
}
