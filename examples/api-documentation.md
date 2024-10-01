# API Documentation

**Author**:   
**Base URL**: ``

---

## Table of Contents

- [Services](#services)

  - [UserService](#UserService)
  
    - [GetUser](#GetUser)
  

  - [CustomerService](#CustomerService)
  
    - [GetCustomer](#GetCustomer)
  


- [Messages](#messages)

  - [User](#User)

  - [GetUserRequest](#GetUserRequest)

  - [GetUserResponse](#GetUserResponse)

  - [Customer](#Customer)

  - [GetCustomerRequest](#GetCustomerRequest)

  - [GetCustomerResponse](#GetCustomerResponse)


---

<a name="services"></a>
## Services


<a name="UserService"></a>
### UserService


#### GetUser

- **HTTP Method**: `POST`
- **Endpoint**: `/myapi/mymethod`
- **Summary**: Get User
- **Description**: Get User Description
- **Input Type**: [#GetUserRequest](#GetUserRequest)
- **Output Type**: [#GetUserResponse](#GetUserResponse)


##### Query Parameters

| **Name** | **Type** | **Required** | **Description** |
| -------- | -------- | ------------ | --------------- |

| `id` | `int` | Yes | The ID of the item to fetch. |




---


---

<a name="CustomerService"></a>
### CustomerService


#### GetCustomer

- **HTTP Method**: `POST`
- **Endpoint**: `/myapi/mymethod`
- **Summary**: Get Customer
- **Description**: Get Customer Description
- **Input Type**: [#GetCustomerRequest](#GetCustomerRequest)
- **Output Type**: [#GetCustomerResponse](#GetCustomerResponse)


##### Query Parameters

| **Name** | **Type** | **Required** | **Description** |
| -------- | -------- | ------------ | --------------- |

| `id` | `int` | Yes | The ID of the item to fetch. |




---


---


---

<a name="messages"></a>
## Messages


<a name="User"></a>
### User

#### Fields

| **Field Name** | **Type** |
| -------------- | -------- |

| `name` | `string*` |

| `email` | `string*` |

| `age` | `int32*` |


---

<a name="GetUserRequest"></a>
### GetUserRequest

#### Fields

| **Field Name** | **Type** |
| -------------- | -------- |

| `id` | `int32*` |


---

<a name="GetUserResponse"></a>
### GetUserResponse

#### Fields

| **Field Name** | **Type** |
| -------------- | -------- |

| `user` | `#User*` |


---

<a name="Customer"></a>
### Customer

#### Fields

| **Field Name** | **Type** |
| -------------- | -------- |

| `name` | `string*` |

| `email` | `string*` |

| `age` | `int32*` |


---

<a name="GetCustomerRequest"></a>
### GetCustomerRequest

#### Fields

| **Field Name** | **Type** |
| -------------- | -------- |

| `id` | `int32*` |


---

<a name="GetCustomerResponse"></a>
### GetCustomerResponse

#### Fields

| **Field Name** | **Type** |
| -------------- | -------- |

| `user` | `#Customer[]` |


---


---

**Generated with ❤️ by Protodoc**

Reach out : [lovelyoyrmia.com](https://lovelyoyrmia.com)