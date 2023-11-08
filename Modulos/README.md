## Saludos en GO
este paquete proporciona una forma simple de obtener saludos personales aleatorios


## instalacion 
ejecuta el siguiente comando para instalar el paquete:
```bash
go get -u github.com/alicia-granados/greetings
```
MODULOS o paquetes
 INICIALIZAR modulo
go mod init github.com/alicia-granados/go-experiments/Modulo

remplazar mod para que detecte la ubicacion
go mod edit -replace github.com/alicia-granados/greetings=../greetings


aplique los cambios 
go mod tidy
