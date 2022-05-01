package mqtt

/*
author Lei Wang
*/

// Pub 此示例为发布消息函数
func Pub() {
	clientPointer := GetMqttClient()
	client := *clientPointer
	client.Publish("testTopic", 0x01, false, "testPayload")
}
