package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/tls"
	"github.com/hashicorp/vault/helper/pluginutil"
	"github.com/hashicorp/vault/plugins"
	"os"
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"testing/quick"
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

//noinspection ALL
func mainn() {
	crypto.SHA3_512()
	args := os.Args


}
func main() {
	apiClientMeta := &pluginutil.APIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args)

	plugins.Serve(New().(MyPlugin), apiClientMeta.GetTLSConfig())
}
