package main

import "bufio"

// Scanner
//func main() {
//	scn := bufio.NewScanner(strings.NewReader("123 asd sa ks"))
//	//buf := make([]byte, 20, 20)
//	//scn.Buffer(buf, 19)
//
//	// 注册SplitFunc
//	split := func(data []byte, atEOF bool) (int, []byte, error) {
//		return bufio.ScanWords(data, atEOF)
//	}
//
//	scn.Split(split)
//
//	b := scn.Scan()
//	fmt.Println(b)
//	fmt.Println(string(scn.Bytes()))
//	fmt.Println(scn.Text())
//}

// Reader
func main() {
	reader := bufio.NewReader()
	reader.Read()
}
