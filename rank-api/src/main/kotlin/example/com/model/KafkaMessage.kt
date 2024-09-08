package example.com.model

import com.google.gson.annotations.SerializedName

data class KafkaMessage(
	@SerializedName("todo_id")
	val todoId: String,
	val email: String
)
