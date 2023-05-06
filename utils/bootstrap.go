package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type App struct {
	Validator *Validator
}

type Validator struct {
	Driver     *validator.Validate
	Uni        *ut.UniversalTranslator
	Translator ut.Translator
}

const (
	// MsgSuccess ...
	MsgSuccess = "APP:SUCCESS"

	// MsgErrValidation ...
	MsgErrValidation = "ERR:VALIDATION"

	// MsgEmptyData Data not found ...
	MsgEmptyData = "ERR:EMPTY_DATA"

	// MsgErrParam error parameter argument or anything in query string
	MsgErrParam = "ERR:INVALID_PARAM"

	// MsgBadReq for general bad request
	MsgBadReq = "ERR:BAD_REQUEST"

	// MsgNotfound for not found 404 page
	MsgNotfound = "ERR:NOT_FOUND"

	// MsgAuthErr ..
	MsgAuthErr = "ERR:AUTHENTICATION"

	MsgAuthorizedErr = "ERR:AUTHORIZED"

	MsgForbiddenErr = "ERR:FORBIDDEN"

	// XChannelHeader custom header for determine what the channel is
	XChannelHeader = "X-Channel"

	AuthHeader = "Authorization"

	// AuthBase64Error flag for error base64
	AuthBase64Error = "[base64:Invalid]"
)

// ErrorBase64 give error string of invalid base64
func (h *App) ErrorBase64() error {
	return fmt.Errorf(AuthBase64Error)
}

// EmptyJSONArr ...
func (h *App) EmptyJSONArr() []map[string]interface{} {
	return []map[string]interface{}{}
}

// SendSuccess send success into response with 200 http code.
func (h *App) SendSuccess(w http.ResponseWriter, payload interface{}, pagination interface{}) {
	if pagination == nil {
		pagination = h.EmptyJSONArr()
	}
	h.RespondWithJSON(w, 200, MsgSuccess, payload)
}

// SendBadRequest send bad request into response with 400 http code.
func (h *App) SendBadRequest(w http.ResponseWriter, message string) {
	h.RespondWithJSON(w, 400, message, h.EmptyJSONArr())
}

// SendNotfound send bad request into response with 400 http code.
func (h *App) SendNotfound(w http.ResponseWriter, message string) {
	h.RespondWithJSON(w, 404, message, h.EmptyJSONArr())
}

// SendAuthError send bad request into response with 400 http code.
func (h *App) SendInternalServerErr(w http.ResponseWriter, message string) {
	h.RespondWithJSON(w, 502, message, h.EmptyJSONArr())
}

// SendRequestValidationError Send validation error response to consumers.
func (h *App) SendRequestValidationError(w http.ResponseWriter, validationErrors validator.ValidationErrors) {
	errorResponse := map[string][]string{}
	errorTranslation := validationErrors.Translate(h.Validator.Translator)
	// fmt.Println(errorTranslation)
	// fmt.Println(validationErrors)
	for _, err := range validationErrors {
		errKey := Underscore(err.StructField())
		errorResponse[errKey] = append(
			errorResponse[errKey],
			strings.Replace(errorTranslation[err.Namespace()], err.StructField(), errKey, -1),
		)
	}

	h.RespondWithJSON(w, 400, MsgErrValidation, errorResponse)
}

func (h *App) SendBindAndValidateError(w http.ResponseWriter, err interface{}) {
	if _, ok := err.(validator.ValidationErrors); ok {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
	} else if _, ok := err.(error); ok {
		h.SendBadRequest(w, err.(error).Error())
	} else {
		log.Fatal("wrong type of (err interface{}) value. the value must instance of (validator.ValidationErrors) or (error)")
	}
}

// RespondWithJSON write json response format
func (h *App) RespondWithJSON(w http.ResponseWriter, httpCode int, message string, payload interface{}) {
	respPayload := map[string]interface{}{
		"message": message,
		"data":    payload,
	}

	response, _ := json.Marshal(respPayload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	_, _ = w.Write(response)
}

// ParamOrder ...
type ParamOrder struct {
	Field string
	By    string
}

// GetIntParam Parse the url param to get value as integer.
// for example, we need to get limit and offset param
func (h *App) GetIntParam(r *http.Request, name string) (int, error) {
	param := r.URL.Query().Get(name)
	if len(param) == 0 {
		return 0, nil
	}

	return strconv.Atoi(param)
}

// GetStringParam Parse the url param to get value as string.
func (h *App) GetStringParam(r *http.Request, name string) (string, error) {
	param := r.URL.Query().Get(name)
	if len(param) == 0 {
		return "", nil
	}

	return param, nil
}

// GetBoolParam Parse the url param to get value as boolean
func (h *App) GetBoolParam(r *http.Request, name string) (bool, error) {
	param := r.URL.Query().Get(name)
	if len(param) == 0 {
		return false, nil
	}

	result, err := strconv.ParseBool(param)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (h *App) PingAction(w http.ResponseWriter, r *http.Request) {
	h.SendSuccess(w, h.EmptyJSONArr(), nil)
}
