import { connect, ConsumeMessage }  from 'amqplib';
import Scheduler from './Scheduler';

const QUEUE_NAME = 'scheduler-queue';

export default class Consumer {
  constructor() {
  }

  async consume() {
    try {
      const conn =  await connect('amqp://fitz:fitz@rabbitmq:5672');
      const channel = await conn.createChannel();
      await channel.assertQueue(QUEUE_NAME);

      await channel.consume(QUEUE_NAME, this.messageReceived);
    } catch(err) {
      console.error(`Error connecting to rabbitmq ${err}`);
    }
  }

  messageReceived(msg: ConsumeMessage | null): void {
    if (msg == null) {
      console.log(`Mensagem recebida null ${msg}`);
      return;
    }

    try {
      const content = JSON.parse(msg.content.toString()) as Scheduler;
      console.log(content);
    } catch(err) {
      console.error(err);
    }
  }


}
