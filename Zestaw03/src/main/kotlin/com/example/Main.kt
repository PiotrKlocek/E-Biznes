package com.example

import kotlinx.coroutines.runBlocking

fun main() = runBlocking {
    val webhookUrl = System.getenv("DISCORD_WEBHOOK_URL")
        ?: error("Brak zmiennej środowiskowej DISCORD_WEBHOOK_URL")

    val discordClient = DiscordClient(webhookUrl)

    discordClient.sendMessage("Cześć z aplikacji Kotlin + Ktor!")
    println("Wiadomość została wysłana na Discord.")
}