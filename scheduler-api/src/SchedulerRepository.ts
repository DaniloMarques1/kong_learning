import Scheduler from './Scheduler';

interface SchedulerRepository {
  save(scheduler: Scheduler): void;
  list(): Array<Scheduler>;
}

export default SchedulerRepository;
