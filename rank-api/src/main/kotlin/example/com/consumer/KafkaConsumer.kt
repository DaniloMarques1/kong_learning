package example.com.consumer

import org.apache.kafka.clients.consumer.KafkaConsumer
import java.util.Properties
import org.apache.kafka.common.serialization.StringDeserializer
import org.apache.kafka.clients.consumer.ConsumerConfig
import java.time.Duration
import com.google.gson.Gson

import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.launch

import com.google.gson.annotations.SerializedName

import example.com.service.RankService

data class KafkaMessage(
	@SerializedName("todo_id")
	val todoId: String,
	val email: String
)

interface Consumer {
	suspend fun consume() 
}

class KafkaConsumerImpl(private val consumer: KafkaConsumer<String, String>, private val rankService: RankService) : Consumer {
	suspend override fun consume() {
		while(true) {
			val records = consumer.poll(Duration.ofMillis(1000))
			for (record in records) {
				val value = record.value()
				val gson = Gson()
				val msg = gson.fromJson(value, KafkaMessage::class.java)
				rankService.saveRank(msg)
			}
		}
	}
}

suspend fun configureConsumer(rankService: RankService) {
	val properties = Properties().apply {
		put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:9092")
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
