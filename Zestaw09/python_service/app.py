from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
import requests
import random

app = FastAPI(title="Serwis AI dla Bota E-biznes")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

class ChatRequest(BaseModel):
    message: str

MODEL = "ibm/granite-4.1-3b"

OPENINGS = [
    "Cześć! W czym mogę pomóc w naszym sklepie odzieżowym?",
    "Witaj! Szukasz czegoś konkretnego?",
    "Dzień dobry! Jak mogę pomóc?",
    "Hej! Interesują Cię ubrania, dostawa czy zwroty?",
    "Witamy w sklepie odzieżowym!",
    "Miło Cię widzieć! Jakiego produktu szukasz?",
    "W czym mogę dzisiaj pomóc przy zakupach?",
    "Zapraszamy do naszego sklepu! Jak mogę pomóc?",
    "Czy interesuje Cię konkretny rodzaj odzieży?",
    "Jestem asystentem sklepu odzieżowego. W czym mogę pomóc?"
]

CLOSINGS = [
    "Dziękuję za rozmowę!",
    "Miłego dnia!",
    "Zapraszam ponownie!",
    "Do zobaczenia!",
    "Udanych zakupów!",
    "Dziękuję za skorzystanie z pomocy naszego sklepu.",
    "Mam nadzieję, że pomogłem. Do usłyszenia!",
    "Zapraszamy ponownie na zakupy.",
    "Życzę udanego wyboru produktów.",
    "W razie kolejnych pytań chętnie pomogę."
]

END_WORDS = [
    "koniec",
    "bye",
    "do widzenia",
    "zakończ",
    "zakoncz",
    "wyjdź",
    "wyjdz"
]

ALLOWED_KEYWORDS = [
    "ubranie", "ubrania", "odzież", "moda", "sklep",
    "kurtka", "kurtki", "kurtek",
    "spodnie", "koszula", "bluza",
    "sukienka", "buty",
    "rozmiar", "rozmiary",
    "kolor", "cena", "promocja",
    "promocje", "rabat",
    "dostawa", "zwrot",
    "reklamacja", "zamówienie"
]

POSITIVE_WORDS = [
    "super",
    "świetnie",
    "dobry",
    "ładny",
    "fajne",
    "idealne",
    "polecam"
]

NEGATIVE_WORDS = [
    "zły",
    "fatalny",
    "problem",
    "drogo",
    "reklamacja",
    "zwrot"
]

def is_store_related(text):
    text = text.lower()
    return any(word in text for word in ALLOWED_KEYWORDS)

def analyze_sentiment(text):
    text = text.lower()

    positive = sum(1 for w in POSITIVE_WORDS if w in text)
    negative = sum(1 for w in NEGATIVE_WORDS if w in text)

    if positive > negative:
        return {
            "label": "pozytywny",
            "score": positive
        }

    if negative > positive:
        return {
            "label": "negatywny",
            "score": -negative
        }

    return {
        "label": "neutralny",
        "score": 0
    }

@app.get("/")
def home():
    return {"status": "ok"}

@app.get("/health")
def health():
    return {
        "status": "ok",
        "model": MODEL
    }

@app.get("/opening")
def opening():
    return {
        "message": random.choice(OPENINGS)
    }

@app.get("/closing")
def closing():
    return {
        "message": random.choice(CLOSINGS)
    }

@app.post("/chat")
def chat(request: ChatRequest):

    user_message = request.message.strip()

    if user_message.lower() in END_WORDS:
        return {
            "reply": random.choice(CLOSINGS),
            "conversation_ended": True
        }

    sentiment = analyze_sentiment(user_message)

    if not is_store_related(user_message):
        return {
            "reply": "Przepraszam, mogę odpowiadać tylko na pytania związane ze sklepem odzieżowym.",
            "filtered": True,
            "sentiment": sentiment
        }

    try:

        payload = {
            "model": MODEL,
            "messages": [
                {
                    "role": "system",
                    "content": "Jesteś asystentem sklepu odzieżowego. Odpowiadasz krótko po polsku."
                },
                {
                    "role": "user",
                    "content": user_message
                }
            ],
            "temperature": 0.7,
            "stream": False
        }

        response = requests.post(
            "http://host.docker.internal:1234/v1/chat/completions",
            json=payload,
            timeout=60
        )

        response.raise_for_status()

        result = response.json()

        reply = result["choices"][0]["message"]["content"]

        reply_sentiment = analyze_sentiment(reply)

        if reply_sentiment["label"] == "negatywny":
            reply = (
                "Przepraszam, odpowiedź została przefiltrowana "
                "ze względu na negatywny wydźwięk. "
                "Mogę pomóc w wyborze produktów, rozmiarów, dostawy lub zwrotów."
            )

        return {
            "reply": reply,
            "filtered": False,
            "sentiment": sentiment,
            "reply_sentiment": reply_sentiment,
            "sentiment_filtered": reply_sentiment["label"] == "negatywny"
        }

    except Exception as e:
        return {
            "reply": "Błąd połączenia z LM Studio.",
            "filtered": False,
            "sentiment": sentiment,
            "error": str(e)
        }