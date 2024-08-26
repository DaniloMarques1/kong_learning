import { connect, ConsumeMessage }  from 'amqplib';
import Scheduler from './Scheduler';
import CreateScheduler from './CreateScheduler';
import SchedulerRepository from './SchedulerRepository';

const QUEUE_NAME = 'scheduler-queue';

export default class Consumer {
  private createScheduler: CreateScheduler;

  constructor(schedulerRepository: SchedulerRepository) {
    this.createScheduler = new CreateScheduler(schedulerRepository);
  }

  async consume() {
    try {
      const QUEUE_URL = process.env['QUEUE_URL'];
      if (QUEUE_URL === null || QUEUE_URL === undefined) return;


      const conn =  await connect(QUEUE_URL);
      const channel = await conn.createChannel();
      await channel.assertQueue(QUEUE_NAME);

      await channel.consume(QUEUE_NAME, (msg: ConsumeMessage | null) => this.messageReceived(msg, this.createScheduler));
    } catch(err) {
      console.error(`Error connecting to rabbitmq ${err}`);
    }
  }

  messageReceived(msg: ConsumeMessage | null, createScheduler: CreateScheduler): void {
    console.log('Mensagem sendo consumida');
    if (msg == null) {
      console.log(`Mensagem recebida null ${msg}`);
      return;
    }

    try {
      const content = JSON.parse(msg.content.toString()) as Scheduler;
      createScheduler.execute(content);
    } catch(err) {
      console.error(err);
    }
  }


}
