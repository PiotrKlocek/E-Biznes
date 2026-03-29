FROM sbtscala/scala-sbt:eclipse-temurin-17.0.15_6_1.12.8_3.8.2

WORKDIR /app

COPY . .

EXPOSE 9000

CMD ["sbt", "run", "-Dplay.server.http.address=0.0.0.0"]