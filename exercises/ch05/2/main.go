package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fileSize, err := fileLen("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Size of file: %v\n", fileSize)
	}
}

func fileLen(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	fileInfo, statErr := f.Stat()
	if statErr != nil {
		return 0, errors.New("could not retrieve file stats")
	}

	return int(fileInfo.Size()), nil

	// alternative option
	//
	//data := make([]byte, 2048)
	//totalSize := 0
	//for {
	//	count, err := f.Read(data)
	//	totalSize += count
	//	if err != nil {
	//		if err != io.EOF {
	//			return 0, err
	//		}
	//		break
	//	}
	//}
	//return totalSize, nil
}
