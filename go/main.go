// ToDo
// use https://github.com/oschwald/geoip2-golang instead of geoiplookup command
package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
	"math"
	"math/rand"
	"time"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"log"
	"regexp"
	"github.com/ChimeraCoder/anaconda"
	"./iso3166_1"
	"encoding/json"
)

type Setting struct {
	RootMessage string "json:rootMessage"
	LoginMessage string "json:loginMessage"
	LoginBonusMessage string "json:loginBonusMessage"
	UptimeMessage string "json:uptimeMessage"
	AutoPostMessage string "json:autoPostMessage"
	Bonus []string `json:"loginbonus"`
}

var settings Setting

func main()  {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	settings = loadSettings(filepath.Dir(exe) + "/settings.json")

	var (
		rootflag bool
		verboseflag bool
		status string
	)
	flag.BoolVar(&rootflag, "r", false, "root login notify")
	flag.BoolVar(&verboseflag, "v", false, "show twitted message")
	flag.Parse()

	if rootflag {
		status = settings.RootMessage + settings.AutoPostMessage
	} else {
		ssh_connection := strings.Split(os.Getenv("SSH_CONNECTION"), " ")
		remote_ipaddr := ssh_connection[0]
		country_code := GeoIP(remote_ipaddr)
		country := iso3166_1.Country_name_ja[country_code]

		re := regexp.MustCompile("!!!!")
		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		var bonus_or_uptime string
		if r1.Intn(100) < 10 {
			bonus_or_uptime = re.ReplaceAllString(settings.UptimeMessage, getUptime())
		} else {
			bonus_or_uptime = re.ReplaceAllString(settings.LoginBonusMessage, getLoginBonus())
		}

		status = re.ReplaceAllString(settings.LoginMessage, remote_ipaddr+"("+country+")") + bonus_or_uptime + settings.AutoPostMessage
	}
	tweet(status, verboseflag)
}

func loadSettings(filename string) Setting {
	var data []byte
	var err error
	if data, err = ioutil.ReadFile(filename); err != nil {
		log.Fatal(err)
	}

	var settings Setting
	if err := json.Unmarshal(data, &settings); err !=nil {
		log.Fatal(err)
	}

	return settings
}

func tweet(text string, verboseflag bool)  {
	api := getTwitterApi()

	var tweet anaconda.Tweet
	var err error
	if tweet, err = api.PostTweet(text, nil); err != nil {
		panic(err)
	}

	if verboseflag {
		fmt.Println(tweet.Text)
	}
}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	return api
}

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
	// data, _ := ioutil.ReadFile("loginbonus")
	// bonus := strings.Split(string(data), "\n")

	bonus := settings.Bonus
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
