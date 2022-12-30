package main

import "fmt"

// Declaro minha Struct ( Metodo )
type Carro struct {
	name string
}

// Ao declarar minha função assim: func (c Carro) andou()
//
// Quando a struct Carro ( Method ) fazer o ( Bind / Instancia ) com a func Carro em c
// c.name na real recebe o valor definido no func main que no caso é ( Ka )
// Porém após a execução do method a variavel c.name é substituida por ( Picanto )
// Entao name agora vale ( Picanto ) Obs.: Ela so vale ( Pincanto ) dentro desse escopo.
// No entanto após a execução do method é impresso o conteudo da variável c.name novamente dentro do bloco main.
// LÁ SIM A VARIAVEL c.name vale ( KA)
// Por isso o Resultado....:
// Picanto
// Endereço de memoria: 0xc000014250
// ka

// Ao declarar minha funcao assim: func (c *Carro) andou()
//
// Quando a struct Carro ( Method ) fazer o ( Bind / Instancia ) com a func Carro em c
// c.name na real representa o endereço de memória (0xc000096210)
// Isso quer dizer que quando a func main é iniciada ela recebe ( KA ),
// Porém após a execução do method o endereço de memória é substituida por ( Picanto )
// Portanto após a execução do method a variavel c.name ainda continua com o valor ( Picanto )
// Pois a mudanca foi realizada na memoria.
// Por isso o Resultado....:
// Picanto
// Endereço de memoria: 0xc000096210
// Picanto

func (c *Carro) andou() {

	// O valor da variável name ( Picanto ) só vale dentro desse escopo.
	c.name = "Picanto"
	fmt.Println(c.name)
	fmt.Println("Endereço de memoria:", &c.name)

}

func main() {
	carro := Carro{
		name: "ka",
	}

	carro.andou()
	fmt.Println(carro.name)
}
