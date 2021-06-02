package main

import (
	"fmt"
	"math"
	"github.com/eRez-ws/go-stringUtil"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/google/go-querystring/query"
	//"github.com/go-cc/cc-table/blob/master/cc-pinyin-range"
	"github.com/eladSalti/go_test-t"
	//"github.com/davecgh/go-spew/spew"
	//"github.com/phayes/freeport"
	//"github.com/gokrazy/gokrazy"
	//"github.com/KyleBanks/depth"
)

func main() {
    fmt.Println(stringUtil.Reverse("@#oG ,olleH"));
	fmt.Println(math.Exp2(10))
	println("Beego version:", beego.VERSION);

	// querystring/query
	type Options struct {
	  Query   string `url:"q"`
	  ShowAll bool   `url:"all"`
	  Page    int    `url:"page"`
	}

	opt := Options{ "foo", true, 2 }
	v, _ := query.Values(opt)
	fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"
	
	/*port, err := freeport.GetFreePort()
	if err != nil {
		println(err)
	} else {
		println(port)
	}*/
}



