package ginx

import (
	"encoding/json"

	j "github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

//GetString by key
func GetString(c *gin.Context, key string) (value string) {
	v, ok := c.Get(key)
	if !ok {
		return
	}
	data, ok := v.(map[string]string)
	if !ok {
		return
	}
	value = data[key]
	return
}

//GetJSON by key
func GetJSON(c *gin.Context, key string) (*j.Json, error) {
	value := GetString(c, key)
	return j.NewJson([]byte(value))
}

//GetObject by key
func GetObject(c *gin.Context, key string, obj interface{}) error {
  value := GetString(c, key)
	return json.Unmarshal([]byte(value), obj)
}
