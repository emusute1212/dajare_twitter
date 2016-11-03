package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/kurehajime/dajarep"
	"github.com/mrjones/oauth"
)

type Twitter struct {
	consumer    *oauth.Consumer
	accessToken *oauth.AccessToken
}

func NewTwitter(consumerKey, consumerSecret, accessToken, accessTokenSecret string) *Twitter {
	twitter := new(Twitter)
	twitter.consumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})
	twitter.accessToken = &oauth.AccessToken{accessToken, accessTokenSecret, nil}
	// twitter.accessToken = &oauth.AccessToken{accessToken, accessTokenSecret}
	return twitter
}

func (t *Twitter) Get(url string, params map[string]string) (interface{}, error) {
	response, err := t.consumer.Get(url, params, t.accessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// decode
	var result interface{}
	err = json.Unmarshal(b, &result)
	return result, err
}

func (t *Twitter) Post(url string, params map[string]string) (interface{}, error) {
	response, err := t.consumer.Post(url, params, t.accessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// decode
	var result interface{}
	err = json.Unmarshal(b, &result)
	return result, err
}

// func checkSig(twitter *Twitter) {
//
// 	code := <-exitChan
// 	os.Exit(code)
// }

func main() {
	// fmt.Println("検出開始です。")
	twitter := NewTwitter("3PNBbgWuPYMAPsJ3lHLuu9E29", "eQxU7jrfpcUiV4O4dpRBfLWMAsS8rTTZkqLehyWB2dGHDS5Ta5", "3527859379-naqp2WQMAXOL1gkmJZL6ILaQUxPnnjJLpGFAWUU", "JWd84aTCAenBgqWytq60hzkkMesgwo3qRRMiyoBHVY046")
	oldTweetID := " "
	dajareCount := 0
	var newTweet interface{}
	var tweet map[string]interface{}

	//ダジャレ検出開始通知
	_, error := twitter.Post(
		"https://api.twitter.com/1.1/statuses/update.json",
		map[string]string{"status": "@CaroBays ダジャレ検出開始します。"})

	for i := 1; error != nil; i++ {
		_, error = twitter.Post(
			"https://api.twitter.com/1.1/statuses/update.json",
			map[string]string{"status": "@CaroBays ダジャレ検出開始します。" + strconv.Itoa(i) + "回目"})
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		for {
			s := <-signalChan
			switch s {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				fmt.Println("hungup")

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				fmt.Println("Warikomi")
				_, error := twitter.Post(
					"https://api.twitter.com/1.1/statuses/update.json",
					map[string]string{"status": "@CaroBays ダジャレの検出を終了します。"})
				// _, error := twitter.Post(
				// 	"https://api.twitter.com/1.1/statuses/update.json",
				// 	map[string]string{"status": "ダジャレの検出を終了します。"})

				for i := 1; error != nil; i++ {
					_, error = twitter.Post(
						"https://api.twitter.com/1.1/statuses/update.json",
						map[string]string{"status": "ダジャレの検出を終了します。" + strconv.Itoa(i) + "回目"})
				}
				os.Exit(0)

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				fmt.Println("force stop")
				_, error := twitter.Post(
					"https://api.twitter.com/1.1/statuses/update.json",
					map[string]string{"status": "@CaroBays ダジャレの検出を終了します。"})
				// _, error := twitter.Post(
				// 	"https://api.twitter.com/1.1/statuses/update.json",
				// 	map[string]string{"status": "ダジャレの検出を終了します。"})

				for i := 1; error != nil; i++ {
					_, error = twitter.Post(
						"https://api.twitter.com/1.1/statuses/update.json",
						map[string]string{"status": "ダジャレの検出を終了します。" + strconv.Itoa(i) + "回目"})
				}
				os.Exit(0)

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				fmt.Println("stop and core dump")
				_, error := twitter.Post(
					"https://api.twitter.com/1.1/statuses/update.json",
					map[string]string{"status": "@CaroBays ダジャレの検出を終了します。"})
				// _, error := twitter.Post(
				// 	"https://api.twitter.com/1.1/statuses/update.json",
				// 	map[string]string{"status": "ダジャレの検出を終了します。"})

				for i := 1; error != nil; i++ {
					_, error = twitter.Post(
						"https://api.twitter.com/1.1/statuses/update.json",
						map[string]string{"status": "ダジャレの検出を終了します。" + strconv.Itoa(i) + "回目"})
				}
				os.Exit(0)

			default:
				fmt.Println("Unknown signal.")
				_, error := twitter.Post(
					"https://api.twitter.com/1.1/statuses/update.json",
					map[string]string{"status": "@CaroBays ダジャレの検出を終了します。"})
				// _, error := twitter.Post(
				// 	"https://api.twitter.com/1.1/statuses/update.json",
				// 	map[string]string{"status": "ダジャレの検出を終了します。"})

				for i := 1; error != nil; i++ {
					_, error = twitter.Post(
						"https://api.twitter.com/1.1/statuses/update.json",
						map[string]string{"status": "ダジャレの検出を終了します。" + strconv.Itoa(i) + "回目"})
				}
				os.Exit(1)
			}
		}
	}()

	fmt.Println("検出開始")

	for {
		var myTweet interface{} = " "
		// ホームタイムラインを取得
		res, err := twitter.Get(
			"https://api.twitter.com/1.1/statuses/home_timeline.json", // Resource URL
			map[string]string{})                                       // Parameters
		if err != nil {
			_, error := twitter.Post(
				"https://api.twitter.com/1.1/statuses/update.json",
				map[string]string{"status": "@CaroBays ダジャレの検出を終了します。"})
			// _, error := twitter.Post(
			// 	"https://api.twitter.com/1.1/statuses/update.json",
			// 	map[string]string{"status": "ダジャレの検出を終了します。"})

			for i := 1; error != nil; i++ {
				_, error = twitter.Post(
					"https://api.twitter.com/1.1/statuses/update.json",
					map[string]string{"status": "ダジャレの検出を終了します。" + strconv.Itoa(i) + "回目"})
			}
		}

		//初回起動時にoldTweetIDにIDをセットする
		if oldTweetID == " " {
			// 最新ツイートの入手
			newTweet = res.([]interface{})[0]
			tweet, _ = newTweet.(map[string]interface{})
			oldTweetID = tweet["id_str"].(string)
		}

		//最新ツイート全ての読み込み
		for i := 0; ; i++ {
			//最新ツイートの読み込み
			newTweet = res.([]interface{})[i]
			tweet, _ = newTweet.(map[string]interface{})

			if tweet["id_str"].(string) != oldTweetID {
				//ツイートの内容を変数に格納
				text := tweet["text"]

				//ツイートの表示
				fmt.Println(tweet["text"])

				//interface{}型からstringへの変換が可能な場合、okはtrueとなる
				if str, ok := text.(string); ok {
					dajare, kana := dajarep.Dajarep(str)
					fmt.Println(dajare)
					fmt.Println(kana)
					// ユーザデータを変数に格納
					user := tweet["user"].(map[string]interface{})

					if len(dajare) >= 1 {
						dajareCount++
						output := ""
						for i := 0; i < len(dajare); i++ {
							output += "@" + user["screen_name"].(string) + " @CaroBays " + "\"" + dajare[i] + "\" から \"" + kana[i] + "\""
							if i < len(dajare)-1 {
								output += "と"
							} else {
								output += "を検出しました。\n検出開始から" + strconv.Itoa(dajareCount) + "回目。"
							}
						}
						if len([]rune(output)) <= 140 {
							// ダジャレの検出ツイート
							myTweet, _ = twitter.Post(
								"https://api.twitter.com/1.1/statuses/update.json",
								map[string]string{"status": output, "in_reply_to_status_id": tweet["id_str"].(string)})
						} else {
							// ダジャレの検出ツイート
							myTweet, _ = twitter.Post(
								"https://api.twitter.com/1.1/statuses/update.json",
								map[string]string{"status": "@" + user["screen_name"].(string) + " @CaroBays " + "ダジャレを検出しました。\n検出開始から" + strconv.Itoa(dajareCount) + "回目。", "in_reply_to_status_id": tweet["id_str"].(string)})
						}
					}
				}
			} else {
				if _, ok := myTweet.(string); !ok {
					oldTweetID = myTweet.(map[string]interface{})["id_str"].(string)
				} else {
					oldTweetID = res.([]interface{})[0].(map[string]interface{})["id_str"].(string)
				}
				break
			}
		}
		time.Sleep(65000 * time.Millisecond)
	}

}
