package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o nome da instancia: ")
	instanceName, _ := reader.ReadString('\n')
	instanceName = strings.TrimSpace(instanceName)

	fmt.Print("Digite o nome da struct: ")
	structName, _ := reader.ReadString('\n')
	structName = strings.TrimSpace(structName)

	fmt.Print("Digite o nome do arquivo: ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	nameSplited := strings.Split(fileName, "/")
	packageName := nameSplited[len(nameSplited)-2]

	content := `
package {package}
import (
    "context"

    "lucares.github.com/minicloud/minicloud/shared/utils"
)

var {instance} *{struct}

type {struct} struct {
    
}

func (uc *{struct}) Execute(ctx context.Context) error {
    return nil
}

func New{struct}(ctx context.Context) (*{struct}, error) {
    if instance == nil {
        // repo, err := utils.GetValueFromCTX[ports.UserRepositoryPort](ports.USER_REPOSITORY_KEY_CTX, ctx)
        // if err != nil {
        //     return nil, err
        // }

        instance = &{struct}{
            repo: repo,
        }
    }

    return instance, nil
}`
	content = strings.ReplaceAll(content, "{struct}", structName)
	content = strings.ReplaceAll(content, "{package}", packageName)
	content = strings.ReplaceAll(content, "{instance}", instanceName)

	path, err := filepath.Abs("minicloud/domain/use_cases/" + fileName)
	if err != nil {
		fmt.Println("Erro ao criar o path:", err)
		return
	}
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Arquivo criado com sucesso: " + fileName)
}
