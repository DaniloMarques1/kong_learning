package example.com.service.impl

import example.com.service.RankService
import example.com.service.RankDto
import example.com.consumer.KafkaMessage

import example.com.repository.RankRepository

class RankServiceImpl(val rankRepository: RankRepository) : RankService {
	override suspend fun saveRank(message: KafkaMessage) {
		rankRepository.save(message.email)
	}

	override fun getRank(): Map<String, Int> {
		return rankRepository.getRank()
	}
}
