package dashboard

import (
	"strconv"

	"github.com/GabrielVilarino/gestao-financeira-back.git/exceptions"
	"github.com/gin-gonic/gin"
)

type dashboardFilters struct {
	dataInicio *string
	dataFim    *string
	idGrupo    *int
	idGrupoJWT *int
}

func getDashboardFilters(c *gin.Context) (*dashboardFilters, error) {
	filters := &dashboardFilters{}

	if val := c.Query("data_inicio"); val != "" {
		filters.dataInicio = &val
	}

	if val := c.Query("data_fim"); val != "" {
		filters.dataFim = &val
	}

	if val := c.Query("id_grupo"); val != "" {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return nil, &exceptions.ValidationError{Message: "id_grupo inválido"}
		}
		filters.idGrupo = &parsed
	}

	idGrupoJWT := c.GetInt("id_group")
	if idGrupoJWT != 0 {
		filters.idGrupoJWT = &idGrupoJWT
	}

	return filters, nil
}
