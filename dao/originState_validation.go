package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/originstate.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : OriginState (originstate)
// Endpoint 	        : OriginState (OriginStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:45
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in originstate_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// OriginState_ObjectValidation_impl provides Record/Object level validation for OriginState
func OriginState_ObjectValidation_impl(iAction string, iId string, iRec dm.OriginState) (dm.OriginState, string, error) {
	logs.Callout("OriginState", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("OriginState" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
