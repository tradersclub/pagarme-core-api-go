# Getting started

Pagarme API

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=PagarmeCoreApi-GoLang&projectName=pagarmecoreapi_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=pagarmecoreapi_lib)

## How to Use

The following section explains how to use the PagarmecoreapiLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=pagarmecoreapi_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=pagarmecoreapi_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=pagarmecoreapi_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=pagarmecoreapi_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=pagarmecoreapi_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=PagarmeCoreApi-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=pagarmecoreapi_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=pagarmecoreapi_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| basicAuthUserName | The username to use with basic authentication |
| basicAuthPassword | The password to use with basic authentication |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [plans_pkg](#plans_pkg)
* [subscriptions_pkg](#subscriptions_pkg)
* [invoices_pkg](#invoices_pkg)
* [orders_pkg](#orders_pkg)
* [customers_pkg](#customers_pkg)
* [recipients_pkg](#recipients_pkg)
* [charges_pkg](#charges_pkg)
* [transfers_pkg](#transfers_pkg)
* [tokens_pkg](#tokens_pkg)
* [sellers_pkg](#sellers_pkg)
* [transactions_pkg](#transactions_pkg)

## <a name="plans_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".plans_pkg") plans_pkg

### Get instance

Factory for the ``` PLANS ``` interface can be accessed from the package plans_pkg.

```go
plans := plans_pkg.NewPLANS()
```

### <a name="get_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.GetPlan") GetPlan

> Gets a plan


```go
func (me *PLANS_IMPL) GetPlan(planId string)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |


#### Example Usage

```go
planId := "plan_id"

var result *models_pkg.GetPlanResponse
result,_ = plans.GetPlan(planId)

```


### <a name="delete_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.DeletePlan") DeletePlan

> Deletes a plan


```go
func (me *PLANS_IMPL) DeletePlan(
            planId string,
            idempotencyKey *string)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetPlanResponse
result,_ = plans.DeletePlan(planId, idempotencyKey)

```


### <a name="update_plan_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.UpdatePlanMetadata") UpdatePlanMetadata

> Updates the metadata from a plan


```go
func (me *PLANS_IMPL) UpdatePlanMetadata(
            planId string,
            request *models_pkg.UpdateMetadataRequest,
            idempotencyKey *string)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | The plan id |
| request |  ``` Required ```  | Request for updating the plan metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
var request *models_pkg.UpdateMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetPlanResponse
result,_ = plans.UpdatePlanMetadata(planId, request, idempotencyKey)

```


### <a name="update_plan_item"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.UpdatePlanItem") UpdatePlanItem

> Updates a plan item


```go
func (me *PLANS_IMPL) UpdatePlanItem(
            planId string,
            planItemId string,
            body *models_pkg.UpdatePlanItemRequest,
            idempotencyKey *string)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| planItemId |  ``` Required ```  | Plan item id |
| body |  ``` Required ```  | Request for updating the plan item |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
planItemId := "plan_item_id"
var body *models_pkg.UpdatePlanItemRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetPlanItemResponse
result,_ = plans.UpdatePlanItem(planId, planItemId, body, idempotencyKey)

```


### <a name="create_plan_item"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.CreatePlanItem") CreatePlanItem

> Adds a new item to a plan


```go
func (me *PLANS_IMPL) CreatePlanItem(
            planId string,
            request *models_pkg.CreatePlanItemRequest,
            idempotencyKey *string)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| request |  ``` Required ```  | Request for creating a plan item |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
var request *models_pkg.CreatePlanItemRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetPlanItemResponse
result,_ = plans.CreatePlanItem(planId, request, idempotencyKey)

```


### <a name="get_plan_item"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.GetPlanItem") GetPlanItem

> Gets a plan item


```go
func (me *PLANS_IMPL) GetPlanItem(
            planId string,
            planItemId string)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| planItemId |  ``` Required ```  | Plan item id |


#### Example Usage

```go
planId := "plan_id"
planItemId := "plan_item_id"

var result *models_pkg.GetPlanItemResponse
result,_ = plans.GetPlanItem(planId, planItemId)

```


### <a name="create_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.CreatePlan") CreatePlan

> Creates a new plan


```go
func (me *PLANS_IMPL) CreatePlan(
            body *models_pkg.CreatePlanRequest,
            idempotencyKey *string)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Request for creating a plan |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreatePlanRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetPlanResponse
result,_ = plans.CreatePlan(body, idempotencyKey)

```


### <a name="delete_plan_item"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.DeletePlanItem") DeletePlanItem

> Removes an item from a plan


```go
func (me *PLANS_IMPL) DeletePlanItem(
            planId string,
            planItemId string,
            idempotencyKey *string)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| planItemId |  ``` Required ```  | Plan item id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
planItemId := "plan_item_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetPlanItemResponse
result,_ = plans.DeletePlanItem(planId, planItemId, idempotencyKey)

```


### <a name="get_plans"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.GetPlans") GetPlans

> Gets all plans


```go
func (me *PLANS_IMPL) GetPlans(
            page *int64,
            size *int64,
            name *string,
            status *string,
            billingType *string,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.ListPlansResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| name |  ``` Optional ```  | Filter for Plan's name |
| status |  ``` Optional ```  | Filter for Plan's status |
| billingType |  ``` Optional ```  | Filter for plan's billing type |
| createdSince |  ``` Optional ```  | Filter for plan's creation date start range |
| createdUntil |  ``` Optional ```  | Filter for plan's creation date end range |


#### Example Usage

```go
page,_ := strconv.ParseInt("170", 10, 8)
size,_ := strconv.ParseInt("170", 10, 8)
name := "name"
status := "status"
billingType := "billing_type"
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.ListPlansResponse
result,_ = plans.GetPlans(page, size, name, status, billingType, createdSince, createdUntil)

```


### <a name="update_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.UpdatePlan") UpdatePlan

> Updates a plan


```go
func (me *PLANS_IMPL) UpdatePlan(
            planId string,
            request *models_pkg.UpdatePlanRequest,
            idempotencyKey *string)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| request |  ``` Required ```  | Request for updating a plan |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
var request *models_pkg.UpdatePlanRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetPlanResponse
result,_ = plans.UpdatePlan(planId, request, idempotencyKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="subscriptions_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".subscriptions_pkg") subscriptions_pkg

### Get instance

Factory for the ``` SUBSCRIPTIONS ``` interface can be accessed from the package subscriptions_pkg.

```go
subscriptions := subscriptions_pkg.NewSUBSCRIPTIONS()
```

### <a name="renew_subscription"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.RenewSubscription") RenewSubscription

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) RenewSubscription(
            subscriptionId string,
            idempotencyKey *string)(*models_pkg.GetPeriodResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetPeriodResponse
result,_ = subscriptions.RenewSubscription(subscriptionId, idempotencyKey)

```


### <a name="update_subscription_card"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionCard") UpdateSubscriptionCard

> Updates the credit card from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionCard(
            subscriptionId string,
            request *models_pkg.UpdateSubscriptionCardRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| request |  ``` Required ```  | Request for updating a card |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateSubscriptionCardRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionCard(subscriptionId, request, idempotencyKey)

```


### <a name="delete_usage"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.DeleteUsage") DeleteUsage

> Deletes a usage


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteUsage(
            subscriptionId string,
            itemId string,
            usageId string,
            idempotencyKey *string)(*models_pkg.GetUsageResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| itemId |  ``` Required ```  | The subscription item id |
| usageId |  ``` Required ```  | The usage id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
usageId := "usage_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetUsageResponse
result,_ = subscriptions.DeleteUsage(subscriptionId, itemId, usageId, idempotencyKey)

```


### <a name="create_discount"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateDiscount") CreateDiscount

> Creates a discount


```go
func (me *SUBSCRIPTIONS_IMPL) CreateDiscount(
            subscriptionId string,
            request *models_pkg.CreateDiscountRequest,
            idempotencyKey *string)(*models_pkg.GetDiscountResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| request |  ``` Required ```  | Request for creating a discount |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.CreateDiscountRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetDiscountResponse
result,_ = subscriptions.CreateDiscount(subscriptionId, request, idempotencyKey)

```


### <a name="create_an_usage"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateAnUsage") CreateAnUsage

> Create Usage


```go
func (me *SUBSCRIPTIONS_IMPL) CreateAnUsage(
            subscriptionId string,
            itemId string,
            idempotencyKey *string)(*models_pkg.GetUsageResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| itemId |  ``` Required ```  | Item id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetUsageResponse
result,_ = subscriptions.CreateAnUsage(subscriptionId, itemId, idempotencyKey)

```


### <a name="update_current_cycle_status"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateCurrentCycleStatus") UpdateCurrentCycleStatus

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateCurrentCycleStatus(
            subscriptionId string,
            request *models_pkg.UpdateCurrentCycleStatusRequest,
            idempotencyKey *string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| request |  ``` Required ```  | Request for updating the end date of the subscription current status |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateCurrentCycleStatusRequest
idempotencyKey := "idempotency-key"

var result 
result,_ = subscriptions.UpdateCurrentCycleStatus(subscriptionId, request, idempotencyKey)

```


### <a name="delete_discount"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.DeleteDiscount") DeleteDiscount

> Deletes a discount


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteDiscount(
            subscriptionId string,
            discountId string,
            idempotencyKey *string)(*models_pkg.GetDiscountResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| discountId |  ``` Required ```  | Discount Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
discountId := "discount_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetDiscountResponse
result,_ = subscriptions.DeleteDiscount(subscriptionId, discountId, idempotencyKey)

```


### <a name="get_subscription_items"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptionItems") GetSubscriptionItems

> Get Subscription Items


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptionItems(
            subscriptionId string,
            page *int64,
            size *int64,
            name *string,
            code *string,
            status *string,
            description *string,
            createdSince *string,
            createdUntil *string)(*models_pkg.ListSubscriptionItemsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| name |  ``` Optional ```  | The item name |
| code |  ``` Optional ```  | Identification code in the client system |
| status |  ``` Optional ```  | The item statis |
| description |  ``` Optional ```  | The item description |
| createdSince |  ``` Optional ```  | Filter for item's creation date start range |
| createdUntil |  ``` Optional ```  | Filter for item's creation date end range |


#### Example Usage

```go
subscriptionId := "subscription_id"
page,_ := strconv.ParseInt("170", 10, 8)
size,_ := strconv.ParseInt("170", 10, 8)
name := "name"
code := "code"
status := "status"
description := "description"
createdSince := "created_since"
createdUntil := "created_until"

var result *models_pkg.ListSubscriptionItemsResponse
result,_ = subscriptions.GetSubscriptionItems(subscriptionId, page, size, name, code, status, description, createdSince, createdUntil)

```


### <a name="update_subscription_payment_method"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionPaymentMethod") UpdateSubscriptionPaymentMethod

> Updates the payment method from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionPaymentMethod(
            subscriptionId string,
            request *models_pkg.UpdateSubscriptionPaymentMethodRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| request |  ``` Required ```  | Request for updating the paymentmethod from a subscription |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateSubscriptionPaymentMethodRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionPaymentMethod(subscriptionId, request, idempotencyKey)

```


### <a name="get_subscription_item"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptionItem") GetSubscriptionItem

> Get Subscription Item


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptionItem(
            subscriptionId string,
            itemId string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| itemId |  ``` Required ```  | Item id |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.GetSubscriptionItem(subscriptionId, itemId)

```


### <a name="get_subscriptions"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptions") GetSubscriptions

> Gets all subscriptions


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptions(
            page *int64,
            size *int64,
            code *string,
            billingType *string,
            customerId *string,
            planId *string,
            cardId *string,
            status *string,
            nextBillingSince *time.Time,
            nextBillingUntil *time.Time,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.ListSubscriptionsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| code |  ``` Optional ```  | Filter for subscription's code |
| billingType |  ``` Optional ```  | Filter for subscription's billing type |
| customerId |  ``` Optional ```  | Filter for subscription's customer id |
| planId |  ``` Optional ```  | Filter for subscription's plan id |
| cardId |  ``` Optional ```  | Filter for subscription's card id |
| status |  ``` Optional ```  | Filter for subscription's status |
| nextBillingSince |  ``` Optional ```  | Filter for subscription's next billing date start range |
| nextBillingUntil |  ``` Optional ```  | Filter for subscription's next billing date end range |
| createdSince |  ``` Optional ```  | Filter for subscription's creation date start range |
| createdUntil |  ``` Optional ```  | Filter for subscriptions creation date end range |


#### Example Usage

```go
page,_ := strconv.ParseInt("170", 10, 8)
size,_ := strconv.ParseInt("170", 10, 8)
code := "code"
billingType := "billing_type"
customerId := "customer_id"
planId := "plan_id"
cardId := "card_id"
status := "status"
nextBillingSince := time.Now()
nextBillingUntil := time.Now()
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.ListSubscriptionsResponse
result,_ = subscriptions.GetSubscriptions(page, size, code, billingType, customerId, planId, cardId, status, nextBillingSince, nextBillingUntil, createdSince, createdUntil)

```


### <a name="cancel_subscription"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CancelSubscription") CancelSubscription

> Cancels a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) CancelSubscription(
            subscriptionId string,
            request *models_pkg.CreateCancelSubscriptionRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| request |  ``` Optional ```  | Request for cancelling a subscription |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.CreateCancelSubscriptionRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.CancelSubscription(subscriptionId, request, idempotencyKey)

```


### <a name="create_increment"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateIncrement") CreateIncrement

> Creates a increment


```go
func (me *SUBSCRIPTIONS_IMPL) CreateIncrement(
            subscriptionId string,
            request *models_pkg.CreateIncrementRequest,
            idempotencyKey *string)(*models_pkg.GetIncrementResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| request |  ``` Required ```  | Request for creating a increment |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.CreateIncrementRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetIncrementResponse
result,_ = subscriptions.CreateIncrement(subscriptionId, request, idempotencyKey)

```


### <a name="create_usage"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateUsage") CreateUsage

> Creates a usage


```go
func (me *SUBSCRIPTIONS_IMPL) CreateUsage(
            subscriptionId string,
            itemId string,
            body *models_pkg.CreateUsageRequest,
            idempotencyKey *string)(*models_pkg.GetUsageResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| itemId |  ``` Required ```  | Item id |
| body |  ``` Required ```  | Request for creating a usage |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
var body *models_pkg.CreateUsageRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetUsageResponse
result,_ = subscriptions.CreateUsage(subscriptionId, itemId, body, idempotencyKey)

```


### <a name="get_discount_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetDiscountById") GetDiscountById

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetDiscountById(
            subscriptionId string,
            discountId string)(*models_pkg.GetDiscountResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| discountId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
discountId := "discountId"

var result *models_pkg.GetDiscountResponse
result,_ = subscriptions.GetDiscountById(subscriptionId, discountId)

```


### <a name="create_subscription"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateSubscription") CreateSubscription

> Creates a new subscription


```go
func (me *SUBSCRIPTIONS_IMPL) CreateSubscription(
            body *models_pkg.CreateSubscriptionRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Request for creating a subscription |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateSubscriptionRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.CreateSubscription(body, idempotencyKey)

```


### <a name="get_increment_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetIncrementById") GetIncrementById

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetIncrementById(
            subscriptionId string,
            incrementId string)(*models_pkg.GetIncrementResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription Id |
| incrementId |  ``` Required ```  | The increment Id |


#### Example Usage

```go
subscriptionId := "subscription_id"
incrementId := "increment_id"

var result *models_pkg.GetIncrementResponse
result,_ = subscriptions.GetIncrementById(subscriptionId, incrementId)

```


### <a name="update_subscription_affiliation_id"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionAffiliationId") UpdateSubscriptionAffiliationId

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionAffiliationId(
            subscriptionId string,
            request *models_pkg.UpdateSubscriptionAffiliationIdRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| request |  ``` Required ```  | Request for updating a subscription affiliation id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateSubscriptionAffiliationIdRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionAffiliationId(subscriptionId, request, idempotencyKey)

```


### <a name="update_subscription_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionMetadata") UpdateSubscriptionMetadata

> Updates the metadata from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionMetadata(
            subscriptionId string,
            request *models_pkg.UpdateMetadataRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| request |  ``` Required ```  | Request for updating the subscrption metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionMetadata(subscriptionId, request, idempotencyKey)

```


### <a name="delete_increment"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.DeleteIncrement") DeleteIncrement

> Deletes a increment


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteIncrement(
            subscriptionId string,
            incrementId string,
            idempotencyKey *string)(*models_pkg.GetIncrementResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| incrementId |  ``` Required ```  | Increment id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
incrementId := "increment_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetIncrementResponse
result,_ = subscriptions.DeleteIncrement(subscriptionId, incrementId, idempotencyKey)

```


### <a name="get_subscription_cycles"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptionCycles") GetSubscriptionCycles

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptionCycles(
            subscriptionId string,
            page string,
            size string)(*models_pkg.ListCyclesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| page |  ``` Required ```  | Page number |
| size |  ``` Required ```  | Page size |


#### Example Usage

```go
subscriptionId := "subscription_id"
page := "page"
size := "size"

var result *models_pkg.ListCyclesResponse
result,_ = subscriptions.GetSubscriptionCycles(subscriptionId, page, size)

```


### <a name="get_discounts"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetDiscounts") GetDiscounts

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetDiscounts(
            subscriptionId string,
            page int64,
            size int64)(*models_pkg.ListDiscountsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| page |  ``` Required ```  | Page number |
| size |  ``` Required ```  | Page size |


#### Example Usage

```go
subscriptionId := "subscription_id"
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)

var result *models_pkg.ListDiscountsResponse
result,_ = subscriptions.GetDiscounts(subscriptionId, page, size)

```


### <a name="update_subscription_billing_date"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionBillingDate") UpdateSubscriptionBillingDate

> Updates the billing date from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionBillingDate(
            subscriptionId string,
            request *models_pkg.UpdateSubscriptionBillingDateRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| request |  ``` Required ```  | Request for updating the subscription billing date |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateSubscriptionBillingDateRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionBillingDate(subscriptionId, request, idempotencyKey)

```


### <a name="delete_subscription_item"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.DeleteSubscriptionItem") DeleteSubscriptionItem

> Deletes a subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteSubscriptionItem(
            subscriptionId string,
            subscriptionItemId string,
            idempotencyKey *string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| subscriptionItemId |  ``` Required ```  | Subscription item id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
subscriptionItemId := "subscription_item_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.DeleteSubscriptionItem(subscriptionId, subscriptionItemId, idempotencyKey)

```


### <a name="get_increments"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetIncrements") GetIncrements

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetIncrements(
            subscriptionId string,
            page *int64,
            size *int64)(*models_pkg.ListIncrementsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
subscriptionId := "subscription_id"
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)

var result *models_pkg.ListIncrementsResponse
result,_ = subscriptions.GetIncrements(subscriptionId, page, size)

```


### <a name="update_subscription_due_days"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionDueDays") UpdateSubscriptionDueDays

> Updates the boleto due days from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionDueDays(
            subscriptionId string,
            request *models_pkg.UpdateSubscriptionDueDaysRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| request |  ``` Required ```  | TODO: Add a parameter description |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateSubscriptionDueDaysRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionDueDays(subscriptionId, request, idempotencyKey)

```


### <a name="update_subscription_start_at"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionStartAt") UpdateSubscriptionStartAt

> Updates the start at date from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionStartAt(
            subscriptionId string,
            request *models_pkg.UpdateSubscriptionStartAtRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| request |  ``` Required ```  | Request for updating the subscription start date |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateSubscriptionStartAtRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionStartAt(subscriptionId, request, idempotencyKey)

```


### <a name="update_subscription_item"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionItem") UpdateSubscriptionItem

> Updates a subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionItem(
            subscriptionId string,
            itemId string,
            body *models_pkg.UpdateSubscriptionItemRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| itemId |  ``` Required ```  | Item id |
| body |  ``` Required ```  | Request for updating a subscription item |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
var body *models_pkg.UpdateSubscriptionItemRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.UpdateSubscriptionItem(subscriptionId, itemId, body, idempotencyKey)

```


### <a name="create_subscription_item"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateSubscriptionItem") CreateSubscriptionItem

> Creates a new Subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) CreateSubscriptionItem(
            subscriptionId string,
            request *models_pkg.CreateSubscriptionItemRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| request |  ``` Required ```  | Request for creating a subscription item |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.CreateSubscriptionItemRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.CreateSubscriptionItem(subscriptionId, request, idempotencyKey)

```


### <a name="get_subscription"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscription") GetSubscription

> Gets a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscription(subscriptionId string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |


#### Example Usage

```go
subscriptionId := "subscription_id"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.GetSubscription(subscriptionId)

```


### <a name="get_usages"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetUsages") GetUsages

> Lists all usages from a subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) GetUsages(
            subscriptionId string,
            itemId string,
            page *int64,
            size *int64,
            code *string,
            group *string,
            usedSince *time.Time,
            usedUntil *time.Time)(*models_pkg.ListUsagesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| itemId |  ``` Required ```  | The subscription item id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| code |  ``` Optional ```  | Identification code in the client system |
| group |  ``` Optional ```  | Identification group in the client system |
| usedSince |  ``` Optional ```  | TODO: Add a parameter description |
| usedUntil |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)
code := "code"
group := "group"
usedSince := time.Now()
usedUntil := time.Now()

var result *models_pkg.ListUsagesResponse
result,_ = subscriptions.GetUsages(subscriptionId, itemId, page, size, code, group, usedSince, usedUntil)

```


### <a name="update_latest_period_end_at"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateLatestPeriodEndAt") UpdateLatestPeriodEndAt

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateLatestPeriodEndAt(
            subscriptionId string,
            request *models_pkg.UpdateCurrentCycleEndDateRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| request |  ``` Required ```  | Request for updating the end date of the current signature cycle |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateCurrentCycleEndDateRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateLatestPeriodEndAt(subscriptionId, request, idempotencyKey)

```


### <a name="update_subscription_minium_price"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionMiniumPrice") UpdateSubscriptionMiniumPrice

> Atualização do valor mínimo da assinatura


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionMiniumPrice(
            subscriptionId string,
            request *models_pkg.UpdateSubscriptionMinimumPriceRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| request |  ``` Required ```  | Request da requisição com o valor mínimo que será configurado |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateSubscriptionMinimumPriceRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionMiniumPrice(subscriptionId, request, idempotencyKey)

```


### <a name="get_subscription_cycle_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptionCycleById") GetSubscriptionCycleById

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptionCycleById(
            subscriptionId string,
            cycleId string)(*models_pkg.GetPeriodResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| cycleId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
cycleId := "cycleId"

var result *models_pkg.GetPeriodResponse
result,_ = subscriptions.GetSubscriptionCycleById(subscriptionId, cycleId)

```


### <a name="get_usage_report"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetUsageReport") GetUsageReport

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetUsageReport(
            subscriptionId string,
            periodId string)(*models_pkg.GetUsageReportResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription Id |
| periodId |  ``` Required ```  | The period Id |


#### Example Usage

```go
subscriptionId := "subscription_id"
periodId := "period_id"

var result *models_pkg.GetUsageReportResponse
result,_ = subscriptions.GetUsageReport(subscriptionId, periodId)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="invoices_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".invoices_pkg") invoices_pkg

### Get instance

Factory for the ``` INVOICES ``` interface can be accessed from the package invoices_pkg.

```go
invoices := invoices_pkg.NewINVOICES()
```

### <a name="update_invoice_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.UpdateInvoiceMetadata") UpdateInvoiceMetadata

> Updates the metadata from an invoice


```go
func (me *INVOICES_IMPL) UpdateInvoiceMetadata(
            invoiceId string,
            request *models_pkg.UpdateMetadataRequest,
            idempotencyKey *string)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | The invoice id |
| request |  ``` Required ```  | Request for updating the invoice metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
invoiceId := "invoice_id"
var request *models_pkg.UpdateMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetInvoiceResponse
result,_ = invoices.UpdateInvoiceMetadata(invoiceId, request, idempotencyKey)

```


### <a name="get_partial_invoice"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.GetPartialInvoice") GetPartialInvoice

> TODO: Add a method description


```go
func (me *INVOICES_IMPL) GetPartialInvoice(subscriptionId string)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |


#### Example Usage

```go
subscriptionId := "subscription_id"

var result *models_pkg.GetInvoiceResponse
result,_ = invoices.GetPartialInvoice(subscriptionId)

```


### <a name="cancel_invoice"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.CancelInvoice") CancelInvoice

> Cancels an invoice


```go
func (me *INVOICES_IMPL) CancelInvoice(
            invoiceId string,
            idempotencyKey *string)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | Invoice id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
invoiceId := "invoice_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetInvoiceResponse
result,_ = invoices.CancelInvoice(invoiceId, idempotencyKey)

```


### <a name="create_invoice"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.CreateInvoice") CreateInvoice

> Create an Invoice


```go
func (me *INVOICES_IMPL) CreateInvoice(
            subscriptionId string,
            cycleId string,
            request *models_pkg.CreateInvoiceRequest,
            idempotencyKey *string)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| cycleId |  ``` Required ```  | Cycle Id |
| request |  ``` Optional ```  | TODO: Add a parameter description |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
cycleId := "cycle_id"
var request *models_pkg.CreateInvoiceRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetInvoiceResponse
result,_ = invoices.CreateInvoice(subscriptionId, cycleId, request, idempotencyKey)

```


### <a name="get_invoices"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.GetInvoices") GetInvoices

> Gets all invoices


```go
func (me *INVOICES_IMPL) GetInvoices(
            page *int64,
            size *int64,
            code *string,
            customerId *string,
            subscriptionId *string,
            createdSince *time.Time,
            createdUntil *time.Time,
            status *string,
            dueSince *time.Time,
            dueUntil *time.Time,
            customerDocument *string)(*models_pkg.ListInvoicesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| code |  ``` Optional ```  | Filter for Invoice's code |
| customerId |  ``` Optional ```  | Filter for Invoice's customer id |
| subscriptionId |  ``` Optional ```  | Filter for Invoice's subscription id |
| createdSince |  ``` Optional ```  | Filter for Invoice's creation date start range |
| createdUntil |  ``` Optional ```  | Filter for Invoices creation date end range |
| status |  ``` Optional ```  | Filter for Invoice's status |
| dueSince |  ``` Optional ```  | Filter for Invoice's due date start range |
| dueUntil |  ``` Optional ```  | Filter for Invoice's due date end range |
| customerDocument |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)
code := "code"
customerId := "customer_id"
subscriptionId := "subscription_id"
createdSince := time.Now()
createdUntil := time.Now()
status := "status"
dueSince := time.Now()
dueUntil := time.Now()
customerDocument := "customer_document"

var result *models_pkg.ListInvoicesResponse
result,_ = invoices.GetInvoices(page, size, code, customerId, subscriptionId, createdSince, createdUntil, status, dueSince, dueUntil, customerDocument)

```


### <a name="get_invoice"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.GetInvoice") GetInvoice

> Gets an invoice


```go
func (me *INVOICES_IMPL) GetInvoice(invoiceId string)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | Invoice Id |


#### Example Usage

```go
invoiceId := "invoice_id"

var result *models_pkg.GetInvoiceResponse
result,_ = invoices.GetInvoice(invoiceId)

```


### <a name="update_invoice_status"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.UpdateInvoiceStatus") UpdateInvoiceStatus

> Updates the status from an invoice


```go
func (me *INVOICES_IMPL) UpdateInvoiceStatus(
            invoiceId string,
            request *models_pkg.UpdateInvoiceStatusRequest,
            idempotencyKey *string)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | Invoice Id |
| request |  ``` Required ```  | Request for updating an invoice's status |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
invoiceId := "invoice_id"
var request *models_pkg.UpdateInvoiceStatusRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetInvoiceResponse
result,_ = invoices.UpdateInvoiceStatus(invoiceId, request, idempotencyKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="orders_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".orders_pkg") orders_pkg

### Get instance

Factory for the ``` ORDERS ``` interface can be accessed from the package orders_pkg.

```go
orders := orders_pkg.NewORDERS()
```

### <a name="get_orders"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.GetOrders") GetOrders

> Gets all orders


```go
func (me *ORDERS_IMPL) GetOrders(
            page *int64,
            size *int64,
            code *string,
            status *string,
            createdSince *time.Time,
            createdUntil *time.Time,
            customerId *string)(*models_pkg.ListOrderResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| code |  ``` Optional ```  | Filter for order's code |
| status |  ``` Optional ```  | Filter for order's status |
| createdSince |  ``` Optional ```  | Filter for order's creation date start range |
| createdUntil |  ``` Optional ```  | Filter for order's creation date end range |
| customerId |  ``` Optional ```  | Filter for order's customer id |


#### Example Usage

```go
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)
code := "code"
status := "status"
createdSince := time.Now()
createdUntil := time.Now()
customerId := "customer_id"

var result *models_pkg.ListOrderResponse
result,_ = orders.GetOrders(page, size, code, status, createdSince, createdUntil, customerId)

```


### <a name="update_order_item"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.UpdateOrderItem") UpdateOrderItem

> TODO: Add a method description


```go
func (me *ORDERS_IMPL) UpdateOrderItem(
            orderId string,
            itemId string,
            request *models_pkg.UpdateOrderItemRequest,
            idempotencyKey *string)(*models_pkg.GetOrderItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order Id |
| itemId |  ``` Required ```  | Item Id |
| request |  ``` Required ```  | Item Model |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
orderId := "orderId"
itemId := "itemId"
var request *models_pkg.UpdateOrderItemRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetOrderItemResponse
result,_ = orders.UpdateOrderItem(orderId, itemId, request, idempotencyKey)

```


### <a name="delete_all_order_items"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.DeleteAllOrderItems") DeleteAllOrderItems

> TODO: Add a method description


```go
func (me *ORDERS_IMPL) DeleteAllOrderItems(
            orderId string,
            idempotencyKey *string)(*models_pkg.GetOrderResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
orderId := "orderId"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetOrderResponse
result,_ = orders.DeleteAllOrderItems(orderId, idempotencyKey)

```


### <a name="delete_order_item"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.DeleteOrderItem") DeleteOrderItem

> TODO: Add a method description


```go
func (me *ORDERS_IMPL) DeleteOrderItem(
            orderId string,
            itemId string,
            idempotencyKey *string)(*models_pkg.GetOrderItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order Id |
| itemId |  ``` Required ```  | Item Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
orderId := "orderId"
itemId := "itemId"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetOrderItemResponse
result,_ = orders.DeleteOrderItem(orderId, itemId, idempotencyKey)

```


### <a name="close_order"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.CloseOrder") CloseOrder

> TODO: Add a method description


```go
func (me *ORDERS_IMPL) CloseOrder(
            id string,
            request *models_pkg.UpdateOrderStatusRequest,
            idempotencyKey *string)(*models_pkg.GetOrderResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | Order Id |
| request |  ``` Required ```  | Update Order Model |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
id := "id"
var request *models_pkg.UpdateOrderStatusRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetOrderResponse
result,_ = orders.CloseOrder(id, request, idempotencyKey)

```


### <a name="create_order"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.CreateOrder") CreateOrder

> Creates a new Order


```go
func (me *ORDERS_IMPL) CreateOrder(
            body *models_pkg.CreateOrderRequest,
            idempotencyKey *string)(*models_pkg.GetOrderResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Request for creating an order |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateOrderRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetOrderResponse
result,_ = orders.CreateOrder(body, idempotencyKey)

```


### <a name="create_order_item"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.CreateOrderItem") CreateOrderItem

> TODO: Add a method description


```go
func (me *ORDERS_IMPL) CreateOrderItem(
            orderId string,
            request *models_pkg.CreateOrderItemRequest,
            idempotencyKey *string)(*models_pkg.GetOrderItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order Id |
| request |  ``` Required ```  | Order Item Model |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
orderId := "orderId"
var request *models_pkg.CreateOrderItemRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetOrderItemResponse
result,_ = orders.CreateOrderItem(orderId, request, idempotencyKey)

```


### <a name="get_order_item"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.GetOrderItem") GetOrderItem

> TODO: Add a method description


```go
func (me *ORDERS_IMPL) GetOrderItem(
            orderId string,
            itemId string)(*models_pkg.GetOrderItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order Id |
| itemId |  ``` Required ```  | Item Id |


#### Example Usage

```go
orderId := "orderId"
itemId := "itemId"

var result *models_pkg.GetOrderItemResponse
result,_ = orders.GetOrderItem(orderId, itemId)

```


### <a name="update_order_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.UpdateOrderMetadata") UpdateOrderMetadata

> Updates the metadata from an order


```go
func (me *ORDERS_IMPL) UpdateOrderMetadata(
            orderId string,
            request *models_pkg.UpdateMetadataRequest,
            idempotencyKey *string)(*models_pkg.GetOrderResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | The order id |
| request |  ``` Required ```  | Request for updating the order metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
orderId := "order_id"
var request *models_pkg.UpdateMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetOrderResponse
result,_ = orders.UpdateOrderMetadata(orderId, request, idempotencyKey)

```


### <a name="get_order"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.GetOrder") GetOrder

> Gets an order


```go
func (me *ORDERS_IMPL) GetOrder(orderId string)(*models_pkg.GetOrderResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order id |


#### Example Usage

```go
orderId := "order_id"

var result *models_pkg.GetOrderResponse
result,_ = orders.GetOrder(orderId)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="customers_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".customers_pkg") customers_pkg

### Get instance

Factory for the ``` CUSTOMERS ``` interface can be accessed from the package customers_pkg.

```go
customers := customers_pkg.NewCUSTOMERS()
```

### <a name="update_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateCard") UpdateCard

> Updates a card


```go
func (me *CUSTOMERS_IMPL) UpdateCard(
            customerId string,
            cardId string,
            request *models_pkg.UpdateCardRequest,
            idempotencyKey *string)(*models_pkg.GetCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| cardId |  ``` Required ```  | Card id |
| request |  ``` Required ```  | Request for updating a card |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"
var request *models_pkg.UpdateCardRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetCardResponse
result,_ = customers.UpdateCard(customerId, cardId, request, idempotencyKey)

```


### <a name="update_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateAddress") UpdateAddress

> Updates an address


```go
func (me *CUSTOMERS_IMPL) UpdateAddress(
            customerId string,
            addressId string,
            request *models_pkg.UpdateAddressRequest,
            idempotencyKey *string)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| addressId |  ``` Required ```  | Address Id |
| request |  ``` Required ```  | Request for updating an address |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
addressId := "address_id"
var request *models_pkg.UpdateAddressRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetAddressResponse
result,_ = customers.UpdateAddress(customerId, addressId, request, idempotencyKey)

```


### <a name="delete_access_token"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteAccessToken") DeleteAccessToken

> Delete a customer's access token


```go
func (me *CUSTOMERS_IMPL) DeleteAccessToken(
            customerId string,
            tokenId string,
            idempotencyKey *string)(*models_pkg.GetAccessTokenResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| tokenId |  ``` Required ```  | Token Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
tokenId := "token_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetAccessTokenResponse
result,_ = customers.DeleteAccessToken(customerId, tokenId, idempotencyKey)

```


### <a name="create_customer"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateCustomer") CreateCustomer

> Creates a new customer


```go
func (me *CUSTOMERS_IMPL) CreateCustomer(
            request *models_pkg.CreateCustomerRequest,
            idempotencyKey *string)(*models_pkg.GetCustomerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| request |  ``` Required ```  | Request for creating a customer |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var request *models_pkg.CreateCustomerRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetCustomerResponse
result,_ = customers.CreateCustomer(request, idempotencyKey)

```


### <a name="create_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateAddress") CreateAddress

> Creates a new address for a customer


```go
func (me *CUSTOMERS_IMPL) CreateAddress(
            customerId string,
            request *models_pkg.CreateAddressRequest,
            idempotencyKey *string)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| request |  ``` Required ```  | Request for creating an address |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var request *models_pkg.CreateAddressRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetAddressResponse
result,_ = customers.CreateAddress(customerId, request, idempotencyKey)

```


### <a name="delete_access_tokens"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteAccessTokens") DeleteAccessTokens

> Delete a Customer's access tokens


```go
func (me *CUSTOMERS_IMPL) DeleteAccessTokens(customerId string)(*models_pkg.ListAccessTokensResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |


#### Example Usage

```go
customerId := "customer_id"

var result *models_pkg.ListAccessTokensResponse
result,_ = customers.DeleteAccessTokens(customerId)

```


### <a name="get_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAddress") GetAddress

> Get a customer's address


```go
func (me *CUSTOMERS_IMPL) GetAddress(
            customerId string,
            addressId string)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| addressId |  ``` Required ```  | Address Id |


#### Example Usage

```go
customerId := "customer_id"
addressId := "address_id"

var result *models_pkg.GetAddressResponse
result,_ = customers.GetAddress(customerId, addressId)

```


### <a name="delete_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteAddress") DeleteAddress

> Delete a Customer's address


```go
func (me *CUSTOMERS_IMPL) DeleteAddress(
            customerId string,
            addressId string,
            idempotencyKey *string)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| addressId |  ``` Required ```  | Address Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
addressId := "address_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetAddressResponse
result,_ = customers.DeleteAddress(customerId, addressId, idempotencyKey)

```


### <a name="create_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateCard") CreateCard

> Creates a new card for a customer


```go
func (me *CUSTOMERS_IMPL) CreateCard(
            customerId string,
            request *models_pkg.CreateCardRequest,
            idempotencyKey *string)(*models_pkg.GetCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| request |  ``` Required ```  | Request for creating a card |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var request *models_pkg.CreateCardRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetCardResponse
result,_ = customers.CreateCard(customerId, request, idempotencyKey)

```


### <a name="get_customers"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCustomers") GetCustomers

> Get all Customers


```go
func (me *CUSTOMERS_IMPL) GetCustomers(
            name *string,
            document *string,
            page *int64,
            size *int64,
            email *string,
            code *string)(*models_pkg.ListCustomersResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| name |  ``` Optional ```  | Name of the Customer |
| document |  ``` Optional ```  | Document of the Customer |
| page |  ``` Optional ```  ``` DefaultValue ```  | Current page the the search |
| size |  ``` Optional ```  ``` DefaultValue ```  | Quantity pages of the search |
| email |  ``` Optional ```  | Customer's email |
| code |  ``` Optional ```  | Customer's code |


#### Example Usage

```go
name := "name"
document := "document"
page,_ := strconv.ParseInt("1", 10, 8)
size,_ := strconv.ParseInt("10", 10, 8)
email := "email"
code := "Code"

var result *models_pkg.ListCustomersResponse
result,_ = customers.GetCustomers(name, document, page, size, email, code)

```


### <a name="update_customer"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateCustomer") UpdateCustomer

> Updates a customer


```go
func (me *CUSTOMERS_IMPL) UpdateCustomer(
            customerId string,
            request *models_pkg.UpdateCustomerRequest,
            idempotencyKey *string)(*models_pkg.GetCustomerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| request |  ``` Required ```  | Request for updating a customer |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var request *models_pkg.UpdateCustomerRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetCustomerResponse
result,_ = customers.UpdateCustomer(customerId, request, idempotencyKey)

```


### <a name="create_access_token"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateAccessToken") CreateAccessToken

> Creates a access token for a customer


```go
func (me *CUSTOMERS_IMPL) CreateAccessToken(
            customerId string,
            request *models_pkg.CreateAccessTokenRequest,
            idempotencyKey *string)(*models_pkg.GetAccessTokenResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| request |  ``` Required ```  | Request for creating a access token |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var request *models_pkg.CreateAccessTokenRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetAccessTokenResponse
result,_ = customers.CreateAccessToken(customerId, request, idempotencyKey)

```


### <a name="get_access_tokens"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAccessTokens") GetAccessTokens

> Get all access tokens from a customer


```go
func (me *CUSTOMERS_IMPL) GetAccessTokens(
            customerId string,
            page *int64,
            size *int64)(*models_pkg.ListAccessTokensResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
customerId := "customer_id"
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)

var result *models_pkg.ListAccessTokensResponse
result,_ = customers.GetAccessTokens(customerId, page, size)

```


### <a name="get_cards"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCards") GetCards

> Get all cards from a customer


```go
func (me *CUSTOMERS_IMPL) GetCards(
            customerId string,
            page *int64,
            size *int64)(*models_pkg.ListCardsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
customerId := "customer_id"
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)

var result *models_pkg.ListCardsResponse
result,_ = customers.GetCards(customerId, page, size)

```


### <a name="renew_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.RenewCard") RenewCard

> Renew a card


```go
func (me *CUSTOMERS_IMPL) RenewCard(
            customerId string,
            cardId string,
            idempotencyKey *string)(*models_pkg.GetCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| cardId |  ``` Required ```  | Card Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetCardResponse
result,_ = customers.RenewCard(customerId, cardId, idempotencyKey)

```


### <a name="get_access_token"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAccessToken") GetAccessToken

> Get a Customer's access token


```go
func (me *CUSTOMERS_IMPL) GetAccessToken(
            customerId string,
            tokenId string)(*models_pkg.GetAccessTokenResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| tokenId |  ``` Required ```  | Token Id |


#### Example Usage

```go
customerId := "customer_id"
tokenId := "token_id"

var result *models_pkg.GetAccessTokenResponse
result,_ = customers.GetAccessToken(customerId, tokenId)

```


### <a name="update_customer_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateCustomerMetadata") UpdateCustomerMetadata

> Updates the metadata a customer


```go
func (me *CUSTOMERS_IMPL) UpdateCustomerMetadata(
            customerId string,
            request *models_pkg.UpdateMetadataRequest,
            idempotencyKey *string)(*models_pkg.GetCustomerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | The customer id |
| request |  ``` Required ```  | Request for updating the customer metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var request *models_pkg.UpdateMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetCustomerResponse
result,_ = customers.UpdateCustomerMetadata(customerId, request, idempotencyKey)

```


### <a name="delete_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteCard") DeleteCard

> Delete a customer's card


```go
func (me *CUSTOMERS_IMPL) DeleteCard(
            customerId string,
            cardId string,
            idempotencyKey *string)(*models_pkg.GetCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| cardId |  ``` Required ```  | Card Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetCardResponse
result,_ = customers.DeleteCard(customerId, cardId, idempotencyKey)

```


### <a name="get_addresses"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAddresses") GetAddresses

> Gets all adressess from a customer


```go
func (me *CUSTOMERS_IMPL) GetAddresses(
            customerId string,
            page *int64,
            size *int64)(*models_pkg.ListAddressesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
customerId := "customer_id"
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)

var result *models_pkg.ListAddressesResponse
result,_ = customers.GetAddresses(customerId, page, size)

```


### <a name="get_customer"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCustomer") GetCustomer

> Get a customer


```go
func (me *CUSTOMERS_IMPL) GetCustomer(customerId string)(*models_pkg.GetCustomerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |


#### Example Usage

```go
customerId := "customer_id"

var result *models_pkg.GetCustomerResponse
result,_ = customers.GetCustomer(customerId)

```


### <a name="get_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCard") GetCard

> Get a customer's card


```go
func (me *CUSTOMERS_IMPL) GetCard(
            customerId string,
            cardId string)(*models_pkg.GetCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| cardId |  ``` Required ```  | Card id |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"

var result *models_pkg.GetCardResponse
result,_ = customers.GetCard(customerId, cardId)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="recipients_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".recipients_pkg") recipients_pkg

### Get instance

Factory for the ``` RECIPIENTS ``` interface can be accessed from the package recipients_pkg.

```go
recipients := recipients_pkg.NewRECIPIENTS()
```

### <a name="update_recipient"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.UpdateRecipient") UpdateRecipient

> Updates a recipient


```go
func (me *RECIPIENTS_IMPL) UpdateRecipient(
            recipientId string,
            request *models_pkg.UpdateRecipientRequest,
            idempotencyKey *string)(*models_pkg.GetRecipientResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| request |  ``` Required ```  | Recipient data |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var request *models_pkg.UpdateRecipientRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetRecipientResponse
result,_ = recipients.UpdateRecipient(recipientId, request, idempotencyKey)

```


### <a name="create_anticipation"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.CreateAnticipation") CreateAnticipation

> Creates an anticipation


```go
func (me *RECIPIENTS_IMPL) CreateAnticipation(
            recipientId string,
            request *models_pkg.CreateAnticipationRequest,
            idempotencyKey *string)(*models_pkg.GetAnticipationResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| request |  ``` Required ```  | Anticipation data |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var request *models_pkg.CreateAnticipationRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetAnticipationResponse
result,_ = recipients.CreateAnticipation(recipientId, request, idempotencyKey)

```


### <a name="get_anticipation_limits"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetAnticipationLimits") GetAnticipationLimits

> Gets the anticipation limits for a recipient


```go
func (me *RECIPIENTS_IMPL) GetAnticipationLimits(
            recipientId string,
            timeframe string,
            paymentDate *time.Time)(*models_pkg.GetAnticipationLimitResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| timeframe |  ``` Required ```  | Timeframe |
| paymentDate |  ``` Required ```  | Anticipation payment date |


#### Example Usage

```go
recipientId := "recipient_id"
timeframe := "timeframe"
paymentDate := time.Now()

var result *models_pkg.GetAnticipationLimitResponse
result,_ = recipients.GetAnticipationLimits(recipientId, timeframe, paymentDate)

```


### <a name="get_recipients"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetRecipients") GetRecipients

> Retrieves paginated recipients information


```go
func (me *RECIPIENTS_IMPL) GetRecipients(
            page *int64,
            size *int64)(*models_pkg.ListRecipientResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)

var result *models_pkg.ListRecipientResponse
result,_ = recipients.GetRecipients(page, size)

```


### <a name="get_withdraw_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetWithdrawById") GetWithdrawById

> TODO: Add a method description


```go
func (me *RECIPIENTS_IMPL) GetWithdrawById(
            recipientId string,
            withdrawalId string)(*models_pkg.GetWithdrawResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | TODO: Add a parameter description |
| withdrawalId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
withdrawalId := "withdrawal_id"

var result *models_pkg.GetWithdrawResponse
result,_ = recipients.GetWithdrawById(recipientId, withdrawalId)

```


### <a name="update_recipient_default_bank_account"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.UpdateRecipientDefaultBankAccount") UpdateRecipientDefaultBankAccount

> Updates the default bank account from a recipient


```go
func (me *RECIPIENTS_IMPL) UpdateRecipientDefaultBankAccount(
            recipientId string,
            request *models_pkg.UpdateRecipientBankAccountRequest,
            idempotencyKey *string)(*models_pkg.GetRecipientResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| request |  ``` Required ```  | Bank account data |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var request *models_pkg.UpdateRecipientBankAccountRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetRecipientResponse
result,_ = recipients.UpdateRecipientDefaultBankAccount(recipientId, request, idempotencyKey)

```


### <a name="update_recipient_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.UpdateRecipientMetadata") UpdateRecipientMetadata

> Updates recipient metadata


```go
func (me *RECIPIENTS_IMPL) UpdateRecipientMetadata(
            recipientId string,
            request *models_pkg.UpdateMetadataRequest,
            idempotencyKey *string)(*models_pkg.GetRecipientResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| request |  ``` Required ```  | Metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var request *models_pkg.UpdateMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetRecipientResponse
result,_ = recipients.UpdateRecipientMetadata(recipientId, request, idempotencyKey)

```


### <a name="get_transfers"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetTransfers") GetTransfers

> Gets a paginated list of transfers for the recipient


```go
func (me *RECIPIENTS_IMPL) GetTransfers(
            recipientId string,
            page *int64,
            size *int64,
            status *string,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.ListTransferResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| status |  ``` Optional ```  | Filter for transfer status |
| createdSince |  ``` Optional ```  | Filter for start range of transfer creation date |
| createdUntil |  ``` Optional ```  | Filter for end range of transfer creation date |


#### Example Usage

```go
recipientId := "recipient_id"
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)
status := "status"
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.ListTransferResponse
result,_ = recipients.GetTransfers(recipientId, page, size, status, createdSince, createdUntil)

```


### <a name="get_transfer"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetTransfer") GetTransfer

> Gets a transfer


```go
func (me *RECIPIENTS_IMPL) GetTransfer(
            recipientId string,
            transferId string)(*models_pkg.GetTransferResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| transferId |  ``` Required ```  | Transfer id |


#### Example Usage

```go
recipientId := "recipient_id"
transferId := "transfer_id"

var result *models_pkg.GetTransferResponse
result,_ = recipients.GetTransfer(recipientId, transferId)

```


### <a name="create_withdraw"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.CreateWithdraw") CreateWithdraw

> TODO: Add a method description


```go
func (me *RECIPIENTS_IMPL) CreateWithdraw(
            recipientId string,
            request *models_pkg.CreateWithdrawRequest)(*models_pkg.GetWithdrawResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | TODO: Add a parameter description |
| request |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var request *models_pkg.CreateWithdrawRequest

var result *models_pkg.GetWithdrawResponse
result,_ = recipients.CreateWithdraw(recipientId, request)

```


### <a name="update_automatic_anticipation_settings"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.UpdateAutomaticAnticipationSettings") UpdateAutomaticAnticipationSettings

> Updates recipient metadata


```go
func (me *RECIPIENTS_IMPL) UpdateAutomaticAnticipationSettings(
            recipientId string,
            request *models_pkg.UpdateAutomaticAnticipationSettingsRequest,
            idempotencyKey *string)(*models_pkg.GetRecipientResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| request |  ``` Required ```  | Metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var request *models_pkg.UpdateAutomaticAnticipationSettingsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetRecipientResponse
result,_ = recipients.UpdateAutomaticAnticipationSettings(recipientId, request, idempotencyKey)

```


### <a name="get_anticipation"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetAnticipation") GetAnticipation

> Gets an anticipation


```go
func (me *RECIPIENTS_IMPL) GetAnticipation(
            recipientId string,
            anticipationId string)(*models_pkg.GetAnticipationResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| anticipationId |  ``` Required ```  | Anticipation id |


#### Example Usage

```go
recipientId := "recipient_id"
anticipationId := "anticipation_id"

var result *models_pkg.GetAnticipationResponse
result,_ = recipients.GetAnticipation(recipientId, anticipationId)

```


### <a name="update_recipient_transfer_settings"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.UpdateRecipientTransferSettings") UpdateRecipientTransferSettings

> TODO: Add a method description


```go
func (me *RECIPIENTS_IMPL) UpdateRecipientTransferSettings(
            recipientId string,
            request *models_pkg.UpdateTransferSettingsRequest,
            idempotencyKey *string)(*models_pkg.GetRecipientResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient Identificator |
| request |  ``` Required ```  | TODO: Add a parameter description |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var request *models_pkg.UpdateTransferSettingsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetRecipientResponse
result,_ = recipients.UpdateRecipientTransferSettings(recipientId, request, idempotencyKey)

```


### <a name="get_anticipations"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetAnticipations") GetAnticipations

> Retrieves a paginated list of anticipations from a recipient


```go
func (me *RECIPIENTS_IMPL) GetAnticipations(
            recipientId string,
            page *int64,
            size *int64,
            status *string,
            timeframe *string,
            paymentDateSince *time.Time,
            paymentDateUntil *time.Time,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.ListAnticipationResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| status |  ``` Optional ```  | Filter for anticipation status |
| timeframe |  ``` Optional ```  | Filter for anticipation timeframe |
| paymentDateSince |  ``` Optional ```  | Filter for start range for anticipation payment date |
| paymentDateUntil |  ``` Optional ```  | Filter for end range for anticipation payment date |
| createdSince |  ``` Optional ```  | Filter for start range for anticipation creation date |
| createdUntil |  ``` Optional ```  | Filter for end range for anticipation creation date |


#### Example Usage

```go
recipientId := "recipient_id"
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)
status := "status"
timeframe := "timeframe"
paymentDateSince := time.Now()
paymentDateUntil := time.Now()
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.ListAnticipationResponse
result,_ = recipients.GetAnticipations(recipientId, page, size, status, timeframe, paymentDateSince, paymentDateUntil, createdSince, createdUntil)

```


### <a name="get_recipient"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetRecipient") GetRecipient

> Retrieves recipient information


```go
func (me *RECIPIENTS_IMPL) GetRecipient(recipientId string)(*models_pkg.GetRecipientResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipiend id |


#### Example Usage

```go
recipientId := "recipient_id"

var result *models_pkg.GetRecipientResponse
result,_ = recipients.GetRecipient(recipientId)

```


### <a name="get_balance"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetBalance") GetBalance

> Get balance information for a recipient


```go
func (me *RECIPIENTS_IMPL) GetBalance(recipientId string)(*models_pkg.GetBalanceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |


#### Example Usage

```go
recipientId := "recipient_id"

var result *models_pkg.GetBalanceResponse
result,_ = recipients.GetBalance(recipientId)

```


### <a name="get_withdrawals"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetWithdrawals") GetWithdrawals

> Gets a paginated list of transfers for the recipient


```go
func (me *RECIPIENTS_IMPL) GetWithdrawals(
            recipientId string,
            page *int64,
            size *int64,
            status *string,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.ListWithdrawals,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | TODO: Add a parameter description |
| page |  ``` Optional ```  | TODO: Add a parameter description |
| size |  ``` Optional ```  | TODO: Add a parameter description |
| status |  ``` Optional ```  | TODO: Add a parameter description |
| createdSince |  ``` Optional ```  | TODO: Add a parameter description |
| createdUntil |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)
status := "status"
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.ListWithdrawals
result,_ = recipients.GetWithdrawals(recipientId, page, size, status, createdSince, createdUntil)

```


### <a name="create_transfer"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.CreateTransfer") CreateTransfer

> Creates a transfer for a recipient


```go
func (me *RECIPIENTS_IMPL) CreateTransfer(
            recipientId string,
            request *models_pkg.CreateTransferRequest,
            idempotencyKey *string)(*models_pkg.GetTransferResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient Id |
| request |  ``` Required ```  | Transfer data |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var request *models_pkg.CreateTransferRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetTransferResponse
result,_ = recipients.CreateTransfer(recipientId, request, idempotencyKey)

```


### <a name="create_recipient"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.CreateRecipient") CreateRecipient

> Creates a new recipient


```go
func (me *RECIPIENTS_IMPL) CreateRecipient(
            request *models_pkg.CreateRecipientRequest,
            idempotencyKey *string)(*models_pkg.GetRecipientResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| request |  ``` Required ```  | Recipient data |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var request *models_pkg.CreateRecipientRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetRecipientResponse
result,_ = recipients.CreateRecipient(request, idempotencyKey)

```


### <a name="get_recipient_by_code"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetRecipientByCode") GetRecipientByCode

> Retrieves recipient information


```go
func (me *RECIPIENTS_IMPL) GetRecipientByCode(code string)(*models_pkg.GetRecipientResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| code |  ``` Required ```  | Recipient code |


#### Example Usage

```go
code := "code"

var result *models_pkg.GetRecipientResponse
result,_ = recipients.GetRecipientByCode(code)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="charges_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".charges_pkg") charges_pkg

### Get instance

Factory for the ``` CHARGES ``` interface can be accessed from the package charges_pkg.

```go
charges := charges_pkg.NewCHARGES()
```

### <a name="update_charge_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.UpdateChargeMetadata") UpdateChargeMetadata

> Updates the metadata from a charge


```go
func (me *CHARGES_IMPL) UpdateChargeMetadata(
            chargeId string,
            request *models_pkg.UpdateMetadataRequest,
            idempotencyKey *string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | The charge id |
| request |  ``` Required ```  | Request for updating the charge metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.UpdateMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetChargeResponse
result,_ = charges.UpdateChargeMetadata(chargeId, request, idempotencyKey)

```


### <a name="update_charge_payment_method"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.UpdateChargePaymentMethod") UpdateChargePaymentMethod

> Updates a charge's payment method


```go
func (me *CHARGES_IMPL) UpdateChargePaymentMethod(
            chargeId string,
            request *models_pkg.UpdateChargePaymentMethodRequest,
            idempotencyKey *string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| request |  ``` Required ```  | Request for updating the payment method from a charge |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.UpdateChargePaymentMethodRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetChargeResponse
result,_ = charges.UpdateChargePaymentMethod(chargeId, request, idempotencyKey)

```


### <a name="get_charge_transactions"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.GetChargeTransactions") GetChargeTransactions

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) GetChargeTransactions(
            chargeId string,
            page *int64,
            size *int64)(*models_pkg.ListChargeTransactionsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge Id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
chargeId := "charge_id"
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)

var result *models_pkg.ListChargeTransactionsResponse
result,_ = charges.GetChargeTransactions(chargeId, page, size)

```


### <a name="update_charge_due_date"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.UpdateChargeDueDate") UpdateChargeDueDate

> Updates the due date from a charge


```go
func (me *CHARGES_IMPL) UpdateChargeDueDate(
            chargeId string,
            request *models_pkg.UpdateChargeDueDateRequest,
            idempotencyKey *string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge Id |
| request |  ``` Required ```  | Request for updating the due date |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.UpdateChargeDueDateRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetChargeResponse
result,_ = charges.UpdateChargeDueDate(chargeId, request, idempotencyKey)

```


### <a name="get_charges"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.GetCharges") GetCharges

> Lists all charges


```go
func (me *CHARGES_IMPL) GetCharges(
            page *int64,
            size *int64,
            code *string,
            status *string,
            paymentMethod *string,
            customerId *string,
            orderId *string,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.ListChargesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| code |  ``` Optional ```  | Filter for charge's code |
| status |  ``` Optional ```  | Filter for charge's status |
| paymentMethod |  ``` Optional ```  | Filter for charge's payment method |
| customerId |  ``` Optional ```  | Filter for charge's customer id |
| orderId |  ``` Optional ```  | Filter for charge's order id |
| createdSince |  ``` Optional ```  | Filter for the beginning of the range for charge's creation |
| createdUntil |  ``` Optional ```  | Filter for the end of the range for charge's creation |


#### Example Usage

```go
page,_ := strconv.ParseInt("7", 10, 8)
size,_ := strconv.ParseInt("7", 10, 8)
code := "code"
status := "status"
paymentMethod := "payment_method"
customerId := "customer_id"
orderId := "order_id"
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.ListChargesResponse
result,_ = charges.GetCharges(page, size, code, status, paymentMethod, customerId, orderId, createdSince, createdUntil)

```


### <a name="capture_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.CaptureCharge") CaptureCharge

> Captures a charge


```go
func (me *CHARGES_IMPL) CaptureCharge(
            chargeId string,
            request *models_pkg.CreateCaptureChargeRequest,
            idempotencyKey *string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| request |  ``` Optional ```  | Request for capturing a charge |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.CreateCaptureChargeRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetChargeResponse
result,_ = charges.CaptureCharge(chargeId, request, idempotencyKey)

```


### <a name="update_charge_card"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.UpdateChargeCard") UpdateChargeCard

> Updates the card from a charge


```go
func (me *CHARGES_IMPL) UpdateChargeCard(
            chargeId string,
            request *models_pkg.UpdateChargeCardRequest,
            idempotencyKey *string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| request |  ``` Required ```  | Request for updating a charge's card |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.UpdateChargeCardRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetChargeResponse
result,_ = charges.UpdateChargeCard(chargeId, request, idempotencyKey)

```


### <a name="get_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.GetCharge") GetCharge

> Get a charge from its id


```go
func (me *CHARGES_IMPL) GetCharge(chargeId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |


#### Example Usage

```go
chargeId := "charge_id"

var result *models_pkg.GetChargeResponse
result,_ = charges.GetCharge(chargeId)

```


### <a name="get_charges_summary"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.GetChargesSummary") GetChargesSummary

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) GetChargesSummary(
            status string,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.GetChargesSummaryResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| status |  ``` Required ```  | TODO: Add a parameter description |
| createdSince |  ``` Optional ```  | TODO: Add a parameter description |
| createdUntil |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
status := "status"
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.GetChargesSummaryResponse
result,_ = charges.GetChargesSummary(status, createdSince, createdUntil)

```


### <a name="retry_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.RetryCharge") RetryCharge

> Retries a charge


```go
func (me *CHARGES_IMPL) RetryCharge(
            chargeId string,
            idempotencyKey *string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetChargeResponse
result,_ = charges.RetryCharge(chargeId, idempotencyKey)

```


### <a name="cancel_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.CancelCharge") CancelCharge

> Cancel a charge


```go
func (me *CHARGES_IMPL) CancelCharge(
            chargeId string,
            request *models_pkg.CreateCancelChargeRequest,
            idempotencyKey *string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| request |  ``` Optional ```  | Request for cancelling a charge |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.CreateCancelChargeRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetChargeResponse
result,_ = charges.CancelCharge(chargeId, request, idempotencyKey)

```


### <a name="create_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.CreateCharge") CreateCharge

> Creates a new charge


```go
func (me *CHARGES_IMPL) CreateCharge(
            request *models_pkg.CreateChargeRequest,
            idempotencyKey *string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| request |  ``` Required ```  | Request for creating a charge |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var request *models_pkg.CreateChargeRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetChargeResponse
result,_ = charges.CreateCharge(request, idempotencyKey)

```


### <a name="confirm_payment"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.ConfirmPayment") ConfirmPayment

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) ConfirmPayment(
            chargeId string,
            request *models_pkg.CreateConfirmPaymentRequest,
            idempotencyKey *string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | TODO: Add a parameter description |
| request |  ``` Optional ```  | Request for confirm payment |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.CreateConfirmPaymentRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetChargeResponse
result,_ = charges.ConfirmPayment(chargeId, request, idempotencyKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="transfers_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".transfers_pkg") transfers_pkg

### Get instance

Factory for the ``` TRANSFERS ``` interface can be accessed from the package transfers_pkg.

```go
transfers := transfers_pkg.NewTRANSFERS()
```

### <a name="get_transfer_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".transfers_pkg.GetTransferById") GetTransferById

> TODO: Add a method description


```go
func (me *TRANSFERS_IMPL) GetTransferById(transferId string)(*models_pkg.GetTransfer,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| transferId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
transferId := "transfer_id"

var result *models_pkg.GetTransfer
result,_ = transfers.GetTransferById(transferId)

```


### <a name="create_transfer"></a>![Method: ](https://apidocs.io/img/method.png ".transfers_pkg.CreateTransfer") CreateTransfer

> TODO: Add a method description


```go
func (me *TRANSFERS_IMPL) CreateTransfer(request *models_pkg.CreateTransfer)(*models_pkg.GetTransfer,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| request |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var request *models_pkg.CreateTransfer

var result *models_pkg.GetTransfer
result,_ = transfers.CreateTransfer(request)

```


### <a name="get_transfers"></a>![Method: ](https://apidocs.io/img/method.png ".transfers_pkg.GetTransfers") GetTransfers

> Gets all transfers


```go
func (me *TRANSFERS_IMPL) GetTransfers()(*models_pkg.ListTransfers,error)
```

#### Example Usage

```go

var result *models_pkg.ListTransfers
result,_ = transfers.GetTransfers()

```


[Back to List of Controllers](#list_of_controllers)

## <a name="tokens_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".tokens_pkg") tokens_pkg

### Get instance

Factory for the ``` TOKENS ``` interface can be accessed from the package tokens_pkg.

```go
tokens := tokens_pkg.NewTOKENS()
```

### <a name="create_token"></a>![Method: ](https://apidocs.io/img/method.png ".tokens_pkg.CreateToken") CreateToken

> *Tags:*  ``` Skips Authentication ``` 

> TODO: Add a method description


```go
func (me *TOKENS_IMPL) CreateToken(
            publicKey string,
            request *models_pkg.CreateTokenRequest,
            idempotencyKey *string)(*models_pkg.GetTokenResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| publicKey |  ``` Required ```  | Public key |
| request |  ``` Required ```  | Request for creating a token |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
publicKey := "public_key"
var request *models_pkg.CreateTokenRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetTokenResponse
result,_ = tokens.CreateToken(publicKey, request, idempotencyKey)

```


### <a name="get_token"></a>![Method: ](https://apidocs.io/img/method.png ".tokens_pkg.GetToken") GetToken

> *Tags:*  ``` Skips Authentication ``` 

> Gets a token from its id


```go
func (me *TOKENS_IMPL) GetToken(
            id string,
            publicKey string)(*models_pkg.GetTokenResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | Token id |
| publicKey |  ``` Required ```  | Public key |


#### Example Usage

```go
id := "id"
publicKey := "public_key"

var result *models_pkg.GetTokenResponse
result,_ = tokens.GetToken(id, publicKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="sellers_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".sellers_pkg") sellers_pkg

### Get instance

Factory for the ``` SELLERS ``` interface can be accessed from the package sellers_pkg.

```go
sellers := sellers_pkg.NewSELLERS()
```

### <a name="create_seller"></a>![Method: ](https://apidocs.io/img/method.png ".sellers_pkg.CreateSeller") CreateSeller

> TODO: Add a method description


```go
func (me *SELLERS_IMPL) CreateSeller(
            request *models_pkg.CreateSellerRequest,
            idempotencyKey *string)(*models_pkg.GetSellerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| request |  ``` Required ```  | Seller Model |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var request *models_pkg.CreateSellerRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSellerResponse
result,_ = sellers.CreateSeller(request, idempotencyKey)

```


### <a name="update_seller_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".sellers_pkg.UpdateSellerMetadata") UpdateSellerMetadata

> TODO: Add a method description


```go
func (me *SELLERS_IMPL) UpdateSellerMetadata(
            sellerId string,
            request *models_pkg.UpdateMetadataRequest,
            idempotencyKey *string)(*models_pkg.GetSellerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| sellerId |  ``` Required ```  | Seller Id |
| request |  ``` Required ```  | Request for updating the charge metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
sellerId := "seller_id"
var request *models_pkg.UpdateMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSellerResponse
result,_ = sellers.UpdateSellerMetadata(sellerId, request, idempotencyKey)

```


### <a name="update_seller"></a>![Method: ](https://apidocs.io/img/method.png ".sellers_pkg.UpdateSeller") UpdateSeller

> TODO: Add a method description


```go
func (me *SELLERS_IMPL) UpdateSeller(
            id string,
            request *models_pkg.UpdateSellerRequest,
            idempotencyKey *string)(*models_pkg.GetSellerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | TODO: Add a parameter description |
| request |  ``` Required ```  | Update Seller model |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
id := "id"
var request *models_pkg.UpdateSellerRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSellerResponse
result,_ = sellers.UpdateSeller(id, request, idempotencyKey)

```


### <a name="delete_seller"></a>![Method: ](https://apidocs.io/img/method.png ".sellers_pkg.DeleteSeller") DeleteSeller

> TODO: Add a method description


```go
func (me *SELLERS_IMPL) DeleteSeller(
            sellerId string,
            idempotencyKey *string)(*models_pkg.GetSellerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| sellerId |  ``` Required ```  | Seller Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
sellerId := "sellerId"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSellerResponse
result,_ = sellers.DeleteSeller(sellerId, idempotencyKey)

```


### <a name="get_seller_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".sellers_pkg.GetSellerById") GetSellerById

> TODO: Add a method description


```go
func (me *SELLERS_IMPL) GetSellerById(id string)(*models_pkg.GetSellerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | Seller Id |


#### Example Usage

```go
id := "id"

var result *models_pkg.GetSellerResponse
result,_ = sellers.GetSellerById(id)

```


### <a name="get_sellers"></a>![Method: ](https://apidocs.io/img/method.png ".sellers_pkg.GetSellers") GetSellers

> TODO: Add a method description


```go
func (me *SELLERS_IMPL) GetSellers(
            page *int64,
            size *int64,
            name *string,
            document *string,
            code *string,
            status *string,
            mtype *string,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.ListSellerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| name |  ``` Optional ```  | TODO: Add a parameter description |
| document |  ``` Optional ```  | TODO: Add a parameter description |
| code |  ``` Optional ```  | TODO: Add a parameter description |
| status |  ``` Optional ```  | TODO: Add a parameter description |
| mtype |  ``` Optional ```  | TODO: Add a parameter description |
| createdSince |  ``` Optional ```  | TODO: Add a parameter description |
| createdUntil |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("98", 10, 8)
size,_ := strconv.ParseInt("98", 10, 8)
name := "name"
document := "document"
code := "code"
status := "status"
mtype := "type"
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.ListSellerResponse
result,_ = sellers.GetSellers(page, size, name, document, code, status, mtype, createdSince, createdUntil)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="transactions_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".transactions_pkg") transactions_pkg

### Get instance

Factory for the ``` TRANSACTIONS ``` interface can be accessed from the package transactions_pkg.

```go
transactions := transactions_pkg.NewTRANSACTIONS()
```

### <a name="get_transaction"></a>![Method: ](https://apidocs.io/img/method.png ".transactions_pkg.GetTransaction") GetTransaction

> TODO: Add a method description


```go
func (me *TRANSACTIONS_IMPL) GetTransaction(transactionId string)(*models_pkg.GetTransactionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| transactionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
transactionId := "transaction_id"

var result *models_pkg.GetTransactionResponse
result,_ = transactions.GetTransaction(transactionId)

```


[Back to List of Controllers](#list_of_controllers)



