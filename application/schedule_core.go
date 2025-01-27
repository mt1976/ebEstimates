package application
// ----------------------------------------------------------------
// Automatically generated  "/application/schedule.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
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

//Schedule_Publish annouces the endpoints available for this object
func Schedule_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Schedule_Path, Schedule_Handler)
	mux.HandleFunc(dm.Schedule_PathList, Schedule_HandlerList)
	mux.HandleFunc(dm.Schedule_PathView, Schedule_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.Schedule_PathSave, Schedule_HandlerSave)
	//Cannot Delete via GUI
    core.API = core.API.AddRoute(dm.Schedule_Title, dm.Schedule_Path, "", dm.Schedule_QueryString, "Application")
	logs.Publish("Application", dm.Schedule_Title)
}

//Schedule_HandlerList is the handler for the Schedule list page
func Schedule_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Schedule

	objectName := dao.Translate("ObjectName", "Schedule")
	reqField := "Base"
	usage := "Defines a filter for the list of Schedule records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Schedule_GetListFiltered(filter)

	pageDetail := dm.Schedule_PageList{
		Title:            CardTitle(dm.Schedule_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Schedule_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Schedule", "List", dm.Schedule_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Schedule_HandlerView is the handler used to View a Schedule database record
func Schedule_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Schedule_QueryString)
	_, rD, _ := dao.Schedule_GetByID(searchID)

	pageDetail := dm.Schedule_Page{
		Title:       CardTitle(dm.Schedule_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Schedule_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = schedule_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Schedule", "View", dm.Schedule_TemplateView)
	nextTemplate = schedule_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Schedule_HandlerSave is the handler used process the saving of an Schedule database record, either new or existing, referenced by Edit & New Handlers.
func Schedule_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("Id")
	logs.Servicing(r.URL.Path+itemID)

	item := schedule_DataFromRequest(r)
	
	item, errStore := dao.Schedule_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Schedule", "Save", dm.Schedule_Redirect)
		nextTemplate = schedule_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Schedule_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Schedule_QueryString,itemID,item)
	}
}
//schedule_PopulatePage Builds/Populates the Schedule Page from an instance of Schedule from the Data Model
func schedule_PopulatePage(rD dm.Schedule, pageDetail dm.Schedule_Page) dm.Schedule_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Name = rD.Name
	pageDetail.Description = rD.Description
	pageDetail.Schedule = rD.Schedule
	pageDetail.Started = rD.Started
	pageDetail.Lastrun = rD.Lastrun
	pageDetail.Message = rD.Message
	pageDetail.Type = rD.Type
	pageDetail.Human = rD.Human
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.Schedule_props = rD.Schedule_props
	pageDetail.Started_props = rD.Started_props
	pageDetail.Lastrun_props = rD.Lastrun_props
	pageDetail.Message_props = rD.Message_props
	pageDetail.Type_props = rD.Type_props
	pageDetail.Human_props = rD.Human_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	return pageDetail
}
//schedule_DataFromRequest is used process the content of an HTTP Request and return an instance of an Schedule
func schedule_DataFromRequest(r *http.Request) dm.Schedule {
	var item dm.Schedule
	item.SYSId = r.FormValue(dm.Schedule_SYSId_scrn)
	item.Id = r.FormValue(dm.Schedule_Id_scrn)
	item.Name = r.FormValue(dm.Schedule_Name_scrn)
	item.Description = r.FormValue(dm.Schedule_Description_scrn)
	item.Schedule = r.FormValue(dm.Schedule_Schedule_scrn)
	item.Started = r.FormValue(dm.Schedule_Started_scrn)
	item.Lastrun = r.FormValue(dm.Schedule_Lastrun_scrn)
	item.Message = r.FormValue(dm.Schedule_Message_scrn)
	item.Type = r.FormValue(dm.Schedule_Type_scrn)
	item.Human = r.FormValue(dm.Schedule_Human_scrn)
	item.SYSCreated = r.FormValue(dm.Schedule_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.Schedule_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.Schedule_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.Schedule_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.Schedule_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.Schedule_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.Schedule_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.Schedule_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.Schedule_SYSDeletedHost_scrn)
	item.SYSDbVersion = r.FormValue(dm.Schedule_SYSDbVersion_scrn)
	return item
}
//schedule_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Schedule Data Model
func schedule_URIQueryData(queryPath string,item dm.Schedule,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Schedule_Id_scrn), item.Id)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Schedule_Name_scrn), item.Name)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Schedule_Description_scrn), item.Description)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Schedule_Schedule_scrn), item.Schedule)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Schedule_Started_scrn), item.Started)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Schedule_Lastrun_scrn), item.Lastrun)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Schedule_Message_scrn), item.Message)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Schedule_Type_scrn), item.Type)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Schedule_Human_scrn), item.Human)
	return queryPath
}