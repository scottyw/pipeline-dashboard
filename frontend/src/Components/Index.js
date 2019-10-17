import React, { Component } from 'react'
import { inject, observer } from 'mobx-react'

import Product from './Product';

import { OverlayTrigger } from 'react-bootstrap';

import Timeline from './Timeline';

const renderWallClockTimeTooltip = props => (
  <div
    {...props}
    style={{
      backgroundColor: 'rgba(0, 0, 0, 0.85)',
      padding: '2px 10px',
      color: 'white',
      borderRadius: 3,
      ...props.style,
    }}
  >
    The time the first job starts and the last job stops.  Wall Clock time is not aggregated for products because there's too much variability
  </div>
);

const renderTotalTimeTooltip = props => (
  <div
    {...props}
    style={{
      backgroundColor: 'rgba(0, 0, 0, 0.85)',
      padding: '2px 10px',
      color: 'white',
      borderRadius: 3,
      ...props.style,
    }}
  >
    The time it would take to run all jobs consecutively.  This is meant to be a measure of how much "work" your job is doing.
  </div>
);



@inject('rootStore')
@observer
class Index extends Component {
  static isPrivate = true

  componentWillMount() {
    var _this = this;
    this.props.rootStore.dataStore.fetchProducts(function() {
      console.log(_this.props.rootStore.dataStore.products);
    });
  }

  render () {
    const { username } = this.props
    console.log(username)

    var productsTable = [];
    console.log(this.props.rootStore.dataStore.products)

    if (this.props.rootStore.dataStore.products) {
      this.props.rootStore.dataStore.products.forEach((product, index) => {
        productsTable.push(
          <Product product={product} key={index} trains={this.props.rootStore.dataStore.trains} jobs={this.props.rootStore.dataStore.jobs} />
        )
      });

    }

    return (
      <div className="container">
        <div className="row">
          <div className="col-3">
            <b>Name</b>
          </div>
          <div className="col-1">
            <b>Build Number</b>
          </div>
          <div className="col-1">
            <b>Version</b>
          </div>
          <div className="col-1">
                <OverlayTrigger
        placement="right-start"
        delay={{ show: 250, hide: 400 }}
        overlay={renderWallClockTimeTooltip}
      >
            <span><b>Wall Clock Time</b><i className="fas fa-info-circle"></i></span>
            </OverlayTrigger>
          </div>
          <div className="col-1">
          <OverlayTrigger
            placement="right-start"
            delay={{ show: 250, hide: 400 }}
            overlay={renderTotalTimeTooltip}
>
            <span><b>Total Time</b><i className="fas fa-info-circle"></i></span>
            </OverlayTrigger>
          </div>
          <div className="col-1">
            <b>Start Time</b>
          </div>
          <div className="col-1">
            <b>End Time</b>
          </div>
          <div className="col-1">
            <b>Errors</b>
          </div>
          <div className="col-1">
            <b>Transients</b>
          </div>
          <div className="col-1">
            <b>Detail</b>
          </div>
        </div>

        {productsTable}

        <Timeline products={this.props.rootStore.dataStore.products}/>
      </div>
    )
  }
}

export default Index
