package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/projectstate.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : ProjectState (projectstate)
// Endpoint 	        : ProjectState (ProjectStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in projectstate_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// ProjectState_ObjectValidation_impl provides Record/Object level validation for ProjectState
func ProjectState_ObjectValidation_impl(iAction string, iId string, iRec dm.ProjectState) (dm.ProjectState, string, error) {
	logs.Callout("ProjectState", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("ProjectState" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
