package main

import (
	"fmt"
	"os"

	listmonk "gomodules.xyz/listmonk-client-go"
)

func main() {
	c := listmonk.New(listmonk.ListmonkTesting, os.Getenv("LISTMONK_USERNAME"), os.Getenv("LISTMONK_PASSWORD"))
	lists, err := c.GetAllLists()
	if err != nil {
		panic(err)
	}
	for _, l := range lists {
		fmt.Printf("%+v\n", l)
	}

	ml, err := c.CreateListIfMissing(listmonk.MailingListRequest{
		Name:  "webinar-2021-04-15",
		Type:  listmonk.ListTypePublic,
		Optin: listmonk.OptinModeDouble,
		Tags: []string{
			"webinar",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", ml)

	err = c.SubscribeToList(listmonk.SubscribeRequest{
		Email:        "tamal@appscode.com",
		Name:         "Tamal Saha",
		MailingLists: []string{ml.UUID},
	})
	if err != nil {
		panic(err)
	}
}
