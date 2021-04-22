package socket

import (
	"bufio"
	"fmt"
	"github.com/willianSteffler/libsrv/data"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"log"
	"net"
	"runtime/debug"
)

func handleConnection(c net.Conn, dao data.LibraryController) {
	defer func() {
		c.Close()
		if r := recover(); r != nil {
			log.Printf("panic no tratamento da conexão !! %v , %s\n",r,string(debug.Stack()))
		}
	}()
	var in, out []byte
	var msg data.SocketMsg
	var err error
	if _, err = bufio.NewReader(c).Read(in); err != nil {
		log.Printf("erro ao ler do socket %v\n ", err)
		return
	}

	if err = proto.Unmarshal(in, &msg); err != nil {
		log.Printf("mensagem não reconhecida %v\n", err)
		return
	}

	resp := &data.SocketMsg{
		Operacao: msg.GetOperacao() + ".resp",
	}
	switch msg.Operacao {
	case "criarLivro":
		if l := msg.GetLivro(); l != nil {
			if err = dao.CriarLivro(l); err == nil {
				resp.OpArgs = &data.SocketMsg_Livro{Livro: l}
			}
		} else {
			err = fmt.Errorf("argumento livro deve ser fornecido")
		}

	case "consultarLivro":
		if c := msg.GetConsulta(); c != nil {
			var livros []data.Livro
			if err, livros = dao.ConsultarLivro(*c); err == nil {
				var lr []*data.Livro
				for _, livro := range livros {
					lr = append(lr, &livro)
				}
				resp.OpArgs = &data.SocketMsg_ConsultaResp{
					ConsultaResp: &data.ConsultarLivrosResp{Livros: lr},
				}
			}
		} else {
			err = fmt.Errorf("argumento livro deve ser fornecido")
		}
	case "removerLivro":
		if l := msg.GetLivro(); l != nil {
			if err = dao.RemoverLivro(l); err == nil {
				resp.OpArgs = &data.SocketMsg_Livro{Livro: l}
			}
		} else {
			err = fmt.Errorf("argumento livro deve ser fornecido")
		}
	case "alterarLivro":
		if l := msg.GetLivro(); l != nil {
			if err = dao.AlterarLivro(l); err == nil {
				resp.OpArgs = &data.SocketMsg_Livro{Livro: l}
			}
		} else {
			err = fmt.Errorf("argumento livro deve ser fornecido")
		}
	default:
		err = fmt.Errorf("operação %s desconhecida", msg.Operacao)
	}

	if err != nil {
		out, _ = proto.Marshal(&data.SocketMsg{
			Operacao: msg.GetOperacao() + ".resp",
			OpArgs:   &data.SocketMsg_Erro{Erro: &data.Erro{Erro: err.Error()}},
		})
	}

	c.Write(out)
	c.Close()
}

func Listen(port int, db *gorm.DB) error {
	dao := data.LibraryController{}
	dao.Init(db)
	l, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	defer l.Close()

	log.Printf("socket aguarando conexões em %s",l.Addr().String())
	for {
		c, err := l.Accept()
		log.Printf("socket conectado %s",c.RemoteAddr().String())
		if err != nil {
			return err
		}
		go handleConnection(c, dao)
	}
}
