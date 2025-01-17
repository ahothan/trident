// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewChassisCollectionGetParams creates a new ChassisCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChassisCollectionGetParams() *ChassisCollectionGetParams {
	return &ChassisCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChassisCollectionGetParamsWithTimeout creates a new ChassisCollectionGetParams object
// with the ability to set a timeout on a request.
func NewChassisCollectionGetParamsWithTimeout(timeout time.Duration) *ChassisCollectionGetParams {
	return &ChassisCollectionGetParams{
		timeout: timeout,
	}
}

// NewChassisCollectionGetParamsWithContext creates a new ChassisCollectionGetParams object
// with the ability to set a context for a request.
func NewChassisCollectionGetParamsWithContext(ctx context.Context) *ChassisCollectionGetParams {
	return &ChassisCollectionGetParams{
		Context: ctx,
	}
}

// NewChassisCollectionGetParamsWithHTTPClient creates a new ChassisCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewChassisCollectionGetParamsWithHTTPClient(client *http.Client) *ChassisCollectionGetParams {
	return &ChassisCollectionGetParams{
		HTTPClient: client,
	}
}

/* ChassisCollectionGetParams contains all the parameters to send to the API endpoint
   for the chassis collection get operation.

   Typically these are written to a http.Request.
*/
type ChassisCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* FrusID.

	   Filter by frus.id
	*/
	FrusIDQueryParameter *string

	/* FrusState.

	   Filter by frus.state
	*/
	FrusStateQueryParameter *string

	/* FrusType.

	   Filter by frus.type
	*/
	FrusTypeQueryParameter *string

	/* ID.

	   Filter by id
	*/
	IDQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* NodesName.

	   Filter by nodes.name
	*/
	NodesNameQueryParameter *string

	/* NodesPcisCardsDevice.

	   Filter by nodes.pcis.cards.device
	*/
	NodesPcisCardsDeviceQueryParameter *string

	/* NodesPcisCardsInfo.

	   Filter by nodes.pcis.cards.info
	*/
	NodesPcisCardsInfoQueryParameter *string

	/* NodesPcisCardsSlot.

	   Filter by nodes.pcis.cards.slot
	*/
	NodesPcisCardsSlotQueryParameter *string

	/* NodesPosition.

	   Filter by nodes.position
	*/
	NodesPositionQueryParameter *string

	/* NodesUsbsEnabled.

	   Filter by nodes.usbs.enabled
	*/
	NodesUsbsEnabledQueryParameter *bool

	/* NodesUsbsPortsConnected.

	   Filter by nodes.usbs.ports.connected
	*/
	NodesUsbsPortsConnectedQueryParameter *bool

	/* NodesUsbsSupported.

	   Filter by nodes.usbs.supported
	*/
	NodesUsbsSupportedQueryParameter *bool

	/* NodesUUID.

	   Filter by nodes.uuid
	*/
	NodesUUIDQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* ShelvesUID.

	   Filter by shelves.uid
	*/
	ShelvesUIDQueryParameter *string

	/* State.

	   Filter by state
	*/
	StateQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the chassis collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChassisCollectionGetParams) WithDefaults() *ChassisCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the chassis collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChassisCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := ChassisCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the chassis collection get params
