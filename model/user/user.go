// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldRegTime holds the string denoting the reg_time field in the database.
	FieldRegTime = "reg_time"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldNickname,
	FieldAccount,
	FieldPassword,
	FieldRegTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// AccountValidator is a validator for the "account" field. It is called by the builders before save.
	AccountValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultRegTime holds the default value on creation for the "reg_time" field.
	DefaultRegTime func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)
