package example.com.repository

interface RankRepository {
	fun save(email: String)
	fun getRank(): Map<String, Int>
}
