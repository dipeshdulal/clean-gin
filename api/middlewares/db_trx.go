package middlewares

import (
	"net/http"

	"github.com/dipeshdulal/clean-gin/constants"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/gin-gonic/gin"
)

// DatabaseTrx middleware for transactions support for database
type DatabaseTrx struct {
	handler lib.RequestHandler
	logger  lib.Logger
	db      lib.Database
}

// NewDatabaseTrx creates new database transactions middleware
func NewDatabaseTrx(
	handler lib.RequestHandler,
	logger lib.Logger,
	env lib.Env,
	db lib.Database,
) DatabaseTrx {
	return DatabaseTrx{
		handler: handler,
		logger:  logger,
		db:      db,
	}
}

// Setup sets up cors middleware
func (m DatabaseTrx) Setup() {
	m.logger.Zap.Info("Setting up cors middleware")

	m.handler.Gin.Use(func(c *gin.Context) {
		txHandle := m.db.DB.Begin()
		m.logger.Zap.Info("Build transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set(constants.DBTransaction, txHandle)
		c.Next()

		// rollback transaction when error
		if c.Writer.Status() == http.StatusInternalServerError {
			m.logger.Zap.Info("rolling back transaction due to status code: 500")
			txHandle.Rollback()
		}

		// commit transaction on ok status
		if c.Writer.Status() == http.StatusOK {
			m.logger.Zap.Info("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				m.logger.Zap.Error("trx commit error: ", err)
			}
		}
	})
}
