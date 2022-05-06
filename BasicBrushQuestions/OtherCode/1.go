/*
只有两个键的键盘
最初在一个记事本上只有一个字符 'A' 。你每次可以对这个记事本进行两种操作：Copy All (复制全
部) : 你可以复制这个记事本中的所有字符(部分的复制是不允许的)。Paste (粘贴) : 你可以粘贴你上
一次复制的字符。
给定一个数字 n 。你需要使用最少的操作次数，在记事本中打印出恰好 n 个 'A'。输出能够打印出 n 个
'A' 的最少操作次数。
输入: 3
输出: 3
解释:
最初, 我们只有一个字符 'A'。
第 1 步, 我们使用 Copy All 操作。
第 2 步, 我们使用 Paste 操作来获得 'AA'。
第 3 步, 我们使用 Paste 操作来获得 'AAA'。




A  A  AA    4       1+1 +1+1
A  A  AA  AA   6    1+1 +1+1 +1
A  A  AA  AA  AA  8 1+1 +1+1 +1 +1     +1 +1 +1 +1
A  A  AA      AAAA        8    1+1 1+1  1+1
A A A AAA AAA             9
A  A  AA      AAAA      AAAAAAAA     16    1+1 1+1  1+1 1+1

*/

package main

import "fmt"

func keyboardwithonlytwokeys(n int) (result int) {
	for i := 2; i <= n; {

		if n%i == 0 {

			result = result + i
			n = n / i
			i = 2
			continue

		} else {
			i++
		}

		if n == 1 {
			break
		}

	}

	return result

}
func main() {
	fmt.Println(keyboardwithonlytwokeys(3))

}
