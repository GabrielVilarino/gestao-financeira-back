package schemas

type GetDashboardTotalGanhosResponse struct {
	TotalGanhos float64 `json:"total_ganhos"`
}

type GetDashboardTotalDespesasResponse struct {
	TotalDespesas float64 `json:"total_despesas"`
}

type GetDashboardSaldoLiquidoResponse struct {
	SaldoLiquido float64 `json:"saldo_liquido"`
}

type GetDashboardEvolucaoMensalItemResponse struct {
	Mes           string  `json:"mes"`
	TotalGanhos   float64 `json:"total_ganhos"`
	TotalDespesas float64 `json:"total_despesas"`
	SaldoLiquido  float64 `json:"saldo_liquido"`
}
