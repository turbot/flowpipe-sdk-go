# FpIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bcc** | Pointer to **[]string** |  | [optional] 
**Cc** | Pointer to **[]string** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EndLineNumber** | Pointer to **int32** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SigningSecret** | Pointer to **string** |  | [optional] 
**SmtpHost** | Pointer to **string** | email | [optional] 
**SmtpPassword** | Pointer to **string** |  | [optional] 
**SmtpPort** | Pointer to **int32** |  | [optional] 
**SmtpTls** | Pointer to **string** |  | [optional] 
**SmtpUsername** | Pointer to **string** |  | [optional] 
**SmtpsPort** | Pointer to **int32** |  | [optional] 
**StartLineNumber** | Pointer to **int32** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**To** | Pointer to **[]string** |  | [optional] 
**Token** | Pointer to **string** | slack | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **string** | slack &amp; msteams | [optional] 

## Methods

### NewFpIntegration

`func NewFpIntegration() *FpIntegration`

NewFpIntegration instantiates a new FpIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFpIntegrationWithDefaults

`func NewFpIntegrationWithDefaults() *FpIntegration`

NewFpIntegrationWithDefaults instantiates a new FpIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBcc

`func (o *FpIntegration) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *FpIntegration) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *FpIntegration) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *FpIntegration) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetCc

`func (o *FpIntegration) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *FpIntegration) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *FpIntegration) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *FpIntegration) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetChannel

`func (o *FpIntegration) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *FpIntegration) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *FpIntegration) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *FpIntegration) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDescription

`func (o *FpIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FpIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FpIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FpIntegration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndLineNumber

`func (o *FpIntegration) GetEndLineNumber() int32`

GetEndLineNumber returns the EndLineNumber field if non-nil, zero value otherwise.

### GetEndLineNumberOk

`func (o *FpIntegration) GetEndLineNumberOk() (*int32, bool)`

GetEndLineNumberOk returns a tuple with the EndLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLineNumber

`func (o *FpIntegration) SetEndLineNumber(v int32)`

SetEndLineNumber sets EndLineNumber field to given value.

### HasEndLineNumber

`func (o *FpIntegration) HasEndLineNumber() bool`

HasEndLineNumber returns a boolean if a field has been set.

### GetFileName

`func (o *FpIntegration) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FpIntegration) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FpIntegration) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FpIntegration) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFrom

`func (o *FpIntegration) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *FpIntegration) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *FpIntegration) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *FpIntegration) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetName

`func (o *FpIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FpIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FpIntegration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FpIntegration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSigningSecret

`func (o *FpIntegration) GetSigningSecret() string`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *FpIntegration) GetSigningSecretOk() (*string, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *FpIntegration) SetSigningSecret(v string)`

SetSigningSecret sets SigningSecret field to given value.

### HasSigningSecret

`func (o *FpIntegration) HasSigningSecret() bool`

HasSigningSecret returns a boolean if a field has been set.

### GetSmtpHost

`func (o *FpIntegration) GetSmtpHost() string`

GetSmtpHost returns the SmtpHost field if non-nil, zero value otherwise.

### GetSmtpHostOk

`func (o *FpIntegration) GetSmtpHostOk() (*string, bool)`

GetSmtpHostOk returns a tuple with the SmtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpHost

`func (o *FpIntegration) SetSmtpHost(v string)`

SetSmtpHost sets SmtpHost field to given value.

### HasSmtpHost

`func (o *FpIntegration) HasSmtpHost() bool`

HasSmtpHost returns a boolean if a field has been set.

### GetSmtpPassword

`func (o *FpIntegration) GetSmtpPassword() string`

GetSmtpPassword returns the SmtpPassword field if non-nil, zero value otherwise.

### GetSmtpPasswordOk

`func (o *FpIntegration) GetSmtpPasswordOk() (*string, bool)`

GetSmtpPasswordOk returns a tuple with the SmtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPassword

`func (o *FpIntegration) SetSmtpPassword(v string)`

SetSmtpPassword sets SmtpPassword field to given value.

### HasSmtpPassword

`func (o *FpIntegration) HasSmtpPassword() bool`

HasSmtpPassword returns a boolean if a field has been set.

### GetSmtpPort

`func (o *FpIntegration) GetSmtpPort() int32`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *FpIntegration) GetSmtpPortOk() (*int32, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *FpIntegration) SetSmtpPort(v int32)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *FpIntegration) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpTls

`func (o *FpIntegration) GetSmtpTls() string`

GetSmtpTls returns the SmtpTls field if non-nil, zero value otherwise.

### GetSmtpTlsOk

`func (o *FpIntegration) GetSmtpTlsOk() (*string, bool)`

GetSmtpTlsOk returns a tuple with the SmtpTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpTls

`func (o *FpIntegration) SetSmtpTls(v string)`

SetSmtpTls sets SmtpTls field to given value.

### HasSmtpTls

`func (o *FpIntegration) HasSmtpTls() bool`

HasSmtpTls returns a boolean if a field has been set.

### GetSmtpUsername

`func (o *FpIntegration) GetSmtpUsername() string`

GetSmtpUsername returns the SmtpUsername field if non-nil, zero value otherwise.

### GetSmtpUsernameOk

`func (o *FpIntegration) GetSmtpUsernameOk() (*string, bool)`

GetSmtpUsernameOk returns a tuple with the SmtpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpUsername

`func (o *FpIntegration) SetSmtpUsername(v string)`

SetSmtpUsername sets SmtpUsername field to given value.

### HasSmtpUsername

`func (o *FpIntegration) HasSmtpUsername() bool`

HasSmtpUsername returns a boolean if a field has been set.

### GetSmtpsPort

`func (o *FpIntegration) GetSmtpsPort() int32`

GetSmtpsPort returns the SmtpsPort field if non-nil, zero value otherwise.

### GetSmtpsPortOk

`func (o *FpIntegration) GetSmtpsPortOk() (*int32, bool)`

GetSmtpsPortOk returns a tuple with the SmtpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpsPort

`func (o *FpIntegration) SetSmtpsPort(v int32)`

SetSmtpsPort sets SmtpsPort field to given value.

### HasSmtpsPort

`func (o *FpIntegration) HasSmtpsPort() bool`

HasSmtpsPort returns a boolean if a field has been set.

### GetStartLineNumber

`func (o *FpIntegration) GetStartLineNumber() int32`

GetStartLineNumber returns the StartLineNumber field if non-nil, zero value otherwise.

### GetStartLineNumberOk

`func (o *FpIntegration) GetStartLineNumberOk() (*int32, bool)`

GetStartLineNumberOk returns a tuple with the StartLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLineNumber

`func (o *FpIntegration) SetStartLineNumber(v int32)`

SetStartLineNumber sets StartLineNumber field to given value.

### HasStartLineNumber

`func (o *FpIntegration) HasStartLineNumber() bool`

HasStartLineNumber returns a boolean if a field has been set.

### GetSubject

`func (o *FpIntegration) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *FpIntegration) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *FpIntegration) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *FpIntegration) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTitle

`func (o *FpIntegration) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FpIntegration) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FpIntegration) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FpIntegration) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTo

`func (o *FpIntegration) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *FpIntegration) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *FpIntegration) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *FpIntegration) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetToken

`func (o *FpIntegration) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FpIntegration) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FpIntegration) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FpIntegration) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *FpIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FpIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FpIntegration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FpIntegration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *FpIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FpIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FpIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FpIntegration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *FpIntegration) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *FpIntegration) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *FpIntegration) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *FpIntegration) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


