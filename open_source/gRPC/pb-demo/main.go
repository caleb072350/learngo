package main

import (
	"os"
	"strconv"

	pb "pb-demo/address/addressbook"

	"google.golang.org/protobuf/proto"
)

func main() {
	// 初始化AddressBook
	p := pb.Person{
		Id:    1234,
		Name:  "Isaac",
		Email: "123@123.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	addressBook := &pb.AddressBook{
		People: []*pb.Person{&p},
	}

	//将AddressBook对象序列化
	out, err := proto.Marshal(addressBook)
	if err != nil {
		panic(err)
	}

	// 将序列化后的数据写入文件
	if err = os.WriteFile("address_obj", out, 0644); err != nil {
		panic(err)
	}

	// 从文件读取数据
	in, err := os.ReadFile("address_obj")
	if err != nil {
		panic(err)
	}

	// 将数据反序列化为对象
	address := &pb.AddressBook{}

	if err = proto.Unmarshal(in, address); err != nil {
		panic(err)
	}

	// println(address.People[0].Id)
	// println(address.People[0].Name)

	for index, person := range address.People {
		print(strconv.Itoa(index) + ": ")
		print(person.Id)
		print(" ")
		print(person.Name)
		print(" ")
		print(person.Email)
		print(" ")

		for j, phone := range person.Phones {
			print("Phone" + strconv.Itoa(j) + ": " + phone.Number)
		}
		println()
	}

}
