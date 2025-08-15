package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"sovylang/internal/evaluator"
	"sovylang/internal/formatter"
	"sovylang/internal/lexer"
	"sovylang/internal/library"
	"sovylang/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: sovy <comando> [argumentos]")
		fmt.Println("Comandos:")
		fmt.Println("  <arquivo.sl>        Executar arquivo")
		fmt.Println("  install <biblioteca> Instalar biblioteca")
		fmt.Println("  list                Listar bibliotecas instaladas")
		fmt.Println("  --format <arquivo>  Formatar arquivo")
		fmt.Println("  --help              Mostrar esta ajuda")
		fmt.Println("  --version           Mostrar versão")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "--help":
		showHelp()
	case "--version":
		fmt.Println("Sovy - Interpretador da linguagem Solara v2.0.0")
	case "install":
		if len(os.Args) < 3 {
			fmt.Println("Uso: sovy install <biblioteca>")
			fmt.Println("Bibliotecas disponíveis:")
			fmt.Println("  smath - Biblioteca de matemática avançada")
			os.Exit(1)
		}
		installLibrary(os.Args[2])
	case "list":
		listLibraries()
	case "--format":
		if len(os.Args) < 3 {
			fmt.Println("Uso: sovy --format <arquivo.sl>")
			os.Exit(1)
		}
		formatFile(os.Args[2])
	default:
		if strings.HasSuffix(command, ".sl") {
			if len(os.Args) > 2 && os.Args[2] == "--format" {
				formatFile(command)
			} else {
				runFile(command)
			}
		} else {
			fmt.Printf("Comando desconhecido: %s\n", command)
			fmt.Println("Use 'sovy --help' para ver os comandos disponíveis")
			os.Exit(1)
		}
	}
}

func showHelp() {
	fmt.Println("Sovy - Interpretador da linguagem Solara v2.0.0")
	fmt.Println()
	fmt.Println("Uso:")
	fmt.Println("  sovy <arquivo.sl>          Executar arquivo")
	fmt.Println("  sovy install <biblioteca>  Instalar biblioteca")
	fmt.Println("  sovy list                  Listar bibliotecas instaladas")
	fmt.Println("  sovy --format <arquivo>    Formatar arquivo")
	fmt.Println("  sovy --help                Mostrar ajuda")
	fmt.Println("  sovy --version             Mostrar versão")
	fmt.Println()
	fmt.Println("Bibliotecas disponíveis:")
	fmt.Println("  smath  - Matemática avançada")
	fmt.Println()
	fmt.Println("Exemplos:")
	fmt.Println("  sovy programa.sl")
	fmt.Println("  sovy install smath")
	fmt.Println("  sovy list")
	fmt.Println()
	fmt.Println("Sintaxe para importar bibliotecas:")
	fmt.Println("  sovy <biblioteca> include")
}

func runFile(filename string) {
	if !fileExists(filename) {
		fmt.Printf("Erro: Arquivo '%s' não encontrado\n", filename)
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo: %v\n", err)
		os.Exit(1)
	}

	l := lexer.New(string(content))

	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println("Erros de sintaxe encontrados:")
		for _, err := range p.Errors() {
			fmt.Printf("  %s\n", err)
		}

		if len(program.Statements) > 0 {
			fmt.Println("Tentando executar o que foi possível...")
		} else {
			os.Exit(1)
		}
	}

	eval := evaluator.New()
	result := eval.Eval(program)

	if result != nil && result.Type() == "ERROR" {
		fmt.Printf("Erro de execução: %s\n", result.Inspect())
		os.Exit(1)
	}
}

func formatFile(filename string) {
	if !fileExists(filename) {
		fmt.Printf("Erro: Arquivo '%s' não encontrado\n", filename)
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo: %v\n", err)
		os.Exit(1)
	}

	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println("Erros de sintaxe encontrados:")
		for _, err := range p.Errors() {
			fmt.Printf("  %s\n", err)
		}
		os.Exit(1)
	}

	formatted := formatter.Format(program)

	err = ioutil.WriteFile(filename, []byte(formatted), 0644)
	if err != nil {
		fmt.Printf("Erro ao salvar arquivo formatado: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Arquivo '%s' formatado com sucesso!\n", filename)
}

func installLibrary(libraryName string) {
	libManager := library.NewLibraryManager()

	err := libManager.InstallLibrary(libraryName)
	if err != nil {
		fmt.Printf("Erro ao instalar biblioteca '%s': %v\n", libraryName, err)
		os.Exit(1)
	}

	fmt.Printf("Biblioteca '%s' instalada com sucesso!\n", libraryName)
}

func listLibraries() {
	libManager := library.NewLibraryManager()
	libraries := libManager.ListInstalledLibraries()

	if len(libraries) == 0 {
		fmt.Println("Nenhuma biblioteca instalada.")
		fmt.Println("Use 'sovy install <biblioteca>' para instalar bibliotecas.")
		return
	}

	fmt.Println("Bibliotecas instaladas:")
	for _, lib := range libraries {
		fmt.Printf("  %s\n", lib)
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
