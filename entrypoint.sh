#!/bin/sh

echo "Activamos CUPS"
/activa_cups.sh

echo "Activamos microservicio"
/go/CUPSasMicroservice
