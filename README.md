# VK API for Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/severecloud/vksdk)](https://goreportcard.com/report/github.com/severecloud/vksdk)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSevereCloud%2Fvksdk.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FSevereCloud%2Fvksdk?ref=badge_shield)

Внимание - этот репозиторий в **очень ранней разработке**.

## Установка

```shell
go get -u github.com/severecloud/vkapi
```

## Readme

- [5.92](https://github.com/SevereCloud/vksdk/tree/master/5.92)

## Пример

```go
package main

import (
	"fmt"
	"log"

	"github.com/severecloud/vksdk/5.92/api"
)

func main() {
	vk := vksdk.Init("<TOKEN>")

	params := make(map[string]string)
	params["user_ids"] = "1"

	users, vkErr := vk.UsersGet(params)
	if vkErr.Code != 0 {
		log.Fatal(vkErr.Message)
	}

	for _, user := range users {
		fmt.Printf("Пользователя с id%d зовут %s %s\n", user.ID, user.FirstName, user.LastName)
	}
}
```

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSevereCloud%2Fvksdk.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FSevereCloud%2Fvksdk?ref=badge_large)