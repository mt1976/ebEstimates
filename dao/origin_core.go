package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/origin.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:49
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"
	"errors"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/google/uuid"
	das  "github.com/mt1976/ebEstimates/das"
	dm   "github.com/mt1976/ebEstimates/datamodel"
	logs   "github.com/mt1976/ebEstimates/logs"
)

var Origin_SQLbase string
var Origin_QualifiedName string
func init(){
	Origin_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Origin_SQLTable)
	Origin_SQLbase =  das.SELECTALL + das.FROM + Origin_QualifiedName
}

// Origin_GetList() returns a list of all Origin records
func Origin_GetList() (int, []dm.Origin, error) {
	count, originList, err := Origin_GetListFiltered("")
	return count, originList, err
}

// Origin_GetListFiltered() returns a filtered list of all Origin records
func Origin_GetListFiltered(filter string) (int, []dm.Origin, error) {
	
	tsql := Origin_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
	count, originList, _, _ := origin_Fetch(tsql)
	
	return count, originList, nil
}


// Origin_GetLookup() returns a lookup list of all Origin items in lookup format
func Origin_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, originList, _ := Origin_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: originList[i].Code, Name: originList[i].FullName})
	}
	return returnList
}

// Origin_GetFilteredLookup() returns a lookup list of all Origin items in lookup format
func Origin_GetFilteredLookup(requestObject string,requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	objectName := Translate("ObjectName", requestObject)
	reqField := requestField+"_Origin_Filter"
	
	usage := "Defines a filter for a lookup list of Origin records, when requested by "+requestField+"." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"

	filter,_ := Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("Origin_GetFilteredLookup() - No filter found : " + reqField + " for Object: " + objectName)
	} 
	count, originList, _ := Origin_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: originList[i].Code, Name: originList[i].FullName})
	}
	return returnList
}



// Origin_GetByID() returns a single Origin record
func Origin_GetByID(id string) (int, dm.Origin, error) {


	tsql := Origin_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Origin_SQLSearchID + das.EQ + das.ID(id)
	_, _, originItem, _ := origin_Fetch(tsql)


	originItem = Origin_PostGet(originItem,id)

	return 1, originItem, nil
}

func Origin_PostGet(originItem dm.Origin,id string) dm.Origin {
	originItem.StateID,originItem.StateID_props = Origin_StateID_validate_impl (GET,id,originItem.StateID,originItem,originItem.StateID_props)
	originItem.DocTypeID,originItem.DocTypeID_props = Origin_DocTypeID_validate_impl (GET,id,originItem.DocTypeID,originItem,originItem.DocTypeID_props)
	originItem.Code,originItem.Code_props = Origin_Code_validate_impl (GET,id,originItem.Code,originItem,originItem.Code_props)
	originItem.FullName,originItem.FullName_props = Origin_FullName_validate_impl (GET,id,originItem.FullName,originItem,originItem.FullName_props)
	originItem.Rate,originItem.Rate_props = Origin_Rate_validate_impl (GET,id,originItem.Rate,originItem,originItem.Rate_props)
	originItem.StartDate,originItem.StartDate_props = Origin_StartDate_validate_impl (GET,id,originItem.StartDate,originItem,originItem.StartDate_props)
	originItem.Currency,originItem.Currency_props = Origin_Currency_validate_impl (GET,id,originItem.Currency,originItem,originItem.Currency_props)
	originItem.NoActiveProjects,originItem.NoActiveProjects_props = Origin_NoActiveProjects_validate_impl (GET,id,originItem.NoActiveProjects,originItem,originItem.NoActiveProjects_props)
	originItem.RateOnLoad,originItem.RateOnLoad_props = Origin_RateOnLoad_validate_impl (GET,id,originItem.RateOnLoad,originItem,originItem.RateOnLoad_props)
	originItem.StatusOnLoad,originItem.StatusOnLoad_props = Origin_StatusOnLoad_validate_impl (GET,id,originItem.StatusOnLoad,originItem,originItem.StatusOnLoad_props)

	return originItem
}




// Origin_DeleteByID() deletes a single Origin record
func Origin_Delete(id string) {
		Origin_SoftDelete(id)
	}
