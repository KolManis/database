// Координирует работу парсера и подготавливает запрос для storage
package compute

import (
	"errors"

	"go.uber.org/zap"
)

type Compute struct {
	parser *Parser
	logger *zap.Logger
}

func NewCompute(logger *zap.Logger) (*Compute, error) {
	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	return &Compute{
		parser: NewParser(),
		logger: logger,
	}, nil
}

func (c *Compute) Process(input string) (Query, error) {
	// 1. Логируем входящий запрос
	c.logger.Info("Received query", zap.String("input", input))

	// 2. Парсим запрос
	query, err := c.parser.Parse(input)
	if err != nil {
		c.logger.Error("Failed to parse query",
			zap.String("input", input),
			zap.Error(err))
		return Query{}, err
	}

	return query, nil
}
