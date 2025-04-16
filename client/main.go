package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	log.Println("Soy un log")

	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	if globals.ClientConfig == nil {
		log.Fatal("No se pudo cargar la configuración")
	}

	log.Println(globals.ClientConfig.Mensaje)
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	mensajes := utils.LeerConsolaHastaLineaVacia()
	log.Printf("Se ingresaron %d mensajes", len(mensajes))
	log.Println(mensajes)

	utils.GenerarYEnviarPaquete(mensajes)
}
