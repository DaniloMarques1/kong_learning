package example.com

import example.com.plugins.*
import io.ktor.server.application.*
import example.com.consumer.configureConsumer
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.launch
import kotlinx.coroutines.Dispatchers

fun main(args: Array<String>) {
	io.ktor.server.netty.EngineMain.main(args)
}

fun Application.module() {
    configureSerialization()
    configureRouting()

	launch(Dispatchers.IO) {
		configureConsumer()
	}
}
