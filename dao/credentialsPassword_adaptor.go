package dao

import (
	"fmt"
	"time"

	"github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
	"github.com/pkg/errors"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/credentialspassword.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : CredentialsPassword (credentialspassword)
// Endpoint 	        : CredentialsPassword (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 24/02/2023 at 16:16:46
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in credentialspassword_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// CredentialsPassword_GetList_impl provides the implementation to get a list of records
func CredentialsPassword_GetList_impl(filter string) (int, []dm.CredentialsPassword, error) {
	return 0, nil, nil
}

// ----------------------------------------------------------------
// CredentialsPassword_GetByID_impl provides the implementation to get a record by ID
func CredentialsPassword_GetByID_impl(id string) (int, dm.CredentialsPassword, error) {
	return 0, dm.CredentialsPassword{}, nil
}

//

// ----------------------------------------------------------------

// ----------------------------------------------------------------
// CredentialsPassword_NewID_impl provides the implementation to generate a new ID for a new record
func CredentialsPassword_NewID_impl(rec dm.CredentialsPassword) string { return rec.ID }

// ----------------------------------------------------------------
// CredentialsPassword_New_impl provides the implementation to delete a new record
func CredentialsPassword_Delete_impl(id string) error { return nil }

// ----------------------------------------------------------------

// ----------------------------------------------------------------
// CredentialsPassword_Store_impl provides the implementation for the store adaptor
func CredentialsPassword_Update_impl(id string, rec dm.CredentialsPassword, usr string) error {
	logs.Storing("CredentialsPassword", id+"-"+rec.UserName)
	// This will update the relevant credentials, updateding the password, and the password expiry date
	objectName := Translate("Credentials", "EstimationSession")
	usage := "The number of days a password is valid for." + core.TEXTAREA_CR
	usage += "This is used to determine when a password needs to be changed."
	usage += "The default is 30 days."
	passwordLife, _ := Data_GetInt(objectName, "Life", dm.Data_Category_Setting, usage)
	fmt.Printf("passwordLife: %v\n", passwordLife)
	if passwordLife == 0 {
		logs.Warning("CredentialsPassword - Unable to get Password Life from Data - Using 30 days")
		passwordLife = 30
	}
	fmt.Printf("passwordLife: %v\n", passwordLife)
	// Get the credentials record
	_, cred, err := Credentials_GetByID(rec.UserName)
	if err != nil {
		logs.Error("Unable to get credentials record for user: "+rec.ID, err)
		return err
	}

	if cred.Password != "" {
		// Check the old password
		if !core.CheckPasswordHash(rec.PasswordOld, cred.Password) {
			logs.Warning("Unable to update credentials record for user (Old Passwords Do No Match): " + rec.ID)
			return errors.New("old passwords do not match")
		}
	}

	//Hash the password
	hash, err := core.HashPassword(rec.PasswordNew)
	if err != nil {
		logs.Error("Unable to hash password for user: "+rec.ID, err)
		return err
	}
	// Update the password
	cred.Password = hash

	// Update the password expiry date
	expiry := time.Now().AddDate(0, 0, passwordLife)
	expiryString := expiry.Format(core.DATEFORMAT)
	cred.PasswordExpiry = expiryString

	cred.Notes = core.AddActivity_ForUser(cred.Notes, "Password changed", usr)
	logs.Information("CredentialsPassword - Password expiry date set to: ", expiryString)
	// Update the credentials record
	_, Sterr := Credentials_StoreSystem(cred)
	if Sterr != nil {
		logs.Error("Unable to update credentials record for user: "+rec.ID, Sterr)
		return Sterr
	}
	usage = "The date the users password was last changed."
	Data_Put("Password", cred.Username, "LastChanged", time.Now().Format(core.DATEMSG), usage)

	return nil
}

// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN CredentialsPassword_UserName
// BEGIN CredentialsPassword_UserName
// BEGIN CredentialsPassword_UserName
// ----------------------------------------------------------------
// CredentialsPassword_UserName_OnStore_impl provides the implementation for the callout
func CredentialsPassword_UserName_OnStore_impl(fieldval string, rec dm.CredentialsPassword, usr string) (string, error) {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_UserName_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// CredentialsPassword_UserName_OnFetch_impl provides the implementation for the callout
func CredentialsPassword_UserName_OnFetch_impl(rec dm.CredentialsPassword) string {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_UserName_scrn, GET, rec.ID)
	return rec.UserName
}

// ----------------------------------------------------------------
// END   CredentialsPassword_UserName
// END   CredentialsPassword_UserName
// END   CredentialsPassword_UserName
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN CredentialsPassword_PasswordOld
// BEGIN CredentialsPassword_PasswordOld
// BEGIN CredentialsPassword_PasswordOld
// ----------------------------------------------------------------
// CredentialsPassword_PasswordOld_OnStore_impl provides the implementation for the callout
func CredentialsPassword_PasswordOld_OnStore_impl(fieldval string, rec dm.CredentialsPassword, usr string) (string, error) {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_PasswordOld_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// CredentialsPassword_PasswordOld_OnFetch_impl provides the implementation for the callout
func CredentialsPassword_PasswordOld_OnFetch_impl(rec dm.CredentialsPassword) string {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_PasswordOld_scrn, GET, rec.ID)
	return rec.PasswordOld
}

// ----------------------------------------------------------------
// END   CredentialsPassword_PasswordOld
// END   CredentialsPassword_PasswordOld
// END   CredentialsPassword_PasswordOld
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN CredentialsPassword_PasswordNew
// BEGIN CredentialsPassword_PasswordNew
// BEGIN CredentialsPassword_PasswordNew
// ----------------------------------------------------------------
// CredentialsPassword_PasswordNew_OnStore_impl provides the implementation for the callout
func CredentialsPassword_PasswordNew_OnStore_impl(fieldval string, rec dm.CredentialsPassword, usr string) (string, error) {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_PasswordNew_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// CredentialsPassword_PasswordNew_OnFetch_impl provides the implementation for the callout
func CredentialsPassword_PasswordNew_OnFetch_impl(rec dm.CredentialsPassword) string {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_PasswordNew_scrn, GET, rec.ID)
	return rec.PasswordNew
}

// ----------------------------------------------------------------
// END   CredentialsPassword_PasswordNew
// END   CredentialsPassword_PasswordNew
// END   CredentialsPassword_PasswordNew
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN CredentialsPassword_PasswordConfirm
// BEGIN CredentialsPassword_PasswordConfirm
// BEGIN CredentialsPassword_PasswordConfirm
// ----------------------------------------------------------------
// CredentialsPassword_PasswordConfirm_OnStore_impl provides the implementation for the callout
func CredentialsPassword_PasswordConfirm_OnStore_impl(fieldval string, rec dm.CredentialsPassword, usr string) (string, error) {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_PasswordConfirm_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// CredentialsPassword_PasswordConfirm_OnFetch_impl provides the implementation for the callout
func CredentialsPassword_PasswordConfirm_OnFetch_impl(rec dm.CredentialsPassword) string {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_PasswordConfirm_scrn, GET, rec.ID)
	return rec.PasswordConfirm
}

// ----------------------------------------------------------------
// END   CredentialsPassword_PasswordConfirm
// END   CredentialsPassword_PasswordConfirm
// END   CredentialsPassword_PasswordConfirm
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
