package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

/*
author Lei Wang
*/

// EMQServerProtocolAddressPort protocol为协议类型 IpAddress为Mqtt服务器IP地址 Port为Mqtt服务器端口
const EMQServerProtocolAddressPort = "protocol://IpAddress:Port"

// ConnectHandler 成功连接到Mqtt服务器的回调函数
var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connect with EMQServer whose Address: '" + EMQServerProtocolAddressPort + "'")
}

// ConnectLostHandler 从Mqtt服务器断开连接的回调函数
var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

// MessagePubHandler 发布消息的回调函数
func MessagePubHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("Handle a 'Publish' message:")
	fmt.Printf("Topic: '%s' \n", msg.Topic())
	fmt.Printf("Payload: '%s' \n", msg.Payload())
	fmt.Println("---------------------------------------------------------------------------------------")
}

// MessageSubHandler 订阅消息的回调函数
func MessageSubHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("Handle a 'Subscribe' message:")
	fmt.Printf("Topic : '%s' \n", msg.Topic())
	fmt.Printf("Payload : '%s' \n", msg.Payload())
	fmt.Println("---------------------------------------------------------------------------------------")
}

func init() {
	// 配置错误提示
	mqtt.DEBUG = log.New(os.Stdout, "		[mqttDEBUG]", 0)
	mqtt.ERROR = log.New(os.Stdout, " 	[mqttERROR]", 0)
}

// GetMqttClient 获得并返回Mqtt客户端
func GetMqttClient() *mqtt.Client {
	options := mqtt.NewClientOptions().AddBroker(EMQServerProtocolAddressPort).SetClientID("Provider-Go!")
	options.SetUsername("admin")
	options.SetPassword("nchud777")
	options.SetKeepAlive(60 * time.Second)
	//duration equals '1' second
	options.SetPingTimeout(1 * time.Second)
	options.OnConnect = ConnectHandler
	options.OnConnectionLost = ConnectLostHandler
	options.DefaultPublishHandler = MessagePubHandler
	client := mqtt.NewClient(options)
	return &client
}
