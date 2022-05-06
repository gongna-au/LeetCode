package main
import "fmt"
"net/http"

"github.com/gorilla/sessions"
var(
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)
