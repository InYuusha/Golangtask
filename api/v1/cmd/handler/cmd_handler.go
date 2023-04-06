package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	c "github.com/InYuusha/api/v1/cmd/controller"
	"github.com/InYuusha/api/v1/cmd/dto"
	"github.com/InYuusha/api/v1/cmd/store"
	"github.com/gin-gonic/gin"
)

var kvs *c.KeyValueStore = store.NewKeyValueStore()

func ExecuteQuery(c *gin.Context) {

	var query dto.Query

	if err := c.BindJSON(&query); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	qString := query.QueryString
	fieldsArr := strings.Split(qString, " ")

	var (
		cmd                string
		key                string
		value              string
		expiryStr          string
		expiry             int64
		condition          string
		values             []string
		blockForSeconds    float64
		blockForSecondsStr string
	)
	cmd = fieldsArr[0]

	if cmd == "SET" {
		switch len(fieldsArr) {
		case 6:
			condition = fieldsArr[5]
			fallthrough
		case 5:
			expiryStr = fieldsArr[4]
			fallthrough
		case 3:
			key = fieldsArr[1]
			value = fieldsArr[2]
		default:
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Invalid Query"))

		}
	} else if cmd == "GET" {
		key = fieldsArr[1]
	} else if cmd == "QPUSH" {
		key = fieldsArr[1]
		values = fieldsArr[2:]
	} else if cmd == "QPOP" {
		key = fieldsArr[1]
	} else if cmd == "BQPOP" {
		key = fieldsArr[1]
		blockForSecondsStr = fieldsArr[2]
	}

	if expiryStr != "" {
		var err error
		expiry, err = strconv.ParseInt(expiryStr, 10, 64)
		if err != nil {
			// Handle error
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	if blockForSecondsStr != "" {
		var err error
		blockForSeconds, err = strconv.ParseFloat(blockForSecondsStr, 64)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	var res interface{}
	var err error

	switch cmd {
	case "SET":
		res, err = kvs.SetCmd(key, &value, &expiry, condition)
		log.Println("Res ", res)
		if err != nil {
			fmt.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	case "GET":
		res, err = kvs.GetCmd(key)
		log.Println("Res ", res)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	case "QPUSH":
		res, err = kvs.Qpush(key, values)
		log.Println("Res ", res)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	case "QPOP":
		res, err = kvs.Qpop(key)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	case "BQPOP":
		res, err = kvs.Bqpop(key, blockForSeconds)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}

	c.JSON(200, dto.Response{
		Status:  200,
		Message: "Query executed successfully",
		Data: map[string]interface{}{
			"data": res,
		},
	})

}
