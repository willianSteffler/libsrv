package data

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type LibraryController struct {
	db *gorm.DB
}

func (c *LibraryController) CriarLivro(livro *Livro)(err error){

	if livro.Titulo == "" {
		return fmt.Errorf("deve ser informado o titulo do livro")
	}

	if len(livro.Autores) == 0 {
		return fmt.Errorf("o livro deve conter ao menos um autor")
	} else {
		for i, e := range livro.Autores {
			if e.Nome == ""  {
				return fmt.Errorf("deve ser informado o nome do autor [%d]",i)
			}
		}
	}

	if len(livro.Edicoes) == 0 {
		return fmt.Errorf("o livro deve conter ao menos uma edição")
	} else {
		for i, e := range livro.Edicoes {
			if e.Numero == 0 || e.Ano == 0 {
				return fmt.Errorf("deve ser informado ano e numero da edição [%d]",i)
			}
		}
	}

	return c.db.Create(&livro).Error
}

func (c *LibraryController) ConsultarLivro(consulta ConsultarLivroArgs)(err error,livros []Livro){
	var db = c.db
	var q string
	var offset,limit int
	offset = int(consulta.Pagina * consulta.Itens)
	limit = int(consulta.Itens)

	if consulta.Nome != "" {
		nomes := strings.Split(consulta.Nome,",")
		q += "( "
		for _, nome := range nomes {
			q += fmt.Sprintf(`autores.nome like '%s' or `,nome)
		}

		q = strings.TrimRight(q," or ")
		q += ") and "
	}

	if consulta.Ano != 0 || consulta.Edicao != 0 {
		if consulta.Ano != 0{
			q += fmt.Sprintf(`e.ano = %d and `,consulta.Ano)
		}
		if consulta.Edicao != 0{
			q += fmt.Sprintf(`e.numero = %d and `,consulta.Edicao)
		}
	}

	if consulta.Titulo != "" {
		q += fmt.Sprintf(`livros.titulo like '%s'`,consulta.Titulo)
	}

	sql := fmt.Sprintf( "SELECT DISTINCT livros.* "+
		"FROM livros, autores "+
		"JOIN livroautor l on l.codigolivro = livros.codigo and l.codigoautor = autores.codigo "+
		"JOIN edicoes e on e.codigolivro = livros.codigo ")
	if q != "" {
		q = strings.TrimRight(q," and ")
		sql = sql + " WHERE " + q
	}

	sql += fmt.Sprintf( "LIMIT %d " +
		"OFFSET %d ",limit,offset);


	db.Preload("Edicoes").Preload("Autores").Raw(sql).Find(&livros)

	return db.Error, livros
}

func (c *LibraryController) RemoverLivro(livro *Livro)(err error){
	return c.db.Delete(&livro).Error
}

func (c *LibraryController) AlterarLivro(livro *Livro)(err error){
	return c.db.Model(livro).Updates(livro).Error
}

func (c *LibraryController) Init(db *gorm.DB) *LibraryController {
	c.db = db
	return c
}

func (Edicao) TableName() string {
	return "edicoes"
}

func (Autor) TableName() string {
	return "autores"
}

func (Livro) TableName() string {
	return "livros"
}