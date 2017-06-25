import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import { getResponse } from '../reducers/reducer-channel';
import Nav from '../components/Nav';
import Player from '../components/Player';
import '../stylesheets/channel.scss';

class Channel extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="channel">
        <Nav title={this.props.response.name} />
        <Player />
      </div>
    );
  }
}

Channel.defaultProps = {
  response: {}
};

Channel.propTypes = {
  response: PropTypes.object
};

const mapStateToProps = (state) => ({
  response: getResponse(state)
});

export default connect(mapStateToProps)(Channel);
