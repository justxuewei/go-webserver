package util

import "github.com/gin-gonic/gin"

func ErrorPage(title string, msg string) (int, string, *gin.H) {
	h := gin.H{
		"title": title,
		"error": msg,
	}
	return 500, "basic/error.html", &h
}