// Origin_SoftDelete(id string) soft deletes a single Origin record
func Origin_SoftDelete(id string) {
	//Uses Soft Delete
		_,originItem,_ := Origin_GetByID(id)
		originItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		originItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		originItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		_,err := Origin_StoreSystem(originItem)
		if err != nil {
			logs.Error("Origin_SoftDelete()",err)
		}
}
	

// Origin_HardDelete(id string) soft deletes a single Origin record
func Origin_HardDelete(id string) {
			// Uses Hard Delete
		object_Table := Origin_QualifiedName
		tsql := das.DELETE + das.FROM + object_Table
		tsql = tsql + " " + das.WHERE + dm.Origin_SQLSearchID + das.EQ + das.ID(id)
		das.Execute(tsql)
		//if err != nil {
		//	logs.Error("Origin_SoftDelete()",err)
		//}
}


// Origin_Store() saves/stores a Origin record to the database
func Origin_Store(r dm.Origin,req *http.Request) (dm.Origin,error) {

	r, err := Origin_Validate(r)
	if err == nil {
		err = origin_Save(r, Audit_User(req))
	} else {
		logs.Information("Origin_Store()", err.Error())
	}

	return r, err
}

// Origin_StoreSystem() saves/stores a Origin record to the database
func Origin_StoreSystem(r dm.Origin) (dm.Origin,error) {
	
	r, err := Origin_Validate(r)
	if err == nil {
		err = origin_Save(r, Audit_Host())
	} else {
		logs.Information("Origin_Store()", err.Error())
	}

	return r, err
}

// Origin_StoreProcess() saves/stores a Origin record to the database
func Origin_StoreProcess(r dm.Origin, operator string) (dm.Origin,error) {
	
	r, err := Origin_Validate(r)
	if err == nil {
		err = origin_Save(r, operator)
	} else {
		logs.Information("Origin_StoreProcess()", err.Error())
	}

	return r, err
}

// Origin_Validate() validates for saves/stores a Origin record to the database
func Origin_Validate(r dm.Origin) (dm.Origin, error) {
	var err error
	r.StateID,r.StateID_props = Origin_StateID_validate_impl (PUT,r.OriginID,r.StateID,r,r.StateID_props)
	if r.StateID_props.MsgMessage != "" {
		err = errors.New(r.StateID_props.MsgMessage)
	}
	r.DocTypeID,r.DocTypeID_props = Origin_DocTypeID_validate_impl (PUT,r.OriginID,r.DocTypeID,r,r.DocTypeID_props)
	if r.DocTypeID_props.MsgMessage != "" {
		err = errors.New(r.DocTypeID_props.MsgMessage)
	}
	r.Code,r.Code_props = Origin_Code_validate_impl (PUT,r.OriginID,r.Code,r,r.Code_props)
	if r.Code_props.MsgMessage != "" {
		err = errors.New(r.Code_props.MsgMessage)
	}
	r.FullName,r.FullName_props = Origin_FullName_validate_impl (PUT,r.OriginID,r.FullName,r,r.FullName_props)
	if r.FullName_props.MsgMessage != "" {
		err = errors.New(r.FullName_props.MsgMessage)
	}
	r.Rate,r.Rate_props = Origin_Rate_validate_impl (PUT,r.OriginID,r.Rate,r,r.Rate_props)
	if r.Rate_props.MsgMessage != "" {
		err = errors.New(r.Rate_props.MsgMessage)
	}
	r.StartDate,r.StartDate_props = Origin_StartDate_validate_impl (PUT,r.OriginID,r.StartDate,r,r.StartDate_props)
	if r.StartDate_props.MsgMessage != "" {
		err = errors.New(r.StartDate_props.MsgMessage)
	}
	r.Currency,r.Currency_props = Origin_Currency_validate_impl (PUT,r.OriginID,r.Currency,r,r.Currency_props)
	if r.Currency_props.MsgMessage != "" {
		err = errors.New(r.Currency_props.MsgMessage)
	}
	r.NoActiveProjects,r.NoActiveProjects_props = Origin_NoActiveProjects_validate_impl (PUT,r.OriginID,r.NoActiveProjects,r,r.NoActiveProjects_props)
	if r.NoActiveProjects_props.MsgMessage != "" {
		err = errors.New(r.NoActiveProjects_props.MsgMessage)
	}
	r.RateOnLoad,r.RateOnLoad_props = Origin_RateOnLoad_validate_impl (PUT,r.OriginID,r.RateOnLoad,r,r.RateOnLoad_props)
	if r.RateOnLoad_props.MsgMessage != "" {
		err = errors.New(r.RateOnLoad_props.MsgMessage)
	}
	r.StatusOnLoad,r.StatusOnLoad_props = Origin_StatusOnLoad_validate_impl (PUT,r.OriginID,r.StatusOnLoad,r,r.StatusOnLoad_props)
	if r.StatusOnLoad_props.MsgMessage != "" {
		err = errors.New(r.StatusOnLoad_props.MsgMessage)
	}

	// Cross Validation
	var errVal error
	r, _, errVal = Origin_ObjectValidation_impl(PUT, r.OriginID, r)
	if errVal != nil {
		err = errVal
	}
	return r,err
}
//

