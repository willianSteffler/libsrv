package grpcctrl

import (
	"context"
	"github.com/willianSteffler/libsrv/data"
	"gorm.io/gorm"
)

type GrpcLibServer struct {
	data.UnimplementedLibraryServer
	dao data.LibraryController
}

func(c*GrpcLibServer) ConsultarLivros(a *data.ConsultarLivroArgs, s data.Library_ConsultarLivrosServer) error{
	var err error
	var l []data.Livro
	if err,l = c.dao.ConsultarLivro(*a); err == nil {
		for i := 0; i < len(l) && err == nil; i++ {
			err = s.Send(&l[i])
		}
	}

	return err
}
func(c*GrpcLibServer) CriarLivro(ctx context.Context,l *data.Livro) (*data.SingleResp, error){

	var res = &data.SingleResp{}
	var err error
	if err = c.dao.CriarLivro(l); err == nil{
		res.Resp = &data.SingleResp_Livro{Livro: l}
	} else {
		res.Resp = &data.SingleResp_Err{Err: &data.Erro{Erro: err.Error()}}
	}

	return res, nil
}
func(c*GrpcLibServer) DeletarLivro( ctx context.Context,l *data.Livro) (*data.SingleResp, error){
	var res = &data.SingleResp{}
	var err error
	if err = c.dao.RemoverLivro(l); err == nil{
		res.Resp = &data.SingleResp_Livro{Livro: l}
	} else {
		res.Resp = &data.SingleResp_Err{Err: &data.Erro{Erro: err.Error()}}
	}

	return res, nil
}
func(c*GrpcLibServer) Alterar(ctx context.Context,l *data.Livro) (*data.SingleResp, error){
	var res = &data.SingleResp{}
	var err error
	if err = c.dao.AlterarLivro(l); err == nil{
		res.Resp = &data.SingleResp_Livro{Livro: l}
	} else {
		res.Resp = &data.SingleResp_Err{Err: &data.Erro{Erro: err.Error()}}
	}

	return res, nil
}

func(c*GrpcLibServer) Init(db *gorm.DB){
	c.dao.Init(db)
}