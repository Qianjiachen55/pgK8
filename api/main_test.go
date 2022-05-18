package api

import (
	db "github.com/Qianjiachen55/pgK8/db/sqlc"
	"github.com/Qianjiachen55/pgK8/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func newTestServer(t *testing.T,store db.Store) *Server{
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config,store)
	require.NoError(t, err)


	return server
}

func TestMain(m *testing.M){

	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

