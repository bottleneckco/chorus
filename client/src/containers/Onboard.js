import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Header from '../components/Header';
import Form from '../components/Form';

class Onboard extends Component {
  constructor(props) {
    super(props);

    this.state = { nickname: '' };

    this.handleInput = this.handleInput.bind(this);
  }

  handleInput(e) {
    this.setState({
      [e.target.name]: e.target.value
    });
  }

  render() {
    const fields = [{
      label: 'Set your nickname',
      name: 'nickname',
      value: this.state.nickname
    }];

    return (
      <div className="onboard">
        <Header large={false} />
        <div className="onboard--content">
          <h1 className="onboard--heading">Join the fun!</h1>
          <Form
            fields={fields}
            handleInput={this.handleInput}
            submitForm={(e) => this.props.submitForm(e, this.state)}
            submitButtonText="Join"
          />
        </div>
      </div>
    );
  }
}

Onboard.propTypes = {
  submitForm: PropTypes.func.isRequired
};

export default Onboard;
