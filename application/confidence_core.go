package application
// ----------------------------------------------------------------
// Automatically generated  "/application/confidence.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Confidence (confidence)
// Endpoint 	        : Confidence (ConfidenceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:47
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

//Confidence_Publish annouces the endpoints available for this object
func Confidence_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Confidence_Path, Confidence_Handler)
	mux.HandleFunc(dm.Confidence_PathList, Confidence_HandlerList)
	mux.HandleFunc(dm.Confidence_PathView, Confidence_HandlerView)
	mux.HandleFunc(dm.Confidence_PathEdit, Confidence_HandlerEdit)
	mux.HandleFunc(dm.Confidence_PathNew, Confidence_HandlerNew)
	mux.HandleFunc(dm.Confidence_PathSave, Confidence_HandlerSave)
	mux.HandleFunc(dm.Confidence_PathDelete, Confidence_HandlerDelete)
    core.API = core.API.AddRoute(dm.Confidence_Title, dm.Confidence_Path, "", dm.Confidence_QueryString, "Application")
	logs.Publish("Application", dm.Confidence_Title)
}

//Confidence_HandlerList is the handler for the Confidence list page
func Confidence_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Confidence

	objectName := dao.Translate("ObjectName", "Confidence")
	reqField := "Base"
	usage := "Defines a filter for the list of Confidence records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Confidence_GetListFiltered(filter)

	pageDetail := dm.Confidence_PageList{
		Title:            CardTitle(dm.Confidence_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Confidence_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Confidence", "List", dm.Confidence_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Confidence_HandlerView is the handler used to View a Confidence database record
func Confidence_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Confidence_QueryString)
	_, rD, _ := dao.Confidence_GetByID(searchID)

	pageDetail := dm.Confidence_Page{
		Title:       CardTitle(dm.Confidence_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Confidence_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = confidence_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Confidence", "View", dm.Confidence_TemplateView)
	nextTemplate = confidence_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Confidence_HandlerEdit is the handler used to Edit of an existing Confidence database record
func Confidence_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Confidence_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Confidence
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Confidence)
	} else {
		_, rD, _ = dao.Confidence_GetByID(searchID)
	}
	
	pageDetail := dm.Confidence_Page{
		Title:       CardTitle(dm.Confidence_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Confidence_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = confidence_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Confidence", "Edit", dm.Confidence_TemplateEdit)
	nextTemplate = confidence_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Confidence_HandlerSave is the handler used process the saving of an Confidence database record, either new or existing, referenced by Edit & New Handlers.
func Confidence_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("ConfidenceID")
	logs.Servicing(r.URL.Path+itemID)

	item := confidence_DataFromRequest(r)
	
	item, errStore := dao.Confidence_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Confidence", "Save", dm.Confidence_Redirect)
		nextTemplate = confidence_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Confidence_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Confidence_QueryString,itemID,item)
	}
}

//Confidence_HandlerNew is the handler used process the creation of an new Confidence database record, then redirect to Edit
func Confidence_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Confidence_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Confidence
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Confidence)
	} else {
		_, _, rD, _ = dao.Confidence_New()
	}

	pageDetail := dm.Confidence_Page{
		Title:       CardTitle(dm.Confidence_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Confidence_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = confidence_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Confidence", "New", dm.Confidence_TemplateNew)
	nextTemplate = confidence_URIQueryData(nextTemplate,dm.Confidence{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//Confidence_HandlerDelete is the handler used process the deletion of an Confidence database record. May be Hard or SoftDelete.
func Confidence_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Confidence_QueryString)
	// DAO Call to Delete Confidence Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.Confidence_Delete(searchID)	

	nextTemplate :=  NextTemplate("Confidence", "Delete", dm.Confidence_Redirect)
	nextTemplate = confidence_URIQueryData(nextTemplate,dm.Confidence{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//confidence_PopulatePage Builds/Populates the Confidence Page from an instance of Confidence from the Data Model
func confidence_PopulatePage(rD dm.Confidence, pageDetail dm.Confidence_Page) dm.Confidence_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.ConfidenceID = rD.ConfidenceID
	pageDetail.Code = rD.Code
	pageDetail.Name = rD.Name
	pageDetail.Perc = rD.Perc
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.Notes = rD.Notes
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.Comments = rD.Comments
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.ConfidenceID_props = rD.ConfidenceID_props
	pageDetail.Code_props = rD.Code_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Perc_props = rD.Perc_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.Notes_props = rD.Notes_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.Comments_props = rD.Comments_props
	return pageDetail
}
//confidence_DataFromRequest is used process the content of an HTTP Request and return an instance of an Confidence
func confidence_DataFromRequest(r *http.Request) dm.Confidence {
	var item dm.Confidence
	item.SYSId = r.FormValue(dm.Confidence_SYSId_scrn)
	item.ConfidenceID = r.FormValue(dm.Confidence_ConfidenceID_scrn)
	item.Code = r.FormValue(dm.Confidence_Code_scrn)
	item.Name = r.FormValue(dm.Confidence_Name_scrn)
	item.Perc = r.FormValue(dm.Confidence_Perc_scrn)
	item.SYSCreated = r.FormValue(dm.Confidence_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.Confidence_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.Confidence_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.Confidence_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.Confidence_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.Confidence_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.Confidence_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.Confidence_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.Confidence_SYSDeletedHost_scrn)
	item.Notes = r.FormValue(dm.Confidence_Notes_scrn)
	item.SYSDbVersion = r.FormValue(dm.Confidence_SYSDbVersion_scrn)
	item.Comments = r.FormValue(dm.Confidence_Comments_scrn)
	return item
}
//confidence_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Confidence Data Model
func confidence_URIQueryData(queryPath string,item dm.Confidence,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Confidence_ConfidenceID_scrn), item.ConfidenceID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Confidence_Code_scrn), item.Code)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Confidence_Name_scrn), item.Name)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Confidence_Perc_scrn), item.Perc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Confidence_Notes_scrn), item.Notes)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Confidence_Comments_scrn), item.Comments)
	return queryPath
}