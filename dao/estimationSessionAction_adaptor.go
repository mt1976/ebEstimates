package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationsessionaction.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : EstimationSessionAction (estimationsessionaction)
// Endpoint 	        : EstimationSessionAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in estimationsessionaction_impl.go

func EstimationSessionAction_GetList_impl(filter string) (int, []dm.EstimationSessionAction, error) {
	return 0, nil, nil
}
func EstimationSessionAction_GetByID_impl(id string) (int, dm.EstimationSessionAction, error) {
	return 0, dm.EstimationSessionAction{}, nil
}

func EstimationSessionAction_NewID_impl(rec dm.EstimationSessionAction) string { return rec.ID }
func EstimationSessionAction_Delete_impl(id string) error                      { return nil }

func EstimationSessionAction_Update_impl(id string, rec dm.EstimationSessionAction, usr string) error {
	return nil
}

// If there are fields below, create the methods in adaptor\estimationsessionaction_impl.go
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
