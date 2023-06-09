package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	flogging "github.com/jhyehuang/fil-cmd/intern/log"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
)

type httpClient interface {
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
	Do(req *http.Request) (*http.Response, error)
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err Error) Error() string {
	return fmt.Sprintf("Error %d (%s)", err.Code, err.Message)
}

type Response struct {
	ID      int             `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *Error          `json:"error"`
}

type Request struct {
	ID      int           `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type HTTPRpc struct {
	Url    string
	Token  string
	Client httpClient
	Log    *zap.SugaredLogger
}

func NewRPC(url string, token string, options ...func(rpc *HTTPRpc)) *HTTPRpc {
	rpc := &HTTPRpc{
		Url:    url,
		Token:  token,
		Client: http.DefaultClient,
		Log:    flogging.Logger,
	}
	for _, option := range options {
		option(rpc)
	}

	return rpc
}

func (rpc *HTTPRpc) URL() string {
	return rpc.Url
}

func (rpc *HTTPRpc) CallWithResult(method string, target interface{}, params ...interface{}) error {
	result, err := rpc.Call(method, params...)
	if err != nil {
		return err
	}

	if target == nil {
		return nil
	}

	return json.Unmarshal(result, target)
}

func (rpc *HTTPRpc) Call(method string, params ...interface{}) (json.RawMessage, error) {
	request := Request{
		ID:      1,
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := rpc.Client.Post(rpc.Url, "application/json", bytes.NewBuffer(body))
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	fmt.Printf("response status: %+v\n", response.Status)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resp := new(Response)
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, *resp.Error
	}

	return resp.Result, nil

}

func (rpc *HTTPRpc) CallDo(method string, params ...interface{}) (json.RawMessage, error) {
	request := Request{
		ID:      1,
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	body, err := json.Marshal(request)
	if err != nil {
		log.Errorf("err: %v\n", err)
		return nil, err
	}
	bearer := "Bearer " + rpc.Token
	req, err := http.NewRequest("POST", rpc.Url, bytes.NewBuffer(body))
	if err != nil {
		log.Errorf("err: %v\n", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	response, err := rpc.Client.Do(req)
	if response != nil {
		defer response.Body.Close()
	}
	log.Warnf("response: %+v\n", response)
	if err != nil {
		log.Errorf("err: %v\n", err)
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Infof("data: %v\n", string(data))
	resp := new(Response)
	if err := json.Unmarshal(data, resp); err != nil {
		log.Errorf("Unmarshal err: %v\n", err)
		return nil, err
	}

	if resp.Error != nil {
		log.Errorf("resp err: %v\n", resp.Error)
		return nil, *resp.Error
	}

	return resp.Result, nil

}

func (rpc *HTTPRpc) CallDoBody(method string, body []byte) (json.RawMessage, error) {

	bearer := "Bearer " + rpc.Token
	req, err := http.NewRequest("POST", rpc.Url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("POST err: %v\n", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	response, err := rpc.Client.Do(req)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		fmt.Printf("Do err: %v\n", err)
		return nil, err
	}
	//fmt.Printf("response: %v\n", response)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", response.StatusCode)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ReadAll err: %v\n", err)
		return nil, err
	}

	resp := new(Response)
	if err := json.Unmarshal(data, resp); err != nil {
		fmt.Printf("Unmarshal err: %v\n", err)
		return nil, err
	}

	if resp.Error != nil {
		return nil, *resp.Error
	}

	return resp.Result, nil

}
