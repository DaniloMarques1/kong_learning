FROM openjdk:17-jdk-slim

WORKDIR /app

COPY . .

EXPOSE 8080

RUN ./gradlew build

ENV KAFKA_URL="kafka:9092"

CMD ["java", "-jar", "./build/libs/com.example.ktor-sample-all.jar"]
