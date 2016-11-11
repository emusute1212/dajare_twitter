package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/emusute1212/dajare_twitter/queue"
	"github.com/kurehajime/dajarep"
)

var dajareCount = 0

type ApiConf struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var apiConf ApiConf
	{
		apiConfPath := flag.String("conf", "config.json", "API Config File")
		flag.Parse()
		data, err_file := ioutil.ReadFile(*apiConfPath)
		check(err_file)
		err_json := json.Unmarshal(data, &apiConf)
		check(err_json)
	}

	anaconda.SetConsumerKey(apiConf.ConsumerKey)
	anaconda.SetConsumerSecret(apiConf.ConsumerSecret)
	api := anaconda.NewTwitterApi(apiConf.AccessToken, apiConf.AccessTokenSecret)
	api.SetLogger(anaconda.BasicLogger) // logger を設定

	v := url.Values{}

	stream := api.UserStream(v)
	//検出ツイート
	_, err := api.PostTweet("@CaroBays ダジャレ検出を開始します。", nil)

	//すでに検出ツイートをしていた場合
	for i := 2; err != nil; i++ {
		_, err = api.PostTweet("@CaroBays "+strconv.Itoa(i)+"回目のダジャレ検出開始通知。", nil)
	}

	//変数の準備
	var tweet string
	var userName string
	var tweetID string
	var userID string
	//キューのリセット
	queue.Init()

	//ツイート検出のためのfor
	for {
		select {
		case item := <-stream.C:
			switch status := item.(type) {
			case anaconda.Tweet:
				//ツイートしたユーザー
				userName = status.User.ScreenName
				//ツイートのID
				tweetID = status.IdStr
				//ツイートしたユーザーのID
				userID = status.User.IdStr

				//自分のアカウント以外を読み込むようにする
				if userID != "3527859379" {
					tweet = status.Text
					// Tweet を受信
					fmt.Println(userName + ":" + tweet)

					// ダジャレがあるか確かめる
					result, daj, wad := checkDajare(tweet)

					//ダジャレがある時
					if result {
						//ダジャレ検出ツイート
						text := makeTweet(userName, daj, wad)

						//キューにツイートする内容を格納
						queue.Enqueue(queue.TweetData{tweetID, text})
					}
					for temp, e := queue.Dequeue(); e == nil; temp, e = queue.Dequeue() {
						//リプライを送る相手を指定
						v.Set("in_reply_to_status_id", temp.ID)
						//ツイート送信
						_, error := api.PostTweet(temp.Tweet, v)
						if error != nil {
							queue.Enqueue(temp)
							fmt.Println(error)
						}
					}
				}

			default:
			}
		}
	}
}

//ダジャレがあるか確かめるメソッド
func checkDajare(text string) (bool, []string, []string) {
	//ダジャレとかかっているワードの検出
	dajare, ward := dajarep.Dajarep(text)
	//かかっているワードが1以上の時trueとなる
	result := len(dajare) >= 1

	return result, dajare, ward
}

//ダジャレ検出ツイートを作るメソッド
func makeTweet(toName string, dajare []string, ward []string) string {
	//変数の準備
	var content string
	var output string

	//検出ツイートの際に巻き込まれるユーザー
	targetUser := " @CaroBays"

	//検出開始してから検出したダジャレをカウントする
	dajareCount++

	//検出したダジャレ全てを通知するために一つのツイートに全てまとめる
	for i := 0; i < len(dajare); i++ {
		//ダジャレをコンソールへ出力
		fmt.Println(dajare[i] + "," + ward[i])

		//検出内容の格納
		content += " \"" + dajare[i] + "\" から \"" + ward[i] + "\""

		if i < len(dajare)-1 {
			content += "と"
		} else {
			content += "を検出しました。\n検出開始から" + strconv.Itoa(dajareCount) + "回目。"
		}
	}

	//ツイート制限140文字を守るための条件文(初めの@の分も考慮して140以上としてある)
	if len([]rune(content))+len([]rune(targetUser))+len([]rune(toName)) >= 140 {
		content = " ダジャレを検出しました。\n検出開始から" + strconv.Itoa(dajareCount) + "回目。"
	}

	//実際に出力される文章
	output = "@" + toName + targetUser + content

	return output
}
