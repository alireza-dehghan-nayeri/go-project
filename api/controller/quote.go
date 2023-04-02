package controller

import (
	"net/http"
	"strconv"

	"github.com/alireza-dehghan-nayeri/go-project/api/service"
	"github.com/alireza-dehghan-nayeri/go-project/models"
	"github.com/alireza-dehghan-nayeri/go-project/util"
	"github.com/gin-gonic/gin"
)

// QuoteController -> QuoteController
type QuoteController struct {
	service      service.QuoteService
	countService service.CountService
}

// NewQuoteController : NewQuoteController
func NewQuoteController(s service.QuoteService, countService service.CountService) QuoteController {
	return QuoteController{
		service:      s,
		countService: countService,
	}
}

// GetQuotes : GetQuotes controller
func (p QuoteController) GetQuotes(ctx *gin.Context) {
	var quotes models.Quote

	keyword := ctx.Query("keyword")

	data, total, err := p.service.FindAll(quotes, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Quote result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddQuote : AddQuote controller
func (p *QuoteController) AddQuote(ctx *gin.Context) {
	var quote models.Quote
	ctx.ShouldBindJSON(&quote)

	if quote.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if quote.Body == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	count, err := p.countService.Count(quote.Body)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to count")
		return
	}

	quote.CharCount = count
	response := count

	err1 := p.service.Save(quote)
	if err1 != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create quote")
		return
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Created quote",
		Data:    &response})
}

// GetQuote : get quote by id
func (p *QuoteController) GetQuote(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var quote models.Quote
	quote.ID = id
	foundQuote, err := p.service.Find(quote)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Quote")
		return
	}

	response := foundQuote.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Quote",
		Data:    &response})

}

// DeleteQuote : Deletes Quote
func (p *QuoteController) DeleteQuote(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.Delete(id)

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Quote")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

// UpdateQuote : get update by id
func (p QuoteController) UpdateQuote(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var quote models.Quote
	quote.ID = id

	quoteRecord, err := p.service.Find(quote)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Quote with given id not found")
		return
	}
	ctx.ShouldBindJSON(&quoteRecord)

	if quoteRecord.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if quoteRecord.Body == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	if err := p.service.Update(quoteRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Quote")
		return
	}
	response := quoteRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Quote",
		Data:    response,
	})
}
