package main

import 
(
	"fmt"
	"math/rand"
	"sort"
	"time"
	"io/ioutil"
    "strings"
    "strconv"

)

type Lista struct 
{
	indice int
	lburst int
	ltempo int
}

func main() {
	var nProcesso, escalonador int
	var burst string
	lista := make([]Lista, 0, nProcesso)
	sair := 0
	for sair == 0 {
		fmt.Print("Escolha um algoritmos de escalonamento ou aperte 0 para sair \n1. FCFS \n2. SJF \n3. SRTF \n4. Round Robin \n5. Multinível \n0. Sair \n")
		fmt.Scanf("%d\n", &escalonador)
		if escalonador == 0 {
			fmt.Print("Saindo...")
			sair = 1
			break
		} else {
			lista, nProcesso, burst := MENU(lista, nProcesso, burst)
			switch escalonador {
			case 1:
				FCFS(lista, nProcesso)
				break
			case 2:
				SJF(lista, nProcesso)
				break
			case 3:
				SRTF(lista, nProcesso)
			case 4:
				RR(burst, lista, nProcesso)
			case 5:
				Multinivel(lista, nProcesso)
			default:
				fmt.Print("ERRO!Insira um valor válido")
				break
			}

		}

	}

}

func MENU(lista []Lista, nProcesso int, burst string) ([]Lista, int, string) {
	fmt.Print("\nBurst manual ou aleatório(M/A)? ")
	fmt.Scanln(&burst)
	proc := Lista{}
	var tburst, tempo, p int
	if burst == "M" || burst == "m" || burst == "Manual" || burst == "manual" {
		fmt.Print("Informe a qtde de processos")
		fmt.Scanf("%d\n", &nProcesso)
		for i := 0; i < nProcesso; i++ {
			fmt.Print("informe o tamanho do Burst \n")
			fmt.Scanf("%d\n", &tburst)
			fmt.Print("informe o tempo de Chegada \n")
			fmt.Scanf("%d\n", &tempo)
			p = i
			proc.indice = p
			proc.lburst = tburst
			proc.ltempo = tempo
			lista = append(lista, proc)
			burst = "M"
		}
	} else {
		b, er := ioutil.ReadFile("file.txt") // just pass the file name
    	if er != nil {
        	fmt.Print(er)
    	}
    	str := string(b) // convert content to a 'string'
    	linha := strings.Split(str, "\n")
    	for i := range linha {
        	x := strings.Split(linha[i], ";")
        	nProcesso = len(linha)
        	fmt.Print(nProcesso)
        	if p, err := strconv.Atoi(x[0]); err == nil {
    			proc.indice = p
			}
			if tburst, err := strconv.Atoi(x[1]); err == nil {
    			proc.lburst = tburst
			}
			if tempo, err := strconv.Atoi(x[2]); err == nil {
    			proc.ltempo = tempo
			}
			if tempo, err := strconv.Atoi(x[2]); err == nil {
    			proc.ltempo = tempo
			}
    		lista = append(lista, proc)
			burst = "M"
    	}

	}
	fmt.Print("Todos os processos\n")
	fmt.Print("Processo\tBurst\tTempoChegada\n")
	for i := 0; i < nProcesso; i++ {
		fmt.Print(lista[i].indice, "\t", lista[i].lburst, "\t", lista[i].ltempo, "\n")
	}

	return lista, nProcesso, burst
}

