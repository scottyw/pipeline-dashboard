export default class Train {
  name: ""
  url: ""
  wallClockTime: 0
  totalTimeDuration: 0
  durationHours: ""
  durationMinutes: ""
  durationSortMinutes: ""
  timestamp: 0
  pipeline: ""
  version: ""
  startTime: ""
  endTime: ""

  constructor(train) {
    console.log("Creating Train");
    console.log(train);
    this.name                = train.Name;
    this.url                 = train.URL;
    this.pipeline            = train.Pipeline;
    this.version             = train.Version;
    this.durationHours       = train.DurationHours;
    this.durationMinutes     = train.DurationMinutes;
    this.durationSortMinutes = train.DurationSortMinutes;
    this.timestamp           = train.Timestamp / 1000;
    this.startTime           = train.StartTime;
    this.endTime             = train.EndTime;
  }

  durationFormatted() {
    return `${this.durationHours}H, ${this.durationMinutes}M`
  }

}
