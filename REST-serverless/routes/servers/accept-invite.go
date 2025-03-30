package servers_route

import (
	db "REST-serverless/db"
	"REST-serverless/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"time"
)

func AcceptInviteTransaction(c *gin.Context, dbConn *pgxpool.Pool, queries *db.Queries, code string, userID int64) (error, *db.ServerToUserMapping) {
	ctx := c.Request.Context()

	tx, err := dbConn.Begin(ctx)
	if err != nil {
		return err, nil
	}
	defer func() {
		if err != nil { // Rollback only if an error occurs
			_ = tx.Rollback(ctx)
		}
	}()

	qtx := queries.WithTx(tx)

	invite, err := qtx.GetServerInvite(ctx, code)
	if err != nil {
		return err, nil
	}

	validForever := invite.Uses.Int32 == -1 && invite.Expiry.Int64 == -1
	validUsesLeft := invite.Uses.Int32 > 0
	validExpiry := invite.Expiry.Int64 > time.Now().Unix() || invite.Expiry.Int64 == -1

	if validForever || (validUsesLeft && validExpiry) {
		mapping, err := qtx.CreateServerToUserMapping(ctx, db.CreateServerToUserMappingParams{
			UserID:   userID,
			ServerID: invite.ServerID,
			Role:     "member",
		})
		if err != nil {
			return err, nil
		}

		if invite.Uses.Int32 > 0 {
			err = qtx.DecrementInviteUses(ctx, code)
			if err != nil {
				return err, nil
			}
		}

		err = tx.Commit(ctx)
		if err != nil {
			return err, nil
		}

		return nil, &mapping
	}

	return errors.New("invite expired or has no uses left"), nil
}

func AcceptInvite() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")

		userId, err := utils.ExtractUserIDFromContext(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err, d := AcceptInviteTransaction(c, db.Pool, db.DBQueries, code, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "id": d.UserID})
	}
}
