package main
import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	// f,err:=os.Open("./18.FileHandling/example.txt")
	// if err!=nil{
	// 	//log the error
	// 	panic(err)
	// }
	// fileInfo,err:=f.Stat()
	// if err!=nil{
	// 	//log the error
	// 	panic(err)
	// }
	// fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.Size())
	// fmt.Println(fileInfo.IsDir())
	// fmt.Println(fileInfo.Mode())


	//* reading into buffer*//
	// defer f.Close()

	// buf :=make([]byte,10)
	// d,err:=f.Read(buf)
	// if err!=nil{
	// 	panic(err);
	// }
	// for i:=0;i<len(buf);i++{
	// 	println("data",d,string(buf[i]))
	// }


	//reading a directory/folder
	// dir,err:=os.Open("./")
	// if err!=nil{
	// 	//log the error
	// 	panic(err)
	// }
	// defer dir.Close()
	// fileInfo ,err:=dir.ReadDir(-1)
	// for _,fi:=range fileInfo{
	// 	fmt.Println(fi.Name(),fi.IsDir())
	// }


	//read and write to another file (streaming fashion)
	sourceFile,err:=os.Open("./18.FileHandling/example.txt")
	if err!=nil{
		panic(err)
	}

	defer sourceFile.Close()

	destFile,err:=os.Create("./18.FileHandling/example2.txt")
	if err!=nil{
		panic(err)
	}

	defer destFile.Close()

	reader:=bufio.NewReader(sourceFile)
	writer:=bufio.NewWriter(destFile)

	for{
		b,err:=reader.ReadByte()
		if err!=nil{
			if err.Error() != "EOF"{
				panic(err)
			}
			break
		}

		e:=writer.WriteByte(b)
		if err!=nil{
			panic(e)
		}
	}
	writer.Flush()

	fmt.Println("written to file")

	//delete a file
	err:=os.Remove("example2.txt")
	if err!=nil{
		panic(err)
	}
	fmt.Println("Deleted successfully")

}