package main

import "fmt"

func main() {
	// TODO: Create a NotificationBuilder and use it to set properties
	var bldr = newNotificationBuilder()
	// TODO: Use the builder to set some properties
	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.png")
	bldr.SetPriority(10)
	bldr.SetMessage("This is a basic Notification")
	bldr.SetType("alert")
	// TODO: Use the Build function to create a finished object

	notif, err := bldr.Build()

	if err != nil {
		fmt.Println("Error creatinng the notification:", err)
	} else {
		fmt.Println("Notification: %+V", notif)
	}

}
