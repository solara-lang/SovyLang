# SovyLang – Linguagem de Programação Solara

**SovyLang** é a implementação oficial da linguagem de programação **Solara**, desenvolvida em **GoLang** para oferecer alto desempenho e portabilidade. A linguagem **Solara** possui sintaxe totalmente em português brasileiro, projetada para ser intuitiva, legível e acessível, especialmente para iniciantes, mas sem abrir mão de recursos avançados.**

## Características Principais

- **Sintaxe 100% em português**: Todas as palavras-chave e estruturas utilizam termos em português
- **Tipagem dinâmica**: Suporte nativo para números, texto, booleanos e estruturas de dados
- **Estruturas de controle intuitivas**: Condicionais e loops com sintaxe clara
- **Funções personalizadas**: Definição e chamada de funções próprias
- **Formatador integrado**: Sistema de formatação automática de código
- **CLI profissional**: Interface de linha de comando completa

## Instalação

### Pré-requisitos
- Go 1.21 ou superior
- Windows, macOS ou Linux

### Compilação
```bash
git clone <repository-url>
cd SovyLang
go build -o sovy cmd/sovy/main.go
```

### Configuração do PATH (Opcional)
Para usar o comando `sovy` globalmente:

**Windows:**
1. Copie `sovy.exe` para uma pasta de sua escolha (ex: `C:\Tools\SovyLang`)
2. Adicione essa pasta ao PATH do sistema
3. Reinicie o terminal

**Linux/macOS:**
```bash
sudo cp sovy /usr/local/bin/
```

## Uso

### Execução de arquivos
```bash
# Executar um arquivo .sl
sovy programa.sl

# Ver ajuda
sovy --help

# Ver versão
sovy --version

# Formatar código
sovy programa.sl --format
```

## Sintaxe da Linguagem Solara

### Variáveis e Tipos de Dados

```solara
:: Declaração de variáveis
numero idade = 25
texto nome = "João Silva"
booleano ativo = verdadeiro
```

**Tipos suportados:**
- `numero`: Números inteiros e decimais
- `texto`: Strings de texto
- `booleano`: Valores verdadeiro/falso

### Operações Matemáticas

```solara
numero a = 10
numero b = 5

numero soma = a + b        :: Resultado: 15
numero subtracao = a - b   :: Resultado: 5
numero multiplicacao = a * b :: Resultado: 50
numero divisao = a / b     :: Resultado: 2.0
numero resto = a % b       :: Resultado: 0
```

### Operadores de Comparação

```solara
numero x = 10
numero y = 20

:: Operadores disponíveis: ==, !=, >, <, >=, <=
se x < y
    imprimir "X é menor que Y"
fim
```

### Operadores Lógicos

```solara
numero idade = 25
booleano possui_carteira = verdadeiro

:: Operadores: e, ou, não
se idade >= 18 e possui_carteira
    imprimir "Pode dirigir"
fim

se idade < 18 ou não possui_carteira
    imprimir "Não pode dirigir"
fim
```

### Estruturas Condicionais

```solara
numero nota = 85

:: Estrutura condicional básica
se nota >= 90
    imprimir "Excelente"
fim

se nota >= 70 e nota < 90
    imprimir "Bom"
fim

se nota < 70
    imprimir "Precisa melhorar"
fim
```

### Estruturas de Repetição

```solara
:: Loop com contador
para numero i = 1 ate 10
    imprimir "Contagem: "
    imprimir i
fim

:: Loop com condições
numero contador = 0
para numero j = 0 ate 100
    se j % 2 == 0
        contador = contador + 1
    fim
fim
```

### Funções

```solara
:: Definição de função sem parâmetros
funcao cumprimentar()
    imprimir "Olá, mundo!"
fim

:: Definição de função com parâmetros
funcao saudar(nome)
    imprimir "Olá, "
    imprimir nome
    imprimir "!"
fim

:: Chamadas de função
cumprimentar()
saudar("Maria")
```

### Funções com Retorno

```solara
funcao calcular_area(largura, altura)
    numero area = largura * altura
    retorne area
fim

numero resultado = calcular_area(5, 3)
imprimir "Área calculada: "
imprimir resultado
```

### Comentários

