import express, { Request, Response, Router } from 'express';

import ListScheduler from './ListScheduler';
import SchedulerRepositoryMemoryImpl from './SchedulerRepositoryMemoryImpl';

export default class SchedulerController {
  constructor() {
  }

  register(): Router {
    const router = express.Router();
    router.get('/', (_: Request, res: Response) => {
      const schedulerRepository = new SchedulerRepositoryMemoryImpl();
      const listScheduler = new ListScheduler(schedulerRepository);
      const schedulers = listScheduler.execute();
      res.status(200).json(schedulers);
    });

    return router;
  }

}
