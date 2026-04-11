**Zadanie 1 Docker**

✅ 3.0 obraz ubuntu z Pythonem w wersji 3.10  
✅ 3.5 obraz ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem  
✅ 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC  
SQLite w ramach projektu na Gradle (build.gradle)  
✅ 4.5 stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji
przez CMD oraz gradle  
✅ 5.0 dodać konfigurację docker-compose  


**Docker images:**

- <b>3.0: https://hub.docker.com/repository/docker/piterek111/zadanie3.0
- 3.5: https://hub.docker.com/repository/docker/piterek111/zadanie_3.5
- 4.0: https://hub.docker.com/repository/docker/piterek111/zadanie_4.0
- 4.5: https://hub.docker.com/repository/docker/piterek111/zadanie_4.5 </b>

Link do video: https://ujchmura-my.sharepoint.com/:v:/g/personal/piotr_klocek_student_uj_edu_pl/IQAE4u4BEUdgToMsGSVRqAbwAW6zN3JTeEPpWmr3y1VE1_E?nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJTdHJlYW1XZWJBcHAiLCJyZWZlcnJhbFZpZXciOiJTaGFyZURpYWxvZy1MaW5rIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXcifX0%3D&e=dzvAqJ


***

**Zadanie 2**

✅ 3.0 Należy stworzyć kontroler do Produktów  
✅ 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane  
pobierane z listy  
✅ 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy
zgodnie z CRUD  
✅ 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać
skrypt uruchamiający aplikację via ngrok  
✅ 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD  

Link do filmiku przedstawiającego działanie aplikacji:
https://ujchmura-my.sharepoint.com/:v:/g/personal/piotr_klocek_student_uj_edu_pl/IQDrdmAqrWt5SqpL2RgxMWT0AUiyh-Cew8py7q7ODdiB730?nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJTdHJlYW1XZWJBcHAiLCJyZWZlcnJhbFZpZXciOiJTaGFyZURpYWxvZy1MaW5rIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXcifX0%3D&e=z417cC


***

**Zadanie 3**

✅ 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor, która pozwala na przesyłanie wiadomości na platformę Discord   
✅ 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z platformy Discord skierowane do aplikacji (bota)  
✅ 4.0 Zwróci listę kategorii na określone żądanie użytkownika  
✅ 4.5 Zwróci listę produktów wg żądanej kategorii  
❌ 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack lub Messenger  

Link do video:  
https://ujchmura-my.sharepoint.com/:v:/g/personal/piotr_klocek_student_uj_edu_pl/IQBKgjVoy_FWQqlzeA-p-ks1ARsIPFW9G8rWw_I9qAY7-M4?nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJPbmVEcml2ZUZvckJ1c2luZXNzIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXciLCJyZWZlcnJhbFZpZXciOiJNeUZpbGVzTGlua0NvcHkifX0&e=ifkOHJ

Obraz dockerowy:
https://hub.docker.com/repository/docker/piterek111/ktor-discord-bot/general

***

**Zadanie 4**

Należy stworzyć projekt w echo w Go. Należy wykorzystać gorm do
stworzenia kilka modeli, gdzie pomiędzy dwoma musi być relacja. Należy
zaimplementować proste endpointy do dodawania oraz wyświetlania danych
za pomocą modeli. Jako bazę danych można wybrać dowolną, sugerowałbym
jednak pozostać przy sqlite.  

✅ 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie
miała kontroler Produktów zgodny z CRUD  
✅ 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz
wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast
listy)  
✅ 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint  
✅ 4.5 Należy stworzyć model kategorii i dodać relację między kategorią,
a produktem  
✅ 5.0 pogrupować zapytania w gorm’owe scope'y  

Link do video:  

https://ujchmura-my.sharepoint.com/:v:/g/personal/piotr_klocek_student_uj_edu_pl/IQCjiMkBJZjAR6L50YU3DDAQAfu0NWre-GkGrhY-b67lcTk?nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJPbmVEcml2ZUZvckJ1c2luZXNzIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXciLCJyZWZlcnJhbFZpZXciOiJNeUZpbGVzTGlua0NvcHkifX0&e=yLjd6P

***