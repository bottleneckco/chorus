import React from 'react';
import PropTypes from 'prop-types';
import '../stylesheets/square.scss';

const Square = ({ text }) => (
  <div className="square">{text}</div>
);

Square.propTypes = {
  text: PropTypes.string.isRequired
};

export default Square;
