package main

import "fmt"

func main() {
	fmt.Println("Maps:")

	usuario := map[string]string{
		"Gui":     "Devs",
		"Matheus": "Devs",
		"Rafa":    "Devs",
		"Lucky":   "Design",
		"Jynxx":   "Design",
		"Ppang":   "SM",
		"Kaisa":   "SM",
		"Caique":  "Nerds",
	}

	fmt.Println("esses são os funcionários do Mais Esports:")

	usuario["Noon"] = "Carecas"

	for name := range usuario {
		fmt.Println(name, "é do time de", usuario[name])
	}

}
