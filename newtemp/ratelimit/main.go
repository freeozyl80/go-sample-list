package main

import (
  "fmt"
)

type Limiter struct {
  limit Limit // 放入桶的频率   （Limit 为 float64类型
  burst int // 桶的大小
  
  mu sync.Mutex
  token float64 // 当前桶内剩余令牌个数
  last time.Time  // 最近取走token的时间
  lastEvent time.Time  // 最近限流事件的时间
}

type Reservation struct {
  ok        bool  // 是否满足条件分配到了tokens
  lim       *Limiter // 发送令牌的限流器
  tokens    int   // tokens 的数量
  timeToAct time.Time  //  满足令牌发放的时间
  limit Limit  // 令牌发放速度
}

// 在 now 时间需要拿到n个令牌，最多可以等待的时间为maxFutureResrve
// 结果将返回一个预留令牌的对象
func (lim *Limiter) reserveN(now time.Time, n int, maxFutureReserve time.Duration) Reservation {
  lim.mu.Lock()

  // 首先判断是否放入频次是否为无穷大，如果为无穷大，说明暂时不限流
  if lim.limit == Inf {
    // ...
  }

  // 拿到截至now 时间时，可以获取的令牌tokens数量，上一次拿走令牌的时间last
  now, last, tokens := lim.advance(now)

  // 然后更新 tokens 的数量，把需要拿走的去掉
  tokens -= float64(n)

  // 如果tokens 为负数，说明需要等待，计算等待的时间
  var waitDuration time.Duration
  if tokens < 0 {
    waitDuration = lim.limit.durationFromTokens(-tokens)
  }

  // 计算是否满足分配条件
  // ① 需要分配的大小不超过桶容量
  // ② 等待时间不超过设定的等待时常
  ok := n <= lim.burst && waitDuration <= maxFutureReserve

  // 最后构造一个Reservation对象
  r := Reservation{
    ok:    ok,
    lim:   lim,
    limit: lim.limit,
  }
  if ok {
    r.tokens = n
    r.timeToAct = now.Add(waitDuration)
  }

  // 并更新当前limiter 的值
  if ok {
    lim.last = now
    lim.tokens = tokens
    lim.lastEvent = r.timeToAct
  } else {
    lim.last = last
  }

  lim.mu.Unlock()
  return r
}

func main() {
  fmt.Println("Hello")
}