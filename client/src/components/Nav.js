import React from 'react';
import PropTypes from 'prop-types';
import logo from '../assets/images/logo.svg';
import '../stylesheets/nav.scss';

const Nav = ({ title }) => (
  <nav className="nav">
    <img
      className="nav--logo"
      src={logo}
      alt="chorus"
    />
    <h3 className="nav--text">chorus</h3>
    <h3 className="nav--text nav--text-right">{title}</h3>
  </nav>
);

Nav.propTypes = {
  title: PropTypes.string.isRequired
};

export default Nav;
