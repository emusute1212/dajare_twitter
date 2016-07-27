package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "encoding/json"
    "github.com/mrjones/oauth"
)

type Twitter struct {
    consumer *oauth.Consumer
    accessToken *oauth.AccessToken
}

func NewTwitter(consumerKey, consumerSecret, accessToken, accessTokenSecret string) (*Twitter) {
    twitter := new(Twitter)
    twitter.consumer = oauth.NewConsumer(
        consumerKey,
        consumerSecret,
        oauth.ServiceProvider{
            RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
            AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
            AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
        })
    twitter.accessToken = &oauth.AccessToken{accessToken, accessTokenSecret,nil}
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
    if (err != nil) {
        return nil, err
    }

    // decode
    var result interface{}
    err = json.Unmarshal(b, &result)
    return result, err
}

func main() {
    twitter := NewTwitter("ConsumerKey", "ConsumerSecret", "AccessToken", "AccessTokenSecret")

    // testとつぶやく
    res, err := twitter.Post(
        "https://api.twitter.com/1.1/statuses/update.json", // Resource URL
        map[string]string{"status": "はじめてのついーと"}) // Parameters
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(res)

    // ホームタイムラインを取得
    res, err = twitter.Get(
        "https://api.twitter.com/1.1/statuses/home_timeline.json", // Resource URL
        map[string]string{}) // Parameters
        if err != nil {
            log.Fatal(err)
        }

    for _, val := range res.([]interface{}) {
        tweet, _ := val.(map[string]interface{})
        fmt.Println(tweet["text"])
    }
}
