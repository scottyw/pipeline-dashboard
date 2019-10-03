import React, { Component } from 'react'
import { observer } from 'mobx-react'

import Moment from 'react-moment';

@observer
class PipelineTrain extends Component {

  hms(nanoseconds) {
    var startSeconds = nanoseconds / 1000000000
    var hours = Math.floor(startSeconds / 3600);
    startSeconds = startSeconds - hours * 3600;

    var minutes = Math.floor(startSeconds / 60);
    // var seconds = startSeconds - minutes * 60;

    return `${hours}H, ${minutes}M`
  }

  openPipeline() {

  }

  render () {
    return (
      <div className="row train-row text-left">
        <div className="col-12">
          <a href={this.props.train.url} target="_blank"><b>{this.props.train.name}</b></a>
        </div>
        <div className="col-5">
        </div>
        <div className="col-2">
          {this.props.train.durationFormatted()}
        </div>
        <div className="col-2">
          <Moment format="YYYY/MM/DD HH:mm">{this.props.train.startTime}</Moment>
        </div>
        <div className="col-2">
          <Moment format="YYYY/MM/DD HH:mm">{this.props.train.endTime}</Moment>
        </div>

      </div>
    )
  }
}

export default PipelineTrain
