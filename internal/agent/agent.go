package agent

import (
	"bufio"
	"github.com/firesworder/password_saver/internal/grpcagent"
	"log"
	"os"
	"strings"
)

// Agent экземпляр агента для вызова в cmd/agent
type Agent struct {
	state     *state
	grpcAgent grpcagent.IGRPCAgent
	reader    *bufio.Reader
	writer    *bufio.Writer
	isAuth    bool
}

// NewAgent конструктор агента, формирует пустой стейт польз.данных + создает экземпляр grpc агента.
func NewAgent(grpcAgent grpcagent.IGRPCAgent) (*Agent, error) {
	a := &Agent{state: newState(), grpcAgent: grpcAgent}
	a.reader = bufio.NewReader(os.Stdin)
	a.writer = bufio.NewWriter(os.Stdout)
	return a, nil
}

func (a *Agent) writeString(str string) {
	a.writer.WriteString(str + "\n")
	a.writer.Flush()
}

func (a *Agent) writeErrorString(errStr string) {
	a.writer.WriteString("err: " + errStr + "\n")
	a.writer.Flush()
}

// Serve запуска агента на обработку команд пользователя.
func (a *Agent) Serve() {
	for {
		command, err := a.reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		switch strings.TrimSpace(command) {
		case "register_user":
			a.registerUserCommand()
		case "login_user":
			a.loginUserCommand()
		case "create_record":
			a.createRecordCommand()
		case "open_record":
			a.openRecordCommand()
		case "update_record":
			a.updateRecordCommand()
		case "delete_record":
			a.deleteRecordCommand()
		case "show_all_records":
			a.showAllRecordsCommand()
		case "help":
			a.helpCommand()
		default:
			a.writeErrorString("unknown command")
		}
	}
}
