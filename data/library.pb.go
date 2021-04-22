// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: library.proto

package data

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Autor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"primaryKey"
	Codigo int32 `protobuf:"varint,1,opt,name=codigo,proto3" json:"codigo,omitempty" gorm:"primaryKey"`
	// @inject_tag: gorm:"many2many:livrosautor;foreignKey:codigoautor;References:autores.codigo"
	Livros []*Livro `protobuf:"bytes,2,rep,name=livros,proto3" json:"livros,omitempty" gorm:"many2many:livrosautor;foreignKey:codigoautor;References:autores.codigo"`
	Nome   string   `protobuf:"bytes,3,opt,name=nome,proto3" json:"nome,omitempty"`
}

func (x *Autor) Reset() {
	*x = Autor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Autor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Autor) ProtoMessage() {}

func (x *Autor) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Autor.ProtoReflect.Descriptor instead.
func (*Autor) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{0}
}

func (x *Autor) GetCodigo() int32 {
	if x != nil {
		return x.Codigo
	}
	return 0
}

func (x *Autor) GetLivros() []*Livro {
	if x != nil {
		return x.Livros
	}
	return nil
}

func (x *Autor) GetNome() string {
	if x != nil {
		return x.Nome
	}
	return ""
}

type Edicao struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codigolivro int32 `protobuf:"varint,2,opt,name=codigolivro,proto3" json:"codigolivro,omitempty"`
	Numero      int32 `protobuf:"varint,3,opt,name=numero,proto3" json:"numero,omitempty"`
	Ano         int32 `protobuf:"varint,4,opt,name=ano,proto3" json:"ano,omitempty"`
}

func (x *Edicao) Reset() {
	*x = Edicao{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Edicao) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Edicao) ProtoMessage() {}

func (x *Edicao) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Edicao.ProtoReflect.Descriptor instead.
func (*Edicao) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{1}
}

func (x *Edicao) GetCodigolivro() int32 {
	if x != nil {
		return x.Codigolivro
	}
	return 0
}

func (x *Edicao) GetNumero() int32 {
	if x != nil {
		return x.Numero
	}
	return 0
}

func (x *Edicao) GetAno() int32 {
	if x != nil {
		return x.Ano
	}
	return 0
}

type Livro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"primaryKey"
	Codigo int32  `protobuf:"varint,1,opt,name=codigo,proto3" json:"codigo,omitempty" gorm:"primaryKey"`
	Titulo string `protobuf:"bytes,2,opt,name=titulo,proto3" json:"titulo,omitempty"`
	// @inject_tag: gorm:"many2many:livrosautor;foreignKey:codigolivro;References:livros.codigo"
	Autores []*Autor `protobuf:"bytes,3,rep,name=autores,proto3" json:"autores,omitempty" gorm:"many2many:livrosautor;foreignKey:codigolivro;References:livros.codigo"`
	// @inject_tag: gorm:"foreignKey:codigolivro;references:livros.codigo"
	Edicoes []*Edicao `protobuf:"bytes,4,rep,name=edicoes,proto3" json:"edicoes,omitempty" gorm:"foreignKey:codigolivro;references:livros.codigo"`
}

func (x *Livro) Reset() {
	*x = Livro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Livro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Livro) ProtoMessage() {}

func (x *Livro) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Livro.ProtoReflect.Descriptor instead.
func (*Livro) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{2}
}

func (x *Livro) GetCodigo() int32 {
	if x != nil {
		return x.Codigo
	}
	return 0
}

func (x *Livro) GetTitulo() string {
	if x != nil {
		return x.Titulo
	}
	return ""
}

func (x *Livro) GetAutores() []*Autor {
	if x != nil {
		return x.Autores
	}
	return nil
}

func (x *Livro) GetEdicoes() []*Edicao {
	if x != nil {
		return x.Edicoes
	}
	return nil
}

type ConsultarLivroArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nome   string `protobuf:"bytes,1,opt,name=nome,proto3" json:"nome,omitempty"`
	Titulo string `protobuf:"bytes,2,opt,name=titulo,proto3" json:"titulo,omitempty"`
	Ano    int32  `protobuf:"varint,3,opt,name=ano,proto3" json:"ano,omitempty"`
	Edicao int32  `protobuf:"varint,4,opt,name=edicao,proto3" json:"edicao,omitempty"`
	Pagina int32  `protobuf:"varint,5,opt,name=pagina,proto3" json:"pagina,omitempty"`
	Itens  int32  `protobuf:"varint,6,opt,name=itens,proto3" json:"itens,omitempty"`
}

func (x *ConsultarLivroArgs) Reset() {
	*x = ConsultarLivroArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultarLivroArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultarLivroArgs) ProtoMessage() {}

func (x *ConsultarLivroArgs) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultarLivroArgs.ProtoReflect.Descriptor instead.
func (*ConsultarLivroArgs) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{3}
}

func (x *ConsultarLivroArgs) GetNome() string {
	if x != nil {
		return x.Nome
	}
	return ""
}

func (x *ConsultarLivroArgs) GetTitulo() string {
	if x != nil {
		return x.Titulo
	}
	return ""
}

