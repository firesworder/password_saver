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

type Agent struct {
	state     *state
	grpcAgent *grpcagent.GRPCAgent
	stdin     io.Reader // для тестов
	isAuth    bool
}

func NewAgent(agentEnv *env.Environment) (*Agent, error) {
	a := &Agent{state: newState()}
	grpcAgent, err := grpcagent.NewGRPCAgent(agentEnv.ServerAddress, agentEnv.CACert)
	if err != nil {
		return nil, err
	}
	a.grpcAgent = grpcAgent
	return a, nil
}

func (a *Agent) Serve() error {
	for {
		var command string
		if _, err := fmt.Scan(&command); err != nil {
			return err
		}

		switch command {
		case "register_user":
			a.RegisterUserCommand()
		case "login_user":
			a.LoginUserCommand()
		case "create_record":
			a.CreateRecordCommand()
		case "open_record":
			a.OpenRecordCommand()
		case "update_record":
			a.UpdateRecordCommand()
		case "delete_record":
			a.DeleteRecordCommand()
		case "show_all_records":
			a.ShowAllRecordsCommand()
		case "help":
			a.HelpCommand()
		default:
			fmt.Println("unknown command")
		}
	}
}