```solara
:: Este é um comentário de linha única
:: Comentários começam com dois dois pontos

numero valor = 42  :: Comentário no final da linha
```

## Exemplos Práticos

### Sistema de Classificação de Notas

```solara
:: Sistema de avaliação acadêmica
numero nota_matematica = 87
numero nota_portugues = 92
numero nota_ciencias = 78

numero media = (nota_matematica + nota_portugues + nota_ciencias) / 3

imprimir "=== BOLETIM ESCOLAR ==="
imprimir "Matemática: "
imprimir nota_matematica
imprimir "Português: "
imprimir nota_portugues
imprimir "Ciências: "
imprimir nota_ciencias
imprimir "Média: "
imprimir media

se media >= 90
    imprimir "Classificação: EXCELENTE"
fim

se media >= 80 e media < 90
    imprimir "Classificação: MUITO BOM"
fim

se media >= 70 e media < 80
    imprimir "Classificação: BOM"
fim

se media >= 60 e media < 70
    imprimir "Classificação: REGULAR"
fim

se media < 60
    imprimir "Classificação: INSUFICIENTE"
fim
```

### Calculadora de IMC

```solara
:: Calculadora de Índice de Massa Corporal
numero peso = 70
numero altura = 1.75

numero imc = peso / (altura * altura)

imprimir "=== CALCULADORA DE IMC ==="
imprimir "Peso: "
imprimir peso
imprimir "Altura: "
imprimir altura
imprimir "IMC: "
imprimir imc

se imc < 18.5
    imprimir "Classificação: Abaixo do peso"
fim

se imc >= 18.5 e imc < 25
    imprimir "Classificação: Peso normal"
fim

se imc >= 25 e imc < 30
    imprimir "Classificação: Sobrepeso"
fim

se imc >= 30
    imprimir "Classificação: Obesidade"
fim
```

### Sistema de Controle de Estoque

```solara
:: Sistema simples de estoque
funcao verificar_estoque(produto, quantidade)
    imprimir "=== CONTROLE DE ESTOQUE ==="
    imprimir "Produto: "
    imprimir produto
    imprimir "Quantidade: "
    imprimir quantidade
    
    se quantidade > 50
        imprimir "Status: ESTOQUE ALTO"
    fim
    
    se quantidade >= 10 e quantidade <= 50
        imprimir "Status: ESTOQUE NORMAL"
    fim
    
    se quantidade < 10 e quantidade > 0
        imprimir "Status: ESTOQUE BAIXO - REABASTECER"
    fim
    
    se quantidade == 0
        imprimir "Status: SEM ESTOQUE - URGENTE"
    fim
fim

:: Testando o sistema
verificar_estoque("Notebook Dell", 5)
verificar_estoque("Mouse Wireless", 25)
verificar_estoque("Teclado Mecânico", 0)
```

## Funções Built-in

A linguagem Solara inclui funções pré-definidas:

- `imprimir(valor)`: Exibe um valor na tela
- `tamanho(lista)`: Retorna o tamanho de uma lista ou string
- `primeiro(lista)`: Retorna o primeiro elemento de uma lista
- `ultimo(lista)`: Retorna o último elemento de uma lista

## Estrutura de Arquivos

```
SovyLang/
├── cmd/sovy/           # Aplicação principal (CLI)
├── internal/
│   ├── lexer/          # Análise léxica
│   ├── parser/         # Análise sintática
│   ├── ast/            # Árvore de sintaxe abstrata
│   ├── evaluator/      # Motor de interpretação
│   ├── object/         # Sistema de objetos
│   ├── token/          # Definições de tokens
│   └── formatter/      # Formatador de código
├── examples/           # Exemplos de código
└── README.md
```

## Desenvolvimento

### Executando Testes
```bash
go test ./...
```

### Compilando para Diferentes Plataformas
```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o sovy.exe cmd/sovy/main.go

# Linux
GOOS=linux GOARCH=amd64 go build -o sovy cmd/sovy/main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o sovy cmd/sovy/main.go
```

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## Contribuição

Contribuições são bem-vindas! Por favor, leia as diretrizes de contribuição antes de enviar pull requests.

## Suporte

Para relatar bugs ou solicitar recursos, abra uma issue no repositório do projeto.

---

**SovyLang** - Programação em português, simples e poderosa.
