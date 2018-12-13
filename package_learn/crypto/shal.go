package main

//Go 在多个 crypto/* 包中实现了一系列散列函数。
import "crypto/sha1"
import "fmt"

func main() {
	s := "sha1 this string"
	//产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
	h := sha1.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(s))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

/*
运行程序计算散列值并以可读 16 进制格式输出。
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94
你可以使用和上面相似的方式来计算其他形式的散列值。例如，计算 MD5 散列，引入 crypto/md5 并使用 md5.New()方法。
注意，如果你需要密码学上的安全散列，你需要小心的研究一下哈希强度。
*/
