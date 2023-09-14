package agent

import (
	"bufio"
	"fmt"
	"github.com/firesworder/password_saver/internal/grpcagent"
	"os"
)

// Agent экземпляр агента для вызова в cmd/agent
type Agent struct {
	state     *state
	grpcAgent grpcagent.IGRPCAgent
	reader    *bufio.Reader
	isAuth    bool
}

// NewAgent конструктор агента, формирует пустой стейт польз.данных + создает экземпляр grpc агента.
func NewAgent(grpcAgent grpcagent.IGRPCAgent) (*Agent, error) {
	a := &Agent{state: newState(), grpcAgent: grpcAgent}
	a.reader = bufio.NewReader(os.Stdin)
	return a, nil
}

// Serve запуска агента на обработку команд пользователя.
func (a *Agent) Serve() error {
	for {
		var command string
		if _, err := fmt.Scan(&command); err != nil {
			return err
		}

		switch command {
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
			fmt.Println("unknown command")
		}
	}
}
