# API Documentation

### Message: QueryParameter
| Field Name | Type |
|------------|------|
| name | string* |
| type | string* |
| description | string* |
| required | bool* |

### Message: APIOptions
| Field Name | Type |
|------------|------|
| path | string* |
| method | string* |
| summary | string* |
| description | string* |
| query_params | #.QueryParameter[] |

### Message: User
| Field Name | Type |
|------------|------|
| name | string* |
| email | string* |
| age | int32* |

### Message: GetUserRequest
| Field Name | Type |
|------------|------|
| id | int32* |

### Message: GetUserResponse
| Field Name | Type |
|------------|------|
| user | #User* |

### Service: UserService
### Method: GetUser
Endpoint: /UserService/GetUser
- **Input Type:** #GetUserRequest
- **Output Type:** #GetUserResponse

### Message: Customer
| Field Name | Type |
|------------|------|
| name | string* |
| email | string* |
| age | int32* |

### Message: GetCustomerRequest
| Field Name | Type |
|------------|------|
| id | int32* |

### Message: GetCustomerResponse
| Field Name | Type |
|------------|------|
| user | #Customer[] |

### Service: CustomerService
### Method: GetCustomer
Endpoint: /CustomerService/GetCustomer
- **Input Type:** #GetCustomerRequest
- **Output Type:** #GetCustomerResponse


GENERATED CODE WITH ❤️ BY lovelyoyrmia

PLEASE MODIFY THIS FILE AND RUN ON YOUR SERVER.

