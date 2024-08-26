import express, { Request, Response, Router } from 'express';

import ListScheduler from './ListScheduler';
import SchedulerRepository from './SchedulerRepository';

export default class SchedulerController {
  private schedulerRepository: SchedulerRepository;
  constructor(schedulerRepository: SchedulerRepository) {
    this.schedulerRepository = schedulerRepository;
  }

  register(): Router {
    const router = express.Router();
    router.get('/', (_: Request, res: Response) => {
      const listScheduler = new ListScheduler(this.schedulerRepository);
      const schedulers = listScheduler.execute();
      res.status(200).json(schedulers);
    });

    return router;
  }

}
