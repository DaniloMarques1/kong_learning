import Scheduler from './Scheduler';
import SchedulerRepository from './SchedulerRepository';

export default class CreateScheduler {
  private repository: SchedulerRepository;
  constructor(repository: SchedulerRepository) {
    this.repository = repository;
  }

  execute(scheduler: Scheduler) {
    this.repository.save(scheduler);
  }
}
