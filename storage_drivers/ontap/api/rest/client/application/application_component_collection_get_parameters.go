// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewApplicationComponentCollectionGetParams creates a new ApplicationComponentCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationComponentCollectionGetParams() *ApplicationComponentCollectionGetParams {
	return &ApplicationComponentCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationComponentCollectionGetParamsWithTimeout creates a new ApplicationComponentCollectionGetParams object
// with the ability to set a timeout on a request.
func NewApplicationComponentCollectionGetParamsWithTimeout(timeout time.Duration) *ApplicationComponentCollectionGetParams {
	return &ApplicationComponentCollectionGetParams{
		timeout: timeout,
	}
}

// NewApplicationComponentCollectionGetParamsWithContext creates a new ApplicationComponentCollectionGetParams object
// with the ability to set a context for a request.
func NewApplicationComponentCollectionGetParamsWithContext(ctx context.Context) *ApplicationComponentCollectionGetParams {
	return &ApplicationComponentCollectionGetParams{
		Context: ctx,
	}
}

// NewApplicationComponentCollectionGetParamsWithHTTPClient creates a new ApplicationComponentCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationComponentCollectionGetParamsWithHTTPClient(client *http.Client) *ApplicationComponentCollectionGetParams {
	return &ApplicationComponentCollectionGetParams{
		HTTPClient: client,
	}
}

/* ApplicationComponentCollectionGetParams contains all the parameters to send to the API endpoint
   for the application component collection get operation.

   Typically these are written to a http.Request.
*/
type ApplicationComponentCollectionGetParams struct {

	/* ApplicationUUID.

	   Application UUID
	*/
	ApplicationUUIDPathParameter string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* StorageServiceName.

	   Filter by storage_service.name
	*/
	StorageServiceNameQueryParameter *string

	/* StorageServiceUUID.

	   Filter by storage_service.uuid
	*/
	StorageServiceUUIDQueryParameter *string

	/* UUID.

	   Filter by UUID
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application component collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationComponentCollectionGetParams) WithDefaults() *ApplicationComponentCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application component collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationComponentCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := ApplicationComponentCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithTimeout(timeout time.Duration) *ApplicationComponentCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithContext(ctx context.Context) *ApplicationComponentCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithHTTPClient(client *http.Client) *ApplicationComponentCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationUUIDPathParameter adds the applicationUUID to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithApplicationUUIDPathParameter(applicationUUID string) *ApplicationComponentCollectionGetParams {
	o.SetApplicationUUIDPathParameter(applicationUUID)
	return o
}

// SetApplicationUUIDPathParameter adds the applicationUuid to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetApplicationUUIDPathParameter(applicationUUID string) {
	o.ApplicationUUIDPathParameter = applicationUUID
}

// WithFields adds the fields to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithFields(fields []string) *ApplicationComponentCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithMaxRecords(maxRecords *int64) *ApplicationComponentCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNameQueryParameter adds the name to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithNameQueryParameter(name *string) *ApplicationComponentCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderBy adds the orderBy to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithOrderBy(orderBy []string) *ApplicationComponentCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithReturnRecords(returnRecords *bool) *ApplicationComponentCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *ApplicationComponentCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithStorageServiceNameQueryParameter adds the storageServiceName to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithStorageServiceNameQueryParameter(storageServiceName *string) *ApplicationComponentCollectionGetParams {
	o.SetStorageServiceNameQueryParameter(storageServiceName)
	return o
}

// SetStorageServiceNameQueryParameter adds the storageServiceName to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetStorageServiceNameQueryParameter(storageServiceName *string) {
	o.StorageServiceNameQueryParameter = storageServiceName
}

// WithStorageServiceUUIDQueryParameter adds the storageServiceUUID to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithStorageServiceUUIDQueryParameter(storageServiceUUID *string) *ApplicationComponentCollectionGetParams {
	o.SetStorageServiceUUIDQueryParameter(storageServiceUUID)
	return o
}

// SetStorageServiceUUIDQueryParameter adds the storageServiceUuid to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetStorageServiceUUIDQueryParameter(storageServiceUUID *string) {
	o.StorageServiceUUIDQueryParameter = storageServiceUUID
}

// WithUUIDQueryParameter adds the uuid to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) WithUUIDQueryParameter(uuid *string) *ApplicationComponentCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the application component collection get params
func (o *ApplicationComponentCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationComponentCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.uuid
	if err := r.SetPathParam("application.uuid", o.ApplicationUUIDPathParameter); err != nil {
		return err
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.StorageServiceNameQueryParameter != nil {

		// query param storage_service.name
		var qrStorageServiceName string

		if o.StorageServiceNameQueryParameter != nil {
			qrStorageServiceName = *o.StorageServiceNameQueryParameter
		}
		qStorageServiceName := qrStorageServiceName
		if qStorageServiceName != "" {

			if err := r.SetQueryParam("storage_service.name", qStorageServiceName); err != nil {
				return err
			}
		}
	}

	if o.StorageServiceUUIDQueryParameter != nil {

		// query param storage_service.uuid
		var qrStorageServiceUUID string

		if o.StorageServiceUUIDQueryParameter != nil {
			qrStorageServiceUUID = *o.StorageServiceUUIDQueryParameter
		}
		qStorageServiceUUID := qrStorageServiceUUID
		if qStorageServiceUUID != "" {

			if err := r.SetQueryParam("storage_service.uuid", qStorageServiceUUID); err != nil {
				return err
			}
		}
	}

	if o.UUIDQueryParameter != nil {

		// query param uuid
		var qrUUID string

		if o.UUIDQueryParameter != nil {
			qrUUID = *o.UUIDQueryParameter
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamApplicationComponentCollectionGet binds the parameter fields
func (o *ApplicationComponentCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamApplicationComponentCollectionGet binds the parameter order_by
func (o *ApplicationComponentCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}