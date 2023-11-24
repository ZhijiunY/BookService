package helper

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func SessionSet(c *gin.Context, userID uuid.UUID) {
	session := sessions.Default(c)
	var idInterface interface{} = &userID
	session.Set("id", idInterface)
	session.Save()
}

func SessionGet(c *gin.Context) uint64 {
	session := sessions.Default(c)
	return session.Get("id").(uint64)
}

func SessionClear(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
