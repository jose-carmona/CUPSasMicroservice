# CUPS as a Microservice

Prueba de concepto.

Consta de:
* CUPS
* microservicio en go

Todo montado alpinelinux en docker

### ejecución

Build:

```
build -t josecarmona/cups-m .
```

La prueba de concepto se lanza con:

```
docker run --rm --name cups-m -it -v $PWD/files:/files -p 18080:8080 -p 1631:631 josecarmona/cups-m
```

Instalación de impresoras:

```
docker exec cups-m /activa_impresoras.sh
```

ping:

```
curl "127.0.0.1:18080"
```

lpstat:

```
curl "127.0.0.1:18080/cups/lpstat"
```

lp:

```
curl "127.0.0.1:18080/cups/lp/impresora/fichero"
```
