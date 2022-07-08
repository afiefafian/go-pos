package helper

import (
	"strconv"

	"github.com/afiefafian/go-pos/src/model"
	"github.com/gofiber/fiber/v2"
)

// Pagination default config
const defaultPaginationLimit = 10
const defaultPaginationSkip = 0

// NewPaginationQueryFromCtx return a model.PaginationQuery struct.
// The source data of pagination is from limit and skip query string data
// get from gofiber context and mapped to model.PaginationQuery struct.
func NewPaginationQueryFromCtx(ctx *fiber.Ctx) *model.PaginationQuery {
	var (
		limit int = defaultPaginationLimit
		skip  int = defaultPaginationSkip
	)

	limitStr := ctx.Query("limit")
	skipStr := ctx.Query("skip")

	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	if skipStr != "" {
		skip, _ = strconv.Atoi(skipStr)
	}

	return &model.PaginationQuery{
		Limit: limit,
		Skip:  skip,
	}
}
