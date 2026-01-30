// Преобразовать сырую строку в структурированный Query
package compute

import (
	"errors"
	"strings"
)

var (
	errInvalidQuery     = errors.New("empty query")
	errInvalidCommand   = errors.New("invalid command")
	errInvalidArguments = errors.New("invalid arguments")
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) (Query, error) {
	// 1. Разбиваем на слова по пробелам
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return Query{}, errInvalidQuery
	}

	// 2. Определяем команду (первое слово)
	commandName := tokens[0]
	commandID := commandNameToCommandID(commandName)
	if commandID == UnknownCommandID {
		return Query{}, errInvalidCommand
	}

	// 3. Проверяем количество аргументов
	arguments := tokens[1:] // все слова кроме первого
	// токенс не будет очищен, т.к на этот же участок ссылается arguments
	// если в будущем делать более сложные комманды, мб копировать, чтобы сборщик убрал мусор
	expectedArguments := commandArgumentsNumber(commandID)
	if len(arguments) != expectedArguments {
		return Query{}, errInvalidArguments
	}

	// 4. Упаковываем в Query
	return NewQuery(commandID, arguments), nil
}
