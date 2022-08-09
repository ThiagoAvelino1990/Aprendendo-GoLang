package trabalhador

var Nome string
var SalarioBruto float64
var SalarioLiquido float64

//CalcularIR tem como objetivo calcular o descondo do IR do salário
func CalcularIR(SalarioBruto float64) float64 {
	aliquotaUm := 0.075
	aliquotaDois := 0.15
	aliquotaTres := 0.225
	aliquotaQuatro := 0.275

	//Faixa um
	if SalarioBruto >= 1903.99 && SalarioBruto <= 2829.65 {
		return ((SalarioBruto * aliquotaUm) - 142.80)
		//Faixa dois
	} else if SalarioBruto >= 2826.66 && SalarioBruto <= 3751.05 {
		return ((SalarioBruto * aliquotaDois) - 354.80)
		//Faixa três
	} else if SalarioBruto >= 3751.06 && SalarioBruto <= 4664.68 {
		return ((SalarioBruto * aliquotaTres) - 636.13)
		//Faixa quatro
	} else if SalarioBruto >= 4664.68 {
		return ((SalarioBruto * aliquotaQuatro) - 869.36)
	}

	return 0
}

//CalcularINSS tem como objetivo calcular o desconto de INSS do salário
func CalcularINSS(SalarioBruto float64) float64 {

	faixaUm := 0.075
	faixaDois := 0.09
	faixaTres := 0.12
	faixaQuatro := 0.14

	if SalarioBruto <= 1212.00 {
		return (SalarioBruto * faixaUm)
	}
	if SalarioBruto >= 1212.01 && SalarioBruto <= 2427.35 {
		return (faixaUm * 1212.00) + (faixaDois * (SalarioBruto - 1212.00))
	}
	if SalarioBruto >= 2427.36 && SalarioBruto <= 3641.03 {
		return (faixaUm * 1212.00) + (faixaDois * 2427.35) + (faixaTres * (SalarioBruto - 2427.35))
	}
	if SalarioBruto >= 3641.04 && SalarioBruto <= 7087.22 {
		return (faixaUm * 1212.00) + (faixaDois * 2427.35) + (faixaTres * 3641.03) + (faixaQuatro * (SalarioBruto - 3641.03))
	}
	return (SalarioBruto * faixaUm) + (SalarioBruto * faixaDois) + (SalarioBruto * faixaTres) + (SalarioBruto * faixaQuatro)
}

//CalculaSalarioLiquido tem como objetivo calcular o salário líquido
func CalculaSalarioLiquido(SalarioBruto float64) float64 {

	SalarioLiquido = SalarioBruto - (CalcularIR(SalarioBruto) + CalcularINSS(SalarioBruto))

	return SalarioLiquido
}
