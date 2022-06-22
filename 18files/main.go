package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("\nWelcome to files!!!")

	content := "This needs to go in a file - nladridasofficial.42web.io"

	file, err := os.Create("./nladridasofficialfile.txt")

	//Checking for error
	checkNilErr(err)

	length, err := io.WriteString(file, content) //io.Wrtetring returns leength if ssuccesfully executed else error

	//checkiing for error
	checkNilErr(err)

	fmt.Println("Length is: ", length)

	//Close the flle
	defer file.Close() //Defer keyword is used to make sure it will close the file only at the last

	readFile("./nladridasofficialfile.txt")

}

//**** Remember creation is the os operation & for reading the file and doing more manipulation, there is a separate utility given to us, which is "ioutil" *****

//function to read file
func readFile(fileName string) {

	//****One important thing to know here is; whenever we read a file, it is not being read in string format, even if we are reading from internet as well, it is always in byte format.
	//Hense we named it "dataByte" to make it siimple to understand what kind of data we are receving*****

	dataByte, err := ioutil.ReadFile(fileName)

	//Checking for error
	checkNilErr(err)

	fmt.Println("Text data inside the file is: ", string(dataByte)) //Parssed in the tring format

}


//Check error funcion
func checkNilErr(err error)  {
	if err != nil {
		panic(err)
	}
}
