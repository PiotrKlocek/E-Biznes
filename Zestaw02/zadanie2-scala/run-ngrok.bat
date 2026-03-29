@echo off
docker build -t zadanie2-scala .
start cmd /k "docker run -p 9000:9000 --name zadanie2-app zadanie2-scala"
timeout /t 20
ngrok http 9000