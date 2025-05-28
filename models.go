package main

type Agendamentos struct {
	DataHora      string `json:"data_hora"`
	NomeExame     string `json:"nome_exame"`
	Instrucoes    string `json:"instrucoes"`
	Paciente      string `json:"nome_paciente"`
	EmailPaciente string `json:"email_paciente"`
}

