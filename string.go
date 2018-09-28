package main

import "fmt"

func main() {
	s := "Hello World"
	s = s + "songweishuai"
	fmt.Println(s, len(s))

	//s = fmt.Sprintf("")

	s1 := s[:5]
	s2 := s[7:]
	s3 := s[1:5]

	fmt.Println("s1 is ", s1)
	fmt.Println("s2 is ", s2)
	fmt.Println("s3 is ", s3)

	/*Unicode Code Point*/
	fmt.Printf("%T\n", 'a')

	bs := []rune(s)
	bs[0] = 'w'
	fmt.Printf("s:%s\n", string(bs))

	bss := []byte(s)
	bss[1] = 'w'
	fmt.Printf("bss:%s\n", string(bss))
	s = string(bss)
	fmt.Printf("s:%s\n", s)
}
