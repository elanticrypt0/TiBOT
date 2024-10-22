#!/bin/bash

# Verifica si se ha pasado un argumento
if [ -z "$1" ]; then
  echo "Por favor, proporciona un nombre."
else
  echo "Hi, $1"
fi
