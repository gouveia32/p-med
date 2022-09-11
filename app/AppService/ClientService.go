package AppService

import (
	"pm/models"
)

var (
	ModelClientLists models.Client
	ModelClientAdd   models.Client
	ModelClientDel   models.Client
	ModelClientOne   models.Client
)

//A combinação torna-se um menu mais espesso
func GetClientLists(offset int, limit int, data_type string) (lists []map[string]interface{}, err error) {
	//ModelMenuLists.ModuleBelong = module_id
	var clients []models.Client
	if data_type == "All" {

		clients, err = ModelClientLists.AllClient(offset, limit)
	} else {
		//Encontre o top primeiro

		clients, err = ModelClientLists.QueryClientLists(0)
		if err != nil {
			return lists, err
		}
	}
	lists = getClientData(clients, data_type)

	return lists, nil
}

//Consultar dados com base nas condições da estrutura
func GetClientByMap(client *models.Client) (data models.Client, err error) {
	ModelClientOne = *client
	return ModelClientOne, ModelClientOne.OneClient()
}

//Client Infinitus para obter um subconjunto
func getClientData(clients []models.Client, data_type string) []map[string]interface{} {
	var ClientsData = make([]map[string]interface{}, len(clients))

	for key, item := range clients {
		row := make(map[string]interface{})
		row["mid"] = item.Id
		row["nome"] = item.Nome
		row["contato_funcao"] = item.ContatoFuncao
		row["contato_nome"] = item.ContatoNome
		row["cgc_cpf"] = item.CgcCpf
		row["razao_social"] = item.RazaoSocial
		row["inscr_estadual"] = item.InscrEstadual
		row["endereco"] = item.Endereco
		row["cidade"] = item.Cidade
		row["uf"] = item.Uf
		row["cep"] = item.Cep
		row["email"] = item.Email
		row["telefone1"] = item.Telefone1
		row["telefone2"] = item.Telefone2
		row["telefone3"] = item.Telefone3
		row["obs"] = item.Obs
		row["estado"] = item.Estado
		row["preco_base"] = item.PrecoBase

		ClientsData[key] = row
	}

	return ClientsData
}

//Adicionar e editar dados de operação
func SaveClient(client *models.Client) (err error) {
	ModelClientAdd = *client
	if ModelClientAdd.Id == 0 {
		err = ModelClientAdd.AddClientData()
	} else {
		err = ModelClientAdd.EditClientData()
	}
	return err
}

//Exclua os dados com base nas condições da estrutura
func DelClientByMap(client *models.Client) (err error) {
	ModelClientDel = *client
	//Adicione uma transação para evitar erros na exclusão
	tx := models.Db.Begin()
	//
	err = ModelClientDel.DelClient()
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}
