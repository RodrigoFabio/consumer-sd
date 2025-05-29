package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func EnviarEmail(agendamentoJson string) {

	var agendamento Agendamentos

	err := json.Unmarshal([]byte(agendamentoJson), &agendamento)
	if err != nil {
		log.Println("Erro ao fazer o unmarshall do JSON:", err)
	}

	// Configurações do servidor SMTP
	from := os.Getenv("FROM_APP_GMAIL")
	senha_app := os.Getenv("SENHA_APP")

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

}

func MontaMensagem(agendamento Agendamentos) string {

	msg := fmt.Sprintf(`Olá, %s! Tudo bem?

	Somos da EXAMED e gostaríamos de informar que o seu exame de %s foi agendado com sucesso para o dia %s.
	
	
	Para que tudo ocorra bem, recomendamos que voce siga as seguintes instruções e preparativos:
	%s`, agendamento.Paciente, agendamento.NomeExame, agendamento.DataHora, agendamento.Instrucoes)

	return msg
}

func GetStringConnFila() string {
	ip_fila := os.Getenv("IP_FILA")
	porta_fila := os.Getenv("PORTA_FILA")
	usuario_fila := os.Getenv("USUARIO_FILA")
	senha_fila := os.Getenv("PASS_FILA")

	return fmt.Sprintf("amqp://%s:%s@%s:%s/", usuario_fila,
		senha_fila, ip_fila, porta_fila)
}
