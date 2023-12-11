package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gofiber/fiber"
)

var cache = memcache.New("localhost:11211")

type Photo struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

func toJson(val []byte) Photo {
	photo := Photo{}
	err := json.Unmarshal(val, &photo)
	if err != nil {
		panic(err)
	}
	return photo
}

func verifyCache(c *fiber.Ctx) error {
	id := c.Params("id")
	val, err := cache.Get(id)
	if err != nil {
		return c.Next()
	}

	data := toJson(val.Value)
	return c.JSON(fiber.Map{"Cached": data})
}

func main() {
	app := fiber.New()

	app.Get("/:id", verifyCache, func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := http.Get("https://jsonplaceholder.typicode.com/photos/" + id)
		if err != nil {
			return err
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cacheErr := cache.Set(&memcache.Item{Key: id, Value: body, Expiration: 10})
		if cacheErr != nil {
			panic(cacheErr)
		}

		data := toJson(body)

		return c.JSON(fiber.Map{"Data": data})
	})

	app.Listen(":8080")
}
