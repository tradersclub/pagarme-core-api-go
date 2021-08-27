/*
 * pagarmecoreapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package sellers_pkg


import(
	"time"
	"encoding/json"
	"github.com/apimatic/unirest-go"
	"pagarmecoreapi_lib/apihelper_pkg"
	"pagarmecoreapi_lib/configuration_pkg"
	"pagarmecoreapi_lib/models_pkg"
)
/*
 * Client structure as interface implementation
 */
type SELLERS_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * TODO: type endpoint description here
 * @param    *models_pkg.CreateSellerRequest        request             parameter: Required
 * @param    *string                                idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.GetSellerResponse response from the API call
 */
func (me *SELLERS_IMPL) CreateSeller (
            request *models_pkg.CreateSellerRequest,
            idempotencyKey *string) (*models_pkg.GetSellerResponse, error) {
    //the endpoint path uri
    _pathUrl := "/sellers/"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "PagarmeCoreApi - Go 5.0.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, request, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetSellerResponse = &models_pkg.GetSellerResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * TODO: type endpoint description here
 * @param    string                                   sellerId            parameter: Required
 * @param    *models_pkg.UpdateMetadataRequest        request             parameter: Required
 * @param    *string                                  idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.GetSellerResponse response from the API call
 */
func (me *SELLERS_IMPL) UpdateSellerMetadata (
            sellerId string,
            request *models_pkg.UpdateMetadataRequest,
            idempotencyKey *string) (*models_pkg.GetSellerResponse, error) {
    //the endpoint path uri
    _pathUrl := "/sellers/{seller_id}/metadata"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "seller_id" : sellerId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "PagarmeCoreApi - Go 5.0.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.PatchWithAuth(_queryBuilder, headers, request, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetSellerResponse = &models_pkg.GetSellerResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * TODO: type endpoint description here
 * @param    string                                 id                  parameter: Required
 * @param    *models_pkg.UpdateSellerRequest        request             parameter: Required
 * @param    *string                                idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.GetSellerResponse response from the API call
 */
func (me *SELLERS_IMPL) UpdateSeller (
            id string,
            request *models_pkg.UpdateSellerRequest,
            idempotencyKey *string) (*models_pkg.GetSellerResponse, error) {
    //the endpoint path uri
    _pathUrl := "/sellers/{id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "PagarmeCoreApi - Go 5.0.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.PutWithAuth(_queryBuilder, headers, request, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetSellerResponse = &models_pkg.GetSellerResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * TODO: type endpoint description here
 * @param    string         sellerId            parameter: Required
 * @param    *string        idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.GetSellerResponse response from the API call
 */
func (me *SELLERS_IMPL) DeleteSeller (
            sellerId string,
            idempotencyKey *string) (*models_pkg.GetSellerResponse, error) {
    //the endpoint path uri
    _pathUrl := "/sellers/{sellerId}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "sellerId" : sellerId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "PagarmeCoreApi - Go 5.0.0",
        "accept" : "application/json",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.DeleteWithAuth(_queryBuilder, headers, nil, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetSellerResponse = &models_pkg.GetSellerResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * TODO: type endpoint description here
 * @param    string        id     parameter: Required
 * @return	Returns the *models_pkg.GetSellerResponse response from the API call
 */
func (me *SELLERS_IMPL) GetSellerById (
            id string) (*models_pkg.GetSellerResponse, error) {
    //the endpoint path uri
    _pathUrl := "/sellers/{id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "PagarmeCoreApi - Go 5.0.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetSellerResponse = &models_pkg.GetSellerResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * TODO: type endpoint description here
 * @param    *int64            page              parameter: Optional
 * @param    *int64            size              parameter: Optional
 * @param    *string           name              parameter: Optional
 * @param    *string           document          parameter: Optional
 * @param    *string           code              parameter: Optional
 * @param    *string           status            parameter: Optional
 * @param    *string           mtype             parameter: Optional
 * @param    *time.Time        createdSince      parameter: Optional
 * @param    *time.Time        createdUntil      parameter: Optional
 * @return	Returns the *models_pkg.ListSellerResponse response from the API call
 */
func (me *SELLERS_IMPL) GetSellers (
            page *int64,
            size *int64,
            name *string,
            document *string,
            code *string,
            status *string,
            mtype *string,
            createdSince *time.Time,
            createdUntil *time.Time) (*models_pkg.ListSellerResponse, error) {
    //the endpoint path uri
    _pathUrl := "/sellers"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "page" : page,
        "size" : size,
        "name" : name,
        "document" : document,
        "code" : code,
        "status" : status,
        "type" : mtype,
        "created_Since" : createdSince,
        "created_Until" : createdUntil,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "PagarmeCoreApi - Go 5.0.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.ListSellerResponse = &models_pkg.ListSellerResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

