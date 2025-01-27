package application
// ----------------------------------------------------------------
// Automatically generated  "/application/resource.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Resource (resource)
// Endpoint 	        : Resource (ResourceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:50
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

//Resource_Publish annouces the endpoints available for this object
func Resource_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Resource_Path, Resource_Handler)
	mux.HandleFunc(dm.Resource_PathList, Resource_HandlerList)
	mux.HandleFunc(dm.Resource_PathView, Resource_HandlerView)
	mux.HandleFunc(dm.Resource_PathEdit, Resource_HandlerEdit)
	mux.HandleFunc(dm.Resource_PathNew, Resource_HandlerNew)
	mux.HandleFunc(dm.Resource_PathSave, Resource_HandlerSave)
	mux.HandleFunc(dm.Resource_PathDelete, Resource_HandlerDelete)
    core.API = core.API.AddRoute(dm.Resource_Title, dm.Resource_Path, "", dm.Resource_QueryString, "Application")
	logs.Publish("Application", dm.Resource_Title)
}

//Resource_HandlerList is the handler for the Resource list page
func Resource_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Resource

	objectName := dao.Translate("ObjectName", "Resource")
	reqField := "Base"
	usage := "Defines a filter for the list of Resource records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Resource_GetListFiltered(filter)

	pageDetail := dm.Resource_PageList{
		Title:            CardTitle(dm.Resource_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Resource_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Resource", "List", dm.Resource_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Resource_HandlerView is the handler used to View a Resource database record
func Resource_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Resource_QueryString)
	_, rD, _ := dao.Resource_GetByID(searchID)

	pageDetail := dm.Resource_Page{
		Title:       CardTitle(dm.Resource_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Resource_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = resource_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Resource", "View", dm.Resource_TemplateView)
	nextTemplate = resource_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Resource_HandlerEdit is the handler used to Edit of an existing Resource database record
func Resource_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Resource_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Resource
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Resource)
	} else {
		_, rD, _ = dao.Resource_GetByID(searchID)
	}
	
	pageDetail := dm.Resource_Page{
		Title:       CardTitle(dm.Resource_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Resource_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = resource_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Resource", "Edit", dm.Resource_TemplateEdit)
	nextTemplate = resource_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Resource_HandlerSave is the handler used process the saving of an Resource database record, either new or existing, referenced by Edit & New Handlers.
func Resource_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("ResourceID")
	logs.Servicing(r.URL.Path+itemID)

	item := resource_DataFromRequest(r)
	
	item, errStore := dao.Resource_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Resource", "Save", dm.Resource_Redirect)
		nextTemplate = resource_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Resource_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Resource_QueryString,itemID,item)
	}
}

//Resource_HandlerNew is the handler used process the creation of an new Resource database record, then redirect to Edit
func Resource_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Resource_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Resource
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Resource)
	} else {
		_, _, rD, _ = dao.Resource_New()
	}

	pageDetail := dm.Resource_Page{
		Title:       CardTitle(dm.Resource_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Resource_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = resource_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Resource", "New", dm.Resource_TemplateNew)
	nextTemplate = resource_URIQueryData(nextTemplate,dm.Resource{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//Resource_HandlerDelete is the handler used process the deletion of an Resource database record. May be Hard or SoftDelete.
func Resource_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Resource_QueryString)
	// DAO Call to Delete Resource Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.Resource_Delete(searchID)	

	nextTemplate :=  NextTemplate("Resource", "Delete", dm.Resource_Redirect)
	nextTemplate = resource_URIQueryData(nextTemplate,dm.Resource{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//resource_PopulatePage Builds/Populates the Resource Page from an instance of Resource from the Data Model
func resource_PopulatePage(rD dm.Resource, pageDetail dm.Resource_Page) dm.Resource_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.ResourceID = rD.ResourceID
	pageDetail.Code = rD.Code
	pageDetail.Name = rD.Name
	pageDetail.Email = rD.Email
	pageDetail.Class = rD.Class
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.UserActive = rD.UserActive
	pageDetail.SYSActivity = rD.SYSActivity
	pageDetail.Notes = rD.Notes
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.Comments = rD.Comments
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.Class_lookup = dao.StubLists_Get("resourcetype")
	pageDetail.UserActive_lookup = dao.StubLists_Get("tf")
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.ResourceID_props = rD.ResourceID_props
	pageDetail.Code_props = rD.Code_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Email_props = rD.Email_props
	pageDetail.Class_props = rD.Class_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.UserActive_props = rD.UserActive_props
	pageDetail.SYSActivity_props = rD.SYSActivity_props
	pageDetail.Notes_props = rD.Notes_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.Comments_props = rD.Comments_props
	return pageDetail
}
//resource_DataFromRequest is used process the content of an HTTP Request and return an instance of an Resource
func resource_DataFromRequest(r *http.Request) dm.Resource {
	var item dm.Resource
	item.SYSId = r.FormValue(dm.Resource_SYSId_scrn)
	item.ResourceID = r.FormValue(dm.Resource_ResourceID_scrn)
	item.Code = r.FormValue(dm.Resource_Code_scrn)
	item.Name = r.FormValue(dm.Resource_Name_scrn)
	item.Email = r.FormValue(dm.Resource_Email_scrn)
	item.Class = r.FormValue(dm.Resource_Class_scrn)
	item.SYSCreated = r.FormValue(dm.Resource_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.Resource_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.Resource_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.Resource_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.Resource_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.Resource_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.Resource_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.Resource_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.Resource_SYSDeletedHost_scrn)
	item.UserActive = r.FormValue(dm.Resource_UserActive_scrn)
	item.SYSActivity = r.FormValue(dm.Resource_SYSActivity_scrn)
	item.Notes = r.FormValue(dm.Resource_Notes_scrn)
	item.SYSDbVersion = r.FormValue(dm.Resource_SYSDbVersion_scrn)
	item.Comments = r.FormValue(dm.Resource_Comments_scrn)
	return item
}
//resource_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Resource Data Model
func resource_URIQueryData(queryPath string,item dm.Resource,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Resource_ResourceID_scrn), item.ResourceID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Resource_Code_scrn), item.Code)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Resource_Name_scrn), item.Name)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Resource_Email_scrn), item.Email)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Resource_Class_scrn), item.Class)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Resource_UserActive_scrn), item.UserActive)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Resource_Notes_scrn), item.Notes)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Resource_Comments_scrn), item.Comments)
	return queryPath
}