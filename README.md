# EventHub 🎉

Sistema completo de gerenciamento de eventos desenvolvido com **Go** (backend) e **React + TypeScript** (frontend).

## 📋 Funcionalidades

### 🔐 Autenticação
- Registro de usuários (Participantes e Organizadores)
- Login com JWT
- Logout com confirmação via modal

### 📅 Gerenciamento de Eventos
- **Organizadores podem:**
  - Criar eventos com limite de participantes
  - Editar seus eventos
  - Visualizar lista de participantes
  - Gerenciar inscrições

- **Participantes podem:**
  - Visualizar todos os eventos
  - Se inscrever em eventos (respeitando limites)
  - Cancelar inscrições
  - Ver histórico de eventos inscritos

### 🎨 Interface
- Design responsivo e moderno
- Componentes reutilizáveis com Tailwind CSS
- Feedback visual em tempo real
- Modais de confirmação
- Sistema de cores por status dos eventos

## 🚀 Tecnologias

### Backend (Go)
- **Framework**: Gin
- **Arquitetura**: Clean Architecture
- **Autenticação**: JWT
- **Banco de Dados**: PostgreSQL
- **ORM**: GORM

### Frontend (React)
- **Framework**: React 18 + TypeScript
- **Build Tool**: Vite
- **Styling**: Tailwind CSS
- **Componentes**: Custom UI Components
- **Roteamento**: React Router
- **Icons**: Lucide React

## 📦 Estrutura do Projeto

```
EventHub/
├── api-go/                 # Backend em Go
│   ├── cmd/               # Ponto de entrada da aplicação
│   ├── internal/          # Código interno da aplicação
│   │   ├── application/   # Casos de uso e DTOs
│   │   ├── domain/        # Modelos e interfaces
│   │   ├── infra/         # Implementações de infraestrutura
│   │   └── server/        # Configuração do servidor
│   └── go.mod
│
├── event-frontend/        # Frontend em React
│   ├── src/
│   │   ├── components/    # Componentes reutilizáveis
│   │   ├── pages/         # Páginas da aplicação
│   │   ├── services/      # Serviços de API
│   │   ├── types/         # Tipos TypeScript
│   │   └── utils/         # Utilitários
│   └── package.json
│
└── README.md
```

## 🛠️ Como Executar

### Pré-requisitos
- **Go** 1.21+
- **Node.js** 18+
- **PostgreSQL** 12+

### 1. Backend (Go)

```bash
# Navegar para o backend
cd api-go

# Instalar dependências
go mod tidy

# Configurar banco de dados (PostgreSQL)
# Criar banco: eventhub

# Executar
go run cmd/main.go
```

O backend estará rodando em `http://localhost:8080`

### 2. Frontend (React)

```bash
# Navegar para o frontend
cd event-frontend

# Instalar dependências
npm install

# Executar em modo desenvolvimento
npm run dev
```

O frontend estará rodando em `http://localhost:5173`

## 🔧 Configuração

### Variáveis de Ambiente (Backend)
Crie um arquivo `.env` no diretório `api-go/`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=eventhub
JWT_SECRET=sua_chave_secreta_jwt
```

### Banco de Dados
O sistema irá criar as tabelas automaticamente na primeira execução.

## 📱 Páginas Principais

- **Home**: Lista de todos os eventos
- **Login/Registro**: Autenticação de usuários
- **Dashboard**: Gerenciamento pessoal de eventos e inscrições
- **Criar Evento**: Formulário para novos eventos (organizadores)
- **Detalhes do Evento**: Informações completas e inscrição
- **Participantes**: Lista de inscritos (apenas organizadores)

## 🤝 Contribuindo

1. Faça o fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 🎯 Próximas Funcionalidades

- [ ] Sistema de notificações
- [ ] Upload de imagens para eventos
- [ ] Sistema de avaliações
- [ ] Integração com calendário
- [ ] Chat em tempo real
- [ ] Relatórios para organizadores

---

Desenvolvido com ❤️ usando Go e React
