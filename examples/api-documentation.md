# API Documentation

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
Endpoint: /myapi/mymethod
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
Endpoint: 
- **Input Type:** #GetCustomerRequest
- **Output Type:** #GetCustomerResponse


GENERATED CODE WITH ❤️ BY lovelyoyrmia

