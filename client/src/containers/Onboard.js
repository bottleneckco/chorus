import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

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

  createChannel(e) {
    e.preventDefault();
    console.log(this.state.title, this.state.nickname);
    this.props.createChannel({});
    // create new channel and redirect to channel page
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

          <form className="onboard--form" onSubmit={this.createChannel}>
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

const mapDispatchToProps = (dispatch) => (
  bindActionCreators({
    createChannel
  }, dispatch)
);

Onboard.propTypes = {
  createChannel: PropTypes.func.isRequired
};

export default connect(null, mapDispatchToProps)(Onboard);
