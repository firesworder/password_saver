package agentcommands

func (ac *AgentCommands) RegisterUser() {
	ac.writer.WriteString(enterAuthData)
	fields, err := ac.reader.GetUserFields()
	if len(fields) != 2 {
		ac.writer.WriteErrorString("input error, required 2 fields")
		return
	}
	login, password := fields[0], fields[1]

	if err = ac.grpcAgent.RegisterUser(login, password); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	ac.isAuthorized = true
}

func (ac *AgentCommands) LoginUser() {
	ac.writer.WriteString(enterAuthData)
	fields, err := ac.reader.GetUserFields()
	if len(fields) != 2 {
		ac.writer.WriteErrorString("input error, required 2 fields")
		return
	}
	login, password := fields[0], fields[1]

	if err = ac.grpcAgent.LoginUser(login, password); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	ac.isAuthorized = true
}
