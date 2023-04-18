package main

import (
  "fmt"
  "log"
)

func f(n int) (r int) {
  defer func() {
    r += n
    recover()
  }()

  var f func()

  defer f()
  f = func() {
    r += 2
  }

  return n + 1
}

func main() {
  fmt.Println(f(3)) // 7

  /**
  考察了defer的调用，通过上面的defer f()可知，f()是没有定义的？(存疑)所以报了panic，而在defer func()里面的recover()处理了这个panic，保证了整个程序能够执行完
  然后看f()函数返回的值为r，而不是n+1，这里考察的是函数的返回这一块的知识，所以最后的结果不是return n+1而是n+1+n的结果（第三点详细解释）
  为什么会等于7？看看最f()函数的最后一行为return n+1,也就是return 3+1,函数f()的返回参数为r int，所以r = 3 + 1,再到defer func()的r+=n也就是r = n + 1 + n即r = 3 + 1 + 3，然后返回r的值，就是7
   */
}