// origin_Save() saves/stores a Origin record to the database
func origin_Save(r dm.Origin,usr string) error {

    var err error

	if len(r.OriginID) == 0 {
		r.OriginID = Origin_NewID(r)
	}

// If there are fields below, create the methods in dao\origin_impl.go
    r.StateID,err = Origin_StateID_OnStore_impl (r.StateID,r,usr)
    r.DocTypeID,err = Origin_DocTypeID_OnStore_impl (r.DocTypeID,r,usr)
    r.Code,err = Origin_Code_OnStore_impl (r.Code,r,usr)
    r.FullName,err = Origin_FullName_OnStore_impl (r.FullName,r,usr)
    r.Rate,err = Origin_Rate_OnStore_impl (r.Rate,r,usr)
    r.StartDate,err = Origin_StartDate_OnStore_impl (r.StartDate,r,usr)
    r.Currency,err = Origin_Currency_OnStore_impl (r.Currency,r,usr)
    r.NoActiveProjects,err = Origin_NoActiveProjects_OnStore_impl (r.NoActiveProjects,r,usr)
    r.RateOnLoad,err = Origin_RateOnLoad_OnStore_impl (r.RateOnLoad,r,usr)
    r.StatusOnLoad,err = Origin_StatusOnLoad_OnStore_impl (r.StatusOnLoad,r,usr)

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("Origin",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Origin_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Origin_OriginID_sql, r.OriginID)
	ts = addData(ts, dm.Origin_StateID_sql, r.StateID)
	ts = addData(ts, dm.Origin_DocTypeID_sql, r.DocTypeID)
	ts = addData(ts, dm.Origin_Code_sql, r.Code)
	ts = addData(ts, dm.Origin_FullName_sql, r.FullName)
	ts = addData(ts, dm.Origin_Rate_sql, r.Rate)
	ts = addData(ts, dm.Origin_Notes_sql, r.Notes)
	ts = addData(ts, dm.Origin_StartDate_sql, r.StartDate)
	ts = addData(ts, dm.Origin_EndDate_sql, r.EndDate)
	ts = addData(ts, dm.Origin_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Origin_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Origin_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Origin_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Origin_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Origin_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Origin_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Origin_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Origin_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Origin_SYSActivity_sql, r.SYSActivity)
	ts = addData(ts, dm.Origin_Currency_sql, r.Currency)
	ts = addData(ts, dm.Origin_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.Origin_Comments_sql, r.Comments)
	ts = addData(ts, dm.Origin_ProjectManager_sql, r.ProjectManager)
	ts = addData(ts, dm.Origin_AccountManager_sql, r.AccountManager)
	
	
	
		
	// 
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := das.INSERT + das.INTO + Origin_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+das.VALUES +"(" + values(ts) + ")"

	Origin_HardDelete(r.OriginID)
	das.Execute(tsql)

	



	return err

}



