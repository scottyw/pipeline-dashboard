import { observable } from 'mobx'
import Train from './train';

const axios = require('axios');


class Product {
  name: ""
  pipeline: ""
  wallClockTime: 0
  totalTimeDuration: 0
  startTime: ""
  errors: 0
  transients: 0
  allJobs: []

  constructor(product, jobs) {
    console.log(product);
    this.name = product.Name;
    this.pipeline = product.Pipeline;
    this.startTime = product.StartTime;
    this.wallClockTime = product.WallClockTime;
    this.totalTimeDuration = product.TotalTimeDuration;
    this.errors = product.Errors;
    this.transients = product.Transients;
    this.allJobs = jobs;
  }

  GetPipelines() {
    var retVal = [];
    retVal = this.allJobs.filter((job) => {
      return (this.pipeline === job.pipelineJob);
    });

    console.log(retVal);

    return retVal;
  }
}


class Job {
  url: ""
  pipeline: ""
  wallClockTime: 0
  totalTimeDuration: 0
  allTrains: []
  pipelineJob: ""
  startTime: ""
  endTime: ""
  errors: 0
  transients: 0
  version: ""
  buildNumber: 0

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
    this.startTime = Date.parse(job.JobDataStrings.StartTime);
    this.endTime = Date.parse(job.JobDataStrings.EndTime);
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

  GetTrains() {
    var retVal = [];
    retVal = this.allTrains.filter((train) => {
      return ((this.pipelineJob === train.pipeline) && (this.version === train.version));
    });

    return retVal;

  }
}


class DataStore {
  @observable data = 'supercalifragilisticexpialidocious'
  @observable products = []
  @observable jobs = []
  @observable trains = []
  @observable title = ""
  @observable state = ""

  fetchProducts(cb) {
    var store = this;
    axios.get('/api/1/products')
      .then((res: any) => res.data)
      .then(function(res: any) {
        store.trains   = res.Trains.map((train)     => new Train(train))
        store.jobs     = res.Jobs.map((job)         => new Job(job, store.trains))
        store.products = res.Products.map((product) => new Product(product, store.jobs))
        store.title    = res.Title
        store.state    = "done"
        cb()
      })
      .catch((err: any) => {
        console.log("in axios ", err)
        store.state = "error"
      })

  }
}

export default DataStore
