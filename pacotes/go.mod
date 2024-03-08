module modulo

go 1.22.1

//para iniciar um modulo usamos o go mod init nomedomodulo
//nesse caso o nome do module é modulo
//dentro desse arquivo vai conter todas as dependencias do projeto, como se fosse o package.json em node
//foi criado um arquivo chamado modulo, ele é um arquivo compilado de todos os arquivos dentro da pasta raiz, dessa forma o computador pode rodar esse arquivo invés de escrever o arquivo main na mão
//para baixar pacotes externos basta digitar go get urldopacote
//go mod tidy - comando usado para remover todas as dependencias do projeto 

require github.com/badoux/checkmail v1.2.4 // indirect
