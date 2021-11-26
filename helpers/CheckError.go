package helpers

import (
	"fmt"
	"os"
)


func CheckError(err error) {

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func CheckDbCon(e error) {
	if e != nil {
		fmt.Println("failed to connect database")
		return
	}
}
