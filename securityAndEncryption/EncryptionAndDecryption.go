package main

import (
	"fmt"
	"encoding/base64"
	"os"
	"crypto/aes"
	"crypto/cipher"
)

func init() {
	 fmt.Println("加密解密的过程   一般的应用可以采用base64算法，更加高级的话可以采用aes或者des算法   ")
	//CSRF攻击、XSS攻击、SQL注入攻击等一些Web应用中典型的攻击手法，它们都是由于应用对用户的输入没有很好的过滤引起的，所以除了介绍攻击的方法外，我们也介绍了了如何有效的进行数据过滤，以防止这些攻击的发生的方法。然后针对日异严重的密码泄漏事件，介绍了在设计Web应用中可采用的从基本到专家的加密方案。最后针对敏感数据的加解密简要介绍了，Go语言提供三种对称加密算法：base64、aes和des的实现

}

func main() {

	//base64加解密

	base64Demo()


	//高级加解密

	advancedEncryption()

}
/*
Go语言的crypto里面支持对称加密的高级加解密包有：
crypto/aes包：AES(Advanced Encryption Standard)，又称Rijndael加密法，是美国联邦政府采用的一种区块加密标准。
crypto/des包：DES(Data Encryption Standard)，是一种对称加密标准，是目前使用最广泛的密钥系统，特别是在保护金融数据的安全中。曾是美国联邦政府的加密标准，但现已被AES所替代。
 */
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
func advancedEncryption() {

	fmt.Println("****************************************")
	//需要加密的本来的原来的密码
    str :=[]byte("i am 仕明 ！！！")

    //加密的串是传入的字符串  这个输入的字符串必然是大于1
	if len(os.Args) > 1 {
		str= []byte(os.Args[1])
	}
    fmt.Println("str=",str)
	//aes的加密字符串 如果有输入的话，就用输入的
	key_text := "kghubghjnkmiljhuijnkhuioiuyhujuy"
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}
	fmt.Println("aes的加密字符串的长度= ",len(key_text))
     // 创建加密算法aes
     /*
      通过调用函数aes.NewCipher(参数key必须是16、24或者32位的[]byte，分别对应AES-128, AES-192或AES-256算法),返回了一个cipher.Block接口
      */
	//NewCipher创建并返回新的密码块。
	//关键参数应该是AES密钥，
	//或16, 24，或32字节选择
	//AES-128，AES-192，或AES-256。
    c,err:= aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Println("发生错误了 err=",err)
		os.Exit(1)
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(str))
	cfb.XORKeyStream(ciphertext, str)
	fmt.Printf("%s=>%x\n", str, ciphertext)


	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(str))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)

}
//   todo    三个函数实现了加解密操作
//type Block interface {
//	// BlockSize returns the cipher's block size.
//	BlockSize() int
//
//	// Encrypt encrypts the first block in src into dst.
//	// Dst and src may point at the same memory.
//	Encrypt(dst, src []byte)
//
//	// Decrypt decrypts the first block in src into dst.
//	// Dst and src may point at the same memory.
//	Decrypt(dst, src []byte)
//}


/*
如果Web应用足够简单，数据的安全性没有那么严格的要求，那么可以采用一种比较简单的加解密方法是base64，这种方式实现起来比较简单，Go语言的base64包已经很好的支持了这个
 */
func base64Demo() {

	str:= "你好 shiming "
	b1:=base64Encoding([]byte(str))
	//加密完了的 的str= [53 76 50 103 53 97 87 57 73 72 78 111 97 87 49 112 98 109 99 103]
	//加密完了的 的str= 5L2g5aW9IHNoaW1pbmcg
	fmt.Println("加密完了的 的str=",b1)
	fmt.Println("加密完了的 的str=",string(b1))


	//b2= 你好 shiming
	//err= <nil>
	b2,err:=base64Decode(b1)
	fmt.Println("b2=",string(b2))
	fmt.Println("err=",err)


}
//解码字符串返回由Base64字符串S表示的字节。
func base64Decode(bytes []byte) ([]byte,error) {
     //解码字符串返回由Base64字符串S表示的字节。
	return base64.StdEncoding.DecodeString(string(bytes))
}
func base64Encoding(bytes []byte) []byte{
	return []byte(base64.StdEncoding.EncodeToString(bytes))
}