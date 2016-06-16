package ginx

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

//GetString by key
func GetString(c *gin.Context, key string) (value string) {
	v, ok := c.Get(key)
	if !ok {
		return
	}
	value, ok = v.(string)
	if !ok {
		return
	}
	return value
}

//GetObject by key
func GetObject(c *gin.Context, key string, obj interface{}) error {
	value := GetString(c, key)
	return json.Unmarshal([]byte(value), obj)
}
