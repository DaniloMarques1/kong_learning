package example.com.service

import example.com.model.KafkaMessage

interface RankService {
	suspend fun saveRank(message: KafkaMessage)
	fun getRank(): Map<String, Int>
}
