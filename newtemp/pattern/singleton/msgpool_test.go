package test

import (
    "testing"
    "pattern/singleton/msgpool"
)

func TestMessagePool(t *testing.T) {
    msg0 := msgpool.Instance().GetMsg()
    if msg0.Count != 0 {
        t.Errorf("expect msg count %d, but actual %d.", 0, msg0.Count)
    } else {
        t.Logf("expect msg count %d", msg0.Count)
    }
    msg0.Count = 1
    msgpool.Instance().AddMsg(msg0)
    msg1 := msgpool.Instance().GetMsg()
    if msg1.Count != 1 {
        t.Errorf("expect msg count %d, but actual %d.", 1, msg1.Count)
    } else {
        t.Logf("expect msg count %d", msg1.Count)
    }
}