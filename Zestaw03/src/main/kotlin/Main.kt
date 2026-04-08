import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.application.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import io.ktor.server.plugins.contentnegotiation.*
import io.ktor.serialization.kotlinx.json.*

import kotlinx.serialization.Serializable

import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.hooks.ListenerAdapter
import net.dv8tion.jda.api.events.message.MessageReceivedEvent
import net.dv8tion.jda.api.requests.GatewayIntent

@Serializable
data class Product(val name: String, val category: String)

val categories = listOf("electronics", "books", "clothes")

val products = listOf(
    Product("Laptop", "electronics"),
    Product("Phone", "electronics"),
    Product("Book", "books"),
    Product("T-shirt", "clothes")
)

class BotListener : ListenerAdapter() {
    override fun onMessageReceived(event: MessageReceivedEvent) {
        if (event.author.isBot) return

        val msg = event.message.contentRaw
        println("ODEBRANO: $msg od ${event.author.name}")

        when {
            msg == "!categories" -> {
                event.channel.sendMessage("Kategorie: ${categories.joinToString()}").queue()
            }

            msg.startsWith("!products") -> {
                val category = msg.split(" ").getOrNull(1)
                val filtered = products.filter { it.category == category }

                if (filtered.isEmpty()) {
                    event.channel.sendMessage("Brak produktów").queue()
                } else {
                    event.channel.sendMessage(filtered.joinToString { it.name }).queue()
                }
            }

            else -> {
                event.channel.sendMessage("Nieznana komenda").queue()
            }
        }
    }
}

fun main() {
    val token = System.getenv("DISCORD_TOKEN")

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
                call.respondText("API działa")
            }

            get("/categories") {
                call.respond(categories)
            }

            get("/products/{category}") {
                val category = call.parameters["category"]
                val filtered: List<Product> = products.filter { it.category == category }
                call.respond<List<Product>>(filtered)
            }
        }
    }.start(wait = true)
}