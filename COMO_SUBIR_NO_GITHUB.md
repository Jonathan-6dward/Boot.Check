# 🚀 Como Subir o BootCheck no GitHub

Este guia mostra o passo-a-passo para colocar o projeto no GitHub como
repositório **público** com licença **MIT**.

---

## 📋 Pré-requisitos

- [ ] Conta no GitHub
- [ ] Git instalado (`git --version`)
- [ ] Go 1.22+ instalado (para rodar testes)
- [ ] Python 3.8+ com `jsonschema` (para validação de schemas)

```bash
# Verificar instalações
git --version
go version
python --version
pip install jsonschema
```

---

## 🪜 Passo 1: Criar repositório vazio no GitHub

1. Acesse https://github.com/new
2. Preencha:
   - **Repository name:** `Boot.Check`
   - **Description:** `Ferramenta defensiva e somente-leitura de triagem forense para Windows`
   - **Visibilidade:** Public ✅
   - **NÃO** marque "Add a README file", "Add .gitignore" nem "Choose a license"
     (já temos tudo isso)
3. Clique em **Create repository**
4. **Copie a URL** que aparece (ex: `https://github.com/Jonathan-6dward/Boot.Check.git`)

---

## 🪜 Passo 2: Inicializar o repositório local

Abra o PowerShell ou Git Bash na pasta do projeto:

```powershell
# Entrar na pasta do projeto
cd C:\caminho\para\bootcheck

# Inicializar repositório
git init

# Configurar branch principal
git branch -M main

# Adicionar remote (substitua pela sua URL)
git remote add origin https://github.com/Jonathan-6dward/Boot.Check.git
```

---

## 🪜 Passo 3: Verificar o que vai entrar

```bash
# Ver status
git status

# Ver o que será commitado (deve ser ~55 arquivos, SEM .exe, SEM chaves)
git status --short
```

**Esperado ver:**
```
M  .github/ISSUE_TEMPLATE/bug_report.md
M  .github/ISSUE_TEMPLATE/feature_request.md
M  .github/PULL_REQUEST_TEMPLATE.md
M  .github/SECURITY.md
M  .github/workflows/ci.yml
M  .gitignore
M  CHANGELOG.md
M  CONTRIBUTING.md
M  LICENSE
M  README.md
... (etc)
```

**NÃO deve aparecer:**
- ❌ Nenhum `.exe` ou binário
- ❌ Nenhum arquivo com chave, token ou segredo
- ❌ Nenhum dado pessoal
- ❌ Nenhum `node_modules/`, `__pycache__/`, etc.

---

## 🪜 Passo 4: Fazer o commit inicial

```bash
# Adicionar tudo
git add .

# Verificar uma última vez
git status

# Commit inicial
git commit -m "feat: scaffold inicial v0.1.0

- Especificação completa do produto (MVP, persona, jornada)
- Arquitetura técnica com decisão Go, schema e prompt 5W2H/MITRE
- Rascunho de termo de uso, disclaimer e questões LGPD
- Comparação de três modelos comerciais
- Handoff para agente de implementação local
- Scaffold Go do coletor com stubs para 8 categorias
- JSON Schema 2020-12 canônico do pacote de evidências
- Contratos de redaction, consentimento, integridade SHA-256
- Cliente HTTP-placeholder de API LLM
- Prompt de sistema/usuário com boundary de dados
- Gerador HTML de relatório com seção leiga e apêndice técnico
- Três fixtures sintéticas (benign, malicious, ambiguous) e testes
- Script de validação de schemas
- Decisões aprovadas (BootCheck, MIT, mock LLM, local-only)
- CI workflow, CONTRIBUTING, SECURITY, templates de issue/PR"
```

---

## 🪜 Passo 5: Push para o GitHub

```bash
# Push inicial
git push -u origin main
```

**Autenticação:** o GitHub não aceita mais senha. Use:
- **Personal Access Token (PAT)** — GitHub > Settings > Developer settings > Personal access tokens
- **SSH key** — `git@github.com:Jonathan-6dward/Boot.Check.git`
- **GitHub CLI** — `gh auth login`

Se pedir credenciais:
```bash
# Configurar PAT
git config --global credential.helper store
# (git vai pedir usuário e PAT uma vez)
```

---

## 🪜 Passo 6: Configurar o repositório no GitHub

