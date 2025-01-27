package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/session.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : Session (session)
// Endpoint 	        : Session (SessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:41:21
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in session_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Session_ObjectValidation_impl provides Record/Object level validation for Session
func Session_ObjectValidation_impl(iAction string, iId string, iRec dm.Session) (dm.Session, string, error) {
	logs.Callout("Session", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Session" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
