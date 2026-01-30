// Централизованное хранение информации о командах
package compute

// 1. ID команд (для внутреннего использования)
const (
	UnknownCommandID = iota //0
	SetCommandID
	GetCommandID
	DelCommandID
)

// 2. Имена команд (как их вводит пользователь)
const (
	UnkwownCommand = "UNKNOWN"
	SetCommand     = "SET"
	GetCommand     = "GET"
	DelCommand     = "DEL"
)

// 3. Маппинг: имя команды → её ID
var commandNameToID = map[string]int{
	SetCommand: SetCommandID,
	GetCommand: GetCommandID,
	DelCommand: DelCommandID,
}

func commandNameToCommandID(command string) int {
	status, found := commandNameToID[command]
	if !found {
		return UnknownCommandID
	}

	return status
}

// 4. Количество аргументов
const (
	setCommandArgumentsNumber = 2
	getCommandArgumentsNumber = 1
	delCommandArgumentsNumber = 1
)

// 5. Маппинг ID команды → количество аргументов
var argumentsNumber = map[int]int{
	SetCommandID: 2, // SET key value
	GetCommandID: 1, // GET key
	DelCommandID: 1, // DEL key
}

//мб сделать возврат ошибки или оставить так
//Если не найдет вернет 0 т.е вроде норм
func commandArgumentsNumber(commandID int) int {
	return argumentsNumber[commandID]
}
