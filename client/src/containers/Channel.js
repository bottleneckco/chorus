import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { getRehydrationStatus } from '../reducers/reducer-persistence';
import { getChannelData, getChannelIsFetching } from '../reducers/reducer-channel';
import { getQueue, getQueueIsFetching } from '../reducers/reducer-queue';
import { fetchChannel, addUserToChannel } from '../actions/action-channel';
import { fetchQueue, skipQueueItem } from '../actions/action-queue';

import WSPlayer from './WSPlayer';
import Onboard from './Onboard';
import '../stylesheets/channel.scss';

class Channel extends Component {
  constructor(props) {
    super(props);

    this.state = {
      rehydrated: false,
      playing: true,
      shouldOnboard: false
    };

    this.fetchQueueWrapper = this.fetchQueueWrapper.bind(this)
    this.submitForm = this.submitForm.bind(this);
  }

  componentWillMount() {
    if (this.props.rehydrated) {
      this.setState({ rehydrated: true });
    }
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.rehydrated && !this.state.rehydrated) {
      this.setState({ rehydrated: true });

      const { data, match } = this.props;

      if (Object.prototype.hasOwnProperty.call(data, 'id')) {
        this.props.fetchQueue(data.id);
      } else {
        this.props.fetchChannel(match.params.hash);
        this.setState({ shouldOnboard: true });
      }
    }
  }

  submitForm(e, data) {
    e.preventDefault();

    this.props.addUserToChannel(this.props.data.id, data);
    this.setState({ shouldOnboard: false });
  }

  fetchQueueWrapper() {
    this.props.fetchQueue(this.props.data.id);
  }

  render() {
    const { data, channelIsFetching, queue, queueIsFetching } = this.props;

    if (channelIsFetching || queueIsFetching
      || Object.keys(data).length === 0) {
      return <div>Loading...</div>;
    }

    if (this.state.shouldOnboard) {
      return <Onboard submitForm={this.submitForm} />;
    }

    return (
      <WSPlayer
        data={data}
        queue={queue}
        fetchQueue={this.fetchQueueWrapper}
        skipQueueItem={this.props.skipQueueItem}
      />
    );
  }
}

Channel.defaultProps = {
  data: {}
};

Channel.propTypes = {
  match: PropTypes.object.isRequired,
  rehydrated: PropTypes.bool.isRequired,

  data: PropTypes.object,
  channelIsFetching: PropTypes.bool.isRequired,
  queue: PropTypes.array.isRequired,
  queueIsFetching: PropTypes.bool.isRequired,

  addUserToChannel: PropTypes.func.isRequired,
  skipQueueItem: PropTypes.func.isRequired,
  fetchChannel: PropTypes.func.isRequired,
  fetchQueue: PropTypes.func.isRequired
};

const mapStateToProps = (state) => ({
  rehydrated: getRehydrationStatus(state),
  data: getChannelData(state),
  channelIsFetching: getChannelIsFetching(state),
  queue: getQueue(state),
  queueIsFetching: getQueueIsFetching(state)
});

const mapDispatchToProps = (dispatch) => (
  bindActionCreators({
    addUserToChannel,
    skipQueueItem,
    fetchChannel,
    fetchQueue
  }, dispatch)
);

export default connect(mapStateToProps, mapDispatchToProps)(Channel);
