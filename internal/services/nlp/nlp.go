package nlp

import (
	"context"
	"log/slog"
	answgen "trood-test/internal/answer_generator"
	"trood-test/internal/domain/models"
	embgen "trood-test/internal/embedding_generator"
	eventdispatcher "trood-test/internal/event_dispatcher"
	"trood-test/internal/repository"
)


type NLPService struct{
	logger *slog.Logger
	knowledgeDB repository.IKnowledgeRepo
	embgen  embgen.EmbeddingGenerator
	answgen answgen.AnswerGenerator
	eventDispatcher eventdispatcher.EventDispatcher
}

const (
	maxSearchResults = 10
)

func New(
	log *slog.Logger,
	knowledgeDB repository.IKnowledgeRepo,
	embgen  embgen.EmbeddingGenerator,
	answgen answgen.AnswerGenerator,
) *NLPService {
	logger := log.With("service", "NPL")
	return &NLPService{
		logger: logger,
		knowledgeDB: knowledgeDB,
		embgen: embgen,
		answgen: answgen,
	}
}

func (s *NLPService) ProcessQuestion(ctx context.Context, chatID int64, questionText string) (string, error) {
	const operationName = "ProcessQuestion"
	logger := s.logger.With(operationName, operationName)

	questionEmbedding, err := s.embgen.GenerateEmbedding(ctx, questionText)
	if err != nil {
		logger.Error("failed to generate embedding", "error", err)
		return "", ErrEmbeddingGeneration
	}

	similarKnowledgeItems, err := s.knowledgeDB.SearchSimilar(ctx, questionEmbedding, maxSearchResults)
	if err != nil {
		logger.Error("failed to search similar knowledge items", "error", err)
		return "", ErrKnowledgeSearch
	}

	if len(similarKnowledgeItems) == 0{
		logger.Debug("No data on the current user question ")

		unresolvedIntentEvent := eventdispatcher.UnresolvedIntentEvent{
			ChatID: chatID,
		}

		if err := s.eventDispatcher.Dispatch(ctx, &unresolvedIntentEvent); err != nil{
			logger.Error("failed to send kafka event", "error", err)
			return "", ErrUnresolvedIntent
		}
		return "Unfortunately I can't answer your question. A member of our staff will be in touch with you in this chat shortly. Stay in touch ", nil
	}

	knowledgeContext := models.ExtractContents(similarKnowledgeItems)

	answerText, err := s.answgen.GenerateAnswer(ctx, questionText, knowledgeContext)
	if err != nil {
		logger.Error("failed to generate completion", "error", err)
		return "", ErrAnswerGeneration
	}

	return answerText, nil
}