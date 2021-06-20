package web

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	//samlmw "go-saml-dev/samlMiddleware"
	//"github.com/spf13/viper"
	//log "github.com/sirupsen/logrus"
)

func init() {
	assertAvailablePRNG()
}

type Stash struct {
	AppPrefix       string
	IsAuthenticated bool
	Username        string
	Roles           []string
	Values          map[string]string
	Nonce           string
}

func (x *WebServerPlugin) initStash(r *http.Request) *Stash {
	vals := make(map[string]string)
	//prefix := viper.GetString("app.prefix")
	prefix := ""
	//vals["username"]    = samlmw.GetAttribute(r, "username")
	vals["username"] = ""
	isAuthed := false
	username := ""

	roles := []string{}
	//if val, ok := vals["username"]; ok && val != "" {
	//   isAuthed = true
	//    username = val
	//    roles = samlmw.GetAttributeAsArray(r, "roles")
	//}
	nonce, err := GenerateRandomStringURLSafe(14)
	if err != nil {
		panic(err)
	}
	return &Stash{prefix, isAuthed, username, roles, vals, nonce}
}

func assertAvailablePRNG() {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand is unavailable: Read() failed with %#v", err))
	}
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomStringURLSafe(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}
