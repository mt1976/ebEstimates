package application

// ----------------------------------------------------------------
// Automatically generated  "/application/feature_api.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Feature (feature)
// Endpoint 	        : Feature (FeatureID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"encoding/json"

	"net/http"


	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
)

//Feature_Handler is the handler for the api calls
func Feature_Handler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	//TODO: Add your security validation here
	//        new => POST
	//  	 read => GET
	// 	   update => PUT
	//     delete => DELETE

	httpMethod := r.Method
	
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	//responseStatus := http.StatusOK

	switch httpMethod {
	case http.MethodGet:
		feature_MethodGet(w, r)

	case http.MethodPost:
		feature_MethodPost(w, r)

	case http.MethodPut:
		feature_MethodPost(w, r)
	case http.MethodDelete:

		feature_MethodDelete(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// Handles GET requests for Feature
func feature_MethodGet(w http.ResponseWriter, r *http.Request) {
	//logs.Information("PATH", r.URL.Path)
	searchID := core.GetURLparam(r, dm.Feature_QueryString)
	//logs.Information("GET", searchID)
	w.Header().Set("Content-Type", "application/json")

	if searchID == "" {
		//Get All Entities
		//logs.Information("GET", "All")
		noRecs, records, _ := dao.Feature_GetList()
		var ci core.ContentList
		ci.Count = noRecs
		ci.Key = dm.Feature_QueryString
		for _, v := range records {
			ciContent := core.ContentListItem{ID:v.FeatureID,Query:"?" + ci.Key +"="+ v.FeatureID}
			ci.Items= append(ci.Items, ciContent)
		}
		json_data, _ := json.Marshal(ci)
		w.Write(json_data)

		if noRecs == 0 {
			w.WriteHeader(int(http.StatusNotFound))
		} else {
			w.WriteHeader(int(http.StatusOK))
		}


	} else {
		//Get a specific entity
		_, record, _ := dao.Feature_GetByID(searchID)
		//spew.Dump(record)
		json_data, _ := json.Marshal(record)
		w.Write(json_data)

		if record.FeatureID == "" {
		    w.WriteHeader(int(http.StatusNotFound))
		} else {
			w.WriteHeader(int(http.StatusOK))
		}
	}


}

//Handles POST & PUT requests for Feature
func feature_MethodPost(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("POST")
	//fmt.Printf("r.Body: %v\n", r.Body)

	decoder := json.NewDecoder(r.Body)
	var t dm.Feature
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(int(http.StatusNotFound))
		panic(err)
	} else {
		w.WriteHeader(int(http.StatusOK))
	}
	//spew.Dump(t)
	
	_,err = dao.Feature_StoreSystem(t)
	if err != nil {

		//	panic(err)

		w.WriteHeader(int(http.StatusNotFound))

	} else {

		w.WriteHeader(int(http.StatusOK))

	}
	
	//logs.Processing("POST BACK")
	//logs.Information("POST", err.Error())
	
	//logs.Success("POST")
}
//Handles DELETE requests for Feature
func feature_MethodDelete(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("DELETE")
	//logs.Information("DELETE", deleteID)
		w.Header().Set("Content-Type", "application/json")

	deleteID := core.GetURLparam(r, dm.Feature_QueryString)

	dao.Feature_Delete(deleteID)
		w.WriteHeader(int(http.StatusOK))


	//fmt.Printf("json_data: %v\n", json_data)

	//logs.Success("DELETE")
}
