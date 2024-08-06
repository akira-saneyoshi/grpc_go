package main

import (
	"fmt"
	"log"
	"os"
	"udemy_study/go_grpc/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "suzuki@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki.",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't serialize.", err)
	}

	if err := os.WriteFile("test.bin", binData, 0666); err != nil {
		log.Fatalln("Can't Write.", err)
	}

	in, err := os.ReadFile("test.bin")
	if err != nil {
		log.Fatalln("Can't read file.", err)
	}

	readEmployee := &pb.Employee{}

	err = proto.Unmarshal(in, readEmployee)
	if err != nil {
		log.Fatalln("Can't deserialize", err)
	}

	fmt.Println(readEmployee)
}
