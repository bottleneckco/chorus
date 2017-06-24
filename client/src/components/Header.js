import React from 'react';
import PropTypes from 'prop-types';
import logo from '../assets/images/logo.svg';
import '../stylesheets/header.scss';

const Header = ({ large }) => {
  const renderTitle = () => (
    large ? <h2>chorus</h2> : <h3>chorus</h3>
  );

  return (
    <div className="header">
      <img
        className={`header--logo-${large ? 'large' : 'small'}`}
        src={logo}
        alt="chorus"
      />
      {renderTitle()}
    </div>
  );
};

Header.defaultProps = {
  large: false
};

Header.propTypes = {
  large: PropTypes.bool
};

export default Header;
