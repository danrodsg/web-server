# Servidor Web Simples (Wiki) em Go

Esta é uma aplicação de Servidor Web simples, desenvolvida em Go (Golang), que funciona como um sistema de Wiki ou bloco de notas básico. Ele permite que usuários visualizem, editem e salvem páginas de texto diretamente no sistema de arquivos.


## 🚀 Funcionalidades Principais
Páginas (Pages): Estrutura de dados para armazenar o título (Title) e o corpo (Body) de uma página.

Persistência: O conteúdo das páginas é salvo e carregado de arquivos .txt no diretório do servidor.

Rotas de Servidor:

/view/: Visualiza o conteúdo de uma página.

/edit/: Exibe um formulário para edição.

/save/: Processa o formulário e salva a página.

Redirecionamento: Após salvar uma página, o usuário é redirecionado automaticamente para a rota /view/ correspondente.

## ⚙️ Tecnologias Utilizadas
- Go (Golang)	/ Linguagem de programação principal. /	Alta performance e concorrência para a criação do servidor web.

- net/http	/ Pacote padrão do Go.	/ Criação do servidor, manipulação de rotas (Handlers) e redirecionamento.

- os	/ Pacote padrão do Go.	/ Leitura e escrita de arquivos (.txt) para persistência das páginas.

  

## ▶️ Como Executar o Projeto
Pré-requisito: Certifique-se de ter o Go (versão 1.18+) instalado.
```

Clone o Repositório:

Bash

git clone github.com/danrodsg/web-server.git
cd web-server
Execute:

Bash

go run main.go
O servidor iniciará na porta 8080.

Acesso: Acesse http://localhost:8080/view/TestPage para visualizar uma das páginas de exemplo.
