package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"unsafe"
)

type Feishu struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

// GetTenantAccessToken 获取tenant_access_token
func (feishu Feishu) GetTenantAccessToken() string {
	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	marshal, _ := json.Marshal(feishu)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(marshal))
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	tokenResp := &TenantAccessTokenResponse{}
	err = json.Unmarshal(respBytes, tokenResp)
	if err != nil {
		return ""
	}
	return tokenResp.TenantAccessToken
}

func (feishu Feishu) GetUserInfo(userId string) string {
	url := "https://open.feishu.cn/open-apis/contact/v3/users/" + userId
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer([]byte("")))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", feishu.GetTenantAccessToken()))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	return *str
}

// GetUsersByDept 获取公司组织下全部用户信息
func (feishu Feishu) GetUsersByDept(deptId string) UserInfos {
	httpUrl := "https://open.feishu.cn/open-apis/contact/v3/users/find_by_department"
	params := url.Values{}
	parseURL, err := url.Parse(httpUrl)
	if err != nil {
		log.Println("err")
	}
	params.Set("department_id", deptId)
	//如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	req, _ := http.NewRequest("GET", urlPathWithParams, bytes.NewBuffer([]byte("")))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", feishu.GetTenantAccessToken()))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	userInfos := &UserInfos{}
	err = json.Unmarshal(respBytes, userInfos)
	if err != nil {
	}
	return *userInfos
}

func (feishu Feishu) SendAlertMessage(msgType string, chatID string) error {
	var err error
	var createResp *MessageItem
	var createReq *CreateMessageRequest
	switch msgType {
	case "text":
		content := "{\"text\":\"<at user_id=\\\"all\\\">所有人</at> 请注意，线上服务发生报警，请及时处理。 \\n服务负责人：<at user_id=\\\"ou_ba44c2d64d161c0f12d8548bef215311\\\">张三</at> \"}"
		createReq = feishu.genCreateMessageRequest(chatID, content, msgType)

	case "post":
		content := "{\"zh_cn\":{\"title\":\"线上服务报警通知！\",\"content\":[[{\"tag\":\"at\",\"user_id\":\"all\",\"user_name\":\"所有人\"},{\"tag\":\"text\",\"text\":\"请注意，线上服务发生报警，请及时处理。\"}],[{\"tag\":\"text\",\"text\":\"服务负责人：\"},{\"tag\":\"at\",\"user_id\":\"ou_ba44c2d64d161c0f12d8548bef215311\",\"user_name\":\"张三\"}]]}}"
		createReq = feishu.genCreateMessageRequest(chatID, content, msgType)
	default:
		return nil
	}
	createResp, err = feishu.SendMessage(createReq)
	if err != nil {
		logrus.WithError(err).Errorf("send %v message failed, chat_id: %v", msgType, chatID)
		return err
	}
	msgID := createResp.MessageID
	logrus.Infof("succeed send alert message, msg_id: %v", msgID)
	return nil
}

// SendMessage 发送消息
func (feishu Feishu) SendMessage(createReq *CreateMessageRequest) (*MessageItem, error) {
	createMessageURL := "https://open.feishu.cn/open-apis/im/v1/messages"
	var err error
	token := feishu.GetTenantAccessToken()
	cli := &http.Client{}
	reqBytes, err := json.Marshal(createReq)
	if err != nil {
		logrus.WithError(err).Errorf("failed to marshal")
		return nil, err
	}
	req, err := http.NewRequest("POST", createMessageURL, strings.NewReader(string(reqBytes)))
	if err != nil {
		logrus.WithError(err).Errorf("new request failed")
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	q := req.URL.Query()
	q.Add("receive_id_type", "open_id")
	req.URL.RawQuery = q.Encode()
	var logID string
	resp, err := cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("create message failed, err=%v", err)
	}
	if resp != nil && resp.Header != nil {
		logID = resp.Header.Get("x-tt-logid")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Error("read body failed")
		return nil, err
	}

	createMessageResp := &CreateMessageResponse{}
	err = json.Unmarshal(body, createMessageResp)
	if err != nil {
		logrus.WithError(err).Errorf("failed to unmarshal")
		return nil, err
	}
	if createMessageResp.Code != 0 {
		logrus.Warnf("failed to create message, code: %v, msg: %v, log_id: %v", createMessageResp.Code, createMessageResp.Message, logID)
		return nil, fmt.Errorf("create message failed")
	}
	logrus.Infof("succeed create message, msg_id: %v", createMessageResp.Data.MessageID)
	return createMessageResp.Data, nil
}

//构建消息对象
func (feishu Feishu) genCreateMessageRequest(chatID, content, msgType string) *CreateMessageRequest {
	return &CreateMessageRequest{
		ReceiveID: chatID,
		Content:   content,
		MsgType:   msgType,
	}
}
