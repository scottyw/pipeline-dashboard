import React, { Component } from 'react'
import { observer } from 'mobx-react'

import { Button } from '@puppet/react-components';

import PipelineTrain from './PipelineTrain';
import Moment from 'react-moment';

@observer
class Product extends Component {

  constructor(props) {
    super(props)

    this.state = {
      showTrains: false
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

  toggleTrain() {
    this.setState({
      showTrains: !this.state.showTrains
    })
  }

  pipelineTrains() {
    var trains = []
    if (this.state.showTrains) {
      this.props.pipeline.GetTrains().map((value, index) => {
        return trains.push(<PipelineTrain key={index} train={value} />)
      });
    }

    return trains
  }

  render () {
    return (
      <div className="row text-left">
        <div className="col-12">
          <div className="row pipeline-row">
            <div className="col-3">
              <a href={this.props.pipeline.url} target="_blank" rel="noopener noreferrer">{this.props.pipeline.pipeline}</a>

            </div>
            <div className="col-1 text-left">
              {this.props.pipeline.buildNumber}
            </div>
            <div className="col-1 text-left">
              {this.props.pipeline.version}
            </div>
            <div className="col-1">
              {this.props.pipeline.wallClockFormatted()}
            </div>
            <div className="col-1">
              {this.props.pipeline.totalFormatted()}
            </div>
            <div className="col-1">
              <Moment format="YYYY/MM/DD HH:mm">{this.props.pipeline.startTime}</Moment>
            </div>
            <div className="col-1">
              <Moment format="MM/DD HH:mm">{this.props.pipeline.endTime}</Moment>
            </div>
            <div className="col-1">
              {this.props.pipeline.errors}
            </div>
            <div className="col-1">
              {this.props.pipeline.transients}
            </div>
            <div className="col-1">
              <Button href="#" onClick={() => this.toggleTrain()}>{this.state.showTrains ? 'Close' : 'Jobs'}</Button>
            </div>
          </div>
        <div className="surround-trains"  >
        {this.pipelineTrains()}
        </div>
      </div>
    </div>
    )
  }
}

export default Product
