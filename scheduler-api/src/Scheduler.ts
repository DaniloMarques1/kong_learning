export default class Scheduler {
  email: string;
  notificationDate: Date;

  constructor(email: string, notificationDate: Date) {
    this.email = email;
    this.notificationDate = notificationDate; 
  }
}
