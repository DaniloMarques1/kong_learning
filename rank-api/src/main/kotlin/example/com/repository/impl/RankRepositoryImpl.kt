package example.com.repository.impl

import example.com.repository.RankRepository

class RankRepositoryImpl : RankRepository {
	private val rank: MutableMap<String, Int>

	constructor() {
		rank = mutableMapOf<String, Int>()
	}

	override fun save(email: String) {
		var current = rank.getOrElse(email) {
			0
		}
		rank[email] = ++current
	}

	override fun getRank(): Map<String, Int> {
		return rank
	}
}
