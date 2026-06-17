package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}
type VIPUser struct {
	user     User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "Hwarang", 26},
		3,
		250,
	}
	fmt.Println(user.Name, " ", user.ID, " ", user.Age)
	fmt.Println(vip.user.Name, " ", vip.user.ID, " ", vip.user.ID, " ", vip.VIPLevel, " ", vip.Price)
}
