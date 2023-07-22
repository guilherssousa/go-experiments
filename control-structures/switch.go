package main

import "fmt"

func diaDaSemana(numero uint8) (string, error) {
	switch numero {
	case 1:
		return "Domingo", nil
	case 2:
		return "Segunda-feira", nil
	case 3:
		return "Terça-feira", nil
	case 4:
		return "Quarta-feira", nil
	case 5:
		return "Quinta-feira", nil
	case 6:
		return "Sexta-feira", nil
	case 7:
		return "Sábado", nil
	default:
		return "", fmt.Errorf("Não foi possível compilar o dia.")
	}

}

func main() {
	fmt.Println("Switch")

	var dia, err = diaDaSemana(4)

	if err != nil {
		panic(err)
	}

	fmt.Println("Hoje é", dia)
}
