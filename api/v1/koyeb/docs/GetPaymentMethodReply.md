# GetPaymentMethodReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethod** | Pointer to [**PaymentMethod**](PaymentMethod.md) |  | [optional] 

## Methods

### NewGetPaymentMethodReply

`func NewGetPaymentMethodReply() *GetPaymentMethodReply`

NewGetPaymentMethodReply instantiates a new GetPaymentMethodReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentMethodReplyWithDefaults

`func NewGetPaymentMethodReplyWithDefaults() *GetPaymentMethodReply`

NewGetPaymentMethodReplyWithDefaults instantiates a new GetPaymentMethodReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethod

`func (o *GetPaymentMethodReply) GetPaymentMethod() PaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GetPaymentMethodReply) GetPaymentMethodOk() (*PaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GetPaymentMethodReply) SetPaymentMethod(v PaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GetPaymentMethodReply) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


