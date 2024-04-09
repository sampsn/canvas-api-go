package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var url string = "https://dixietech.instructure.com/api/v1"

func main() {
	app := fiber.New()

	app.Get("/courses", getCourses)
	app.Get("/discussions/:course_id", getDiscussions)

	app.Listen(":8001")
}

func getCourses(c *fiber.Ctx) error {
	body := getJSON(url + "/courses")
	var courses []Course
	err := json.Unmarshal([]byte(body), &courses)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(courses)
}

func getDiscussions(c *fiber.Ctx) error {
	course_id := c.Params("course_id")
	body := getJSON(url + "/courses/" + course_id + "/discussion_topics")
	var discussions []Discussion
	json.Unmarshal([]byte(body), &discussions)
	return c.JSON(discussions)
}

func userAuth(req *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	access_key := os.Getenv("ACCESS_TOKEN")

	headers := map[string]string{"Authorization": "Bearer " + access_key}

	for key, value := range headers {
		req.Header.Add(key, value)
	}
}

func getJSON(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	userAuth(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
