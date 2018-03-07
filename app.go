package main

import (
	"os"
  "os/exec"
  "log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

)

func main() {
	app := iris.New()

	app.Get("/", index)

	apiGroup := app.Party("/cups")
	{
		apiGroup.Get("/lp/{impresora:string}/{fichero:string}", lp)
    apiGroup.Get("/lpstat", lpstat)
	}

	app.Run(iris.Addr(":" + port()))
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if port[0] == ':' {
		port = port[1:]
	}

	return port
}

func index(ctx context.Context) {
	ctx.Writef("Welcome to CUPSasMicroservice")
}

func lpstat(ctx context.Context) {
  cmd := "lpstat -a"
  log.Println("lpstat")
  out, err := exec.Command("bash", "-c", cmd).Output()
  if err != nil {
    ctx.Writef(err.Error())
  }
  ctx.Writef(string(out))

}

func lp(ctx context.Context) {
  directorio := "/files"
  impresora := ctx.Params().Get("impresora")
  fichero := ctx.Params().Get("fichero")
  cmd := "/usr/bin/lp"
  log.Printf("%s impresora=%s fichero=%s\n", cmd, impresora, directorio + "/" + fichero)
  out, err := exec.Command("bash", "-c", cmd + " -d " + impresora + " " + directorio + "/" + fichero).Output()
  if err != nil {
    ctx.Writef(err.Error())
  }
  ctx.Writef(string(out))

}
