#!/usr/bin/env bash
mkdir teste
for i in $(seq 100)
do
  echo "arquivo-${i}.txt"
  touch teste/arquivo-${i}.txt
  [[ $? -ne 0 ]] && break
done
