package answgen

import "context"

type AnswerGenerator interface {
	GenerateAnswer(ctx context.Context, prompt string, info []string) (string, error)
}