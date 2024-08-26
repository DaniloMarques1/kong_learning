import express from 'express';

import SchedulerController from './SchedulerController';
import SchedulerRepositoryMemoryImpl from './SchedulerRepositoryMemoryImpl';
import Consumer from './consumer';

export default class App {
  private port: number;
  private schedulerController: SchedulerController;
  private consumer: Consumer;

  constructor(port: number) {
    this.port = port;
    const schedulerRepository = new SchedulerRepositoryMemoryImpl();
    this.schedulerController = new SchedulerController(schedulerRepository);
    this.consumer = new Consumer(schedulerRepository);
  }

  async start() {
    const app = express();
    app.use(express.json());

    app.use('/scheduler', this.schedulerController.register());

    console.log(`Server running on port ${this.port}`);
    app.listen(this.port);


    await this.consumer.consume();
  }
}
