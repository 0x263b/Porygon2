package bot

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/steveyen/gkvlite"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var (
	kvFile   *os.File
	KV       *gkvlite.Store
	Users    *gkvlite.Collection
	Channels *gkvlite.Collection
)

func initkv() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	kvFile, _ = os.OpenFile(fmt.Sprintf("%s/database.gkvlite", dir), os.O_RDWR|os.O_CREATE, 0666)
	KV, _ = gkvlite.NewStore(kvFile)
	Users = KV.SetCollection("users", nil)
	Channels = KV.SetCollection("channels", nil)
}

func getInterface(bts []byte, data interface{}) error {
	buf := bytes.NewBuffer(bts)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getKeyString(nick string) map[string]string {
	decoded := make(map[string]string)
	user, _ := Users.Get([]byte(nick))
	if user == nil {
		return nil
	}
	err := getInterface(user, &decoded)
	if err != nil {
		return nil
	}
	return decoded
}

func GetUserKey(nick string, key string) string {
	nick = strings.ToLower(nick)
	user := getKeyString(nick)
	return user[key]
}

func SetUserKey(nick string, key string, value string) {
	nick = strings.ToLower(nick)
	user := getKeyString(nick)
	if user == nil {
		user = map[string]string{
			key: value,
		}
	} else {
		user[key] = value
	}

	val := reflect.ValueOf(user)
	in := val.Interface()
	byt, _ := getBytes(in)

	Users.Set([]byte(nick), byt)
	KV.Flush()
}

func DeleteUserKey(nick string, key string) {
	nick = strings.ToLower(nick)
	user := getKeyString(nick)
	if user == nil {
		return
	}
	delete(user, key)

	val := reflect.ValueOf(user)
	in := val.Interface()
	byt, _ := getBytes(in)

	Users.Set([]byte(nick), byt)
	KV.Flush()
}

func getKeyBool(channel string) map[string]bool {
	decoded := make(map[string]bool)
	chann, _ := Channels.Get([]byte(channel))
	if chann == nil {
		return nil
	}
	err := getInterface(chann, &decoded)
	if err != nil {
		return nil
	}
	return decoded
}

func GetChannelKey(channel string, key string) bool {
	channel = strings.ToLower(channel)
	chann := getKeyBool(channel)
	return chann[key]
}

func SetChannelKey(channel string, key string, value bool) {
	channel = strings.ToLower(channel)
	chann := getKeyBool(channel)
	if chann == nil {
		chann = map[string]bool{
			key: value,
		}
	} else {
		chann[key] = value
	}

	val := reflect.ValueOf(chann)
	in := val.Interface()
	byt, _ := getBytes(in)

	Channels.Set([]byte(channel), byt)
	KV.Flush()
}
