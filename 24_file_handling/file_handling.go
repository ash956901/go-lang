package main

import "os"

func main() {
	// Opening the file
	// f, err := os.Open("./example.txt")
	// if err != nil {
	// 	// log the error & panic
	// 	panic(err)
	// }

	// // File stats
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("fileName:", fileInfo.Name())
	// fmt.Println("directory:", fileInfo.IsDir())
	// fmt.Println("size:", fileInfo.Size())
	// fmt.Println("mode(permissions):", fileInfo.Mode())
	// fmt.Println("File modified at:", fileInfo.ModTime())

	// read file
	// f, err := os.Open("./example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // read and store in buffer
	// // temp location in memory
	// buf := make([]byte, 60)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// Another way to read --> not always recommended
	// it gives issue as loads full file into memory at once
	// for small file no issue
	// but big file like videos --> gbs --> issue
	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// reading folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(),fi.IsDir())
	// }

	//create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // This appends a string
	// // f.WriteString("hi go")
	// // f.WriteString("nice language")

	// bytes := []byte("hello golang")
	// f.Write(bytes)

	// read and write to another file (streaming file in chunks )
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}

	// }

	// writer.Flush()

	// fmt.Println("Written to new file successfully")

	// you can also use the copy function

	// Delete a File

	err:=os.Remove("example2.txt")
	if err!=nil{
		panic(err)
	}

}
