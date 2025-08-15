# SovyLang – Linguagem de Programação Solara

**SovyLang** é a implementação oficial da linguagem de programação **Solara**, desenvolvida em **Go** para oferecer desempenho, simplicidade e acessibilidade.  
Com sintaxe totalmente em português brasileiro, Solara foi projetada para ser intuitiva e legível, tornando a programação mais próxima da linguagem natural, sem abrir mão de recursos avançados.

---

## Características

- **Sintaxe 100% em português** – todas as palavras-chave e estruturas utilizam termos do português brasileiro.
- **Tipagem dinâmica** – suporte nativo para números, texto, booleanos e estruturas de dados.
- **Estruturas de controle claras** – condicionais e loops de fácil leitura.
- **Funções personalizadas** – definição e utilização de funções próprias.
- **Formatador automático** – formata o código de forma padronizada.
- **Portabilidade** – executa em Windows, macOS e Linux.

---

## Instalação

### Pré-requisitos
- Go 1.21 ou superior
- Windows, macOS ou Linux

### Compilação
```bash
git clone <repository-url>
cd SovyLang
go build -o sovy cmd/sovy/main.go
````

### Execução

```bash
# Executar um arquivo .sl
./sovy caminho/do/arquivo.sl
```

---

## Sintaxe da Linguagem Solara

### Variáveis e Tipos

```solara
numero idade = 25
texto nome = "João Silva"
booleano ativo = verdadeiro
```

Tipos disponíveis:

* `numero` → números inteiros ou decimais
* `texto` → cadeias de caracteres
* `booleano` → verdadeiro ou falso

---

### Estruturas Condicionais

```solara
numero idade = 20

se idade >= 18
    imprimir "Maior de idade"
fim

se idade < 18
    imprimir "Menor de idade"
fim
```

---

### Operadores

**Aritméticos:** `+`, `-`, `*`, `/`, `%`
**Comparação:** `==`, `!=`, `>`, `<`, `>=`, `<=`
**Lógicos:** `e`, `ou`, `não`

---

### Laços de Repetição

```solara
para numero i = 1 ate 5
    imprimir i
fim
```

---

### Funções

```solara
funcao saudacao(nome)
    imprimir "Olá, " 
    imprimir nome
fim

saudacao("Maria")
```

---

## Exemplos Práticos

### Classificação por Idade

```solara
numero idade = 4

se idade >= 18
    imprimir "Maior de idade"
fim

se idade < 18
    imprimir "Menor de idade"
fim
```

### Calculadora de Área

```solara
funcao calcular_area(largura, altura)
    numero area = largura * altura
    retorne area
fim

numero resultado = calcular_area(5, 3)
imprimir "Área: "
imprimir resultado
```

---

## Estrutura do Projeto

```
SovyLang/
├── cmd/sovy/           # Ponto de entrada da aplicação
├── internal/
│   ├── lexer/          # Analisador léxico
│   ├── parser/         # Analisador sintático
│   ├── ast/            # Árvore de sintaxe abstrata
│   ├── evaluator/      # Interpretador
│   ├── object/         # Sistema de tipos
│   ├── token/          # Definição de tokens
│   └── formatter/      # Formatador de código
├── examples/           # Exemplos de uso
└── README.md
```

---

## Desenvolvimento

Executar testes:

```bash
go test ./...
```

Compilar para diferentes plataformas:

```bash
GOOS=windows GOARCH=amd64 go build -o sovy.exe cmd/sovy/main.go
GOOS=linux   GOARCH=amd64 go build -o sovy cmd/sovy/main.go
GOOS=darwin  GOARCH=amd64 go build -o sovy cmd/sovy/main.go
```

---

## Licença

Projeto sob licença MIT – consulte o arquivo `LICENSE` para mais detalhes.

---

**SovyLang** – Programação em português, simples e eficiente.