/*=====================================================
FCFS (First Come First Serve)
=======================================================*/
func FCFS(lista []Lista, nProcesso int) {
	sort.Slice(lista, func(i, j int) bool {
		return lista[i].ltempo < lista[j].ltempo
	})
	fmt.Print("\nSimulando FCFS\n")
	fmt.Print("Proc\tBurst\tTChegada\n")
	for i := 0; i < nProcesso; i++ {
		fmt.Print(lista[i].indice, "\t", lista[i].lburst, "\t", lista[i].ltempo, "\n")
	}
	fmt.Print("Tempo de Espera\n")
	e := 0
	i := 0
	b := 0
	t := 0
	for nProcesso > i {
		if lista[i].ltempo <= t{
			if b == 0 {
				e = 0
			}
			if e == 0 {
				fmt.Print(t - lista[i].ltempo, "\n")
				e = 1
				b = lista[i].lburst
				i++
			}
		}
		if i > 0{
			if lista[i - 1].ltempo <= t{
				b = b - 1
			}
		}
		t++

	}
	fmt.Print("Tempo de Espera Médio\n")
	fmt.Print(int(t/nProcesso), "\n")
}
/*=====================================================
SJF Shortest Job First
Ordena a fila de processo pelo menor tempo de burst
=======================================================*/
func SJF(lista []Lista, nProcesso int) {
	sort.Slice(lista, func(i, j int) bool {
		return lista[i].lburst < lista[j].lburst
	})
	fmt.Print("\nSimulando SJF\n")
	fmt.Print("Proc\tBurst\tTChegada\n")
	for i := 0; i < nProcesso; i++ {
		fmt.Print(lista[i].indice, "\t", lista[i].lburst, "\t", lista[i].ltempo, "\n")
	}
	fmt.Print("Tempo de Espera\n")
	e := 0
	i := 0
	b := 0
	t := 0
	for nProcesso > i {
		if lista[i].ltempo <= t{
			if b == 0 {
				e = 0
			}
			if e == 0 {
				fmt.Print(t - lista[i].ltempo, "\n")
				e = 1
				b = lista[i].lburst
				i++
			}
		}
		if i > 0{
			if lista[i - 1].ltempo <= t{
				b = b - 1
			}
		}
		t++

	}
	fmt.Print("Tempo de Espera Médio\n")
	fmt.Print(int(t/nProcesso), "\n")
	
}
/*=====================================================
SRTF(Shorttest-Remainning-Time-First)
=======================================================*/
func SRTF(lista []Lista, nProcesso int) {
	lista2 := make([]Lista, 0, nProcesso)
	sort.Slice(lista, func(i, j int) bool {
		return lista[i].ltempo < lista[j].ltempo
	})
	fmt.Print("\nSimulando SRTF\n")
	fmt.Print("Proc\tBurst\tTChegada\n")
	atual := 0
	proximo := 1
	count := lista[atual].ltempo
	process := make([]int, 0, nProcesso)
	burst := make([]int, 0, nProcesso)
	tcheg := make([]int, 0, nProcesso)
	lista2 = append(lista2, lista[atual])
	fmt.Print(lista2[atual].indice, "\t", lista2[atual].lburst, "\t", count, "\n")
	for atual < nProcesso {
		if proximo < nProcesso && count == lista[proximo].ltempo {
			lista2 = append(lista2, lista[proximo])
			sort.Slice(lista, func(i, j int) bool {
				return lista[i].lburst < lista[j].lburst
			})
			process = append(process, lista2[atual].indice)
			burst = append(burst, lista2[atual].lburst)
			tcheg = append(tcheg, lista2[atual].ltempo)
			fmt.Print(lista2[atual].indice, "\t", lista2[atual].lburst, "\t", count, "\n")
			proximo++
		} else if lista2[atual].lburst == 0 {
			atual++
			proximo++
			if atual < len(lista2) {
				process = append(process, lista2[atual].indice)
				burst = append(burst, lista2[atual].lburst)
				tcheg = append(tcheg, lista2[atual].ltempo)
				fmt.Print(lista2[atual].indice, "\t", lista2[atual].lburst, "\t", count, "\n")
			} else {
				break
			}
		} else {
			lista2[atual].lburst--
			count++
		}
	}
	fmt.Print("Tempo de Espera\n")
	e := 0
	i := 0
	b := 0
	t := 0
	for nProcesso > i {
		if lista[i].ltempo <= t{
			if b == 0 {
				e = 0
			}
			if e == 0 {
				fmt.Print(t - lista[i].ltempo, "\n")
				e = 1
				b = lista[i].lburst
				i++
			}
		}
		if i > 0{
			if lista[i - 1].ltempo <= t{
				b = b - 1
			}
		}
		t++

	}
	fmt.Print("Tempo de Espera Médio\n")
	fmt.Print(int(t/nProcesso), "\n")

}

