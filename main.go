package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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

func main() {
	twitter := NewTwitter("3PNBbgWuPYMAPsJ3lHLuu9E29", "eQxU7jrfpcUiV4O4dpRBfLWMAsS8rTTZkqLehyWB2dGHDS5Ta5", "3527859379-naqp2WQMAXOL1gkmJZL6ILaQUxPnnjJLpGFAWUU", "JWd84aTCAenBgqWytq60hzkkMesgwo3qRRMiyoBHVY046")
	dajareCount := 0
	oldTweetID := " "
	for {
		// ホームタイムラインを取得
		res, err := twitter.Get(
			"https://api.twitter.com/1.1/statuses/home_timeline.json", // Resource URL
			map[string]string{})                                       // Parameters
		if err != nil {
			log.Fatal(err)
		}

		// 最新ツイートの入手
		newTweet := res.([]interface{})[0]
		tweet, _ := newTweet.(map[string]interface{})

		if tweet["id_str"].(string) != oldTweetID {
			//ツイートの内容を変数に格納
			text := tweet["text"]

			//ツイートの表示
			fmt.Println(tweet["text"])

			//interface{}型からstringへの変換が可能な場合、okはtrueとなる
			if str, ok := text.(string); ok {
				dajare, _ := dajarep.Dajarep(str)
				if len(dajare) >= 1 {
					// ダジャレカウント
					dajareCount++
					// ユーザデータを変数に格納
					user := tweet["user"].(map[string]interface{})
					// ダジャレの検出ツイート
					twitter.Post(
						"https://api.twitter.com/1.1/statuses/update.json",
						map[string]string{"status": "@" + user["screen_name"].(string) + " " + strconv.Itoa(dajareCount) + "回目のダジャレを検出しました。", "in_reply_to_status_id": tweet["id_str"].(string)})
				}
			}
			oldTweetID = tweet["id_str"].(string)
		}
		time.Sleep(60000 * time.Millisecond)
	}
}
