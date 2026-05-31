import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.application.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import io.ktor.server.plugins.contentnegotiation.*
import io.ktor.serialization.kotlinx.json.*

import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation as ClientCN
import io.ktor.http.*

import kotlinx.serialization.Serializable
import kotlinx.serialization.json.Json

import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.hooks.ListenerAdapter
import net.dv8tion.jda.api.events.message.MessageReceivedEvent
import net.dv8tion.jda.api.requests.GatewayIntent

import kotlinx.coroutines.runBlocking

@Serializable
data class ChatRequest(val message: String)

@Serializable
data class ChatResponse(
    val reply: String,
    val filtered: Boolean = false
)

val client = HttpClient(CIO) {
    install(ClientCN) {
        json(Json { ignoreUnknownKeys = true })
    }
}

val PYTHON = System.getenv("PYTHON_SERVICE_URL") ?: "http://localhost:5000"

class BotListener : ListenerAdapter() {
    override fun onMessageReceived(event: MessageReceivedEvent) {
        if (event.author.isBot) return

        val msg = event.message.contentRaw

        runBlocking {
            try {
                val res: HttpResponse = client.post("$PYTHON/chat") {
                    contentType(ContentType.Application.Json)
                    setBody(ChatRequest(message = msg))
                }

                val body = res.bodyAsText()

                println("STATUS: ${res.status}")
                println("BODY: $body")

                if (!res.status.isSuccess()) {
                    event.channel.sendMessage("Błąd AI: $body").queue()
                    return@runBlocking
                }

                val data = Json { ignoreUnknownKeys = true }
                    .decodeFromString<ChatResponse>(body)

                event.channel.sendMessage(data.reply).queue()

            } catch (e: Exception) {
                e.printStackTrace()
                event.channel.sendMessage("Błąd połączenia z serwisem AI.").queue()
            }
        }
    }
}

fun main() {
    val token = System.getenv("DISCORD_TOKEN") ?: error("Brak DISCORD_TOKEN")

    JDABuilder.createDefault(token)
        .enableIntents(
            GatewayIntent.GUILD_MESSAGES,
            GatewayIntent.MESSAGE_CONTENT
        )
        .addEventListeners(BotListener())
        .build()

    embeddedServer(Netty, port = 8080) {
        install(ContentNegotiation) {
            json()
        }

        routing {
            get("/") {
                call.respondText("Bot działa")
            }
        }
    }.start(wait = true)
}