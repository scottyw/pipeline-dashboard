export default class Train {
  name = ""
  url = ""
  wallClockTime = 0
  totalTimeDuration = 0
  durationHours = ""
  durationMinutes = ""
  durationSortMinutes = ""
  queueTimeHours = ""
  queueTimeMinutes = ""
  queueTimeSortMinutes = ""
  timestamp = 0
  pipeline = ""
  version = ""
  startTime = ""
  endTime = ""
  errors = 0
  transients = 0

  constructor(train) {
    this.name                = train.Name;
    this.url                 = train.URL;
    this.pipeline            = train.Pipeline;
    this.version             = train.Version;
    this.durationHours       = train.DurationHours;
    this.durationMinutes     = train.DurationMinutes;
    this.durationSortMinutes = train.DurationSortMinutes;
    this.queueTimeHours       = train.QueueTimeHours;
    this.queueTimeMinutes     = train.QueueTimeMinutes;
    this.queueTimeSortMinutes = train.QueueTimeSortMinutes;
    this.timestamp           = train.Timestamp / 1000;
    this.startTime           = train.StartTime;
    this.endTime             = train.EndTime;
    this.errors              = train.Errors;
    this.transients          = train.Transients;

  }

  durationFormatted() {
    return `${this.durationHours}H, ${this.durationMinutes}M`
  }

  queueTimeFormatted() {
    return `${this.queueTimeHours}H, ${this.queueTimeMinutes}M`
  }

}
