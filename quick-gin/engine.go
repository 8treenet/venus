package quickgin

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// New .
func New() *QuickEngine {
	engine := &QuickEngine{}
	engine.Engine = gin.New()
	return engine
}

// QuickEngine .
type QuickEngine struct {
	*gin.Engine
}

// RunH2C .
func (engine *QuickEngine) RunH2C(addr ...string) (err error) {
	address := engine.resolveAddress(addr)
	h2cSer := &http2.Server{}
	ser := &http.Server{
		Addr:    address,
		Handler: h2c.NewHandler(engine, h2cSer),
	}
	if gin.IsDebugging() {
		fmt.Fprintf(gin.DefaultWriter, "[GIN-debug] "+fmt.Sprintf("Listening and serving HTTP on %s\n", address))
	}
	return ser.ListenAndServe()
}

func (engine *QuickEngine) resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			return ":" + port
		}
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too much parameters")
	}
}
