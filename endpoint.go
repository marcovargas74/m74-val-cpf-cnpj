package m74bankAPI

import (
	"fmt"
	"log"
	"net/http"
)

const (
	//serverDirDev  = "/home/intelbras/projetos-go/src/github.com/marcovargas74/m74tester/appliance/public"
	//serverDirProd = "/home/iap/appliance/public"
	serverPort = ":8080"
)

//Mode é o modo do teste se "prod" ou "dev"
//var Mode = "dev"

//var Mode = "prod"

//WorkDir é o diretorio de trabalho vai ser diferenta pra Dev ou Prod
//var WorkDir string

//Ifaces interfaces de rede do equipamento
//var Ifaces []net.Interface

func getPlayerPoints(name string) string {

	if name == "Maria" {
		return "20"
	}

	if name == "Pedro" {
		return "10"
	}
	return ""

}

//ServidorJogador teste
func ServidorJogador(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/jogadores/"):]
	fmt.Fprint(w, getPlayerPoints(player))
}

//StartAPI inicia o servidor http do appliance
func StartAPI(modo string) {
	/*
		if modo != "" {
			Mode = modo
		}
	*/
	//serverDir := WorkDir + "/public"
	//serverDir := "/tmp"
	//fmt.Println("serverDir: ", serverDir)

	//server := http.FileServer(http.Dir(serverDir))
	//server := http.FileServer(http.Dir(serverDir))
	//fmt.Println("serverDir: ", server)
	//http.Handle("/", server)

	//HandleFuncions()

	//	log.Fatal(http.ListenAndServe(serverPort, nil))
	tratador := http.HandlerFunc(ServidorJogador)
	if err := http.ListenAndServe(":5000", tratador); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
}

//HandleFuncions Prepara funcoe para serem usadas
func HandleFuncions() {
	//http.HandleFunc("/macrec", MacAddressRec)
}

/*
//Incluir o codigo que grama MAC na Memoria aqui
func setMacInInterface(index int, mac string) error {
	fmt.Printf("OK [%d] Mac[%s]", index, mac)
	//formatMessage(w, "OK MAC GRAVADO COM SUCESSO:%s", iface.Name)
	//formatMessage(w, "OK MAC GRAVADO COM SUCESSO")
	return nil
}
*/

/*
//MacAddressRec grava endereco mac nas interfaces data
func MacAddressRec(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	mac1 := r.FormValue("mac1")
	//fmt.Fprintf(w, "<valor><font color='#2e802e' size='4'>INFO Teste de %s</font></valor>", testName)
	err := setMacInInterface(1, mac1)
	if err != nil {
		fmt.Println(err)
		//formatMessage(w, "ERR Erro ao Gravar MAC1")
		return
	}

	fmt.Println("OK MacAddressRec")
	//formatMessage(w, "OK MAC GRAVADO COM SUCESSO")
}
*/

/*
//func formatMessage(w http.ResponseWriter, message string) string {
func formatMessage(w http.ResponseWriter, format string, a ...interface{}) {

	message := fmt.Sprintf(format, a...)

	//default: PRETA
	color := "#000000"
	erro := 0

	//#INFO Verde
	if strings.Contains(message, "INFO") {
		color = "#2e802e"
	}

	//#OK Azul
	if strings.Contains(message, "OK") {
		color = "#0066FF"
	}

	//#WARN Laranja
	if strings.Contains(message, "WARN") {
		color = "#FF9900"
	}

	//#ERR vermelho
	if strings.Contains(message, "ERR") {
		color = "#FF0000"
		erro = 1
	}

	fmt.Fprintf(w, "<valor><font color='%s' size='3'>\t%s</font></valor>", color, message)

	if erro != 0 {
		fmt.Fprintf(w, "<resposta>1</resposta>")
	} else {
		fmt.Fprintf(w, "<resposta>0</resposta>")
	}

}
*/

/*
//ReadFile Le arquivo
func ReadFile(w http.ResponseWriter, r *http.Request) {
	//filename := r.FormValue("nomeArquivo")
	//Nao da bola para o dado que vem do js
	filename := WorkDir + "/public/static/hard.conf"

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", filename, "Err:", err)
		formatMessage(w, "ERR Erro ao abrir Arquivo de configuracao")
		return
	}

	fmt.Fprintf(w, "%s", body)
	//fmt.Println("ReadFile: ", filename)
	//fmt.Println("Body: ", string(body))
	//fmt.Fprintf(w, "<resposta>0</resposta>")
}

//SelfTestIni inicia os testes
func SelfTestIni(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Println(r.Form)
	//fmt.Println(r.FormValue("aData"))
	//fmt.Fprintf(w, "SelfTestIni OK")
	//fmt.Fprintf(w, "<h1><font color='#2e802e' size='4'>INFO Teste %s %s</font></h1>", "Memoria", "OK")

	erro := 0
	testName := "config"
	if Mode == "dev" {
		testName = "hardware"
	}

	//HoraCerta(w, r)
	fmt.Fprintf(w, "<valor><font color='#2e802e' size='4'>INFO Teste de %s</font></valor>", testName)
	if erro != 0 {
		fmt.Fprintf(w, "<resposta>1</resposta>")
	} else {
		fmt.Fprintf(w, "<resposta>0</resposta>")
	}

	fmt.Println("SelfTestIni OK ")
	//if (ODO_ == "prod" {
	///$test = $_POST['x'];
	///$exec = "sh /usr/bin/selftest.sh $test";

	//showInterfaces()
}


//CheckErr chech if have err and print a log default
func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
		//panic(e)
	}
}
*/
