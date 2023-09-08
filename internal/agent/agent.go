// Package agent реализует взаимодействие между пользователем(консоль) и grpc агентом.
package agent

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/agent/env"
	"github.com/firesworder/password_saver/internal/grpcagent"
	"io"
)

func scanMetaInfo() (string, error) {
	var metaInfo string
	fmt.Println("Enter meta info")
	if _, err := fmt.Scan(&metaInfo); err != nil {
		return "", err
	}
	return metaInfo, nil
}

// Agent экземпляр агента для вызова в cmd/agent
type Agent struct {
	state     *state
	grpcAgent *grpcagent.GRPCAgent
	stdin     io.Reader // для тестов
	isAuth    bool
}

// NewAgent конструктор агента, формирует пустой стейт польз.данных + создает экземпляр grpc агента.
func NewAgent(agentEnv *env.Environment) (*Agent, error) {
	a := &Agent{state: newState()}
	grpcAgent, err := grpcagent.NewGRPCAgent(agentEnv.ServerAddress, agentEnv.CACert)
	if err != nil {
		return nil, err
	}
	a.grpcAgent = grpcAgent
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
