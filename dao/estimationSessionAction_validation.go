package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationsessionaction.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : EstimationSessionAction (estimationsessionaction)
// Endpoint 	        : EstimationSessionAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:43
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in estimationsessionaction_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// EstimationSessionAction_ObjectValidation_impl provides Record/Object level validation for EstimationSessionAction
func EstimationSessionAction_ObjectValidation_impl(iAction string, iId string, iRec dm.EstimationSessionAction) (dm.EstimationSessionAction, string, error) {
	logs.Callout("EstimationSessionAction", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("EstimationSessionAction" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
