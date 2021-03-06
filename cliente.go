package syncnfe

import (
	"fmt"
	"time"
)

// NewCliente Retorna um objeto de cliente vazio
func NewCliente() *Cliente {
	return &Cliente{}
}

// Cliente Estrutura que armazena todos os elementos de um Cliente
type Cliente struct {
	ID1                int
	ID2                int
	RazaoSocial        string
	Endereco           string
	Numero             int
	Complemento        string // Opcional
	Bairro             string
	Cidade             string
	Estado             string
	CEP                string
	CNPJ               string
	IE                 string
	CPF                string
	RG                 string
	DiaDeVencimento    string
	Modelo             string
	CFOP               string
	Telefone           string
	Email              string
	CodigoDoConsumidor string
	TipoAssinante      int
	TipoUtilizacao     int
	DataEmissao        string
	DataPrestacao      string
	NumeroNFe          string
	Observacao         string // Opcional
	CodigoDoMunicipio  string // Opcional
}

// String retornar uma string cliente formatada
func (c Cliente) String() string {
	return fmt.Sprintf("%d|%d|%s|%s|%d|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%d|%d|%s|%s|%s|%s|%s|",
		c.ID1, c.ID2, c.RazaoSocial, c.Endereco, c.Numero, c.Complemento, c.Bairro, c.Cidade, c.Estado, c.CEP, c.CNPJ,
		c.IE, c.CPF, c.RG, c.DiaDeVencimento, c.Modelo, c.CFOP, c.Telefone, c.Email, c.CodigoDoConsumidor, c.TipoAssinante,
		c.TipoUtilizacao, c.DataEmissao, c.DataPrestacao, c.NumeroNFe, c.Observacao, c.CodigoDoMunicipio)
}

/* TiposItensAssinantes
1 - Comercial/Industrial
2 - Poder Publico
3 - Residencial/Pessoa Fisica
4 - Publico
5 - Semi-Publico
6 - Outros
*/
func (c *Cliente) TiposItensAssinantes(tiposAssinatura int) {
	c.TipoAssinante = tiposAssinatura
}

/* TiposUtilizacao
1 - Telefonia
2 - Comunicação de dados
3 - TV por Assinatura
4 - Provimento de internet
5 - Multimídia
6 - Outros
*/
func (c *Cliente) TiposUtilizacao(tiposUtilizacao int) {
	c.TipoUtilizacao = tiposUtilizacao
}

// SetDataEmissao Gravar data da emissão
func (c *Cliente) SetDataEmissao(tempo time.Time) {
	c.DataEmissao = fmt.Sprintf("%d/%d/%d", tempo.Day(), tempo.Month(), tempo.Year())
}

// SetDataPrestacao Gravar data da prestação
func (c *Cliente) SetDataPrestacao(tempo time.Time) {
	c.DataPrestacao = fmt.Sprintf("%d/%d/%d", tempo.Day(), tempo.Month(), tempo.Year())
}
