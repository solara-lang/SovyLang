package library

import (
	"fmt"
	"os"
	"path/filepath"
)

type LibraryManager struct {
	installedLibraries map[string]bool
	libraryPath        string
}

func NewLibraryManager() *LibraryManager {
	homeDir, _ := os.UserHomeDir()
	libPath := filepath.Join(homeDir, ".sovy", "libraries")

	return &LibraryManager{
		installedLibraries: make(map[string]bool),
		libraryPath:        libPath,
	}
}

func (lm *LibraryManager) IsLibraryInstalled(name string) bool {

	libFile := filepath.Join(lm.libraryPath, name+".slib")
	_, err := os.Stat(libFile)
	return err == nil
}

func (lm *LibraryManager) InstallLibrary(name string) error {
	switch name {
	case "smath":
		return lm.installSMathLibrary()
	default:
		return fmt.Errorf("biblioteca '%s' não encontrada", name)
	}
}

func (lm *LibraryManager) installSMathLibrary() error {

	err := os.MkdirAll(lm.libraryPath, 0755)
	if err != nil {
		return fmt.Errorf("erro ao criar diretório de bibliotecas: %v", err)
	}


	smathContent := `{
	"name": "smath",
	"version": "1.0.0",
	"description": "Biblioteca de matemática avançada para Solara",
	"functions": {
		"potencia": "Calcula potência (base, expoente)",
		"raiz": "Calcula raiz quadrada",
		"sin": "Calcula seno",
		"cos": "Calcula cosseno",
		"tan": "Calcula tangente",
		"abs": "Valor absoluto",
		"max": "Valor máximo entre dois números",
		"min": "Valor mínimo entre dois números",
		"pi": "Constante PI (3.14159...)"
	}
}`

	libFile := filepath.Join(lm.libraryPath, "smath.slib")
	err = os.WriteFile(libFile, []byte(smathContent), 0644)
	if err != nil {
		return fmt.Errorf("erro ao instalar biblioteca smath: %v", err)
	}

	fmt.Println("Biblioteca 'smath' instalada com sucesso!")
	fmt.Printf("Local: %s\n", libFile)
	return nil
}

func (lm *LibraryManager) LoadLibrary(name string) error {
	if !lm.IsLibraryInstalled(name) {
		return fmt.Errorf("biblioteca '%s' não está instalada. Use: sovy install %s", name, name)
	}

	lm.installedLibraries[name] = true
	return nil
}

func (lm *LibraryManager) IsLibraryLoaded(name string) bool {
	return lm.installedLibraries[name]
}

func (lm *LibraryManager) RequiresMath() bool {

	return !lm.IsLibraryLoaded("smath")
}

func (lm *LibraryManager) ListInstalledLibraries() []string {
	var libraries []string

	files, err := os.ReadDir(lm.libraryPath)
	if err != nil {
		return libraries
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".slib" {
			name := file.Name()[:len(file.Name())-6]
			libraries = append(libraries, name)
		}
	}

	return libraries
}
