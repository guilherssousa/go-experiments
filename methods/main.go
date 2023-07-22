package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) apresentar() {
	fmt.Printf("Meu nome é %s e tenho %d anos.\n", u.nome, u.idade)
}

func (u *usuario) fazerAniversario() {
	u.idade++
	fmt.Printf("%s fez %d anos! Parabéns!\n", u.nome, u.idade)
}

type grupo struct {
	nome          string
	membros       []usuario
	apenasMaiores bool
}

func (g grupo) listarMembros() {
	membros := ""

	for i, m := range g.membros {
		if i != 0 {
			membros += ", "
		}
		membros += m.nome
	}

	fmt.Println(membros)
}

func (g *grupo) adicionarPessoa(u usuario) (bool, error) {
	if g.apenasMaiores && u.idade < 18 {
		fmt.Printf("%s não tem idade o suficiente para participar do grupo %s.\n", u.nome, g.nome)
		return false, fmt.Errorf("%s não tem idade o suficiente para participar do grupo %s.\n", u.nome, g.nome)
	}

	g.membros = append(g.membros, u)

	fmt.Printf("%s foi adicionado com sucesso no grupo %s.\n", u.nome, g.nome)
	return true, nil
}

func escrever() {
	fmt.Println("Escrevendo...")
}

func main() {
	escrever()

	grupo := grupo{
		nome:          "Lolzinho da zoeira",
		apenasMaiores: true,
	}

	gui := usuario{
		nome:  "Guilherme",
		idade: 18,
	}

	yuyu := usuario{
		nome:  "Sayuri",
		idade: 17,
	}

	le := usuario{
		nome:  "Letícia",
		idade: 24,
	}

	gael := usuario{
		nome:  "Gael",
		idade: 20,
	}

	pedro := usuario{
		nome:  "Pedro",
		idade: 20,
	}

	grupo.adicionarPessoa(gui)
	grupo.adicionarPessoa(yuyu)
	grupo.adicionarPessoa(le)
	grupo.adicionarPessoa(gael)
	grupo.adicionarPessoa(pedro)

	grupo.listarMembros()

	yuyu.fazerAniversario()
	grupo.adicionarPessoa(yuyu)

	grupo.listarMembros()

}
