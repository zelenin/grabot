package client

import (
    "net/http"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "errors"
    "log"
    "github.com/mxmCherry/multipartbuilder"
    "strconv"
    "io"
    "regexp"
)

const baseUrl = "https://api.telegram.org"

type Client struct {
    token      string
    httpClient *http.Client
    logger     *log.Logger
}

type Option func(*Client)

func WithStdLogger(client *Client) {
    client.logger = StdLoger
}

func New(token string, options ...Option) (*Client, error) {
    if !isValidToken(token) {
        return nil, fmt.Errorf("invalid token: %s", token)
    }

    client := &Client{
        token: token,
    }

    for _, option := range options {
        option(client)
    }

    if client.httpClient == nil {
        client.httpClient = http.DefaultClient
    }

    if client.logger == nil {
        client.logger = NullLogger
    }

    return client, nil
}

type ApiResponse struct {
    Ok          bool                `json:"ok"`
    Description *string             `json:"description,omitempty"`
    Result      json.RawMessage     `json:"result,omitempty"`
    ErrorCode   *int64              `json:"error_code,omitempty"`
    Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

type ApiError struct {
    Description string              `json:"description"`
    ErrorCode   int64               `json:"error_code"`
    Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

func (error *ApiError) Error() string {
    return fmt.Sprintf("%d %s", error.ErrorCode, error.Description)
}

func newError(resp *ApiResponse) error {
    return &ApiError{
        Description: *resp.Description,
        ErrorCode:   *resp.ErrorCode,
        Parameters:  resp.Parameters,
    }
}

func (client *Client) Request(method string, params map[string]interface{}) (*ApiResponse, error) {
    uri := fmt.Sprintf("%s/bot%s/%s", baseUrl, client.token, method)

    builder := multipartbuilder.New()

    fileParams := map[string]InputFile{}
    stringParams := map[string]string{}

Loop:
    for key, rawParam := range params {
        if rawParam == nil {
            continue Loop
        }

        var stringParam string

        switch param := rawParam.(type) {
        case InputFile:
            defer param.Close()

            if param.IsStream() {
                fileParams[key] = param
            } else {
                data, _ := ioutil.ReadAll(param.GetReader())
                stringParam = string(data)
            }

        case bool:
            stringParam = strconv.FormatBool(param)

        case *bool:
            stringParam = strconv.FormatBool(*param)

        case int64:
            stringParam = strconv.FormatInt(param, 10)

        case *int64:
            stringParam = strconv.FormatInt(*param, 10)

        case float64:
            stringParam = strconv.FormatFloat(param, 'g', -1, 64)

        case *float64:
            stringParam = strconv.FormatFloat(*param, 'g', -1, 64)

        case string:
            stringParam = param

        case *string:
            stringParam = *param

        default:
            stringerParam, isStringer := param.(fmt.Stringer)
            if isStringer {
                stringParam = stringerParam.String()
            } else {
                byteParam, err := json.Marshal(param)
                if err != nil {
                    return nil, err
                }
                stringParam = string(byteParam)
            }
        }

        stringParams[key] = stringParam
    }

    for key, param := range stringParams {
        builder.AddField(key, param)
    }

    for key, file := range fileParams {
        builder.AddReader(key, file.Name(), file.GetReader())
    }

    var bodyReader io.ReadCloser
    var contentType = "application/json"

    if len(stringParams)+len(fileParams) > 0 {
        contentType, bodyReader = builder.Build()
        defer bodyReader.Close()
    }

    req, err := http.NewRequest("POST", uri, bodyReader)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", contentType)

    resp, err := client.httpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    client.logger.Print(string(data))

    var apiResponse ApiResponse

    err = json.Unmarshal(data, &apiResponse)
    if err != nil {
        return nil, errors.New(string(data))
    }

    return &apiResponse, nil
}

var tokenRegex = regexp.MustCompile(`^[\d]{3,11}:[\w-]{35}$`)

func isValidToken(token string) bool {
    return tokenRegex.MatchString(token)
}
