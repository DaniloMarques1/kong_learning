package example.com.plugins

import io.ktor.server.application.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import io.ktor.http.HttpStatusCode

import example.com.service.RankService

fun Application.configureRouting(rankService: RankService) {
    routing {
		get("/rank") {
			call.response.status(HttpStatusCode.OK)
			call.respond(rankService.getRank())
		}
    }
}
