import Scheduler from './Scheduler';
import SchedulerRepository from './SchedulerRepository';

export default class SchedulerRepositoryMemoryImpl implements SchedulerRepository {
  private schedulers: Array<Scheduler>;
  constructor() {
    this.schedulers = new Array<Scheduler>();
  }
  
  save(scheduler: Scheduler) {
    this.schedulers.push(scheduler);
  }

  list() {
    return this.schedulers;
  }

}
