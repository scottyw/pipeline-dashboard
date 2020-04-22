export default class Product {
  name = ""
  pipeline = ""
  wallClockTime = 0
  queueTimeDuration = 0
  totalTimeDuration = 0
  startTime = ""
  errors = 0
  transients = 0
  allJobs = []

  constructor(product, jobs) {
    console.log(product);
    this.name = product.Name;
    this.pipeline = product.Pipeline;
    this.startTime = product.StartTime;
    this.queueTimeDuration = product.QueueTimeMinutes
    this.wallClockTime = product.WallClockTime;
    this.totalTimeDuration = product.TotalTimeDuration;
    this.errors = product.Errors;
    this.transients = product.Transients;
    this.allJobs = jobs;
  }

  queueTimeFormatted() {
    let hours = this.queueTimeDuration / 60;
    let minutes = this.queueTimeDuration % 60;
    return `${hours.toFixed(0)}H, ${minutes}M`
  }

  GetPipelines() {
    var retVal = [];
    retVal = this.allJobs.filter((job) => {
      return (this.pipeline === job.pipeline);
    });

    return retVal;
  }
}
