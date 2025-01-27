package application
// ----------------------------------------------------------------
// Automatically generated  "/application/profile.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Profile (profile)
// Endpoint 	        : Profile (ProfileID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:49
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"
	"strings"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//Profile_Publish annouces the endpoints available for this object
func Profile_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Profile_Path, Profile_Handler)
	mux.HandleFunc(dm.Profile_PathList, Profile_HandlerList)
	mux.HandleFunc(dm.Profile_PathView, Profile_HandlerView)
	mux.HandleFunc(dm.Profile_PathEdit, Profile_HandlerEdit)
	mux.HandleFunc(dm.Profile_PathNew, Profile_HandlerNew)
	mux.HandleFunc(dm.Profile_PathSave, Profile_HandlerSave)
	mux.HandleFunc(dm.Profile_PathDelete, Profile_HandlerDelete)
    core.API = core.API.AddRoute(dm.Profile_Title, dm.Profile_Path, "", dm.Profile_QueryString, "Application")
	logs.Publish("Application", dm.Profile_Title)
}

//Profile_HandlerList is the handler for the Profile list page
func Profile_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Profile

	objectName := dao.Translate("ObjectName", "Profile")
	reqField := "Base"
	usage := "Defines a filter for the list of Profile records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Profile_GetListFiltered(filter)

	pageDetail := dm.Profile_PageList{
		Title:            CardTitle(dm.Profile_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Profile_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Profile", "List", dm.Profile_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Profile_HandlerView is the handler used to View a Profile database record
func Profile_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Profile_QueryString)
	_, rD, _ := dao.Profile_GetByID(searchID)

	pageDetail := dm.Profile_Page{
		Title:       CardTitle(dm.Profile_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Profile_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = profile_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Profile", "View", dm.Profile_TemplateView)
	nextTemplate = profile_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Profile_HandlerEdit is the handler used to Edit of an existing Profile database record
func Profile_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Profile_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Profile
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Profile)
	} else {
		_, rD, _ = dao.Profile_GetByID(searchID)
	}
	
	pageDetail := dm.Profile_Page{
		Title:       CardTitle(dm.Profile_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Profile_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = profile_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Profile", "Edit", dm.Profile_TemplateEdit)
	nextTemplate = profile_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Profile_HandlerSave is the handler used process the saving of an Profile database record, either new or existing, referenced by Edit & New Handlers.
func Profile_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("ProfileID")
	logs.Servicing(r.URL.Path+itemID)

	item := profile_DataFromRequest(r)
	
	item, errStore := dao.Profile_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Profile", "Save", dm.Profile_Redirect)
		nextTemplate = profile_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Profile_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Profile_QueryString,itemID,item)
	}
}

