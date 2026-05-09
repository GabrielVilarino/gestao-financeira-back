package jobs

import (
	"time"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
)

// StartRecorrenciaJob inicia a goroutine que processa recorrências.
// Roda imediatamente ao iniciar e depois a cada 24 horas.
func StartRecorrenciaJob() {
	go func() {
		for {
			processarRecorrencias()
			esperarProximaMeiaNoite()
		}
	}()
}

func processarRecorrencias() {
	configs.Log.Info("Job de recorrências iniciado")

	recorrencias, err := repository.GetRecorrenciasAtivasRepository()
	if err != nil {
		configs.Log.Error("Erro ao buscar recorrências ativas: ", err)
		return
	}

	hoje := time.Now()
	horizonte := hoje.AddDate(0, 0, 30)

	for _, rec := range recorrencias {
		if err := gerarTransacoesPendentes(rec, hoje, horizonte); err != nil {
			configs.Log.Errorf("Erro ao gerar transações para recorrência %d: %v", rec.ID, err)
		}
	}

	configs.Log.Info("Job de recorrências concluído")
}

func gerarTransacoesPendentes(rec entities.Recorrencia, hoje, horizonte time.Time) error {
	ultimaStr, err := repository.GetUltimaCompetenciaRecorrenciaRepository(rec.ID)
	if err != nil {
		return err
	}

	var ultima time.Time
	if ultimaStr != nil {
		ultima, err = time.Parse("2006-01-02", (*ultimaStr)[:10])
		if err != nil {
			return err
		}
	} else {
		dataInicio, err := time.Parse("2006-01-02", rec.DataInicio)
		if err != nil {
			return err
		}
		ultima = dataInicio.AddDate(0, 0, -1)
	}

	dataFimRec := horizonte
	if rec.DataFim != nil {
		df, err := time.Parse("2006-01-02", (*rec.DataFim)[:10])
		if err == nil && df.Before(dataFimRec) {
			dataFimRec = df
		}
	}

	proxima := proximaCompetencia(ultima, rec.Frequencia, rec.Intervalo)

	for !proxima.After(dataFimRec) {
		competencia := proxima.Format("2006-01-02")

		transacao := entities.Transacao{
			IDUsuario:      rec.IDUsuario,
			IDGrupo:        rec.IDGrupo,
			Escopo:         rec.Escopo,
			IDCategoria:    rec.IDCategoria,
			IDSubcategoria: rec.IDSubcategoria,
			IDRecorrencia:  &rec.ID,
			Tipo:           rec.Tipo,
			Descricao:      rec.Descricao,
			Observacao:     rec.Observacao,
			Valor:          rec.Valor,
			Competencia:    competencia,
			Status:         "PENDENTE",
		}

		if _, err := repository.CreateTransacaoComRecorrenciaRepository(transacao); err != nil {
			return err
		}

		proxima = proximaCompetencia(proxima, rec.Frequencia, rec.Intervalo)
	}

	return nil
}

func proximaCompetencia(base time.Time, frequencia string, intervalo int) time.Time {
	switch frequencia {
	case "DIARIA":
		return base.AddDate(0, 0, intervalo)
	case "SEMANAL":
		return base.AddDate(0, 0, intervalo*7)
	case "MENSAL":
		return base.AddDate(0, intervalo, 0)
	case "ANUAL":
		return base.AddDate(intervalo, 0, 0)
	default:
		return base.AddDate(0, 1, 0)
	}
}

func esperarProximaMeiaNoite() {
	agora := time.Now()
	proximaMeiaNoite := time.Date(agora.Year(), agora.Month(), agora.Day()+1, 0, 0, 0, 0, agora.Location())
	time.Sleep(time.Until(proximaMeiaNoite))
}
