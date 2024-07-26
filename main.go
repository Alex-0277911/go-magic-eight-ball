package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	ID       int    `json:"id"`
	Response string `json:"response"`
}

var responses = []response{
	{ID: 1, Response: "It is certain."},
	{ID: 2, Response: "It is decidedly so."},
	{ID: 3, Response: "Without a doubt."},
	{ID: 4, Response: "Yes definitely."},
	{ID: 5, Response: "You may rely on it."},
	{ID: 6, Response: "As I see it, yes."},
	{ID: 7, Response: "Most likely."},
	{ID: 8, Response: "Outlook good."},
	{ID: 9, Response: "Yes."},
	{ID: 10, Response: "Signs point to yes."},
	{ID: 11, Response: "Reply hazy, try again."},
	{ID: 12, Response: "Ask again later."},
	{ID: 13, Response: "Better not tell you now."},
	{ID: 14, Response: "Cannot predict now."},
	{ID: 15, Response: "Concentrate and ask again."},
	{ID: 16, Response: "Don't count on it."},
	{ID: 17, Response: "My reply is no."},
	{ID: 18, Response: "My sources say no."},
	{ID: 19, Response: "Outlook not so good."},
	{ID: 20, Response: "Very doubtful."},
}

func hello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello with Gooooo!")
}

func getAllAnswers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, responses)
}

func getRandomAnswer(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, responses[rand.Intn(len(responses))])
}

func main() {
	router := gin.Default()
	router.GET("/hello", hello)
	router.GET("/answers", getAllAnswers)
	router.GET("/answer", getRandomAnswer)

	if err := router.Run(":8080"); err != nil {
		log.Panicf("error: %s", err)
	}
}