func RR(comand string, lista []Lista, nProcesso int) {
	var nQuantum int
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	if comand == "M" {
		fmt.Print("Informe o tamanho do Quantum")
		fmt.Scanf("%d\n", &nQuantum)
	} else {
		nQuantum = r.Intn(100)
	}
	lista2 := make([]Lista, 0, nProcesso)
	sort.Slice(lista, func(i, j int) bool {
		return lista[i].ltempo < lista[j].ltempo
	})
	fmt.Print("Simulando RR\n")
	fmt.Print("Processo\tBurst\tTempoChegada\n")
	atual := 0
	proximo := 1
	auxiliar := 0
	count := lista[atual].ltempo
	process := make([]int, 0, nProcesso)
	burst := make([]int, 0, nProcesso)
	tcheg := make([]int, 0, nProcesso)
	lista2 = append(lista2, lista[atual])
	for atual < nProcesso {
		if proximo < nProcesso && count >= lista[proximo].ltempo {
			lista2 = append(lista2, lista[proximo])
			process = append(process, lista2[atual].indice)
			burst = append(burst, lista2[atual].lburst)
			tcheg = append(tcheg, lista2[atual].ltempo)
			proximo++
		} else if atual < len(lista2) && lista2[atual].lburst <= 0 {
			lista2 = append(lista2[:atual], lista2[atual+1:]...)
			atual++
			auxiliar++
			if auxiliar == nProcesso {
				break
			} else {
				atual = 0
			}
		} else {
			burst = append(burst, lista2[atual].lburst)
			tcheg = append(tcheg, lista2[atual].ltempo)
			process = append(process, lista2[atual].indice)
			fmt.Print(lista2[atual].indice, "\t", lista2[atual].lburst, "\t", count, "\n")
			if lista2[atual].lburst <= nQuantum {
				count = count + lista2[atual].lburst
				lista2[atual].lburst = lista2[atual].lburst - lista2[atual].lburst
			} else {
				count = count + nQuantum
				lista2[atual].lburst = lista2[atual].lburst - nQuantum
			}
			atual++
			if atual >= len(lista2) {
				atual = 0
			}
		}
	}
	fmt.Print("Tempo de Espera\n")
	e := 0
	i := 0
	b := 0
	t := 0
	for nProcesso > i {
		if lista[i].ltempo <= t{
			if b == 0 {
				e = 0
			}
			if e == 0 {
				fmt.Print(t - lista[i].ltempo, "\n")
				e = 1
				b = lista[i].lburst
				i++
			}
		}
		if i > 0{
			if lista[i - 1].ltempo <= t{
				b = b - 1
			}
		}
		t++

	}
	fmt.Print("Waiting time\n")
	w := 0
	i := 0
	q := 0
	fmt.Print(w, "\n")
	for nProsses > i {
		q = q + w
		w = w + lista[i].lburst
		i++

		if nProsses > i {
			w = w - lista[i].ltempo
			fmt.Print(w, "\n")
		} else {
			break
		}
	}
	
	fmt.Print("Tempo de Espera\n")
	w := 0
	i := 0
	q := 0
	fmt.Print(process[i], "\t", w, "\n")
	for nProcesso > i {
		q = q + w
		w = w + burst[i]
		i++

		if nProcesso > i {
			w = w - tcheg[i]
			fmt.Print(process[i], "\t", w, "\n")
		} else {
			break
		}
	}
	fmt.Print("Tempo de Espera Médio\n")
	fmt.Print(int(q/nProcesso), "\n")
}

func Multinivel(lista []Lista, nProcesso int) {
	sort.Slice(lista, func(i, j int) bool {
		return lista[i].ltempo < lista[j].ltempo
	})
	fmt.Print("Simulando Multiniveis(3)\n")
	lista1 := make([]Lista, 0, nProcesso)
	lista2 := make([]Lista, 0, nProcesso)
	lista3 := make([]Lista, 0, nProcesso)
	count1 := 0
	count2 := 0
	for i := range lista {
		lista1 = append(lista1, lista[i])
		if lista[i].lburst > 30 {
			lista1[count1].lburst = 30
			lista[i].lburst = lista[i].lburst - 30
			lista2 = append(lista2, lista[i])
			count1++
			if lista[i].lburst > 20 {
				lista2[count2].lburst = 20
				lista[i].lburst = lista[i].lburst - 20
				lista3 = append(lista3, lista[i])
				count2++
			}
		}
	}
	if len(lista1) > 0 {
		FCFS(lista1, len(lista1))
	}
	if len(lista2) > 0 {
		SJF(lista2, len(lista2))
	}
	if len(lista3) > 0 {
		SRTF(lista3, len(lista3))
	}
}
