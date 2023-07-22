package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	gui := pessoa{"Guilherme", "Sousa", 19}
	ruan := pessoa{"Ruan", "Pablo", 20}

	estudanteGui := estudante{gui, "Engenharia de Software", "UNESA"}
	estudanteRuan := estudante{ruan, "Edição de Vídeos", "UGV"}
	estudantePedro := estudante{
		pessoa: pessoa{
			nome:      "Pedro",
			sobrenome: "Fracassi",
			idade:     20,
		},
		curso:     "Ciências da Computação",
		faculdade: "Insper",
	}
	estudanteLe := estudante{
		pessoa: pessoa{
			nome:      "Leticia",
			sobrenome: "Abreu",
			idade:     24,
		},
		curso:     "한국어",
		faculdade: "욘세이대",
	}

	fmt.Println(estudanteGui)
	fmt.Println(estudanteRuan)
	fmt.Println(estudantePedro)
	fmt.Println(estudanteLe)
}
