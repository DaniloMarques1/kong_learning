package example.com.service

import example.com.consumer.KafkaMessage

data class RankDto(val email: String, val tasksDone: Int)

interface RankService {
	suspend fun saveRank(message: KafkaMessage)
	fun getRank(): Map<String, Int>
}
