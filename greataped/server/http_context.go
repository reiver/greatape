package server

import (
	"activitypub"
	"app/models/domain"
	"bytes"
	"config"
	. "contracts"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path"
	"server/mime"
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

func (context *httpServerContext) String(payload any) error {
	return context.underlyingContext.SendString(payload.(string))
}

func (context *httpServerContext) Json(payload interface{}) error {
	return context.underlyingContext.JSON(payload)
}

func (context *httpServerContext) Activity(payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	context.underlyingContext.Response().Header.Add("Content-Type", mime.ActivityJsonUtf8)
	return context.underlyingContext.Send(data)
}

func (context *httpServerContext) File(key string) error {
	// TODO: Compress the response
	filePath := path.Join(config.UPLOAD_PATH, key)
	return context.underlyingContext.SendFile(filePath)
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

func (context *httpServerContext) SaveFile(f *multipart.FileHeader, path string) error {
	return context.underlyingContext.SaveFile(f, path)
}

func (context *httpServerContext) GetUser() uint {
	id, _ := context.underlyingContext.Locals("USER").(uint)
	return id
}

// Error create a server error with status code and message
func (context *httpServerContext) Error(code int, format interface{}, args ...any) IServerError {
	switch arg := format.(type) {
	case string:
		return newError(code, fmt.Sprintf(arg, args...))
	case error:
		return newError(code, arg.Error())
	default:
		return newError(code, fmt.Sprintf("%v", arg))
	}
}

func (context *httpServerContext) BadRequest(format interface{}, args ...any) IServerError {
	return context.Error(StatusBadRequest, format, args...)
}

func (context *httpServerContext) NotFound(format interface{}, args ...any) IServerError {
	return context.Error(StatusNotFound, format, args...)
}

func (context *httpServerContext) InternalServerError(format interface{}, args ...any) IServerError {
	return context.Error(StatusInternalServerError, format, args...)
}

func (context *httpServerContext) Unauthorized(format interface{}, args ...any) IServerError {
	return context.Error(StatusUnauthorized, format, args...)
}

func (context *httpServerContext) Conflict(format interface{}, args ...any) IServerError {
	return context.Error(StatusConflict, format, args...)
}

func (context *httpServerContext) signRequest(keyId, privateKey string, data []byte, req *http.Request) error {
	privKey, err := httpsig.ParseRsaPrivateKeyFromPemStr(privateKey)
	if err != nil {
		return err
	}

	signer := httpsig.NewRSASHA256Signer(keyId, privKey, []string{"Date", "Digest"})
	if data != nil {
		hasher := sha256.New()
		hasher.Write(data)
		sum := hasher.Sum(nil)
		encodedHash := base64.StdEncoding.EncodeToString(sum)
		digest := fmt.Sprintf("sha-256=%s", encodedHash)
		req.Header.Set("Content-Type", mime.ActivityJsonUtf8)
		req.Header.Set("Digest", digest)
	}

	if err := signer.Sign(req); err != nil {
		return err
	}

	return nil
}

func (context *httpServerContext) requestActivityStream(method, url, keyId, privateKey string, data []byte, output interface{}) error {
	var reader io.Reader
	if data != nil {
		reader = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", mime.ActivityJson)

	if privateKey != "" {
		if err := context.signRequest(keyId, privateKey, data, req); err != nil {
			return err
		}
	}

	res, err := context.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != http.StatusOK &&
		res.StatusCode != http.StatusAccepted {
		return fmt.Errorf("%s", res.Status)
	}

	if output != nil {
		if err := json.NewDecoder(res.Body).Decode(output); err != nil {
			return err
		}
	}

	return nil
}

func (context *httpServerContext) GetActivityStream(url string, data []byte, output interface{}) error {
	return context.requestActivityStream(http.MethodGet, url, "", "", data, output)
}

func (context *httpServerContext) PostActivityStream(url string, data []byte, output interface{}) error {
	return context.requestActivityStream(http.MethodPost, url, "", "", data, output)
}

func (context *httpServerContext) GetActivityStreamSigned(url, keyId, privateKey string, data []byte, output interface{}) error {
	return context.requestActivityStream(http.MethodGet, url, keyId, privateKey, data, output)
}

func (context *httpServerContext) PostActivityStreamSigned(url, keyId, privateKey string, data []byte, output interface{}) error {
	return context.requestActivityStream(http.MethodPost, url, keyId, privateKey, data, output)
}

func (context *httpServerContext) GetWebFinger(username domain.Username) (activitypub.Webfinger, error) {
	result, err := context.GetObject(username.Webfinger(), &activitypub.Webfinger{})
	return result.(activitypub.Webfinger), err
}

func (context *httpServerContext) GetActor(webfinger activitypub.Webfinger) (activitypub.Actor, error) {
	result, err := context.GetObject(webfinger.Self(), &activitypub.Actor{})
	return result.(activitypub.Actor), err
}

func (context *httpServerContext) GetOrderedCollection(url string) (activitypub.OrderedCollection, error) {
	result, err := context.GetObject(url, &activitypub.OrderedCollection{})
	return result.(activitypub.OrderedCollection), err
}

func (context *httpServerContext) GetObject(url string, result interface{}) (interface{}, error) {
	if err := context.GetActivityStream(url, nil, &result); err != nil {
		return result, err
	}

	return result, nil
}
