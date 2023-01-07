package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationstate.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : EstimationState (estimationstate)
// Endpoint 	        : EstimationState (EstimationStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in estimationstate_impl.go

// If there are fields below, create the methods in adaptor\estimationstate_impl.go
// START - GET API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func EstimationState_ObjectValidation_impl(iAction string, iId string, iRec dm.EstimationState) (dm.EstimationState, string, error) {
	logs.Callout("EstimationState", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("EstimationState" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}