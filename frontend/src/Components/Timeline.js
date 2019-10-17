import React, { Component } from 'react'
import { inject, observer } from 'mobx-react'

import {Chart} from 'react-google-charts';


@inject('rootStore')
@observer
class Timeline extends Component {

  constructor(props) {
    super(props);

    this.state = {
      days: 100
    }

    this.changeDays = this.changeDays.bind(this);
  }

  componentWillMount() {
    var _this = this;
    this.props.rootStore.dataStore.fetchProducts(function() {
      console.log(_this.props.rootStore.dataStore.products);
    });
  }

  changeDays(ev) {
    console.log(ev.target.value);
    this.setState({
      days: ev.target.value,
    })
  }

  render () {
    var data = [
      [
        { type: 'string', id: 'Product' },
        { type: 'date', id: 'Start' },
        { type: 'date', id: 'End' },
      ],
    ];

    console.log(this.props.rootStore.dataStore.products);

    this.props.rootStore.dataStore.products.forEach((product, index) => {
      return product.GetPipelines().map((value, index) => {
        if (value.startTime > (new Date()).getTime() - (3600 * 1000 * 24) * this.state.days) {
          data.push([
            `${value.pipelineJob} ${value.version}`, value.startTime, value.endTime
          ]);
        }
        return null;
      });

    });

    console.log(data);

    return (
      <div>
        <div className="row timeline">
          <br />
          <br />
          <br />
          <div className="padded-row">
            <h1>Timeline View for Pipelines</h1>
            <b>Max Shown Days:</b> <input onChange={this.changeDays}  value={this.state.days}></input><i className="text-muted">Change to 1 to zoom in</i>
          </div>
          <Chart
            width={'100%'}
            height={'500px'}
            chartType="Timeline"
            loader={<div>Loading Chart</div>}
            data={data}
            options={{
              showRowNumber: true,
            }}
            rootProps={{ 'data-testid': '1' }}
          />
        </div>
      </div>
    )
  }
}

export default Timeline;