//Profile_HandlerNew is the handler used process the creation of an new Profile database record, then redirect to Edit
func Profile_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Profile_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Profile
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Profile)
	} else {
		_, _, rD, _ = dao.Profile_New()
	}

	pageDetail := dm.Profile_Page{
		Title:       CardTitle(dm.Profile_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Profile_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = profile_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Profile", "New", dm.Profile_TemplateNew)
	nextTemplate = profile_URIQueryData(nextTemplate,dm.Profile{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//Profile_HandlerDelete is the handler used process the deletion of an Profile database record. May be Hard or SoftDelete.
func Profile_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Profile_QueryString)
	// DAO Call to Delete Profile Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.Profile_Delete(searchID)	

	nextTemplate :=  NextTemplate("Profile", "Delete", dm.Profile_Redirect)
	nextTemplate = profile_URIQueryData(nextTemplate,dm.Profile{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//profile_PopulatePage Builds/Populates the Profile Page from an instance of Profile from the Data Model
func profile_PopulatePage(rD dm.Profile, pageDetail dm.Profile_Page) dm.Profile_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.ProfileID = rD.ProfileID
	pageDetail.Code = rD.Code
	pageDetail.Name = rD.Name
	pageDetail.StartDate = rD.StartDate
	pageDetail.EndDate = rD.EndDate
	pageDetail.DefaultReleases = rD.DefaultReleases
	pageDetail.DefaultReleaseHours = rD.DefaultReleaseHours
	pageDetail.BlendedRate = rD.BlendedRate
	pageDetail.Rounding = rD.Rounding
	pageDetail.HoursPerDay = rD.HoursPerDay
	pageDetail.REQPerc = rD.REQPerc
	pageDetail.ANAPerc = rD.ANAPerc
	pageDetail.DOCPerc = rD.DOCPerc
	pageDetail.PMPerc = rD.PMPerc
	pageDetail.UATPerc = rD.UATPerc
	pageDetail.GTMPerc = rD.GTMPerc
	pageDetail.SupportUplift = rD.SupportUplift
	pageDetail.ContigencyPerc = rD.ContigencyPerc
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSActivity = rD.SYSActivity
	pageDetail.Notes = rD.Notes
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.Comments = rD.Comments
	pageDetail.TrainingPerc = rD.TrainingPerc
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.ProfileID_props = rD.ProfileID_props
	pageDetail.Code_props = rD.Code_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.StartDate_props = rD.StartDate_props
	pageDetail.EndDate_props = rD.EndDate_props
	pageDetail.DefaultReleases_props = rD.DefaultReleases_props
	pageDetail.DefaultReleaseHours_props = rD.DefaultReleaseHours_props
	pageDetail.BlendedRate_props = rD.BlendedRate_props
	pageDetail.Rounding_props = rD.Rounding_props
	pageDetail.HoursPerDay_props = rD.HoursPerDay_props
	pageDetail.REQPerc_props = rD.REQPerc_props
	pageDetail.ANAPerc_props = rD.ANAPerc_props
	pageDetail.DOCPerc_props = rD.DOCPerc_props
	pageDetail.PMPerc_props = rD.PMPerc_props
	pageDetail.UATPerc_props = rD.UATPerc_props
	pageDetail.GTMPerc_props = rD.GTMPerc_props
	pageDetail.SupportUplift_props = rD.SupportUplift_props
	pageDetail.ContigencyPerc_props = rD.ContigencyPerc_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSActivity_props = rD.SYSActivity_props
	pageDetail.Notes_props = rD.Notes_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.Comments_props = rD.Comments_props
	pageDetail.TrainingPerc_props = rD.TrainingPerc_props
	return pageDetail
}
//profile_DataFromRequest is used process the content of an HTTP Request and return an instance of an Profile
func profile_DataFromRequest(r *http.Request) dm.Profile {
	var item dm.Profile
	item.SYSId = r.FormValue(dm.Profile_SYSId_scrn)
	item.ProfileID = r.FormValue(dm.Profile_ProfileID_scrn)
	item.Code = r.FormValue(dm.Profile_Code_scrn)
	item.Name = r.FormValue(dm.Profile_Name_scrn)
	item.StartDate = r.FormValue(dm.Profile_StartDate_scrn)
	item.EndDate = r.FormValue(dm.Profile_EndDate_scrn)
	item.DefaultReleases = r.FormValue(dm.Profile_DefaultReleases_scrn)
	item.DefaultReleaseHours = r.FormValue(dm.Profile_DefaultReleaseHours_scrn)
	item.BlendedRate = r.FormValue(dm.Profile_BlendedRate_scrn)
	item.Rounding = r.FormValue(dm.Profile_Rounding_scrn)
	item.HoursPerDay = r.FormValue(dm.Profile_HoursPerDay_scrn)
	item.REQPerc = r.FormValue(dm.Profile_REQPerc_scrn)
	item.ANAPerc = r.FormValue(dm.Profile_ANAPerc_scrn)
	item.DOCPerc = r.FormValue(dm.Profile_DOCPerc_scrn)
	item.PMPerc = r.FormValue(dm.Profile_PMPerc_scrn)
	item.UATPerc = r.FormValue(dm.Profile_UATPerc_scrn)
	item.GTMPerc = r.FormValue(dm.Profile_GTMPerc_scrn)
	item.SupportUplift = r.FormValue(dm.Profile_SupportUplift_scrn)
	item.ContigencyPerc = r.FormValue(dm.Profile_ContigencyPerc_scrn)
	item.SYSCreated = r.FormValue(dm.Profile_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.Profile_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.Profile_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.Profile_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.Profile_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.Profile_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.Profile_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.Profile_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.Profile_SYSDeletedHost_scrn)
	item.SYSActivity = r.FormValue(dm.Profile_SYSActivity_scrn)
	item.Notes = r.FormValue(dm.Profile_Notes_scrn)
	item.SYSDbVersion = r.FormValue(dm.Profile_SYSDbVersion_scrn)
	item.Comments = r.FormValue(dm.Profile_Comments_scrn)
	item.TrainingPerc = r.FormValue(dm.Profile_TrainingPerc_scrn)
	return item
}
//profile_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Profile Data Model
func profile_URIQueryData(queryPath string,item dm.Profile,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_ProfileID_scrn), item.ProfileID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_Code_scrn), item.Code)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_Name_scrn), item.Name)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_StartDate_scrn), item.StartDate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_EndDate_scrn), item.EndDate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_DefaultReleases_scrn), item.DefaultReleases)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_DefaultReleaseHours_scrn), item.DefaultReleaseHours)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_BlendedRate_scrn), item.BlendedRate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_Rounding_scrn), item.Rounding)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_HoursPerDay_scrn), item.HoursPerDay)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_REQPerc_scrn), item.REQPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_ANAPerc_scrn), item.ANAPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_DOCPerc_scrn), item.DOCPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_PMPerc_scrn), item.PMPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_UATPerc_scrn), item.UATPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_GTMPerc_scrn), item.GTMPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_SupportUplift_scrn), item.SupportUplift)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_ContigencyPerc_scrn), item.ContigencyPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_Notes_scrn), item.Notes)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_Comments_scrn), item.Comments)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Profile_TrainingPerc_scrn), item.TrainingPerc)
	return queryPath
}