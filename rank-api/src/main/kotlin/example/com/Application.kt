package example.com

import example.com.plugins.*
import io.ktor.server.application.*
import example.com.consumer.configureConsumer
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.launch
import kotlinx.coroutines.Dispatchers

import example.com.service.impl.RankServiceImpl
import example.com.repository.impl.RankRepositoryImpl

fun main(args: Array<String>) {
	io.ktor.server.netty.EngineMain.main(args)
}

fun Application.module() {
	val rankRepository = RankRepositoryImpl()
	val rankService = RankServiceImpl(rankRepository)

    configureSerialization()
    configureRouting(rankService)

	launch(Dispatchers.IO) {
		configureConsumer(rankService)
	}
}
