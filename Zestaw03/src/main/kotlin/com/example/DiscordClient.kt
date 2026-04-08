package com.example

import io.ktor.client.HttpClient
import io.ktor.client.call.body
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.post
import io.ktor.client.request.setBody
import io.ktor.http.ContentType
import io.ktor.http.contentType
import io.ktor.serialization.kotlinx.json.json

class DiscordClient(
    private val webhookUrl: String
) {
    private val client = HttpClient(CIO) {
        install(ContentNegotiation) {
            json()
        }
    }

    suspend fun sendMessage(message: String) {
        val response = client.post(webhookUrl) {
            contentType(ContentType.Application.Json)
            setBody(DiscordMessage(content = message))
        }

        if (response.status.value !in 200..299 && response.status.value != 204) {
            error("Błąd wysyłania wiadomości: ${response.status} ${response.body<String>()}")
        }
    }
}