package example.com.consumer

import org.apache.kafka.clients.consumer.KafkaConsumer
import java.util.Properties
import org.apache.kafka.common.serialization.StringDeserializer
import org.apache.kafka.clients.consumer.ConsumerConfig
import java.time.Duration
import com.google.gson.Gson

import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.launch

data class KafkaMessage(val todo_id: String, val email: String)

interface Consumer {
	suspend fun consume() 
}

class KafkaConsumerImpl(private val consumer: KafkaConsumer<String, String>) : Consumer {
	suspend override fun consume() {
		println("Comecando a consumir mensagens")
		while(true) {
			val records = consumer.poll(Duration.ofMillis(1000))
			for (record in records) {
				val value = record.value()
				val gson = Gson()
				val msg = gson.fromJson(value, KafkaMessage::class.java)
				println(msg)
			}
		}
	}
}

suspend fun configureConsumer() {
	println("Configurando consumer")
	val properties = Properties().apply {
		put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:9092")
		put(ConsumerConfig.GROUP_ID_CONFIG, "rank-api-group")
		put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
		put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
		put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest")
	}
	val consumer = KafkaConsumer<String, String>(properties)
	consumer.subscribe(listOf("rank-topic"))


	val	kafkaConsumer = KafkaConsumerImpl(consumer)
	kafkaConsumer.consume()
}
