package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/userrole.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : UserRole (userrole)
// Endpoint 	        : UserRole (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:41:22
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in userrole_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// UserRole_ObjectValidation_impl provides Record/Object level validation for UserRole
func UserRole_ObjectValidation_impl(iAction string, iId string, iRec dm.UserRole) (dm.UserRole, string, error) {
	logs.Callout("UserRole", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("UserRole" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
