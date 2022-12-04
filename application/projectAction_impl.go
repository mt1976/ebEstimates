package application

// ----------------------------------------------------------------
// Automatically generated  "/application/projectaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ProjectAction (projectaction)
// Endpoint 	        : ProjectAction (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 30/11/2022 at 13:10:09
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ProjectAction_Publish annouces the endpoints available for this object
func ProjectAction_Publish_Impl(mux http.ServeMux) {
	//No API
	//Cannot List via GUI
	//Cannot View via GUI
	//Cannot Edit via GUI
	mux.HandleFunc(dm.ProjectAction_PathNew, ProjectAction_HandlerNew_Impl)
	mux.HandleFunc(dm.ProjectAction_PathSave, ProjectAction_HandlerSave_Impl)
	//Cannot Delete via GUI
	logs.Publish("Implementation", dm.ProjectAction_Title)
	//No API
}

// ProjectAction_HandlerSave is the handler used process the saving of an ProjectAction
func ProjectAction_HandlerSave_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("ProjectID"))

	var item dm.Project
	// START
	// Dynamically generated 30/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	item.SYSId = r.FormValue(dm.ProjectAction_SYSId_scrn)
	item.ProjectID = r.FormValue(dm.ProjectAction_ProjectID_scrn)
	item.OriginID = r.FormValue(dm.ProjectAction_OriginID_scrn)
	item.ProjectStateID = r.FormValue(dm.ProjectAction_ProjectStateID_scrn)
	item.ProfileID = r.FormValue(dm.ProjectAction_ProfileID_scrn)
	item.Name = r.FormValue(dm.ProjectAction_Name_scrn)
	item.Description = r.FormValue(dm.ProjectAction_Description_scrn)
	item.StartDate = r.FormValue(dm.ProjectAction_StartDate_scrn)
	item.EndDate = r.FormValue(dm.ProjectAction_EndDate_scrn)
	item.SYSCreated = r.FormValue(dm.ProjectAction_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.ProjectAction_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.ProjectAction_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.ProjectAction_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.ProjectAction_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.ProjectAction_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.ProjectAction_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.ProjectAction_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.ProjectAction_SYSDeletedHost_scrn)

	//
	// Dynamically generated 30/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	dao.Project_Store(item, r)
	// Get Origin by Code
	_, origin, _ := dao.Origin_GetByCode(item.OriginID)

	REDR := dm.Project_Origin_PathList + "?" + dm.Origin_QueryString + "=" + origin.OriginID
	logs.Information("REDIRECTING TO: ", REDR)
	http.Redirect(w, r, REDR, http.StatusFound)
}

// ProjectAction_HandlerNew is the handler used process the creation of an ProjectAction
func ProjectAction_HandlerNew_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	_, _, rD, _ := dao.ProjectAction_New()

	pageDetail := dm.ProjectAction_Page{
		Title:     CardTitle(dm.ProjectAction_Title, core.Action_New),
		PageTitle: PageTitle(dm.ProjectAction_Title, core.Action_New),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	originID := core.GetURLparam(r, dm.Origin_QueryString)
	if originID == "" {
		ExecuteTemplate("html/Impl_Oops.html", w, r, pageDetail)
		return
	}

	pageDetail = projectaction_PopulatePage(rD, pageDetail)

	pageDetail.OriginID = originID

	spew.Dump(pageDetail)

	ExecuteTemplate(dm.ProjectAction_TemplateNew, w, r, pageDetail)

}

// Builds/Popuplates the ProjectAction Page
func projectaction_PopulatePage(rD dm.ProjectAction, pageDetail dm.ProjectAction_Page) dm.ProjectAction_Page {
	// START
	// Dynamically generated 30/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.ProjectID = rD.ProjectID
	pageDetail.OriginID = rD.OriginID
	pageDetail.ProjectStateID = rD.ProjectStateID
	pageDetail.ProfileID = rD.ProfileID
	pageDetail.Name = rD.Name
	pageDetail.Description = rD.Description
	pageDetail.StartDate = rD.StartDate
	pageDetail.EndDate = rD.EndDate
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost

	//
	// Automatically generated 30/11/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//

	pageDetail.OriginID_lookup = dao.Origin_GetLookup()

	pageDetail.ProjectStateID_lookup = dao.ProjectState_GetLookup()

	pageDetail.ProfileID_lookup = dao.Profile_GetLookup()

	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.ProjectID_props = rD.ProjectID_props
	pageDetail.OriginID_props = rD.OriginID_props
	pageDetail.ProjectStateID_props = rD.ProjectStateID_props
	pageDetail.ProfileID_props = rD.ProfileID_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.StartDate_props = rD.StartDate_props
	pageDetail.EndDate_props = rD.EndDate_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props

	//
	// Dynamically generated 30/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
	return pageDetail
}