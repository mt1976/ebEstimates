package application
// ----------------------------------------------------------------
// Automatically generated  "/application/index.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Index (index)
// Endpoint 	        : Index (IndexID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 24/01/2023 at 13:18:09
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//Index_Publish annouces the endpoints available for this object
//Index_Publish - Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Index_Publish(mux http.ServeMux) {
	// START
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.Index_Path, Index_Handler)
	mux.HandleFunc(dm.Index_PathList, Index_HandlerList)
	mux.HandleFunc(dm.Index_PathView, Index_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.Index_PathSave, Index_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Index_Title)
    core.Catalog_Add(dm.Index_Title, dm.Index_Path, "", dm.Index_QueryString, "Application")
	// 
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Index_HandlerList is the handler for the list page
//Allows Listing of Index records
//Index_HandlerList - Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Index_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Index
	noItems, returnList, _ := dao.Index_GetList()

	pageDetail := dm.Index_PageList{
		Title:            CardTitle(dm.Index_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Index_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Index_TemplateList, w, r, pageDetail)
	// 
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//Index_HandlerView is the handler used to View a page
//Allows Viewing for an existing Index record
//Index_HandlerView - Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Index_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Index_QueryString)
	_, rD, _ := dao.Index_GetByID(searchID)

	pageDetail := dm.Index_Page{
		Title:       CardTitle(dm.Index_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Index_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = index_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Index_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//Index_HandlerSave is the handler used process the saving of an Index
//It is called from the Edit and New pages
//Index_HandlerSave  - Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Index_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("IndexID"))

	item := index_DataFromRequest(r)
	
	dao.Index_Store(item,r)	
	http.Redirect(w, r, dm.Index_Redirect, http.StatusFound)
	// 
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}




//index_PopulatePage Builds/Populates the Index Page 
func index_PopulatePage(rD dm.Index, pageDetail dm.Index_Page) dm.Index_Page {
	// START
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.IndexID = rD.IndexID
	pageDetail.KeyClass = rD.KeyClass
	pageDetail.KeyName = rD.KeyName
	pageDetail.KeyID = rD.KeyID
	pageDetail.Link = rD.Link
	pageDetail.LinkView = rD.LinkView
	pageDetail.LinkEdit = rD.LinkEdit
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
	pageDetail.KeyValue = rD.KeyValue
	
	
	//
	// Automatically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.IndexID_props = rD.IndexID_props
	pageDetail.KeyClass_props = rD.KeyClass_props
	pageDetail.KeyName_props = rD.KeyName_props
	pageDetail.KeyID_props = rD.KeyID_props
	pageDetail.Link_props = rD.Link_props
	pageDetail.LinkView_props = rD.LinkView_props
	pageDetail.LinkEdit_props = rD.LinkEdit_props
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
	pageDetail.KeyValue_props = rD.KeyValue_props
	
	// 
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	


//index_DataFromRequest is used process the content of an HTTP Request and return an instance of an Index
func index_DataFromRequest(r *http.Request) dm.Index {
	// START
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	var item dm.Index
	// FIELD SET START
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Index_SYSId_scrn)
		item.IndexID = r.FormValue(dm.Index_IndexID_scrn)
		item.KeyClass = r.FormValue(dm.Index_KeyClass_scrn)
		item.KeyName = r.FormValue(dm.Index_KeyName_scrn)
		item.KeyID = r.FormValue(dm.Index_KeyID_scrn)
		item.Link = r.FormValue(dm.Index_Link_scrn)
		item.LinkView = r.FormValue(dm.Index_LinkView_scrn)
		item.LinkEdit = r.FormValue(dm.Index_LinkEdit_scrn)
		item.SYSCreated = r.FormValue(dm.Index_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Index_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Index_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Index_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Index_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Index_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Index_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Index_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Index_SYSDeletedHost_scrn)
		item.SYSDbVersion = r.FormValue(dm.Index_SYSDbVersion_scrn)
		item.KeyValue = r.FormValue(dm.Index_KeyValue_scrn)
	
	// 
	// Auto generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return item
}