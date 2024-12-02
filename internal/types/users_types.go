package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateUsersRequest request params
type CreateUsersRequest struct {
	Name         string `json:"name" binding:""`
	Email        string `json:"email" binding:""`
	MobileNumber string `json:"mobileNumber" binding:""`
	Age          int    `json:"age" binding:""`
}

// UpdateUsersByIDRequest request params
type UpdateUsersByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	Name         string `json:"name" binding:""`
	Email        string `json:"email" binding:""`
	MobileNumber string `json:"mobileNumber" binding:""`
	Age          int    `json:"age" binding:""`
}

// UsersObjDetail detail
type UsersObjDetail struct {
	ID uint64 `json:"id"` // convert to uint64 id

	Name         string    `json:"name"`
	Email        string    `json:"email"`
	MobileNumber string    `json:"mobileNumber"`
	Age          int       `json:"age"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// CreateUsersReply only for api docs
type CreateUsersReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateUsersByIDReply only for api docs
type UpdateUsersByIDReply struct {
	Result
}

// GetUsersByIDReply only for api docs
type GetUsersByIDReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Users UsersObjDetail `json:"users"`
	} `json:"data"` // return data
}

// DeleteUsersByIDReply only for api docs
type DeleteUsersByIDReply struct {
	Result
}

// DeleteUserssByIDsReply only for api docs
type DeleteUserssByIDsReply struct {
	Result
}

// ListUserssRequest request params
type ListUserssRequest struct {
	query.Params
}

// ListUserssReply only for api docs
type ListUserssReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Userss []UsersObjDetail `json:"userss"`
	} `json:"data"` // return data
}

// DeleteUserssByIDsRequest request params
type DeleteUserssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// GetUsersByConditionRequest request params
type GetUsersByConditionRequest struct {
	query.Conditions
}

// GetUsersByConditionReply only for api docs
type GetUsersByConditionReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Users UsersObjDetail `json:"users"`
	} `json:"data"` // return data
}

// ListUserssByIDsRequest request params
type ListUserssByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListUserssByIDsReply only for api docs
type ListUserssByIDsReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Userss []UsersObjDetail `json:"userss"`
	} `json:"data"` // return data
}
