package main

import (
	"fmt"
	"log"
 	"os"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {


	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env")
	}
 
	
	ConsumerKey := os.Getenv("CONSUMERKEY")
	ConsumerSecret := os.Getenv("CONSUMERSECRET")
	AccessToken := os.Getenv("ACCESSTOKEN")
	AccessTokenSecret := os.Getenv("ACCESSTOKENSECRET")

	config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	token := oauth1.NewToken(AccessToken, AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("Account :  %s (%s)\n", user.ScreenName, user.Name)

	
	tweet, resp, err := client.Statuses.Update("tst", nil)

	if err != nil {
		log.Printf("err: %v\n",err)
	}
	log.Printf("resp%+v\n", resp)
	log.Printf("twee%+v\n", tweet)

}
