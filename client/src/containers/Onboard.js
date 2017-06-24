import React, { Component, PropTypes } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { createChannel } from '../actions/action-channel';

class Onboard extends Component {
  constructor(props) {
    super(props);

    this.state = {};

    this.submitNewChannel = this.submitNewChannel.bind(this);
  }

  submitNewChannel() {
    this.props.createChannel({});
  }

  render() {
    return (
      <div>Onboard
        <button onClick={this.submitNewChannel}>Test here</button>
      </div>
    );
  }
}

const mapDispatchToProps = (dispatch) => (
  bindActionCreators({
    createChannel
  }, dispatch)
);

Onboard.propTypes = {
  createChannel: PropTypes.func.isRequired
};

export default connect(null, mapDispatchToProps)(Onboard);
