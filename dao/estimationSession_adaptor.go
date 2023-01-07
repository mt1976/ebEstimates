package dao

import (
	"fmt"
	"strconv"
	"time"

	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationsession.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:57:49
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in estimationsession_impl.go

// If there are fields below, create the methods in adaptor\estimationsession_impl.go
// START - GET API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func EstimationSession_EstimationStateID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_EstimationStateID_scrn, PUT, rec.EstimationSessionID)
	// Check the status of the Estimation
	status := rec.EstimationStateID
	// Get EstimationState from DB
	_, est, _ := EstimationState_GetByCode(status)
	if est.Notify == core.TRUE {
		_, err := StatusChangeMessage(est.Name, status, rec, usr)
		if err != nil {
			return fieldval, err
		}
	}
	return fieldval, nil
}

func StatusChangeMessage(action string, val string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "StatusChangeMessage", PUT, rec.EstimationSessionID)
	// Get Project from DB
	_, project, _ := Project_GetByID(rec.ProjectID)

	action = Translate("EstimationSessionStatus", action)
	// Send a notification to the Estimation Owner
	MSG_TXT := Translate("Notification", "%s: Estimation ''%s'' has been updated. %s (%s)")
	MSG_TXT = fmt.Sprintf(MSG_TXT, project.OriginID, rec.Name, action, val)
	MSG_BODY := Translate("NotificationBody", "%s: Estimation ''%s'' has been updated to ''%s'' (%s). Please review and take appropriate action. %s")
	// Get the Estimation from the database
	MSG_BODY = fmt.Sprintf(MSG_BODY, project.OriginID, rec.Name, action, val, rec.Comments)
	URI := "/EstimationSessionFormatted/?EstimationSessionID=" + rec.EstimationSessionID
	core.Notification_URL(MSG_TXT, MSG_BODY, URI)

	// Get the Origin
	_, origin, _ := Origin_GetByCode(rec.Origin)

	// Get the Project Manager
	thisPM := project.ProjectManager
	if thisPM == "" {
		thisPM = origin.ProjectManager
	}

	// Send a notification to the Project Manager
	_, pm, _ := Resource_GetByCode(thisPM)
	core.SendEmail(pm.Email, pm.Name, MSG_TXT, MSG_BODY)

	// Send a notification to the Account Manager
	_, am, _ := Resource_GetByCode(origin.AccountManager)
	core.SendEmail(am.Email, am.Name, MSG_TXT, MSG_BODY)

	return "", nil
}

func EstimationSession_Approver_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_Approver_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_Origin_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_Origin_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginStateID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginStateID_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginState_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginState_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginDocTypeID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocTypeID_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginDocType_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocType_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginCode_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginCode_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginName_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginName_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginRate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginRate_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectProfileID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfileID_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectProfile_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfile_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectDefaultReleases_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleases_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectDefaultReleaseHours_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleaseHours_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectBlendedRate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectBlendedRate_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectStateID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStateID_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectState_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectState_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectName_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectName_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectStartDate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStartDate_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectEndDate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectEndDate_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProfileSupportUpliftPerc_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProfileSupportUpliftPerc_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_CCY_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_CCY_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_CCYCode_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_CCYCode_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_EffortTotal_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_EffortTotal_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_FreshDeskURI_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_FreshDeskURI_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ADOURI_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ADOURI_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_NoActiveFeatures_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_NoActiveFeatures_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}

// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func EstimationSession_EstimationStateID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_EstimationStateID_scrn, GET, rec.EstimationSessionID)
	return rec.EstimationStateID
}
func EstimationSession_Approver_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_Approver_scrn, GET, rec.EstimationSessionID)
	return rec.Approver
}
func EstimationSession_Origin_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_Origin_scrn, GET, rec.EstimationSessionID)
	return rec.Origin
}
func EstimationSession_OriginStateID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginStateID_scrn, GET, rec.EstimationSessionID)
	return rec.OriginStateID
}
func EstimationSession_OriginState_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginState_scrn, GET, rec.EstimationSessionID)
	return rec.OriginState
}
func EstimationSession_OriginDocTypeID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocTypeID_scrn, GET, rec.EstimationSessionID)
	return rec.OriginDocTypeID
}
func EstimationSession_OriginDocType_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocType_scrn, GET, rec.EstimationSessionID)
	return rec.OriginDocType
}
func EstimationSession_OriginCode_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginCode_scrn, GET, rec.EstimationSessionID)
	return rec.OriginCode
}
func EstimationSession_OriginName_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginName_scrn, GET, rec.EstimationSessionID)
	return rec.OriginName
}
func EstimationSession_OriginRate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginRate_scrn, GET, rec.EstimationSessionID)
	return rec.OriginRate
}
func EstimationSession_ProjectProfileID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfileID_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectProfileID
}
func EstimationSession_ProjectProfile_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfile_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectProfile
}
func EstimationSession_ProjectDefaultReleases_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleases_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectDefaultReleases
}
func EstimationSession_ProjectDefaultReleaseHours_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleaseHours_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectDefaultReleaseHours
}
func EstimationSession_ProjectBlendedRate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectBlendedRate_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectBlendedRate
}
func EstimationSession_ProjectStateID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStateID_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectStateID
}
func EstimationSession_ProjectState_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectState_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectState
}
func EstimationSession_ProjectName_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectName_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectName
}
func EstimationSession_ProjectStartDate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStartDate_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectStartDate
}
func EstimationSession_ProjectEndDate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectEndDate_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectEndDate
}
func EstimationSession_ProfileSupportUpliftPerc_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProfileSupportUpliftPerc_scrn, GET, rec.EstimationSessionID)
	return rec.ProfileSupportUpliftPerc
}
func EstimationSession_CCY_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_CCY_scrn, GET, rec.EstimationSessionID)
	return rec.CCY
}
func EstimationSession_CCYCode_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_CCYCode_scrn, GET, rec.EstimationSessionID)
	return rec.CCYCode
}
func EstimationSession_EffortTotal_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_EffortTotal_scrn, GET, rec.EstimationSessionID)
	return rec.EffortTotal
}
func EstimationSession_FreshDeskURI_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "FreshDeskURI", GET, rec.EstimationSessionID)
	rtn := core.ApplicationProperties["freshdeskticketuri"]
	rtn = core.ReplaceWildcard(rtn, "ID", rec.FreshdeskID)
	return rtn
}
func EstimationSession_ADOURI_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ADOURI_scrn, GET, rec.EstimationSessionID)
	rtn := core.ApplicationProperties["adoticketuri"]
	rtn = core.ReplaceWildcard(rtn, "ID", rec.AdoID)
	return rtn
}
func EstimationSession_NoActiveFeatures_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_NoActiveFeatures_scrn, GET, rec.EstimationSessionID)
	return rec.NoActiveFeatures
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//
// EstimationSession_EstimationStateID_impl provides validation/actions for EstimationStateID
func EstimationSession_EstimationStateID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_EstimationStateID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_Approver_impl provides validation/actions for Approver
func EstimationSession_Approver_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_Approver_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_Origin_impl provides validation/actions for Origin
func EstimationSession_Origin_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_Origin_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginStateID_impl provides validation/actions for OriginStateID
func EstimationSession_OriginStateID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginStateID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginState_impl provides validation/actions for OriginState
func EstimationSession_OriginState_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginState_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginDocTypeID_impl provides validation/actions for OriginDocTypeID
func EstimationSession_OriginDocTypeID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocTypeID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginDocType_impl provides validation/actions for OriginDocType
func EstimationSession_OriginDocType_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocType_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginCode_impl provides validation/actions for OriginCode
func EstimationSession_OriginCode_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginCode_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginName_impl provides validation/actions for OriginName
func EstimationSession_OriginName_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginRate_impl provides validation/actions for OriginRate
func EstimationSession_OriginRate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginRate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectProfileID_impl provides validation/actions for ProjectProfileID
func EstimationSession_ProjectProfileID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfileID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectProfile_impl provides validation/actions for ProjectProfile
func EstimationSession_ProjectProfile_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfile_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectDefaultReleases_impl provides validation/actions for ProjectDefaultReleases
func EstimationSession_ProjectDefaultReleases_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleases_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectDefaultReleaseHours_impl provides validation/actions for ProjectDefaultReleaseHours
func EstimationSession_ProjectDefaultReleaseHours_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleaseHours_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectBlendedRate_impl provides validation/actions for ProjectBlendedRate
func EstimationSession_ProjectBlendedRate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectBlendedRate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectStateID_impl provides validation/actions for ProjectStateID
func EstimationSession_ProjectStateID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStateID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectState_impl provides validation/actions for ProjectState
func EstimationSession_ProjectState_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectState_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectName_impl provides validation/actions for ProjectName
func EstimationSession_ProjectName_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectStartDate_impl provides validation/actions for ProjectStartDate
func EstimationSession_ProjectStartDate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStartDate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectEndDate_impl provides validation/actions for ProjectEndDate
func EstimationSession_ProjectEndDate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectEndDate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProfileSupportUpliftPerc_impl provides validation/actions for ProfileSupportUpliftPerc
func EstimationSession_ProfileSupportUpliftPerc_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProfileSupportUpliftPerc_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_CCY_impl provides validation/actions for CCY
func EstimationSession_CCY_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_CCY_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_CCYCode_impl provides validation/actions for CCYCode
func EstimationSession_CCYCode_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_CCYCode_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_EffortTotal_impl provides validation/actions for EffortTotal
func EstimationSession_EffortTotal_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_EffortTotal_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_FreshDeskURI_impl provides validation/actions for FreshDeskURI
func EstimationSession_FreshDeskURI_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_FreshDeskURI_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ADOURI_impl provides validation/actions for ADOURI
func EstimationSession_ADOURI_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ADOURI_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_NoActiveFeatures_impl provides validation/actions for NoActiveFeatures
func EstimationSession_NoActiveFeatures_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_NoActiveFeatures_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func EstimationSession_ObjectValidation_impl(iAction string, iId string, iRec dm.EstimationSession) (dm.EstimationSession, string, error) {
	logs.Callout("EstimationSession", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("EstimationSession" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

func EstimationSession_IssueDate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_IssueDate_scrn, PUT, rec.EstimationSessionID)
	issuedState, _ := Data_GetString("Quote", "IssuedState")
	if fieldval == "" && rec.ExpiryDate == "" && rec.EstimationStateID == issuedState {
		// Get Todays Date yyyy-mm-dd
		today := time.Now().Format(core.DATEFORMAT)
		return today, nil
	}

	return fieldval, nil
}

func EstimationSession_IssueDate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_IssueDate_scrn, GET, rec.EstimationSessionID)
	return rec.IssueDate
}

