package main

import (
  "encoding/json"
  "fmt"
)

type Aluno struct{
  Nome  string
  Idade int
}

type Localidade struct{
  X int `json:"valor_x"`
  Y int `json:"valor_y"`
  Nome string `json:"nome_de_localidade"`
  valor int
}

func(p *Localidade) Soma(l Localidade){
  p.X += l.X
  p.Y += l.Y
}

func main(){
     aluno := Aluno{
       Nome: "Allan",
       Idade: 24,
     }

     fmt.Printf("Aluno:", aluno)

     minhaCasa := Localidade{3, 4, "casa", 500}

     estruturaVazia := Aluno{}

     escolaLocalidade := Localidade{
       Y:	100,
			 X:	200,
			 Nome: "escola",
     }

		 var outraLocalidade Localidade
		 outraLocalidade.X = 10
		 outraLocalidade.Y = 20
		 outraLocalidade.Nome = "trabalho"

		 	fmt.Println("minha casa:", minhaCasa)
 	 		fmt.Println("outra localidade:", outraLocalidade)
 			fmt.Println("escola:", escolaLocalidade)
 			fmt.Println("aluno:", aluno)
 			fmt.Printf("estrutura vazia: %q\r\n", estruturaVazia)

			minhaCasa.Soma(outraLocalidade)

			fmt.Printf("localidade minha casa somada com outra localidade", minhaCasa)

			fmt.Printf("minha casa: %v\r\n",minhaCasa)
			fmt.Printf("minha casa: %+v\r\n",minhaCasa)

			j, err := json.Marshal(minhaCasa)
			if err != nil{
				panic(err)
			}

			fmt.Printf("minha casa json", string(j))

			novaLocalidade := Localidade{}
			err = json.Unmarshal(j, &novaLocalidade)
			if err != nil {
				panic(err)
			}

			fmt.Printf("json depois do Unmarshal:", novaLocalidade)

}
