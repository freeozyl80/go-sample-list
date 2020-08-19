package test

import (
    "testing"
    "pattern/builder/msg"
)
func TestMessageBuilder(t *testing.T) {
  // 使用消息建造者进行对象创建
    message := msg.Builder().
        WithSrcAddr("192.168.0.1").
        WithSrcPort(1234).
        WithDestAddr("192.168.0.2").
        WithDestPort(8080).
        WithHeaderItem("contents", "application/json").
        WithBodyItem("record1").
        WithBodyItem("record2").
        Build()
    if message.Header.SrcAddr != "192.168.0.1" {
        t.Errorf("expect src address 192.168.0.1, but actual %s.", message.Header.SrcAddr)
    }
    if message.Body.Items[0] != "record1" {
        t.Errorf("expect body item0 record1, but actual %s.", message.Body.Items[0])
    }
}