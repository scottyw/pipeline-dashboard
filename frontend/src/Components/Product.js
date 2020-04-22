import React, { Component } from 'react'
import { observer } from 'mobx-react'

import { Link } from '@puppet/react-components';

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
    return (
      <div className="text-left rc-table-row">
        <div className="col-12">
          <div className="row product-row rc-table-row">
            <div className="rc-table-cell col-5">
              <b>{this.props.product.name}</b>
            </div>
            <div className="rc-table-cell col-1">
              {this.props.product.queueTimeFormatted()}
            </div>
            <div className="rc-table-cell col-1">
            </div>
            <div className="rc-table-cell col-1">
              {this.props.product.totalTimeDuration}
            </div>
            <div className="rc-table-cell col-1">
            </div>
            <div className="rc-table-cell col-1">
            </div>
            <div className="rc-table-cell col-1">
              {this.props.product.errors} / {this.props.product.transients}
            </div>
            <div className="rc-table-cell col-1">
              <Link href="#" onClick={() => this.togglePipeline()}>{this.state.showPipelines ? 'Close' : 'Pipelines'}</Link>
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
