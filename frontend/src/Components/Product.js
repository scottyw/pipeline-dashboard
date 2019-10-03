import React, { Component } from 'react'
import { observer } from 'mobx-react'

import { Button } from '@puppet/react-components';

import ProductPipeline from './ProductPipeline';
import './Product.css';

@observer
class Product extends Component {

  constructor(props) {
    super(props)

    this.state = {
      showPipelines: false
    }
  }

  hms(nanoseconds) {
    var startSeconds = nanoseconds / 1000000000
    var hours = Math.floor(startSeconds / 3600);
    startSeconds = startSeconds - hours * 3600;

    var minutes = Math.floor(startSeconds / 60);
    // var seconds = startSeconds - minutes * 60;

    return `${hours}H, ${minutes}M`
  }

  togglePipeline() {
    this.setState({
      showPipelines: !this.state.showPipelines
    })
  }

  productPipelines() {
    var pipelines = []

    if (this.state.showPipelines) {
      pipelines = this.props.product.GetPipelines().map((value, index) => {
        return (<ProductPipeline key={index} pipeline={value} />)
      });
    }

    return pipelines
  }

  render () {
    console.log(this.props.product);
    return (
      <div className="row text-left">
        <div className="col-12">
          <div className="row product-row">
            <div className="col-5">
              <b>{this.props.product.name}</b>
            </div>
            <div className="col-1">
            </div>
            <div className="col-1">
              {this.props.product.totalTimeDuration}
            </div>
            <div className="col-2">
            </div>
            <div className="col-2">
            </div>
            <div className="col-1">
              <Button href="#" onClick={() => this.togglePipeline()}>{this.state.showPipelines ? 'Close' : 'Pipelines'}</Button>
            </div>
          </div>
          <div className="surround-pipelines"  >
            {this.productPipelines()}
          </div>
        </div>
      </div>
    )
  }
}

export default Product