// origin_Fetch read all Origin's
func origin_Fetch(tsql string) (int, []dm.Origin, dm.Origin, error) {

	var recItem dm.Origin
	var recList []dm.Origin

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Origin_SYSId_sql, "0")
	   recItem.OriginID  = get_String(rec, dm.Origin_OriginID_sql, "")
	   recItem.StateID  = get_String(rec, dm.Origin_StateID_sql, "")
	   recItem.DocTypeID  = get_String(rec, dm.Origin_DocTypeID_sql, "")
	   recItem.Code  = get_String(rec, dm.Origin_Code_sql, "")
	   recItem.FullName  = get_String(rec, dm.Origin_FullName_sql, "")
	   recItem.Rate  = get_String(rec, dm.Origin_Rate_sql, "")
	   recItem.Notes  = get_String(rec, dm.Origin_Notes_sql, "")
	   recItem.StartDate  = get_String(rec, dm.Origin_StartDate_sql, "")
	   recItem.EndDate  = get_String(rec, dm.Origin_EndDate_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Origin_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Origin_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Origin_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Origin_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Origin_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Origin_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Origin_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Origin_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Origin_SYSDeletedHost_sql, "")
	   recItem.SYSActivity  = get_String(rec, dm.Origin_SYSActivity_sql, "")
	   recItem.Currency  = get_String(rec, dm.Origin_Currency_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.Origin_SYSDbVersion_sql, "")
	   recItem.Comments  = get_String(rec, dm.Origin_Comments_sql, "")
	   recItem.ProjectManager  = get_String(rec, dm.Origin_ProjectManager_sql, "")
	   recItem.AccountManager  = get_String(rec, dm.Origin_AccountManager_sql, "")
	
	
	
	
	// If there are fields below, create the methods in dao\Origin_adaptor.go
	   recItem.StateID  = Origin_StateID_OnFetch_impl (recItem)
	   recItem.DocTypeID  = Origin_DocTypeID_OnFetch_impl (recItem)
	   recItem.Code  = Origin_Code_OnFetch_impl (recItem)
	   recItem.FullName  = Origin_FullName_OnFetch_impl (recItem)
	   recItem.Rate  = Origin_Rate_OnFetch_impl (recItem)
	   recItem.StartDate  = Origin_StartDate_OnFetch_impl (recItem)
	   recItem.Currency  = Origin_Currency_OnFetch_impl (recItem)
	   recItem.NoActiveProjects  = Origin_NoActiveProjects_OnFetch_impl (recItem)
	   recItem.RateOnLoad  = Origin_RateOnLoad_OnFetch_impl (recItem)
	   recItem.StatusOnLoad  = Origin_StatusOnLoad_OnFetch_impl (recItem)
	// 
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Origin_NewID(r dm.Origin) string {
	
	id := uuid.New().String()
	
	return id
}



// origin_Fetch read all Origin's
func Origin_New() (int, []dm.Origin, dm.Origin, error) {

	var r = dm.Origin{}
	var rList []dm.Origin
	

	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.StateID,r.StateID_props = Origin_StateID_validate_impl (NEW,r.OriginID,r.StateID,r,r.StateID_props)
	r.DocTypeID,r.DocTypeID_props = Origin_DocTypeID_validate_impl (NEW,r.OriginID,r.DocTypeID,r,r.DocTypeID_props)
	r.Code,r.Code_props = Origin_Code_validate_impl (NEW,r.OriginID,r.Code,r,r.Code_props)
	r.FullName,r.FullName_props = Origin_FullName_validate_impl (NEW,r.OriginID,r.FullName,r,r.FullName_props)
	r.Rate,r.Rate_props = Origin_Rate_validate_impl (NEW,r.OriginID,r.Rate,r,r.Rate_props)
	r.StartDate,r.StartDate_props = Origin_StartDate_validate_impl (NEW,r.OriginID,r.StartDate,r,r.StartDate_props)
	r.Currency,r.Currency_props = Origin_Currency_validate_impl (NEW,r.OriginID,r.Currency,r,r.Currency_props)
	r.NoActiveProjects,r.NoActiveProjects_props = Origin_NoActiveProjects_validate_impl (NEW,r.OriginID,r.NoActiveProjects,r,r.NoActiveProjects_props)
	r.RateOnLoad,r.RateOnLoad_props = Origin_RateOnLoad_validate_impl (NEW,r.OriginID,r.RateOnLoad,r,r.RateOnLoad_props)
	r.StatusOnLoad,r.StatusOnLoad_props = Origin_StatusOnLoad_validate_impl (NEW,r.OriginID,r.StatusOnLoad,r,r.StatusOnLoad_props)
	
	// 
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}