package server

import (
	"bytes"
	. "contracts"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"utility"
	"utility/httpsig"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type httpServerContext struct {
	underlyingContext *fiber.Ctx
	request           IRequest
	response          IResponse
	stringUtil        IStringUtil
	httpClient        *http.Client
}

func newContext(underlyingContext *fiber.Ctx) IContext {
	return &httpServerContext{
		underlyingContext: underlyingContext,
		request:           newRequest(underlyingContext),
		response:          newResponse(underlyingContext),
		stringUtil:        utility.NewStringUtil(),
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (context *httpServerContext) GUID() string {
	return uuid.New().String()
}

func (context *httpServerContext) WriteString(data any) error {
	return context.underlyingContext.SendString(data.(string))
}

func (context *httpServerContext) JSON(data interface{}) error {
	return context.underlyingContext.JSON(data)
}

func (context *httpServerContext) Nothing() error {
	return context.underlyingContext.JSON(&struct{}{})
}

func (context *httpServerContext) Redirect(location string, status ...int) error {
	return context.underlyingContext.Redirect(location, status...)
}

func (context *httpServerContext) Render(name string, bind interface{}, layouts ...string) error {
	return context.underlyingContext.Render(name, bind, layouts...)
}

func (context *httpServerContext) Config() IConfig {
	return nil
}

func (context *httpServerContext) Storage() IStorage {
	return nil
}

func (context *httpServerContext) Request() IRequest {
	return context.request
}

func (context *httpServerContext) Response() IResponse {
	return context.response
}

func (context *httpServerContext) StringUtil() IStringUtil {
	return context.stringUtil
}

func (context *httpServerContext) ParseJson(interface{}) IResult {
	return nil
}

func (context *httpServerContext) ParseBodyAndValidate(body interface{}) error {
	if err := context.underlyingContext.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return utility.Validate(body)
}

func (context *httpServerContext) GetUser() uint {
	id, _ := context.underlyingContext.Locals("USER").(uint)
	return id
}

// Error create a server error with status code and message
func (context *httpServerContext) Error(code int, format string, args ...any) IServerError {
	return newError(code, fmt.Sprintf(format, args...))
}

func (context *httpServerContext) BadRequest(format string, args ...any) IServerError {
	return context.Error(StatusBadRequest, format, args...)
}

func (context *httpServerContext) NotFound(format string, args ...any) IServerError {
	return context.Error(StatusNotFound, format, args...)
}

func (context *httpServerContext) InternalServerError(format string, args ...any) IServerError {
	return context.Error(StatusInternalServerError, format, args...)
}

func (context *httpServerContext) Unauthorized(format string, args ...any) IServerError {
	return context.Error(StatusUnauthorized, format, args...)
}

func (context *httpServerContext) Conflict(format string, args ...any) IServerError {
	return context.Error(StatusConflict, format, args...)
}

func (context *httpServerContext) GetActivityStream(url string, keyId, privateKey string, output interface{}) error {
	privKey, err := httpsig.ParseRsaPrivateKeyFromPemStr(privateKey)
	if err != nil {
		return err
	}

	signer := httpsig.NewRSASHA256Signer(keyId, privKey, []string{"Date", "Digest"})

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/activity+json")

	if err := signer.Sign(req); err != nil {
		return err
	}

	res, err := context.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", res.Status)
	}

	if err := json.NewDecoder(res.Body).Decode(output); err != nil {
		return err
	}

	return nil
}

func (context *httpServerContext) PostActivityStream(url, keyId, privateKey string, data []byte, output interface{}) error {
	privKey, err := httpsig.ParseRsaPrivateKeyFromPemStr(privateKey)
	if err != nil {
		return err
	}

	signer := httpsig.NewRSASHA256Signer(keyId, privKey, []string{"Date", "Digest"})

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/activity+json; charset=UTF-8")
	req.Header.Set("Accept", "application/activity+json")

	if err := signer.Sign(req); err != nil {
		return err
	}

	res, err := context.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", res.Status)
	}

	return nil
}
