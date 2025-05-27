package main

import (
	"encoding/json"
	"fmt"
	"net/smtp"
)

type Agendamentos struct {
	DataHora      string `json:"data_hora"`
	NomeExame     string `json:"nome_exame"`
	Instrucoes    string `json:"instrucoes"`
	Paciente      string `json:"nome_paciente"`
	EmailPaciente string `json:"email_paciente"`
}

func EnviarEmail(agendamentoJson string) {

	var agendamento Agendamentos

	err := json.Unmarshal([]byte(agendamentoJson), &agendamento)
	if err != nil {
		fmt.Println("Erro ao fazer o unmarshall do JSON:", err)
	}

	// Configurações do servidor SMTP
	from := "examedsd@gmail.com"
	senha_app := "ywcs uxja mjrm ziit"
	fmt.Print(":::::::::ENVIANDO EMAIL:::::::::")
	body := MontaMensagem(agendamento)
	to := []string{agendamento.EmailPaciente}
	fmt.Print(body)
	errr := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, senha_app, "smtp.gmail.com"),
		from,
		to,
		[]byte(body))

	if errr != nil {
		fmt.Print(errr)
	}

	// Exemplo de impressão dos dados do agendamento
	fmt.Println("Enviando e-mail com os seguintes dados:")
	fmt.Printf("Data e Hora: %s\n", agendamento.DataHora)
	fmt.Printf("ID do Exame: %s\n", agendamento.NomeExame)
	fmt.Printf("Instruções: %s\n", agendamento.Instrucoes)
	fmt.Printf("Nome do Paciente: %s\n", agendamento.Paciente)
	fmt.Printf("Email do Paciente: %s\n", agendamento.EmailPaciente)
}

func MontaMensagem(agendamento Agendamentos) string {

	msg := fmt.Sprintf(`Olá, %s! Tudo bem?
	Somos da EXAMED e gostaríamos de informar que o seu exame de %s foi agendado com sucesso para o dia %s. 
	Para que tudo ocorra bem, recomendamos que voce siga as seguintes instruções e preparativos:
	%s`, agendamento.Paciente, agendamento.NomeExame, agendamento.DataHora, agendamento.Instrucoes)

	return msg
}
