package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/go-counting-stars/pkg/go-badge"
)

func main() {

	type Badge struct {
		Template string `uri:"template" binding:"required"`
		Subject  string `uri:"subject" binding:"required"`
		Status   string `uri:"status" binding:"required"`
		Color    string `uri:"color" binding:"required"`
	}

	route := gin.Default()
	route.GET("/:template/:subject/:status/:color/", func(c *gin.Context) {
		var reqBadge Badge
		if err := c.ShouldBindUri(&reqBadge); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		badgeBytes, err := badge.RenderBytes(reqBadge.Template, reqBadge.Subject, reqBadge.Status, badge.Color(reqBadge.Color))
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		c.Header("Content-Type", "image/svg+xml")
		c.Data(200, "image/svg+xml", badgeBytes)
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		route.Run(":8080")
	}()
	log.Println("Program has been started 🚀")

	<-ctx.Done()

	stop()
	log.Println("The program is closing, to end immediately press CTRL+C")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

}
