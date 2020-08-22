**https://labuladong.gitbook.io/algo/** 
**https://github.com/halfrost/LeetCode-Go**


# servicetmpl1 https://github.com/jfeng45/servicetmpl1/blob/master/README.zh.md 
# go 抢占式调度 https://mp.weixin.qq.com/s/YVJh0wGlkGfymfY5u1iqTA
# go micro 框架 https://github.com/micro-in-cn/starter-kit
# crocordile 框架 https://github.com/labulaka521/crocodile/blob/master/README_ZH.md 

# go的 25种设计模式 https://juejin.im/post/6859015515344633863
（
   这里只出了一期，工厂抽象模式 挺不错：
   要点：
      interface 的 灵活使用
      reflect.Typeof 减少了 新增 对依赖项的改写
 ）
# go的 线程模型
# go断点 续传
# gorm v2 https://studygolang.com/topics/12120#reply0
# 如何避免用动态语言的思维写Go代码 https://mp.weixin.qq.com/s/LqV7ODo7QCRykcFSRMRRgA
(尽量用 *[]prodcut 这种模式，特殊逻辑也是： map[int64]*Product )

# https://juejin.im/post/6861471535878651917 go http1.1 比较好的文章


## 依赖注入的概念： https://www.cyhone.com/articles/facebookgo-inject/ （已读，可以使得 apiHander -> controller -> service 的流程不需要 逐步初始化）