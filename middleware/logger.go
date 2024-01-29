package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	logFileName := "app.log"
	logFile, err := os.OpenFile("logs/"+logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
	}

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// อ่าน body ของ request
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("Error reading request body:", err)
		}

		// เซ็ต body ให้กลับมาเป็นต่างเดิมหลังจากอ่านแล้ว
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		if logFile != nil {
			_, err := logFile.WriteString(fmt.Sprintf("%s - [%s] \"%s %s %s\" %d %s \"%s\" %s\n",
				clientIP,
				endTime.Format("02/Jan/2006:15:04:05 -0700"),
				method,
				path,
				query,
				statusCode,
				latency,
				c.Request.UserAgent(),
				c.Errors.ByType(gin.ErrorTypePrivate).String(),
			))
			if err != nil {
				fmt.Println("Error writing to log file:", err)
			}

			// บันทึก body ของ request
			_, err = logFile.WriteString(fmt.Sprintf("Request Body: %s\n", string(body)))
			if err != nil {
				fmt.Println("Error writing request body to log file:", err)
			}
		}
	}
}
