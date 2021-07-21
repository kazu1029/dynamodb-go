package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	attrs, err := attributevalue.Marshal([]string{})
	if err != nil {
		panic(err)
	}
	spew.Dump(attrs)
	fmt.Printf("%#v\n", attrs)

	e := attributevalue.NewEncoder(func(opt *attributevalue.EncoderOptions) {
		opt.NullEmptySets = true
	})
	attrs, err = e.Encode([]string{})
	if err != nil {
		panic(err)
	}

	var s []string
	err = attributevalue.Unmarshal(attrs, &s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", attrs)
	fmt.Printf("%#v\n", s)
}