func (o *ChassisCollectionGetParams) WithTimeout(timeout time.Duration) *ChassisCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chassis collection get params
func (o *ChassisCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chassis collection get params
func (o *ChassisCollectionGetParams) WithContext(ctx context.Context) *ChassisCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chassis collection get params
func (o *ChassisCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chassis collection get params
func (o *ChassisCollectionGetParams) WithHTTPClient(client *http.Client) *ChassisCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chassis collection get params
func (o *ChassisCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the chassis collection get params
func (o *ChassisCollectionGetParams) WithFieldsQueryParameter(fields []string) *ChassisCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the chassis collection get params
func (o *ChassisCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithFrusIDQueryParameter adds the frusID to the chassis collection get params
func (o *ChassisCollectionGetParams) WithFrusIDQueryParameter(frusID *string) *ChassisCollectionGetParams {
	o.SetFrusIDQueryParameter(frusID)
	return o
}

// SetFrusIDQueryParameter adds the frusId to the chassis collection get params
func (o *ChassisCollectionGetParams) SetFrusIDQueryParameter(frusID *string) {
	o.FrusIDQueryParameter = frusID
}

// WithFrusStateQueryParameter adds the frusState to the chassis collection get params
func (o *ChassisCollectionGetParams) WithFrusStateQueryParameter(frusState *string) *ChassisCollectionGetParams {
	o.SetFrusStateQueryParameter(frusState)
	return o
}

// SetFrusStateQueryParameter adds the frusState to the chassis collection get params
func (o *ChassisCollectionGetParams) SetFrusStateQueryParameter(frusState *string) {
	o.FrusStateQueryParameter = frusState
}

// WithFrusTypeQueryParameter adds the frusType to the chassis collection get params
func (o *ChassisCollectionGetParams) WithFrusTypeQueryParameter(frusType *string) *ChassisCollectionGetParams {
	o.SetFrusTypeQueryParameter(frusType)
	return o
}

// SetFrusTypeQueryParameter adds the frusType to the chassis collection get params
func (o *ChassisCollectionGetParams) SetFrusTypeQueryParameter(frusType *string) {
	o.FrusTypeQueryParameter = frusType
}

// WithIDQueryParameter adds the id to the chassis collection get params
func (o *ChassisCollectionGetParams) WithIDQueryParameter(id *string) *ChassisCollectionGetParams {
	o.SetIDQueryParameter(id)
	return o
}

// SetIDQueryParameter adds the id to the chassis collection get params
func (o *ChassisCollectionGetParams) SetIDQueryParameter(id *string) {
	o.IDQueryParameter = id
}

// WithMaxRecordsQueryParameter adds the maxRecords to the chassis collection get params
func (o *ChassisCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *ChassisCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the chassis collection get params
func (o *ChassisCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNodesNameQueryParameter adds the nodesName to the chassis collection get params
func (o *ChassisCollectionGetParams) WithNodesNameQueryParameter(nodesName *string) *ChassisCollectionGetParams {
	o.SetNodesNameQueryParameter(nodesName)
	return o
}

// SetNodesNameQueryParameter adds the nodesName to the chassis collection get params
func (o *ChassisCollectionGetParams) SetNodesNameQueryParameter(nodesName *string) {
	o.NodesNameQueryParameter = nodesName
}

// WithNodesPcisCardsDeviceQueryParameter adds the nodesPcisCardsDevice to the chassis collection get params
func (o *ChassisCollectionGetParams) WithNodesPcisCardsDeviceQueryParameter(nodesPcisCardsDevice *string) *ChassisCollectionGetParams {
	o.SetNodesPcisCardsDeviceQueryParameter(nodesPcisCardsDevice)
	return o
}

// SetNodesPcisCardsDeviceQueryParameter adds the nodesPcisCardsDevice to the chassis collection get params
func (o *ChassisCollectionGetParams) SetNodesPcisCardsDeviceQueryParameter(nodesPcisCardsDevice *string) {
	o.NodesPcisCardsDeviceQueryParameter = nodesPcisCardsDevice
}

// WithNodesPcisCardsInfoQueryParameter adds the nodesPcisCardsInfo to the chassis collection get params
func (o *ChassisCollectionGetParams) WithNodesPcisCardsInfoQueryParameter(nodesPcisCardsInfo *string) *ChassisCollectionGetParams {
	o.SetNodesPcisCardsInfoQueryParameter(nodesPcisCardsInfo)
	return o
}

// SetNodesPcisCardsInfoQueryParameter adds the nodesPcisCardsInfo to the chassis collection get params
func (o *ChassisCollectionGetParams) SetNodesPcisCardsInfoQueryParameter(nodesPcisCardsInfo *string) {
	o.NodesPcisCardsInfoQueryParameter = nodesPcisCardsInfo
}

// WithNodesPcisCardsSlotQueryParameter adds the nodesPcisCardsSlot to the chassis collection get params
func (o *ChassisCollectionGetParams) WithNodesPcisCardsSlotQueryParameter(nodesPcisCardsSlot *string) *ChassisCollectionGetParams {
	o.SetNodesPcisCardsSlotQueryParameter(nodesPcisCardsSlot)
	return o
}

// SetNodesPcisCardsSlotQueryParameter adds the nodesPcisCardsSlot to the chassis collection get params
func (o *ChassisCollectionGetParams) SetNodesPcisCardsSlotQueryParameter(nodesPcisCardsSlot *string) {
	o.NodesPcisCardsSlotQueryParameter = nodesPcisCardsSlot
}

// WithNodesPositionQueryParameter adds the nodesPosition to the chassis collection get params
func (o *ChassisCollectionGetParams) WithNodesPositionQueryParameter(nodesPosition *string) *ChassisCollectionGetParams {
	o.SetNodesPositionQueryParameter(nodesPosition)
	return o
}

// SetNodesPositionQueryParameter adds the nodesPosition to the chassis collection get params
func (o *ChassisCollectionGetParams) SetNodesPositionQueryParameter(nodesPosition *string) {
	o.NodesPositionQueryParameter = nodesPosition
}

// WithNodesUsbsEnabledQueryParameter adds the nodesUsbsEnabled to the chassis collection get params
func (o *ChassisCollectionGetParams) WithNodesUsbsEnabledQueryParameter(nodesUsbsEnabled *bool) *ChassisCollectionGetParams {
	o.SetNodesUsbsEnabledQueryParameter(nodesUsbsEnabled)
	return o
}

// SetNodesUsbsEnabledQueryParameter adds the nodesUsbsEnabled to the chassis collection get params
func (o *ChassisCollectionGetParams) SetNodesUsbsEnabledQueryParameter(nodesUsbsEnabled *bool) {
	o.NodesUsbsEnabledQueryParameter = nodesUsbsEnabled
}

// WithNodesUsbsPortsConnectedQueryParameter adds the nodesUsbsPortsConnected to the chassis collection get params
func (o *ChassisCollectionGetParams) WithNodesUsbsPortsConnectedQueryParameter(nodesUsbsPortsConnected *bool) *ChassisCollectionGetParams {
	o.SetNodesUsbsPortsConnectedQueryParameter(nodesUsbsPortsConnected)
	return o
}

// SetNodesUsbsPortsConnectedQueryParameter adds the nodesUsbsPortsConnected to the chassis collection get params
func (o *ChassisCollectionGetParams) SetNodesUsbsPortsConnectedQueryParameter(nodesUsbsPortsConnected *bool) {
	o.NodesUsbsPortsConnectedQueryParameter = nodesUsbsPortsConnected
}

// WithNodesUsbsSupportedQueryParameter adds the nodesUsbsSupported to the chassis collection get params
func (o *ChassisCollectionGetParams) WithNodesUsbsSupportedQueryParameter(nodesUsbsSupported *bool) *ChassisCollectionGetParams {
	o.SetNodesUsbsSupportedQueryParameter(nodesUsbsSupported)
	return o
}

// SetNodesUsbsSupportedQueryParameter adds the nodesUsbsSupported to the chassis collection get params
func (o *ChassisCollectionGetParams) SetNodesUsbsSupportedQueryParameter(nodesUsbsSupported *bool) {
	o.NodesUsbsSupportedQueryParameter = nodesUsbsSupported
}

// WithNodesUUIDQueryParameter adds the nodesUUID to the chassis collection get params
func (o *ChassisCollectionGetParams) WithNodesUUIDQueryParameter(nodesUUID *string) *ChassisCollectionGetParams {
	o.SetNodesUUIDQueryParameter(nodesUUID)
	return o
}

// SetNodesUUIDQueryParameter adds the nodesUuid to the chassis collection get params
func (o *ChassisCollectionGetParams) SetNodesUUIDQueryParameter(nodesUUID *string) {
	o.NodesUUIDQueryParameter = nodesUUID
}

// WithOrderByQueryParameter adds the orderBy to the chassis collection get params
func (o *ChassisCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *ChassisCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the chassis collection get params
func (o *ChassisCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the chassis collection get params
func (o *ChassisCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *ChassisCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the chassis collection get params
func (o *ChassisCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the chassis collection get params
func (o *ChassisCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *ChassisCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the chassis collection get params
func (o *ChassisCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithShelvesUIDQueryParameter adds the shelvesUID to the chassis collection get params
func (o *ChassisCollectionGetParams) WithShelvesUIDQueryParameter(shelvesUID *string) *ChassisCollectionGetParams {
	o.SetShelvesUIDQueryParameter(shelvesUID)
	return o
}

// SetShelvesUIDQueryParameter adds the shelvesUid to the chassis collection get params
func (o *ChassisCollectionGetParams) SetShelvesUIDQueryParameter(shelvesUID *string) {
	o.ShelvesUIDQueryParameter = shelvesUID
}

// WithStateQueryParameter adds the state to the chassis collection get params
func (o *ChassisCollectionGetParams) WithStateQueryParameter(state *string) *ChassisCollectionGetParams {
	o.SetStateQueryParameter(state)
	return o
}

// SetStateQueryParameter adds the state to the chassis collection get params
func (o *ChassisCollectionGetParams) SetStateQueryParameter(state *string) {
	o.StateQueryParameter = state
}

// WriteToRequest writes these params to a swagger request
func (o *ChassisCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.FrusIDQueryParameter != nil {

		// query param frus.id
		var qrFrusID string

		if o.FrusIDQueryParameter != nil {
			qrFrusID = *o.FrusIDQueryParameter
		}
		qFrusID := qrFrusID
		if qFrusID != "" {

			if err := r.SetQueryParam("frus.id", qFrusID); err != nil {
				return err
			}
		}
	}

	if o.FrusStateQueryParameter != nil {

		// query param frus.state
		var qrFrusState string

		if o.FrusStateQueryParameter != nil {
			qrFrusState = *o.FrusStateQueryParameter
		}
		qFrusState := qrFrusState
		if qFrusState != "" {

			if err := r.SetQueryParam("frus.state", qFrusState); err != nil {
				return err
			}
		}
	}

	if o.FrusTypeQueryParameter != nil {

		// query param frus.type
		var qrFrusType string

		if o.FrusTypeQueryParameter != nil {
			qrFrusType = *o.FrusTypeQueryParameter
		}
		qFrusType := qrFrusType
		if qFrusType != "" {

			if err := r.SetQueryParam("frus.type", qFrusType); err != nil {
				return err
			}
		}
	}

	if o.IDQueryParameter != nil {

		// query param id
		var qrID string

		if o.IDQueryParameter != nil {
			qrID = *o.IDQueryParameter
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NodesNameQueryParameter != nil {

		// query param nodes.name
		var qrNodesName string

		if o.NodesNameQueryParameter != nil {
			qrNodesName = *o.NodesNameQueryParameter
		}
		qNodesName := qrNodesName
		if qNodesName != "" {

			if err := r.SetQueryParam("nodes.name", qNodesName); err != nil {
				return err
			}
		}
	}

	if o.NodesPcisCardsDeviceQueryParameter != nil {

		// query param nodes.pcis.cards.device
		var qrNodesPcisCardsDevice string

		if o.NodesPcisCardsDeviceQueryParameter != nil {
			qrNodesPcisCardsDevice = *o.NodesPcisCardsDeviceQueryParameter
		}
		qNodesPcisCardsDevice := qrNodesPcisCardsDevice
		if qNodesPcisCardsDevice != "" {

			if err := r.SetQueryParam("nodes.pcis.cards.device", qNodesPcisCardsDevice); err != nil {
				return err
			}
		}
	}

	if o.NodesPcisCardsInfoQueryParameter != nil {

		// query param nodes.pcis.cards.info
		var qrNodesPcisCardsInfo string

		if o.NodesPcisCardsInfoQueryParameter != nil {
			qrNodesPcisCardsInfo = *o.NodesPcisCardsInfoQueryParameter
		}
		qNodesPcisCardsInfo := qrNodesPcisCardsInfo
		if qNodesPcisCardsInfo != "" {

			if err := r.SetQueryParam("nodes.pcis.cards.info", qNodesPcisCardsInfo); err != nil {
				return err
			}
		}
	}

	if o.NodesPcisCardsSlotQueryParameter != nil {

		// query param nodes.pcis.cards.slot
		var qrNodesPcisCardsSlot string

		if o.NodesPcisCardsSlotQueryParameter != nil {
			qrNodesPcisCardsSlot = *o.NodesPcisCardsSlotQueryParameter
		}
		qNodesPcisCardsSlot := qrNodesPcisCardsSlot
		if qNodesPcisCardsSlot != "" {

			if err := r.SetQueryParam("nodes.pcis.cards.slot", qNodesPcisCardsSlot); err != nil {
				return err
			}
		}
	}

	if o.NodesPositionQueryParameter != nil {

		// query param nodes.position
		var qrNodesPosition string

		if o.NodesPositionQueryParameter != nil {
			qrNodesPosition = *o.NodesPositionQueryParameter
		}
		qNodesPosition := qrNodesPosition
		if qNodesPosition != "" {

			if err := r.SetQueryParam("nodes.position", qNodesPosition); err != nil {
				return err
			}
		}
	}

	if o.NodesUsbsEnabledQueryParameter != nil {

		// query param nodes.usbs.enabled
		var qrNodesUsbsEnabled bool

		if o.NodesUsbsEnabledQueryParameter != nil {
			qrNodesUsbsEnabled = *o.NodesUsbsEnabledQueryParameter
		}
		qNodesUsbsEnabled := swag.FormatBool(qrNodesUsbsEnabled)
		if qNodesUsbsEnabled != "" {

			if err := r.SetQueryParam("nodes.usbs.enabled", qNodesUsbsEnabled); err != nil {
				return err
			}
		}
	}

	if o.NodesUsbsPortsConnectedQueryParameter != nil {

		// query param nodes.usbs.ports.connected
		var qrNodesUsbsPortsConnected bool

		if o.NodesUsbsPortsConnectedQueryParameter != nil {
			qrNodesUsbsPortsConnected = *o.NodesUsbsPortsConnectedQueryParameter
		}
		qNodesUsbsPortsConnected := swag.FormatBool(qrNodesUsbsPortsConnected)
		if qNodesUsbsPortsConnected != "" {

			if err := r.SetQueryParam("nodes.usbs.ports.connected", qNodesUsbsPortsConnected); err != nil {
				return err
			}
		}
	}

	if o.NodesUsbsSupportedQueryParameter != nil {

		// query param nodes.usbs.supported
		var qrNodesUsbsSupported bool

		if o.NodesUsbsSupportedQueryParameter != nil {
			qrNodesUsbsSupported = *o.NodesUsbsSupportedQueryParameter
		}
		qNodesUsbsSupported := swag.FormatBool(qrNodesUsbsSupported)
		if qNodesUsbsSupported != "" {

			if err := r.SetQueryParam("nodes.usbs.supported", qNodesUsbsSupported); err != nil {
				return err
			}
		}
	}

	if o.NodesUUIDQueryParameter != nil {

		// query param nodes.uuid
		var qrNodesUUID string

		if o.NodesUUIDQueryParameter != nil {
			qrNodesUUID = *o.NodesUUIDQueryParameter
		}
		qNodesUUID := qrNodesUUID
		if qNodesUUID != "" {

			if err := r.SetQueryParam("nodes.uuid", qNodesUUID); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.ShelvesUIDQueryParameter != nil {

		// query param shelves.uid
		var qrShelvesUID string

		if o.ShelvesUIDQueryParameter != nil {
			qrShelvesUID = *o.ShelvesUIDQueryParameter
		}
		qShelvesUID := qrShelvesUID
		if qShelvesUID != "" {

			if err := r.SetQueryParam("shelves.uid", qShelvesUID); err != nil {
				return err
			}
		}
	}

	if o.StateQueryParameter != nil {

		// query param state
		var qrState string

		if o.StateQueryParameter != nil {
			qrState = *o.StateQueryParameter
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamChassisCollectionGet binds the parameter fields
func (o *ChassisCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamChassisCollectionGet binds the parameter order_by
func (o *ChassisCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
