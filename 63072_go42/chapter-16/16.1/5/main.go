package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"go42/chapter-16/16.1/5/pb"
)

func main() {
	//初始化message UserInfo
	usermsg := &pb.UserInfo{
		UserType: 1,
		UserName: "Jim",
		UserInfo: "I am a woker!",
	}

	//序列化
	userdata, err := proto.Marshal(usermsg)
	if err != nil {
		fmt.Println("Marshaling error: ", err)
	}

	//反序列化
	encodingmsg := &pb.UserInfo{}
	err = proto.Unmarshal(userdata, encodingmsg)

	if err != nil {
		fmt.Println("Unmarshaling error: ", err)
	}

	fmt.Printf("GetUserType: %d\n", encodingmsg.GetUserType())
	fmt.Printf("GetUserName: %s\n", encodingmsg.GetUserName())
	fmt.Printf("GetUserInfo: %s\n", encodingmsg.GetUserInfo())
}
