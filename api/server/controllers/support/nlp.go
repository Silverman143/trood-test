package support

import (
	"log/slog"
	"trood-test/api/server/response"
	"trood-test/api/server/utils"
	"trood-test/internal/services/nlp"

	"github.com/gin-gonic/gin"
)




type SupportContreoller struct{
    log *slog.Logger
    responseBuilder *response.ResponseBuilder
	nlpService *nlp.NLPService

}

func NewController(log *slog.Logger, responseBuilder *response.ResponseBuilder, nlpService *nlp.NLPService) *SupportContreoller{
    return &SupportContreoller{
        log: log,
        responseBuilder: responseBuilder,
        nlpService: nlpService,
    }
}

func (c *SupportContreoller) ProcessQuestion(ctx *gin.Context) {
    const op = "server.controllers.support.ProcessQuestion"
    logger := c.log.With(slog.String("op", op))

    var request ProcessQuestionRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error("failed to bind JSON", slog.Any("err", err))
		c.responseBuilder.BadRequest(ctx, "Invalid input")
		return
	}

    tokensData, err := c.nlpService.ProcessQuestion(utils.BuildInternalContext(ctx), request.ChatID, request.Text)
    if err != nil {
        c.log.Error("%s: processing question: %w", op, err)
        c.responseBuilder.InternalServerError(ctx, "internal server error")
        return
    }

    c.responseBuilder.Ok(ctx, tokensData)
}

