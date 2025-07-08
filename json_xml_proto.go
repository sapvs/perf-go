package perf

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/vsapan/perf-go/gen"
	"google.golang.org/protobuf/proto"
)

type User struct {
	XMLName xml.Name `json:"-" xml:"user"`
	ID      uint64   `json:"userid" xml:"userid"`
	Name    string   `json:"username" xml:"username"`
}

func (u *User) String() string { return fmt.Sprintf("ID=%d Name=%s", u.ID, u.Name) }

func JSONToStruct(userString string) (*User, error) {
	var user = &User{}
	err := json.Unmarshal([]byte(userString), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func StructToJSON(user *User) (string, error) {
	bytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func XMLToStruct(userString string) (*User, error) {
	var user = &User{}

	err := xml.Unmarshal([]byte(userString), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func StructToXML(user *User) (string, error) {
	bytes, err := xml.Marshal(user)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ProtoToStruct(by []byte) (*gen.UserP, error) {
	var user gen.UserP

	err := proto.Unmarshal(by, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func StructToProto(user *gen.UserP) ([]byte, error) {
	by, err := proto.Marshal(user)
	if err != nil {
		return nil, err
	}
	return by, nil
}
