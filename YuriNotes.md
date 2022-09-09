🔗 Link Curso em videos:
https://www.youtube.com/c/AprendaGo/videos

Aula Atual: Operador curto de declaração
🔗 https://www.youtube.com/watch?v=QT7YvrEWMGE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=9

🚀 Cap. 1
🗣️ Uma visão geral do curso e funcionamento.

🚀 Cap. 2 (Variáveis, Valores & Tipos)
🗣️ Go Playground funcionamento.
🔗 https://play.golang.org

         - Run: Executa o Codigo
         - Share: Compartilhar o Codigo via um link (util para foruns)
         - Format: Formata o codigo
         - Imports: Identifica packages no codigo e importa os usatos e não importados e exclui os importados e não utilizados

    🗣️ Estrutura do código:
        - package main: Indica qual o arquivo principal que sera executado.
        - func main: Compreende todas as ações do código, (O loop de C).
        - import: Packages importados.

    🗣️ Packages:
        - Coleções de funções para serem utilizadas. um package possui n funções.

        - Notação: pacote.Identificador.
            -- Exemplo: fmt.Println()

        - Lendo as Docks oficiais:
            --Exemplo func Println: func Println(a ...interface{}) (n int, err error)
                ---func: Função
                ---Println: Identificador dentro do packege.
                ---(a ...): Pode passar quantos parâmetros quiser
                ---(interface{}): Os parâmetros podem ser de qualquer type (types são discrimados dentro da interface, como ela não especifica ( {} ) pode qualque type.
                ---(n int, err error): Oque a função retorna, no caso retorna o numero de elementos e se teve algum erro em sua execussão.

        - Usando a função:
            --A Função Println tanto printa na linha o texto, quanto quanto atribuida a variaveis retorna o tamanho do texto.

            --Como a função me retorna 2 valores, eu preciso de duas variaveis para atribuir cada valor (numero, erros), a ordem de declaração delas corresponde a ordem em que serão atribuidos o valor comforme a Dock.

              func main() {
                  numero, erros:= fmt.Println("texto", "oi")
                  fmt.Println(numero,erros)
               }

            --Caso eu não queira o numero de elementos eu descarto ao atribuir o valor das variaveis:

               func main() {
                   _, erros:= fmt.Println("texto", "oi")
                   fmt.Println(erros)
               }