func (x *ConsultarLivroArgs) GetAno() int32 {
	if x != nil {
		return x.Ano
	}
	return 0
}

func (x *ConsultarLivroArgs) GetEdicao() int32 {
	if x != nil {
		return x.Edicao
	}
	return 0
}

func (x *ConsultarLivroArgs) GetPagina() int32 {
	if x != nil {
		return x.Pagina
	}
	return 0
}

func (x *ConsultarLivroArgs) GetItens() int32 {
	if x != nil {
		return x.Itens
	}
	return 0
}

type Erro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Erro string `protobuf:"bytes,1,opt,name=erro,proto3" json:"erro,omitempty"`
}

func (x *Erro) Reset() {
	*x = Erro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Erro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Erro) ProtoMessage() {}

func (x *Erro) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Erro.ProtoReflect.Descriptor instead.
func (*Erro) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{4}
}

func (x *Erro) GetErro() string {
	if x != nil {
		return x.Erro
	}
	return ""
}

type ConsultarLivrosResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Livros []*Livro `protobuf:"bytes,1,rep,name=livros,proto3" json:"livros,omitempty"`
}

func (x *ConsultarLivrosResp) Reset() {
	*x = ConsultarLivrosResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultarLivrosResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultarLivrosResp) ProtoMessage() {}

func (x *ConsultarLivrosResp) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultarLivrosResp.ProtoReflect.Descriptor instead.
func (*ConsultarLivrosResp) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{5}
}

func (x *ConsultarLivrosResp) GetLivros() []*Livro {
	if x != nil {
		return x.Livros
	}
	return nil
}

type SocketMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operacao string `protobuf:"bytes,1,opt,name=operacao,proto3" json:"operacao,omitempty"`
	// Types that are assignable to OpArgs:
	//	*SocketMsg_Livro
	//	*SocketMsg_Consulta
	//	*SocketMsg_Erro
	//	*SocketMsg_ConsultaResp
	OpArgs isSocketMsg_OpArgs `protobuf_oneof:"opArgs"`
}

func (x *SocketMsg) Reset() {
	*x = SocketMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketMsg) ProtoMessage() {}

func (x *SocketMsg) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketMsg.ProtoReflect.Descriptor instead.
func (*SocketMsg) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{6}
}

func (x *SocketMsg) GetOperacao() string {
	if x != nil {
		return x.Operacao
	}
	return ""
}

func (m *SocketMsg) GetOpArgs() isSocketMsg_OpArgs {
	if m != nil {
		return m.OpArgs
	}
	return nil
}

func (x *SocketMsg) GetLivro() *Livro {
	if x, ok := x.GetOpArgs().(*SocketMsg_Livro); ok {
		return x.Livro
	}
	return nil
}

func (x *SocketMsg) GetConsulta() *ConsultarLivroArgs {
	if x, ok := x.GetOpArgs().(*SocketMsg_Consulta); ok {
		return x.Consulta
	}
	return nil
}

func (x *SocketMsg) GetErro() *Erro {
	if x, ok := x.GetOpArgs().(*SocketMsg_Erro); ok {
		return x.Erro
	}
	return nil
}

func (x *SocketMsg) GetConsultaResp() *ConsultarLivrosResp {
	if x, ok := x.GetOpArgs().(*SocketMsg_ConsultaResp); ok {
		return x.ConsultaResp
	}
	return nil
}

type isSocketMsg_OpArgs interface {
	isSocketMsg_OpArgs()
}

type SocketMsg_Livro struct {
	Livro *Livro `protobuf:"bytes,2,opt,name=livro,proto3,oneof"`
}

type SocketMsg_Consulta struct {
	Consulta *ConsultarLivroArgs `protobuf:"bytes,3,opt,name=consulta,proto3,oneof"`
}

type SocketMsg_Erro struct {
	Erro *Erro `protobuf:"bytes,4,opt,name=erro,proto3,oneof"`
}

type SocketMsg_ConsultaResp struct {
	ConsultaResp *ConsultarLivrosResp `protobuf:"bytes,5,opt,name=consulta_resp,json=consultaResp,proto3,oneof"`
}

func (*SocketMsg_Livro) isSocketMsg_OpArgs() {}

func (*SocketMsg_Consulta) isSocketMsg_OpArgs() {}

func (*SocketMsg_Erro) isSocketMsg_OpArgs() {}

func (*SocketMsg_ConsultaResp) isSocketMsg_OpArgs() {}

var File_library_proto protoreflect.FileDescriptor

