package aliyun

import (
	"TicketSales/pkg/cache"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"math/rand"
	"strconv"
	"time"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func _main(mobile, code string) (_result *dysmsapi20170525.SendSmsResponse, _err error) {
	var client *dysmsapi20170525.Client
	client, _err = createClient(tea.String("key"), tea.String("secret"))
	if _err != nil {
		return
	}
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("阿里云短信测试"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String(mobile),
		TemplateParam: tea.String(`{"code":` + code + `}`),
	}

	// 复制代码运行请自行打印 API 的返回值
	_result, _err = client.SendSms(sendSmsRequest)
	//fmt.Println(_result)
	if _err != nil {
		return
	}
	return
}

func SendMobileCaptcha(mobile string) error {
	captcha := code()
	resp, err := _main(mobile, captcha)
	if err != nil {
		return err
	}
	if *resp.Body.Code != "OK" {
		return errors.New(*resp.Body.Message)
	}

	err = cache.Store.Set(mobile, captcha, 300)
	if err != nil {
		return err
	}
	return nil
}

func CheckMobileCaptcha(mobile, captcha string) (bool, error) {
	val, flag := cache.Store.Get(mobile)
	if !flag {
		return flag, errors.New("验证码失效")
	} else if val.(string) != captcha {
		return flag, errors.New("验证码错误")
	}
	cache.Store.Delete(mobile)
	return flag, nil
}

func code() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(899999) + 100000
	res := strconv.Itoa(code) //转字符串返回
	return res
}
