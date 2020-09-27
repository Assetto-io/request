package example

import (
	"fmt"
	request2 "github.com/assetto-io/request/request"
	"io/ioutil"
)

func exampleUsage() {
	client := request2.Client()
	response, err := client.Get("http://github.com", nil)

	if err != nil {

	}

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Printf(string(bytes))

}
