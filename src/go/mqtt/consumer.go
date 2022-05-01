package mqtt

/*
author Lei Wang
*/

// Sub 此示例为订阅主题函数
func Sub() {
	clientPointer := GetMqttClient()
	client := *clientPointer
	client.Subscribe("testTopic", 0x01, MessageSubHandler)
}
