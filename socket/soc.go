package socket

import (
	"fmt"
	"github.com/willianSteffler/libsrv/data"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"io"
	"log"
	"net"
	"runtime/debug"
	"time"
)

const  (
	CRIAR_LIVRO = "criar"
	REMOVER_LIVRO = "remover"
	ALTERAR_LIVRO = "alterar"
	BUSCAR_LIVRO = "buscar"
)

func handleConnection(c net.Conn, dao data.LibraryController) {
	defer func() {
		c.Close()
		log.Printf("fechando conexão socket em %v\n",c.RemoteAddr().String())
		if r := recover(); r != nil {
			log.Printf("panic no tratamento da conexão !! %v , %s\n",r,string(debug.Stack()))
		}
	}()

	for   {
		var err error
		var msg data.SocketMsg
		if err,msg = GetSocketMessage(c); err != nil {
			log.Println(err)
			return
		}

		resp := data.SocketMsg{
			Operacao: msg.GetOperacao() + ".resp",
		}
		switch msg.Operacao {
		case CRIAR_LIVRO:
			if l := msg.GetLivro(); l != nil {
				if err = dao.CriarLivro(l); err == nil {
					resp.OpArgs = &data.SocketMsg_Livro{Livro: l}
				}
			} else {
				err = fmt.Errorf("argumento livro deve ser fornecido")
			}

		case BUSCAR_LIVRO:
			if c := msg.GetConsulta(); c != nil {
				var livros []data.Livro
				if err, livros = dao.ConsultarLivro(*c); err == nil {
					var lr []*data.Livro
					for i, _ := range livros {
						lr = append(lr, &livros[i])
					}
					resp.OpArgs = &data.SocketMsg_ConsultaResp{
						ConsultaResp: &data.ConsultarLivrosResp{Livros: lr},
					}
				}
			} else {
				err = fmt.Errorf("argumento livro deve ser fornecido")
			}
		case REMOVER_LIVRO:
			if l := msg.GetLivro(); l != nil {
				if err = dao.RemoverLivro(l); err == nil {
					resp.OpArgs = &data.SocketMsg_Livro{Livro: l}
				}
			} else {
				err = fmt.Errorf("argumento livro deve ser fornecido")
			}
		case ALTERAR_LIVRO:
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
			resp.OpArgs =   &data.SocketMsg_Erro{Erro: &data.Erro{Erro: err.Error()}}
		}

		if err = SendSocketMessage(c,resp); err != nil {
			log.Printf("erro ao escrever no socket %v\n",err)
			return
		}
	}
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

func GetSocketMessage(c net.Conn) (error ,data.SocketMsg){
	var msg data.SocketMsg
	var err error
	var n int
	var buf  []byte
	tmp := make([]byte, 1024)
	for {
		n, err = c.Read(tmp)
		if err != nil {
			if err == io.EOF  {
				err = nil
			}
			break
		}
		buf = append(buf, tmp[:n]...)
		if n == 1024{
			// ler buffer grande
			c.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
		} else {
			break
		}
	}

	if err != nil {
		err = fmt.Errorf("erro ao ler do socket %v\n ", err)
	} else if err = proto.Unmarshal(buf, &msg); err != nil {
		err = fmt.Errorf("mensagem não reconhecida %v\n", err)
	}
	return err,msg
}

func SendSocketMessage(c net.Conn,msg data.SocketMsg)(err error){
	var out []byte
	out, err = proto.Marshal(&msg)
	if err != nil {
		return err
	}
	_, err = c.Write(out)
	return err
}
