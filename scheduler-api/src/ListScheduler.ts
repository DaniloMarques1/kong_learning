import SchedulerRepository from './SchedulerRepository';

export default class ListScheduler {
  private schedulerRepository: SchedulerRepository;

  constructor(schedulerRepository: SchedulerRepository) {
    this.schedulerRepository = schedulerRepository;
  }

  execute() {
    return this.schedulerRepository.list();
  }

}