Depois do push, acesse https://github.com/Jonathan-6dward/Boot.Check e:

### 6.1 Adicionar descrição e topics
1. Clique no ⚙️ ao lado de "About"
2. **Description:** `Ferramenta defensiva e somente-leitura de triagem forense para Windows`
3. **Website:** _(deixe vazio por enquanto)_
4. **Topics:** `windows`, `security`, `forensics`, `defensive`, `triage`, `read-only`, `golang`, `lgpd`
5. Salvar

### 6.2 Ativar Issues e Discussions
- Settings > General > Features
  - ✅ Issues
  - ✅ Discussions (recomendado para roadmap)

### 6.3 Configurar branch protection
- Settings > Branches > Add rule
- **Branch name pattern:** `main`
- ✅ Require a pull request before merging
- ✅ Require approvals: 1
- ✅ Dismiss stale pull request approvals
- ✅ Require status checks to pass before merging
  - Selecione: `validate-schemas`, `go-test`, `gofmt`
- ✅ Require linear history
- Salvar

### 6.4 Criar a primeira Release
1. Vá em **Releases** > **Create a new release**
2. **Choose a tag:** `v0.1.0-scaffold` (criar nova)
3. **Release title:** `v0.1.0-scaffold — Pacote inicial de especificação e scaffold`
4. **Description:** (cole do CHANGELOG.md)
5. ✅ Set as the latest release
6. Publicar

### 6.5 Adicionar badges ao README (opcional)

Adicione no topo do `README.md`:

```markdown
[![CI](https://github.com/Jonathan-6dward/Boot.Check/actions/workflows/ci.yml/badge.svg)](https://github.com/Jonathan-6dward/Boot.Check/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/Jonathan-6dward/Boot.Check)](https://goreportcard.com/report/github.com/Jonathan-6dward/Boot.Check)
```

---

## 🪜 Passo 7: Verificar o CI

1. Vá em **Actions** no GitHub
2. Veja o workflow "CI" rodando
3. Se algum job falhar, ajuste o código e faça novo push

---

## 🪜 Passo 8: Divulgar (opcional)

- Tweet sobre o projeto com hashtags: `#golang` `#security` `#windows` `#opensource`
- Poste em https://www.reddit.com/r/golang/
- Submeta para https://github.com/avelino/awesome-go
- Anuncie em https://news.ycombinator.com/ (opcional, com moderação)

---

## ⚠️ Antes do release público

Releases públicos **exigem revisão jurídica**. O arquivo
`docs/RESPONSABILIDADE_LEGAL.md` está marcado como **rascunho** e precisa
de advogado antes de:

- Vender ou distribuir comercialmente
- Coletar dados de terceiros sem consentimento
- Transmitir pacotes a provedores LLM externos
- Anunciar a ferramenta como "antivírus" ou "EDR"

Para a v0.1.0-scaffold, o **mock provider** é usado e nenhum dado sai da
máquina, então o release público do código é seguro.

---

## 🆘 Problemas comuns

### "Repository not found"
- Verifique se a URL está correta
- Confirme que o repositório foi criado no GitHub
- Se for privado, confirme que seu token tem escopo `repo`

### "Permission denied (publickey)"
- Configure uma chave SSH: https://docs.github.com/pt/authentication/connecting-to-github-with-ssh

### CI falha em `go test`
- O scaffold retorna `TODO` deliberado, então os testes devem passar com `go test`
- Se falhar, veja os logs em Actions

### CI falha em `validate-schemas`
- Rode localmente: `python validate_schemas.py`
- Deve mostrar 4 linhas "OK"

### Esqueci de criar `.gitignore` antes do primeiro commit
```bash
# Remover do índice mas manter no disco
git rm --cached -r .
git add .
git commit -m "chore: aplicar .gitignore"
```

---

## ✅ Checklist final

- [ ] Repositório criado no GitHub (público)
- [ ] `git init` + `git remote add` configurados
- [ ] `git add .` + `git commit` feito
- [ ] `git push -u origin main` bem-sucedido
- [ ] CI workflow rodando (Actions > CI)
- [ ] Branch protection ativada em `main`
- [ ] Release v0.1.0-scaffold publicada
- [ ] Descrição e topics adicionados
- [ ] README com badges (opcional)

🎉 **Pronto! Seu projeto BootCheck está no GitHub!**
