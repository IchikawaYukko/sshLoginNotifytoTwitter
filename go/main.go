// ToDo
// use configfile for status message
//    config & loginbonus from json
// use https://github.com/oschwald/geoip2-golang instead of geoiplookup command

// translate from PHP version
// $uptimeMessage = "サーバのUptime:$UPTIME(自動投稿)";
// $rootMessage = "百合子さんがrootに権限昇格しました♪ (自動投稿)";

package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
	"math"
	"math/rand"
	"time"
	"os"
	"os/exec"
	"log"
	// "github.com/ChimeraCoder/anaconda"
	"./iso3166_1"
)

func main()  {
	ssh_connection := strings.Split(os.Getenv("SSH_CONNECTION"), " ")
	remote_ipaddr := ssh_connection[0]
	country_code := GeoIP(remote_ipaddr)
	country := iso3166_1.Country_name_ja[country_code]
	fmt.Println("百合子さんが "+remote_ipaddr+"("+country+") から ConoHa にsshログインしました♪ ログインぼおなす！："+getLoginBonus())
}

// func tweet(text string)  {
// 	api := getTwitterApi()

// 	tweet, err := api.PostTweet(text, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func getTwitterApi() *anaconda.TwitterApi {
// 	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
// 	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
// 	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
// 	return api
// }

func isIPv6(ipaddr string) bool {
	if -1 == strings.Index(ipaddr, ":") {
		return false
	} else {
		return true
	}
}

func GeoIP(ipaddr string) string {
	var out []byte
	var err error

	if isIPv6(ipaddr) {
		out, err = exec.Command("sh", "-c", "geoiplookup6 " + ipaddr).Output()
	} else {
		out, err = exec.Command("sh", "-c", "geoiplookup " + ipaddr).Output()
	}
	if err != nil {
		log.Fatal(err)
	}

	country_code := strings.Split(string(out), " ")
	if isIPv6(ipaddr) {
		return strings.Trim(country_code[4], ",")
	} else {
		return strings.Trim(country_code[3], ",")
	}
}

func getUptime() string {
	data, _ := ioutil.ReadFile("/proc/uptime")
	slice := strings.Split(string(data), " ")
	num , _ := strconv.ParseFloat(slice[0], 64)

	return strconv.FormatFloat(math.Trunc(num / 60 / 60 / 24), 'f', 0, 64) + " days"
}

func getLoginBonus() string {
	data, _ := ioutil.ReadFile("loginbonus")
	bonus := strings.Split(string(data), "\n")

	bonus = append(bonus, getUptime())
	shuffle(bonus)

	return bonus[0]
}


func shuffle(data []string) {
	n := len(data)
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := n - 1; i >= 0; i-- {
		j := r1.Intn(i + 1)
        data[i], data[j] = data[j], data[i]
    }
}