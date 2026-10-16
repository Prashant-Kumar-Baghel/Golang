package main

import (
	"fmt"
	"os"
)

// In this tutorial, we will learn about files in Go. We will learn how to create, read, write, and delete files in Go.
func main() {

	/*//Chapter 1: How to get the information of a file in Go?
	// os.Open ask the OS for read access to this file, give me a handle (*os.File) I can use, or an error if it can't be opened.
	f, err := os.Open("example.text") //os.Open Open a file named "example.txt" in read-only mode. If the file does not exist, it will return an error. inside Open function we provide the path of the file we want to open. it returns a file descriptor(file descriptor is used to interact with the file) and an error. If the file is opened successfully, the error will be nil. If the file does not exist, the error will be of type *os.PathError.
	//In go, if there is any error come inside function then we return the error , we don't throw the error like other programming languages. we have to handle the error in go. if there is any error then we will print the error and exit the program or we can panic the program . panic mean to stop the execution of the program and print the error message.
	if err != nil {
		//if there is any error then we will panic the program and print the error message.
		panic(err)
	}

	fileInfo, err := f.Stat() //Stat function is used to get the file information. it returns a FileInfo interface and an error. if the file does not exist, it will return an error of type *os.PathError. if the file exists, it will return a FileInfo interface which contains the file information like name, size, mode, modification time, etc.
	if err != nil {
		panic(err)
	}
	fmt.Println(fileInfo.Name())    //Name function is used to get the name of the file. it returns a string which is the name of the file.
	fmt.Println(fileInfo.Size())    //Size function is used to get the size of the file. it returns an int64 which is the size of the file in bytes.lets say we  have "Helloji" text then it will return 7 because there are 7 characters in the text. if we have "Hello ji" text then it will return 8 because there are 8 characters in the text including space.each character is 1 byte in size.
	fmt.Println(fileInfo.Mode())    //Mode function is used to get the mode of the file. it returns a FileMode which contains the file permissions.
	fmt.Println(fileInfo.ModTime()) //ModTime function is used to get the modification time of the file. it returns a time.Time which is the modification time of the file.
	fmt.Println(fileInfo.IsDir())   //IsDir function is used to check if the file is a directory. it returns a boolean which indicates if the file is a directory.*/

	//Chapter 2: How to read the content of the file in Go?

	//1st way
	f, err := os.Open("example.text")
	if err != nil {

		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {

		panic(err)
	}
	//write close condition as soon as possible because if we write after err then it will not close the file if there is an error. so we write close condition as soon as possible.
	defer f.Close()                         //defer keyword is used to close the file after the function is executed. it will close the file when the function is exited. we are using defer keyword to close the file because we want to close the file after we are done with it. we close file because we don't want to keep the file open for a long time. it will free up the resources and avoid memory leak.
	buffer := make([]byte, fileInfo.Size()) //buffer is a temporary storage area for data that is being transferred from one place to another. in this case we are creating a buffer of size equal to the size of the file. we are using make function to create a slice of bytes with length equal to the size of the file. we are using fileInfo.Size() to get the size of the file in bytes. we are using []byte to create a slice of bytes.buffer is a slice of bytes which will be used to read the content of the file.
	l, err := f.Read(buffer)                //Read function is used to read the content of the file. it reads the content of the file into the buffer. it returns the number of bytes read and an error. if the file is empty, it will return 0 and an error of type io.EOF. if the file is not empty, it will return the number of bytes read and nil error.
	if err != nil {
		panic(err)
	}
	for i := 0; i < l; i++ { //loop through the buffer and print the content of the file. we are using l to get the number of bytes read from the file. we are using i to loop through the buffer and print the content of the file. we are using string(buffer[i]) to convert the byte to string and print it.
		// fmt.Println(buffer[i])
		fmt.Println(l, string(buffer[i]))
	}
	fmt.Println("file read successfully", string(buffer)) //we are using string(buffer) to convert the slice of bytes to string and print it.

	//2nd way ( 2nd way is easy to write and read file content but it is not efficient because it reads the whole file content into memory in one time . if the file is very large then it will cause memory overflow. its good for small files but not good for large files.)
	data, err := os.ReadFile("example.text") //ReadFile function is used to read the content of the file. it reads the content of the file and returns a slice of bytes and an error. if the file does not exist, it will return an error of type *os.PathError. if the file exists, it will return a slice of bytes which contains the content of the file.
	if err != nil {
		panic(err)
	}
	fmt.Println("data", string(data)) //we are using string(data) to convert the slice of bytes to string and print it.

	/*//Chapter 3: How to read folder in Go ?
	f, err := os.Open(".") //Open function is used to open a directory. it returns a file descriptor and an error. if the directory does not exist, it will return an error of type *os.PathError. if the directory exists, it will return a file descriptor which can be used to read the directory.
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.ReadDir(n) :- if n <= 0 (e.g. 0 or -1) in ReadDir function then it will read ALL the entries of the directory in one slice. if n > 0 then it will read AT MOST n entries.
	dirEntry, err := f.ReadDir(0) //ReadDir function is used to read the content of the directory. it reads the content of the directory and returns a slice of DirEntry and an error. if the directory is empty, it will return an empty slice and nil error. if the directory is not empty, it will return a slice of DirEntry which contains the information of the files and directories in the directory.
	if err != nil {
		panic(err)
	}
	for _, dr := range dirEntry { //loop through the slice of DirEntry and print the name of the files and directories in the directory. we are using range keyword to loop through the slice of DirEntry. we are using _ to ignore the index of the slice. we are using dr to get the value of the slice. we are using dr.Name() to get the name of the file or directory.
		fmt.Println(dr.Name(), dr.IsDir()) //IsDir function is used to check if the file is a directory. it returns a boolean which indicates if the file is a directory. if the file is a directory, it will return true. if the file is not a directory, it will return false.
	}

	//Chapter 4: How to create a new file in Go ?
	/*f, err := os.Create("example5.txt") //Create function is used to create a new file. it creates a new file and returns a file descriptor and an error. if the file already exists, it will truncate the file and return a file descriptor. if the file does not exist, it will create a new file and return a file descriptor. if there is any error in creating the file, it will return an error of type *os.PathError.

	"Truncate" means to cut a file's existing content down to zero length — it empties the file out, keeping the file itself (and its name) but discarding everything that was in it.
	So for os.Create:
	If example5.txt doesn't exist → Go creates a brand new, empty file.
	If example5.txt already exists → Go doesn't error out or refuse. Instead it opens that same file and wipes its contents, leaving you with a 0-byte file that you can now write fresh data into.

	if err != nil {
		panic(err)
	}
	defer f.Close()

	//WriteString does NOT work in append mode. every open file has a hidden "offset" (a cursor position). WriteString writes at the current offset and then moves the offset forward by the number of bytes written. after os.Create the offset is 0, so the first write lands at the beginning. it only LOOKS like append because each write leaves the cursor at the end of what it just wrote.
	_, err = f.WriteString("hello gboy") //WriteString returns the number of bytes written and an error. we must not ignore the error, otherwise a failed write goes unnoticed. _ ignores the byte count because we don't need it here.
	f.WriteString("hello team")          //we get "hello gboyhello team" in the file because the first write wrote "hello gboy" and moved the offset to 10, then the second write wrote "hello team" starting at offset 10. if we want to write at the beginning of the file, we must move the offset back to 0 with Seek(0, 0) before writing again.

	// if err != nil {
	// 	panic(err)
	// }

	//concept for Truncate and Seek(Replacement of file content)
	// //Truncate(0) removes all the content of the file (size becomes 0), BUT it does NOT move the offset back. after writing "hello go" the offset is at 8, and truncate leaves it at 8. if we write again now, the write lands at position 8 and the OS fills positions 0-7 with NUL (zero) bytes — the file becomes 8 invisible NUL bytes + "hello go" instead of just "hello go".
	// err = f.Truncate(0)
	// if err != nil {
	// 	panic(err)
	// }

	// //so after Truncate we must move the offset back to the start with Seek. Seek(offset, whence) sets the cursor: whence 0 = relative to start of file, 1 = relative to current position, 2 = relative to end. Seek(0, 0) means "go to byte 0 from the start".
	// _, err = f.Seek(0, 0)
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = f.WriteString("hello go") //now this writes at position 0 and the file content is exactly "hello go".
	// if err != nil {
	// 	panic(err)
	// }

	//another way to write the content in file is to use Write function. Write function is used to write the content in the file. it writes the content in the file and returns the number of bytes written and an error. if there is any error in writing the content, it will return an error of type *os.PathError.
	// f, err := os.Create("example5.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// bytes := []byte("hello go") //we are using []byte to convert the string to slice of bytes. we are using slice of bytes because Write function takes slice of bytes as input.
	// f.Write(bytes)              //Write function is used to write the content in the file. it writes the content in the file and returns the number of bytes written and an error. if there is any error in writing the content, it will return an error of type *os.PathError.
	*/

	/*//Chapter 5: Reading and then writing files in streaming fashion in Go .streaming fashion means we read the file content in chunks instead of reading the whole file content into memory in one time. it is more efficient than reading the whole file content into memory in one time. we can use bufio package to read the file content in streaming fashion.
		sourceFile, err := os.Open("example.text") //sourceFile is a file descriptor which is used to read the content of the file.
		if err != nil {
			panic(err)
		}
		defer sourceFile.Close()

		destinationFile, err := os.Create("example2.txt") //destinationFile is a file descriptor which is used to write the content of the file.
		if err != nil {
			panic(err)
		}
		defer destinationFile.Close()
		//reader is a buffered reader which is used to read the content of the file in streaming fashion.
		//writer is a buffered writer which is used to write the content of the file in streaming fashion.
		//Buffered I/O means reads and writes go through an in-memory staging area (the "buffer") instead of hitting the file/OS directly on every call.
		reader := bufio.NewReader(sourceFile) //bufio is a package which provides buffered I/O which has default fix size. it provides buffered reader and buffered writer. buffered reader is used to read the content of the file in streaming fashion. buffered writer is used to write the content of the file in streaming fashion. bufio.NewReader function is used to create a new buffered reader which takes a file descriptor as input and returns a buffered reader. bufio.NewWriter function is used to create a new buffered writer which takes a file descriptor as input and returns a buffered writer.
		writer := bufio.NewWriter(destinationFile)

		for {
			b, err := reader.ReadByte() //ReadByte function is used to read a single byte from the buffered reader. it returns a byte and an error. if there is any error in reading the byte, it will return an error of type *os.PathError. if the end of the file is reached, it will return an error of type io.EOF.
			if err != nil {
				if err.Error() != "EOF" { //we write this condition because when we get the end of file then file will end and we don't have anything to read so we get error of type i.oEOF and this is infinite loop then we do panic again and again so we need to break the loop .
					panic(err)
				}
				break
			}

			e := writer.WriteByte(b) //WriteByte function is used to write a single byte to the buffered writer. it takes a byte as input and returns an error. if there is any error in writing the byte, it will return an error of type *os.PathError.
			if e != nil {
				panic(e)
			}

		}
		writer.Flush() //Flush function is used to flush the buffered writer that is it writes the content of the buffer to the file. it returns an error. if there is any error in flushing the buffer, it will return an error of type *os.PathError.
		fmt.Println("written to file successfully in streaming fashion")


		//Why write a manual streaming loop when io.Copy exists? io.Copy IS streaming internally — it loops with a 32KB buffer (read chunk, write chunk) so the whole file is never in memory, same as our byte-by-byte loop.
	//So for a plain copy, io.Copy(destinationFile, sourceFile) is the idiomatic one-liner: no bufio, no loop, no Flush needed.
	//The manual bufio loop is needed only when we must look at or change bytes in transit — e.g. uppercase while copying, count/skip lines, stop after 1 MB, compute a checksum.
	//Rule of thumb: plain copy → io.Copy; read line-by-line → bufio.Scanner; transform bytes mid-stream → manual bufio loop.

	*/

	//Chapter 6: How to delete a file in Go ?
	/*err := os.Remove("example2.txt") //Remove function is used to delete a file. it takes the name of the file as input and returns an error. if there is any error in deleting the file, it will return an error of type *os.PathError.
	if err != nil {
		panic(err)
	}
	fmt.Println("file deleted successfully")
	*/
}