// EstimationSession_IssueDate_impl provides validation/actions for IssueDate
func EstimationSession_IssueDate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_IssueDate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

func EstimationSession_ExpiryDate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ExpiryDate_scrn, PUT, rec.EstimationSessionID)

	issuedState, _ := Data_GetString("Quote", "IssuedState")
	expiredState, _ := Data_GetString("Quote", "ExpiredState")

	switch rec.EstimationStateID {
	case issuedState:
		// Get Expiroty Date 30 days from today
		if fieldval == "" {
			expPeriod, _ := Data_GetInt("Quote", "Expiry")
			logs.Information("Expiry Period: ", strconv.Itoa(expPeriod))
			expiry := time.Now().AddDate(0, 0, expPeriod).Format(core.DATEFORMAT)
			return expiry, nil
		}
	case expiredState:
		logs.Information("Quote is expired", " cannot change Expiry Date")
		return emailRelatedResources(rec)
	default:
	}

	return fieldval, nil
}

func emailRelatedResources(rec dm.EstimationSession) (string, error) {

	_, project, _ := Project_GetByID(rec.ProjectID)
	_, origin, _ := Origin_GetByCode(project.OriginID)

	// Get the email addresses
	var pmEmail dm.Resource

	if rec.ProjectManager == "" {
		// Get the Project Manager from the Project
		_, pmEmail, _ = Resource_GetByCode(project.ProjectManager)
	} else {
		_, pmEmail, _ = Resource_GetByCode(rec.ProjectManager)
	}
	_, pdEmail, _ := Resource_GetByCode(rec.ProductManager)
	_, acEmail, _ := Resource_GetByCode(origin.AccountManager)

	MSG_SUBJECT := "%s Quote Expired - %s"
	MSG_SUBJECT = Translate("QuoteExpiredSubject", MSG_SUBJECT)
	MSG_SUBJECT = fmt.Sprintf(MSG_SUBJECT, project.OriginID, rec.Name, pmEmail.Name, pmEmail.Email)

	MSG_BODY := "Quote for %s %s has expired. Please contact the Project Manager %s (%s) for further information."
	MSG_SUBJECT = Translate("QuoteExpiredBody", MSG_BODY)
	MSG_BODY = fmt.Sprintf(MSG_BODY, project.OriginID, rec.Name, pmEmail.Email)

	// Send the email
	core.SendEmail(pmEmail.Email, pmEmail.Name, MSG_SUBJECT, MSG_BODY)
	core.SendEmail(pdEmail.Email, pdEmail.Name, MSG_SUBJECT, MSG_BODY)
	core.SendEmail(acEmail.Email, acEmail.Name, MSG_SUBJECT, MSG_BODY)

	return "", nil
}

func EstimationSession_ExpiryDate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ExpiryDate_scrn, GET, rec.EstimationSessionID)
	return rec.ExpiryDate
}

// EstimationSession_ExpiryDate_impl provides validation/actions for ExpiryDate
func EstimationSession_ExpiryDate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ExpiryDate_scrn, VAL+"-"+iAction, iId)

	return iValue, fP
}