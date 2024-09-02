package example.com.plugins

import io.ktor.server.application.*
import io.ktor.server.response.*
import io.ktor.server.routing.*

import com.example.model.Priority
import com.example.model.Task

fun Application.configureRouting() {
	println("Configurando as rotas")
    routing {
		get("/tasks") {
			call.respond(mapOf("tasks" to listOf(
				Task("Cleaning my Bedroom", "I need to clean my bedroom", Priority.Medium),
				Task("Study Kong", "Need to learn before my job starts", Priority.Vital),
			)))
        }
    }
}
