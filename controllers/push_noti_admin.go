package controllers

import (
	"encoding/json"
	"log"
	"os"
	"ticketing-system/config"
	"ticketing-system/models"

	"github.com/SherClockHolmes/webpush-go"
)

func SendNotificationsToAdmin(title, body string) {
	var subs []models.Subscription
	if err := config.DB.Where("role=?", "admin").Find(&subs).Error; err != nil {
		log.Printf("Error fetching admin subscriptions: %v", err)
		return
	}
	if len(subs) == 0 {
		log.Println("No admin subscriptions found")
		return
	}
	payload, err := json.Marshal(map[string]string{
		"title": title,
		"body":  body,
	})
	if err != nil {
		log.Printf("Error marshaling push payload: %v", err)
		return
	}
	for _, sub := range subs {
		res, err := webpush.SendNotification(payload, &webpush.Subscription{
			Endpoint: sub.Endpoint,
			Keys: webpush.Keys{
				P256dh: sub.P256dh,
				Auth:   sub.Auth,
			},
		}, &webpush.Options{
			VAPIDPublicKey:  os.Getenv("VAPID_PUBLIC"),
			VAPIDPrivateKey: os.Getenv("VAPID_PRIVATE"),
			TTL:             30,
		})
		if err!=nil{
			log.Printf("Push error for endpoint %s: %v", sub.Endpoint,err)
			continue
		}
		if res!=nil{
			defer res.Body.Close()
		}
	}
	log.Fatal("Push notifications sent successfully to all admin subscriptions.")
}
