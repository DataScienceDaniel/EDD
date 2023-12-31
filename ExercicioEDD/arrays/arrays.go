package arrays

import (
		"fmt"
		"ExercicioEDD/contatos"
)

func AdicionaContato(c contatos.Contato, a *[5]contatos.Contato) {
	for ind, contato := range a {
		if (contato == contatos.Contato{}) {
			a[ind] = c
			break
		}
	}

}

func ExcluiContato(a [5]contatos.Contato) [5]contatos.Contato {
	if (a[0] == contatos.Contato{}) {
		fmt.Println("Lista de contatos vazia!")
		return a
	}

	for ind, contato := range a {
		if (contato == contatos.Contato{}) {
			a[ind-1] = contatos.Contato{}
		}
	}

	return a
}



