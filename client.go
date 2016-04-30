package push

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const baseURL = "https://api.xmpush.xiaomi.com/v3/message/"

// Client is used to send push request.
type Client struct {
	AppSecret string
}

// NewClient function construct a new Client.
func NewClient(appSecret string) *Client {
	c := new(Client)
	c.AppSecret = appSecret

	return c
}

func addIfNotEmpty(data *url.Values, key, value string) {
	if len(value) > 0 {
		data.Add(key, value)
	}
}

func fillData(r *Request) string {
	data := url.Values{}
	data.Add("title", r.Title)
	data.Add("description", r.Description)
	data.Add("payload", r.Payload)
	addIfNotEmpty(&data, "restricted_package_name", r.RestrictedPackageName)
	data.Add("passthrough", strconv.Itoa(r.PassThrough))
	data.Add("notify_type", strconv.Itoa(r.NotifyType))

	switch r.TargetType {
	case "regid":
		data.Add("registration_id", r.TargetValue)
	case "alias":
		data.Add("alias", r.TargetValue)
	case "topic":
		data.Add("topic", r.TargetValue)
	case "multi_topic":
		data.Add("topics", r.TargetValue)
	case "user_account":
		data.Add("user_account", r.TargetValue)
	}
	return data.Encode()
}

// SendRequest method sends out a PushRequest.
func (c Client) SendRequest(r *Request) (string, error) {
	// http://dev.xiaomi.com/doc/?p=533
	client := &http.Client{}

	data := fillData(r)
	req, err := http.NewRequest("POST", baseURL+r.TargetType, strings.NewReader(data))
	req.Header.Add("Authorization", "key="+c.AppSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
