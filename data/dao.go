package data

import (
	"fmt"
	"gorm.io/gorm"
)

type LibraryController struct {
	db *gorm.DB
}

func (c *LibraryController) CriarLivro(livro *Livro)(err error){
	return c.db.Create(livro).Error
}

func (c *LibraryController) ConsultarLivro(consulta ConsultarLivroArgs)(err error,livros []Livro){
	var db = consultaPreloads(consulta,c.db)
	var offset,limit int
	offset = int(consulta.Pagina * consulta.Itens)
	limit = int(consulta.Itens)
	db.Offset(offset).Limit(limit)

	if consulta.Titulo != "" {
		db.Find(&livros,fmt.Sprintf(`titulo like %%%s%%`,consulta.Titulo)).Order("titulo")
	} else {
		db.Find(&livros).Order("titulo")
	}

	return db.Error, livros
}

func consultaPreloads(consulta ConsultarLivroArgs,db *gorm.DB) *gorm.DB{
	var q string
	if consulta.Nome != "" {
		q = fmt.Sprintf(`nome like %%%s%%`,consulta.Nome)
	}
	db.Preload("autores",q)
	q = ""
	if consulta.Ano != 0 || consulta.Edicao != 0 {
		if consulta.Ano != 0{
			q = fmt.Sprintf(`ano = %d`,consulta.Ano)
		}
		if consulta.Edicao != 0{
			if q != "" {
				q += " and "
			}
			q += fmt.Sprintf(`numero = %d`,consulta.Edicao)
		}
	}
	db.Preload("edicoes",q)
	return db
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