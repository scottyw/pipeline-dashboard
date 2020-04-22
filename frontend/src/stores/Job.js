const moment = require('moment');

export default class Job {
    url = ""
    pipeline = ""
    wallClockTime = 0
    totalTimeDuration = 0
    allTrains = []
    pipelineJob = ""
    startTime = ""
    endTime = ""
    errors = 0
    transients = 0
    version = ""
    buildNumber = 0
  
    constructor(job, trains) {
      console.log(job);
      this.url = job.URL;
      this.pipeline = job.Pipeline;
      this.pipelineJob = job.PipelineJob;
      this.wallClockTime = job.WallClockTime;
      this.totalTimeDuration = job.TotalTimeDuration;
      this.version = job.Version;
      this.jobDataStrings = job.JobDataStrings;
      this.buildNumber = job.BuildNumber;
      this.startTime = moment(job.JobDataStrings.StartTime, "YYYY-MM-DD HH:mm:ss Z PDT");
      this.endTime = moment(job.JobDataStrings.EndTime, "YYYY-MM-DD HH:mm:ss Z PDT");
      this.errors = job.Errors;
      this.transients = job.Transients;
      this.allTrains = trains;
    }
  
    totalFormatted() {
      return `${this.jobDataStrings.TotalHours}H, ${this.jobDataStrings.TotalMinutes}M`
    }
  
    wallClockFormatted() {
      return `${this.jobDataStrings.WallClockTimeHours}H, ${this.jobDataStrings.WallClockTimeMinutes}M`
    }

    queueTimeFormatted() {
      return `${this.jobDataStrings.QueueTimeHours}H, ${this.jobDataStrings.QueueTimeMinutes}M`
    }
  
    GetTrains() {
      var retVal = [];
      retVal = this.allTrains.filter((train) => {
        return ((this.pipeline === train.pipeline) && (this.version === train.version));
      });
  
      return retVal;
  
    }
}  