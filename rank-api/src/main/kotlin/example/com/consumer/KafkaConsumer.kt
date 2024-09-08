package example.com.consumer

import org.apache.kafka.clients.consumer.KafkaConsumer
import java.util.Properties
import org.apache.kafka.common.serialization.StringDeserializer
import org.apache.kafka.clients.consumer.ConsumerConfig

import example.com.service.RankService
import example.com.consumer.impl.KafkaConsumerImpl

interface Consumer {
	suspend fun consume() 
}

suspend fun configureConsumer(rankService: RankService) {
	val properties = Properties().apply {
		val kafkaUrl = System.getenv("KAFKA_URL")
		put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, kafkaUrl)
		put(ConsumerConfig.GROUP_ID_CONFIG, "rank-api-group")
		put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
		put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
		put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest")
	}
	val consumer = KafkaConsumer<String, String>(properties)
	consumer.subscribe(listOf("rank-topic"))


	val	kafkaConsumer = KafkaConsumerImpl(consumer, rankService)
	kafkaConsumer.consume()
}
