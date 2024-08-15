package apis

import (
	"net/http"

	"github.com/Evilcmd/Hackup-backend/internal/utilities"
)

func AdminCheck(res http.ResponseWriter, req *http.Request) {
	utilities.RespondWithJson(res, 200, "Hello from Admin")
}
