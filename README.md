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

**Zadanie 5**

Należy stworzyć aplikację kliencką wykorzystując bibliotekę React.js.
W ramach projektu należy stworzyć trzy komponenty: Produkty, Koszyk
oraz Płatności. Koszyk oraz Płatności powinny wysyłać do aplikacji
serwerowej dane, a w Produktach powinniśmy pobierać dane o produktach
z aplikacji serwerowej. Aplikacja serwera w jednym z trzech języków:
Kotlin, Scala, Go. Dane pomiędzy wszystkimi komponentami powinny być
przesyłane za pomocą React hooks.  

✅ 3.0 W ramach projektu należy stworzyć dwa komponenty: Produkty oraz
Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w
Produktach powinniśmy pobierać dane o produktach z aplikacji
serwerowej  
✅ 3.5 Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing  
✅ 4.0 Dane pomiędzy wszystkimi komponentami powinny być przesyłane za
pomocą React hooks  
✅ 4.5 Należy dodać skrypt uruchamiający aplikację serwerową oraz
kliencką na dockerze via docker-compose  
✅ 5.0 Należy wykorzystać axios’a oraz dodać nagłówki pod CORS  

Link do video:  
https://ujchmura-my.sharepoint.com/:v:/g/personal/piotr_klocek_student_uj_edu_pl/IQDl8xKe97JOTqyjEGvd67XQAQbdWC3b4Okm-uK9PZB330g?nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJPbmVEcml2ZUZvckJ1c2luZXNzIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXciLCJyZWZlcnJhbFZpZXciOiJNeUZpbGVzTGlua0NvcHkifX0&e=kQrtHM

***

**Zadanie 6**  

Należy stworzyć 20 przypadków testowych w jednym z rozwiązań:  

- Cypress JS (JS)  
- Selenium (Kotlin, Python, Java, JS, Go, Scala)  

Testy mają w sumie zawierać minimum 50 asercji (3.5). Mają również
uruchamiać się na platformie Browserstack (5.0). Proszę pamiętać o
stworzeniu darmowego konta via https://education.github.com/pack.  

✅ 3.0 Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium
(Kotlin, Python, Java, JS, Go, Scala)  
✅ 3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50
asercji  
❌ 4.0 Należy stworzyć testy jednostkowe do wybranego wcześniejszego
projektu z minimum 50 asercjami  
❌ 4.5 Należy dodać testy API, należy pokryć wszystkie endpointy z
minimum jednym scenariuszem negatywnym per endpoint  
❌ 5.0 Należy uruchomić testy funkcjonalne na Browserstacku  

Link do video:  
https://ujchmura-my.sharepoint.com/:v:/g/personal/piotr_klocek_student_uj_edu_pl/IQBtEVNlQ0fERoKiBAmkuYdDAXLhmr_fDmIpZqZfZIlKVnw?e=gJnRGD&nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJTdHJlYW1XZWJBcHAiLCJyZWZlcnJhbFZpZXciOiJTaGFyZURpYWxvZy1MaW5rIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXcifX0%3D

***

**Zadanie 7**

✅ 3.0 Należy dodać litera do odpowiedniego kodu aplikacji serwerowej w
hookach gita  
✅ 3.5 Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod
aplikacji serwerowej)  
✅ 4.0 Należy wyeliminować wszystkie zapaszki w kodzie w Sonarze (kod
aplikacji serwerowej)  
✅ 4.5 Należy wyeliminować wszystkie podatności oraz błędy bezpieczeństwa
w kodzie w Sonarze (kod aplikacji serwerowej)  
✅ 5.0 Należy wyeliminować wszystkie błędy oraz zapaszki w kodzie
aplikacji klienckiej  

## Repozytoria

### Backend
- GitHub: https://github.com/PiotrKlocek/project-backend
- SonarCloud: https://sonarcloud.io/project/overview?id=PiotrKlocek_project-backend

### Frontend
- GitHub: https://github.com/PiotrKlocek/project-frontend
- SonarCloud: https://sonarcloud.io/project/overview?id=PiotrKlocek_project-frontend


frontend:

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=PiotrKlocek_project-frontend&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=PiotrKlocek_project-frontend)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=PiotrKlocek_project-frontend&metric=bugs)](https://sonarcloud.io/summary/new_code?id=PiotrKlocek_project-frontend)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=PiotrKlocek_project-frontend&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=PiotrKlocek_project-frontend)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=PiotrKlocek_project-frontend&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=PiotrKlocek_project-frontend)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=PiotrKlocek_project-frontend&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=PiotrKlocek_project-frontend)

backend:

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=PiotrKlocek_project-backend&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=PiotrKlocek_project-backend)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=PiotrKlocek_project-backend&metric=bugs)](https://sonarcloud.io/summary/new_code?id=PiotrKlocek_project-backend)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=PiotrKlocek_project-backend&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=PiotrKlocek_project-backend)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=PiotrKlocek_project-backend&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=PiotrKlocek_project-backend)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=PiotrKlocek_project-backend&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=PiotrKlocek_project-backend)



## Commity

### Backend
- https://github.com/PiotrKlocek/project-backend/commits/main/

### Frontend
- https://github.com/PiotrKlocek/project-frontend/commits/main/

***

**Zadanie 8**

Należy skonfigurować klienta Oauth2 (4.0). Dane o użytkowniku wraz z
tokenem powinny być przechowywane po stronie bazy serwera, a nowy
token (inny niż ten od dostawcy) powinien zostać wysłany do klienta
(React). Można zastosować mechanizm sesji lub inny dowolny (5.0).
Zabronione jest tworzenie klientów bezpośrednio po stronie React'a
wyłączając z komunikacji aplikację serwerową.  

Prawidłowa komunikacja: react-sewer-dostawca-serwer(via return
uri)-react.  

✅ 3.0 logowanie przez aplikację serwerową (bez Oauth2)  
✅ 3.5 rejestracja przez aplikację serwerową (bez Oauth2)  
✅ 4.0 logowanie via Google OAuth2  
✅ 4.5 logowanie via Facebook lub Github OAuth2  
✅ 5.0 zapisywanie danych logowania OAuth2 po stronie serwera  

Link do video:  
https://ujchmura-my.sharepoint.com/:v:/g/personal/piotr_klocek_student_uj_edu_pl/IQAVFAsH-VpjTqhGbvOAGG7iAWG3yMWtwWucte4SjiFbpbA?nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJPbmVEcml2ZUZvckJ1c2luZXNzIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXciLCJyZWZlcnJhbFZpZXciOiJNeUZpbGVzTGlua0NvcHkifX0&e=bdxzcQ