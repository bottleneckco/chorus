import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { getChannelIsCreating, getChannelData } from '../reducers/reducer-channel';
import { createChannel } from '../actions/action-channel';
import Header from '../components/Header';
import Form from '../components/Form';
import '../stylesheets/create.scss';

class Create extends Component {
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
      this.props.history.push(nextProps.data.id);
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

    const fields = [{
      label: 'Give your channel a title',
      name: 'title',
      value: title
    }, {
      label: 'Set your nickname',
      name: 'nickname',
      value: nickname
    }];

    return (
      <div className="create">
        <Header large={false} />
        <div className="create--content">
          <h1 className="create--heading">Let&apos;s get this party started</h1>

          <Form
            fields={fields}
            handleInput={this.handleInput}
            submitForm={this.createChannel}
            submitButtonText="Create"
          />
        </div>
      </div>
    );
  }
}

Create.defaultProps = {
  data: {}
};

Create.propTypes = {
  history: PropTypes.object.isRequired,
  createChannel: PropTypes.func.isRequired,
  isCreating: PropTypes.bool.isRequired,
  data: PropTypes.object
};

const mapStateToProps = (state) => ({
  isCreating: getChannelIsCreating(state),
  data: getChannelData(state)
});

const mapDispatchToProps = (dispatch) => (
  bindActionCreators({
    createChannel
  }, dispatch)
);

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Create));
