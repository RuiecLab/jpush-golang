package main

import (
	"fmt"
	jpush "github.com/RuiecLab/jpush-golang"
)

func main() {
	cid := jpush.NewCidRequest(1, "")
	res, err := cid.GetCidList("key", "secret") // 这里的 key 和 secret 需要替换成自己的
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", res.String())
	}
}
