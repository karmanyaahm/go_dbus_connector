package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/xakep666/ecego"
	up "unifiedpush.org/go/dbus_connector/api"
	"unifiedpush.org/go/dbus_connector/definitions"
)

var Endpoint string
var priv *ecdsa.PrivateKey
var auth []byte

type NotificationHandler struct{}

func decryptECE(privKey *ecdsa.PrivateKey, authSecret []byte, ciphertext []byte) (plaintext []byte, err error) {
	return ecego.NewEngine(ecego.SingleKey(privKey), ecego.WithAuthSecret(authSecret[:16])).Decrypt(ciphertext, nil, ecego.OperationalParams{})
}

func (n NotificationHandler) Message(instance string, iMessage []byte, id string) {
	fmt.Println("new message received")

	decoded, err := decryptECE(priv, auth, iMessage) // any salt, ecego will decode actual salt later
	if err != nil {
		log.Println(err)
	}
	message := string(decoded)

	// this message can be in whatever format you like, in this case the title and message body are two strings seperated by a '-'
	parts := strings.Split(message, "-")

	title := "No Title Provided"
	if len(parts) > 1 {
		title = parts[1]
	}

	err = beeep.Notify(title, parts[0], "")
	if err != nil {
		panic(err)
	}
}

func (n NotificationHandler) NewEndpoint(instance, endpoint string) {
	// the endpoint should be sent to whatever server your app is using
	Endpoint = endpoint
	fmt.Println("New endpoint received", Endpoint)
	http.Post(endpoint, "", strings.NewReader("body-title"))
}

func (n NotificationHandler) Unregistered(instance string) {
	Endpoint = ""
	fmt.Println("endpoint unregistered", Endpoint)
}

func main() {

	mycurve := elliptic.P256()
	priv, _ = ecdsa.GenerateKey(mycurve, bytes.NewBufferString("abcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefgh"))
	pubkey := base64.URLEncoding.EncodeToString(elliptic.Marshal(mycurve, priv.X, priv.Y))
	auth = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16} // 16 bytes long auth value buffer
	_, _ = rand.Read(auth)
	fmt.Printf("p256dh key: %s ; Auth Secret: %s\n", pubkey, base64.StdEncoding.EncodeToString(auth))

	connector := NotificationHandler{}
	up.InitializeAndCheck("cc.malhotra.karmanyaah.testapp.golibrary.webpush", connector)

	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "unregister":
			err := up.Unregister("")
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
	}

	if len(up.GetDistributor()) == 0 { // not picked distributor yet
		pickDist()
	}
	// run this for each instance on each application startup to get the most up-to-date info
	result, reason, err := up.Register("")
	if err != nil {
		panic(err)
	}
	switch result {
	case definitions.RegisterStatusFailed:
		fmt.Println("registration failed because", reason)
		return
	case definitions.RegisterStatusRefused:
		fmt.Println("Registration refused", reason)
		return
	default:
		fmt.Println("will receive registration soon", reason)
	}

	// do whatever your app does
	fmt.Println("app waiting now")
	<-make(chan struct{})
}

func pickDist() {
	dist, err := up.GetDistributors()
	if err != nil {
		panic(err)
	}
	fmt.Println(dist)

	var distributor string

	if len(dist) == 0 {
		fmt.Println("No distributor so can't be push notifications, exiting")
		os.Exit(0)
	} else if len(dist) == 1 {
		distributor = dist[0]
		fmt.Println("Picking the only distributor available", distributor)
	} else {
		fmt.Println("avalible distributors")
		for i, j := range dist {
			fmt.Println(i, j)
		}
		fmt.Print("Pick one distributor by number  ")
		var num int
		fmt.Scanln(&num)
		distributor = dist[num]
		fmt.Println("Picked distributor", distributor)
	}
	err = up.SaveDistributor(distributor)
	if err != nil {
		panic(err)
	}
}
