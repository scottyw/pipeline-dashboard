import { observable } from 'mobx'
import Train from './train';
import Product from './Product';
import Job from './Job';

const axios = require('axios');


class Link {
  title = ""
  url = ""

  constructor(product, jobs) {
    this.url = product.URL;
    this.title = product.Title;
  }
}

class DataStore {
  @observable data = 'supercalifragilisticexpialidocious'
  @observable products = []
  @observable jobs = []
  @observable trains = []
  @observable links = []
  @observable title = ""
  @observable state = ""

  fetchProducts(cb) {
    var store = this;
    axios.get('/api/1/products')
      .then((res) => res.data)
      .then(function(res) {
        store.trains   = res.Trains.map((train)     => new Train(train))
        store.jobs     = res.Jobs.map((job)         => new Job(job, store.trains))
        store.products = res.Products.map((product) => new Product(product, store.jobs))
        store.links    = res.Links.map((link)       => new Link(link))
        store.title    = res.Title
        store.state    = "done"
        cb()
      })
      .catch((err) => {
        console.log("in axios ", err)
        store.state = "error"
      })

  }
}

export default DataStore
