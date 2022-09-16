🔗 Link Curso em videos:
https://www.youtube.com/c/AprendaGo/videos

Aula Atual: Cap. 6 – Fluxo de Controle – 1. Entendendo fluxo de controle
🔗 https://www.youtube.com/watch?v=1G-tbQ6UE_A&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=38

🚀 Cap. 1

### Uma visão geral do curso e funcionamento.

🚀 Cap. 2 (Variáveis, Valores & Tipos)

### Go Playground funcionamento.

🔗 https://play.golang.org

- Run: Executa o Codigo
- Share: Compartilhar o Codigo via um link (util para foruns)
- Format: Formata o codigo
- Imports: Identifica packages no codigo e importa os usatos e não importados e exclui os importados e não utilizados

### Estrutura do código:

- package main: Indica qual o arquivo principal que sera executado.
- func main: Compreende todas as ações do código, (O loop de C).
- import: Packages importados.

### Packages:

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

### Operador curto de declaração

- := parece uma marmota (gopher).
- Uso:
  -- Tipagem automática
  -- Só pode repetir se houverem variáveis novas
  -- != do assignment operator (operador de atribuição)
  -- Só funciona dentro de codeblocks
- Terminologia:
  -- statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões
  -- expressão -> qualquer coisa que "produz um resultado"
- Lição principal:
  -- := utilizado pra criar novas variáveis, dentro de code blocks
  -- = para atribuir valores a variáveis já existentes

### A palavra-chave `var`

- Forma de declarar variavel fora de codeblocks
  -exemplo: var y = 10
- Utilizada quando variavel necessita de, package level scope.

### Explorando tipos

- Tipos em Go são estáticos.
  -- Uma variavel de tipo number só pode receber numbers.
- Use := sempre que possível.
- Use var para package-level scope.

### Valor zero

- Os zeros:
  -- ints: 0
  -- floats: 0.0
  -- booleans: false
  -- strings: ""
  -- pointers, functions, interfaces, slices, channels, maps: nil

### O pacote fmt

- Strings:
  --Interpreted string literals
  ---Strings interpretadas, (yuri tem x anos) === yuri tem 27 anos
  --Raw string literals.
  ---Strings literais, (`yuri tem x anos`) === yuri tem x anos
- Print printar valor na tela:
  -- func Print(a ...interface{}) (n int, err error)
  -- func Println(a ...interface{}) (n int, err error)
  -- func Printf(format string, a ...interface{}) (n int, err error)
- Sprint() Atribui o valor em () para alguma váriavel (STRINGS Apenas):
  -- func Sprint(a ...interface{}) string
  -- func Sprintf(format string, a ...interface{}) string
  -- func Sprintln(a ...interface{}) string
- Fprint() Atribui o valor em () para alguma váriavel (Qualquer binário):
  -- func Fprint(w io.Writer, a ...interface{}) (n int, err error)
  -- func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
  -- func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

### Criando seu próprio tipo

- Sintaxe
  --exemplo: type hotdog int
- A vantagem de criar seus proprios tipos eu ainda n'ao sei, mas acima foi criado o tipo hotdog de sub-tipo inteiro

### Conversão

- basta colocar o tipo e oque vocề quer modificar entre ()
  -- exemplo: a = int(b)

🚀 Cap. 4

### Fundamentos da Programação

- Tipo Boolean
  -- Funciona normalmente como qualquer outro boolean
  -- false == nulo

- Tipo Numerico
  -- int e int32 não são a mesma coisa
  -- Para "misturá-los" é necessário conversão
  -- Regra geral: use somente int quando não quiser especificar.
  --- 🔗 https://www.youtube.com/watch?v=3yIKCLSWAHA&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=24

### -- Overflow

     -- Um uint16, por exemplo, vai de 0 a 65535.
     -- Se adicionar um numero a uma variavel declarada com uint16 vai zera-la (rodometro)

- Tipo String
  -- Não muito diferente

### Constantes

- São valores imutáveis.
- Podem ser tipadas ou não:
  -- const oi = "Bom dia"
  -- const oi string = "Bom dia"
- As não tipadas só terão um tipo atribuido a elas quando forem usadas.
- exemplo: const (
    x = 10
    y = 20
    z = 30
)

### Iota

- Numa declaração de constantes, o identificador iota representa números sequenciais.
- Na prática. https://go.dev/play/p/eSrwoQjuYR

### Deslocamento de bits