var file_library_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6c, 0x69, 0x62, 0x73, 0x72, 0x76, 0x22, 0x5a, 0x0a, 0x05, 0x41, 0x75, 0x74, 0x6f, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12, 0x25, 0x0a, 0x06, 0x6c, 0x69, 0x76, 0x72,
	0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x69, 0x62, 0x73, 0x72,
	0x76, 0x2e, 0x4c, 0x69, 0x76, 0x72, 0x6f, 0x52, 0x06, 0x6c, 0x69, 0x76, 0x72, 0x6f, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x6f, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x06, 0x45, 0x64, 0x69, 0x63, 0x61, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x6c, 0x69, 0x76, 0x72, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x6c, 0x69, 0x76, 0x72, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6e, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x6e, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x05, 0x4c, 0x69,
	0x76, 0x72, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x69, 0x74, 0x75, 0x6c, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x74,
	0x75, 0x6c, 0x6f, 0x12, 0x27, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x69, 0x62, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x07,
	0x65, 0x64, 0x69, 0x63, 0x6f, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x6c, 0x69, 0x62, 0x73, 0x72, 0x76, 0x2e, 0x45, 0x64, 0x69, 0x63, 0x61, 0x6f, 0x52, 0x07, 0x65,
	0x64, 0x69, 0x63, 0x6f, 0x65, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x76, 0x72, 0x6f, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x74, 0x75, 0x6c, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x69, 0x74, 0x75, 0x6c, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6e, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x64, 0x69, 0x63, 0x61, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6e,
	0x73, 0x22, 0x1a, 0x0a, 0x04, 0x45, 0x72, 0x72, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x72, 0x72,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x72, 0x72, 0x6f, 0x22, 0x3c, 0x0a,
	0x13, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x76, 0x72, 0x6f, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x06, 0x6c, 0x69, 0x76, 0x72, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x69, 0x62, 0x73, 0x72, 0x76, 0x2e, 0x4c, 0x69,
	0x76, 0x72, 0x6f, 0x52, 0x06, 0x6c, 0x69, 0x76, 0x72, 0x6f, 0x73, 0x22, 0xfa, 0x01, 0x0a, 0x09,
	0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x63, 0x61, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x63, 0x61, 0x6f, 0x12, 0x25, 0x0a, 0x05, 0x6c, 0x69, 0x76, 0x72, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x69, 0x62, 0x73, 0x72, 0x76, 0x2e, 0x4c, 0x69,
	0x76, 0x72, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x76, 0x72, 0x6f, 0x12, 0x38, 0x0a, 0x08,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6c, 0x69, 0x62, 0x73, 0x72, 0x76, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61,
	0x72, 0x4c, 0x69, 0x76, 0x72, 0x6f, 0x41, 0x72, 0x67, 0x73, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x04, 0x65, 0x72, 0x72, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x69, 0x62, 0x73, 0x72, 0x76, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x65, 0x72, 0x72, 0x6f, 0x12, 0x42, 0x0a, 0x0d, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x73, 0x72, 0x76, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x76, 0x72, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x48, 0x00,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x42, 0x08,
	0x0a, 0x06, 0x6f, 0x70, 0x41, 0x72, 0x67, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_library_proto_rawDescOnce sync.Once
	file_library_proto_rawDescData = file_library_proto_rawDesc
)

func file_library_proto_rawDescGZIP() []byte {
	file_library_proto_rawDescOnce.Do(func() {
		file_library_proto_rawDescData = protoimpl.X.CompressGZIP(file_library_proto_rawDescData)
	})
	return file_library_proto_rawDescData
}

var file_library_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_library_proto_goTypes = []interface{}{
	(*Autor)(nil),               // 0: libsrv.Autor
	(*Edicao)(nil),              // 1: libsrv.Edicao
	(*Livro)(nil),               // 2: libsrv.Livro
	(*ConsultarLivroArgs)(nil),  // 3: libsrv.ConsultarLivroArgs
	(*Erro)(nil),                // 4: libsrv.Erro
	(*ConsultarLivrosResp)(nil), // 5: libsrv.ConsultarLivrosResp
	(*SocketMsg)(nil),           // 6: libsrv.SocketMsg
}
var file_library_proto_depIdxs = []int32{
	2, // 0: libsrv.Autor.livros:type_name -> libsrv.Livro
	0, // 1: libsrv.Livro.autores:type_name -> libsrv.Autor
	1, // 2: libsrv.Livro.edicoes:type_name -> libsrv.Edicao
	2, // 3: libsrv.ConsultarLivrosResp.livros:type_name -> libsrv.Livro
	2, // 4: libsrv.SocketMsg.livro:type_name -> libsrv.Livro
	3, // 5: libsrv.SocketMsg.consulta:type_name -> libsrv.ConsultarLivroArgs
	4, // 6: libsrv.SocketMsg.erro:type_name -> libsrv.Erro
	5, // 7: libsrv.SocketMsg.consulta_resp:type_name -> libsrv.ConsultarLivrosResp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_library_proto_init() }
func file_library_proto_init() {
	if File_library_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_library_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Autor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_library_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Edicao); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_library_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Livro); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_library_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultarLivroArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_library_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Erro); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_library_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultarLivrosResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_library_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_library_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*SocketMsg_Livro)(nil),
		(*SocketMsg_Consulta)(nil),
		(*SocketMsg_Erro)(nil),
		(*SocketMsg_ConsultaResp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_library_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_library_proto_goTypes,
		DependencyIndexes: file_library_proto_depIdxs,
		MessageInfos:      file_library_proto_msgTypes,
	}.Build()
	File_library_proto = out.File
	file_library_proto_rawDesc = nil
	file_library_proto_goTypes = nil
	file_library_proto_depIdxs = nil
}