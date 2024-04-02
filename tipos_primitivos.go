package main

import (
	f "tipos_primitivos/functions"
)

func main() {
	f.Print("Estou usando modulos para pegar funções de outros arquivos")

	f.VarType(10)
	f.VarType(+10)
	f.VarType(-10)
	f.VarType(10.9)
	f.VarType(+10.9)
	f.VarType(-10.9)
	f.VarType("Flavio")
	f.VarType(`Ket`)
	f.VarType(true)
	f.VarType(false)
}