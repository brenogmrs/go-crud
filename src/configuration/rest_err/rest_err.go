package rest_err

import "net/http"

type RestErr struct {
  Message   string `json:"message"`
  Err       string `json:"error"`
  Code      int64 `json:"code"`
  Causes    []Causes `json:"causes"`
}

type Causes struct {
  Field     string `json:"field"`
  Message   string `json:"message"`
}

func (r *RestErr) Error() string {

  return r.Message
}


func NewRestErr(message, err string, code int64, causes []Causes) *RestErr {

  return &RestErr{
    Message: message,
    Err: err,
    Code: code,
    Causes: causes,
  }
}

func NewBadRequestError(message string) *RestErr{
   return &RestErr{
    Message: message,
    Err: "BAD_REQUEST",
    Code: http.StatusBadRequest,
  }
}


func NewBadRequestValidationError(message string, causes []Causes) *RestErr{
  return &RestErr{
   Message: message,
   Err: "BAD_REQUEST",
   Code: http.StatusBadRequest,
   Causes: causes,
 }
}

func NewInternalServerError(message string) *RestErr{
  return &RestErr{
   Message: message,
   Err: "INTERNAL_SERVER_ERROR",
   Code: http.StatusInternalServerError,
 }
}

func NewNotFoundError(message string, causes []Causes) *RestErr{
  return &RestErr{
   Message: message,
   Err: "INTERNAL_SERVER_ERROR",
   Code: http.StatusNotFound,
 }
}

func NewForbiddenError(message string, causes []Causes) *RestErr{
  return &RestErr{
   Message: message,
   Err: "INTERNAL_SERVER_ERROR",
   Code: http.StatusForbidden,
 }
}