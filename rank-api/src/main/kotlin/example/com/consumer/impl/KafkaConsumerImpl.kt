package example.com.consumer.impl

import com.google.gson.Gson

import example.com.model.KafkaMessage
import example.com.consumer.Consumer

import example.com.service.RankService

import org.apache.kafka.clients.consumer.KafkaConsumer
import java.time.Duration

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


