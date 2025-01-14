package application

// ----------------------------------------------------------------
// Automatically generated  "/application/origin.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 01/12/2022 at 09:40:01
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// Origin_Publish annouces the endpoints available for this object
func Origin_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc(dm.Origin_PathRemove, Origin_HandlerRemove)
	mux.HandleFunc(dm.Origin_PathStore, Origin_HandlerStore)

	//Cannot Delete via GUI
	logs.Publish("Implmentation", dm.Origin_Title)
}

func Origin_HandlerRemove(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	originID := core.GetURLparam(r, dm.Origin_QueryString)
	_, originREC, _ := dao.Origin_GetByID(originID)

	dao.Origin_SoftDelete(originID)

	err := Project_SoftDeleteByOriginCode(originREC.Code)
	if err != nil {
		logs.Panic("Cannot SoftDelete Projects {"+originID+"}", err)
	}

	REDR := "/home"
	http.Redirect(w, r, REDR, http.StatusFound)
}

// Origin_HandlerStore is the handler used process the saving of an Origin
// It is called from the Edit and New pages
// Origin_HandlerStore  - Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Origin_HandlerStore(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("OriginID"))
	//spew.Dump(r.Form)

	item := origin_DataFromRequest(r)
	//spew.Dump(item)

	item, errStore := dao.Origin_Store(item, r)
	if errStore == nil {
		nextTemplate := NextTemplate("Origin", "Save", dm.Origin_Redirect)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Origin_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r, dm.Origin_QueryString, item.OriginID, item)
	}

	//REDR := "/home"
	//http.Redirect(w, r, REDR, http.StatusFound)
	//
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}
