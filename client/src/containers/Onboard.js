import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { getIsCreating, getResponse } from '../reducers/reducer-channel';
import { createChannel } from '../actions/action-channel';
import Header from '../components/Header';
import '../stylesheets/onboard.scss';

class Onboard extends Component {
  constructor(props) {
    super(props);

    this.state = {
      title: '',
      nickname: ''
    };

    this.createChannel = this.createChannel.bind(this);
    this.handleInput = this.handleInput.bind(this);
  }

  componentWillReceiveProps(nextProps) {
    if (this.props.isCreating && !nextProps.isCreating) {
      this.props.history.push(nextProps.response.access_code);
    }
  }

  createChannel(e) {
    e.preventDefault();

    this.props.createChannel({
      name: this.state.title,
      created_by: this.state.nickname
    });
  }

  handleInput(e) {
    this.setState({
      [e.target.name]: e.target.value
    });
  }

  render() {
    const { title, nickname } = this.state;

    return (
      <div className="onboard">
        <Header large={false} />
        <div className="onboard--content">
          <h1>Let&apos;s get this party started</h1>

          <form
            autoComplete="off"
            className="onboard--form"
            onSubmit={this.createChannel}
          >
            <div className="fields">
              <h3>Give your channel a title</h3>
              <input
                type="text"
                name="title"
                value={title}
                className="textbox-line"
                onChange={this.handleInput}
              />
              <h3>Set your nickname</h3>
              <input
                type="text"
                name="nickname"
                value={nickname}
                className="textbox-line"
                onChange={this.handleInput}
              />
            </div>
            <h3>
              <input
                type="submit"
                value="Create"
                className="button-solid"
              />
            </h3>
          </form>
        </div>
      </div>
    );
  }
}

Onboard.defaultProps = {
  isCreating: false,
  response: {}
};

Onboard.propTypes = {
  createChannel: PropTypes.func.isRequired,
  history: PropTypes.object.isRequired,
  isCreating: PropTypes.bool,
  response: PropTypes.object
};

const mapStateToProps = (state) => ({
  isCreating: getIsCreating(state),
  response: getResponse(state)
});

const mapDispatchToProps = (dispatch) => (
  bindActionCreators({
    createChannel
  }, dispatch)
);

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Onboard));
