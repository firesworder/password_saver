// Package agent реализует контроллер агента, который команды от пользователя отправляет в модуль agentcommands.
// Данный пакет только направляет процесс исполнения в связанный с командой пользователя метод agentcommands или
// возвращает ошибку, если команда не известна. Результат работы команды - выводится agentcommands.
//
// Тип Agent состоит из:
// - reader/writer реализованные через bufio.Reader/bufio.Writer(которые также передаются в agentcommands)
// Такая реализация необходима для тестирования функций контроллера и agentcommands, через эмуляцию ввода пользователя
// и считывания вывода результата обработки команды.
// - commands
// Пакет agentcommands является реализацией логики выполняемых команд, получаемых из контроллера.
// Основная идея такой реализации - разделить контроллер от деталей реализации.
package agent

import (
	"bufio"
	"context"
	"errors"
	"github.com/firesworder/password_saver/internal/agent/agentcommands"
	"github.com/firesworder/password_saver/internal/agent/agentcommands/grpcagent"
	"github.com/firesworder/password_saver/internal/agent/agentreader"
	"github.com/firesworder/password_saver/internal/agent/agentwriter"
	"io"
	"os"
	"strconv"
)

const enterDT = "Choose data type(enter name type): text, bank or binary"
const enterIDaDT = "Enter recordID and dataType"

// Agent тип контроллера агента, распределяющего команды полученные из reader для реализации в
// agentcommands.IAgentCommands и вывода в writer контроллера.
type Agent struct {
	reader   *agentreader.AgentReader
	writer   *agentwriter.AgentWriter
	commands agentcommands.IAgentCommands
}

// NewAgent конструктор агента, инициал. reader/writer в os.Stdin/os.Stdout соотв-но.
// А также инициал. commands(исполнителя) через конструктор agentcommands.NewAgentCommands.
// grpcAgent проксируется в конструктор commands.
func NewAgent(grpcAgent grpcagent.IGRPCAgent) *Agent {
	a := &Agent{
		reader: agentreader.NewAgentReader(bufio.NewReader(os.Stdin)),
		writer: agentwriter.NewAgentWriter(bufio.NewWriter(os.Stdout)),
	}
	a.commands = agentcommands.NewAgentCommands(grpcAgent, a.reader, a.writer)
	return a
}

// Serve функция запуска агента на постоянную обработку команд пользователя.
// Можно прекратить выполнение завершением контекста передаваемого в аргум.функции.
func (a *Agent) Serve(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			a.writer.WriteString("Enter command")
			command, err := a.reader.GetUserInput()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				a.writer.WriteErrorString(err.Error())
			}
			a.controller(command)
		}
	}
}

func (a *Agent) controller(command string) {
	switch command {
	case "register_user":
		a.commands.RegisterUser()
	case "login_user":
		a.commands.LoginUser()
	case "create_record":
		a.createRecordCommand()
	case "open_record":
		a.openRecordCommand()
	case "update_record":
		a.updateRecordCommand()
	case "delete_record":
		a.deleteRecordCommand()
	case "show_all_records":
		a.commands.ShowAllRecords()
	case "help":
		a.helpCommand()
	default:
		a.writer.WriteErrorString("unknown command")
	}
}

func (a *Agent) createRecordCommand() {
	a.writer.WriteString(enterDT)
	dataType, err := a.reader.GetUserInput()
	if err != nil {
		a.writer.WriteErrorString(err.Error())
		return
	}

	switch dataType {
	case "text":
		a.commands.CreateTextData()
	case "bank":
		a.commands.CreateBankData()
	case "binary":
		a.commands.CreateBinaryData()
	default:
		a.writer.WriteErrorString("unknown data type")
		return
	}
}

func (a *Agent) openRecordCommand() {
	a.writer.WriteString(enterIDaDT)
	fields, err := a.reader.GetUserFields()
	if err != nil {
		a.writer.WriteErrorString(err.Error())
		return
	}
	if len(fields) != 2 {
		a.writer.WriteErrorString("input error")
		return
	}
	recordID, err := strconv.Atoi(fields[0])
	if err != nil {
		a.writer.WriteErrorString(err.Error())
		return
	}
	dataType := fields[1]

	switch dataType {
	case "text":
		a.commands.OpenTextData(recordID)
	case "bank":
		a.commands.OpenBankData(recordID)
	case "binary":
		a.commands.OpenBinaryData(recordID)
	default:
		a.writer.WriteErrorString("unknown data type")
		return
	}
}

func (a *Agent) updateRecordCommand() {
	a.writer.WriteString(enterIDaDT)
	fields, err := a.reader.GetUserFields()
	if err != nil {
		a.writer.WriteErrorString(err.Error())
		return
	}
	if len(fields) != 2 {
		a.writer.WriteErrorString("input error")
		return
	}
	recordID, err := strconv.Atoi(fields[0])
	if err != nil {
		a.writer.WriteErrorString(err.Error())
		return
	}
	dataType := fields[1]

	switch dataType {
	case "text":
		a.commands.UpdateTextData(recordID)
	case "bank":
		a.commands.UpdateBankData(recordID)
	case "binary":
		a.commands.UpdateBinaryData(recordID)
	default:
		a.writer.WriteErrorString("unknown data type")
		return
	}
}

func (a *Agent) deleteRecordCommand() {
	a.writer.WriteString(enterIDaDT)
	fields, err := a.reader.GetUserFields()
	if err != nil {
		a.writer.WriteErrorString(err.Error())
		return
	}
	if len(fields) != 2 {
		a.writer.WriteErrorString("input error")
		return
	}
	recordID, err := strconv.Atoi(fields[0])
	if err != nil {
		a.writer.WriteErrorString(err.Error())
		return
	}
	dataType := fields[1]

	switch dataType {
	case "text":
		a.commands.DeleteTextData(recordID)
	case "bank":
		a.commands.DeleteBankData(recordID)
	case "binary":
		a.commands.DeleteBinaryData(recordID)
	default:
		a.writer.WriteErrorString("unknown data type")
		return
	}
}

func (a *Agent) helpCommand() {
	a.writer.WriteString(`Commands:
Auth methods:
- register_user, login_user

User data methods(required auth!):
- create_record, open_record, update_record, delete_record
- show_all_records
`)
}
