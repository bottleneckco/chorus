import React from 'react';
import { Link } from 'react-router-dom';
import logo from '../assets/images/logo.svg';
import '../stylesheets/landing.scss';

const Landing = () => (
  <div className="landing">
    <img className="landing--logo" src={logo} alt="logo" />
    <h2 className="landing--title">chorus</h2>
    <div className="landing--content">
      <h1 className="heading">Music sounds better with friends.</h1>
      <h3>
        <Link className="button-outline" to="/onboard">Get started</Link>
      </h3>
    </div>
    <span className="landing--company">H & Co.</span>
  </div>
);

export default Landing;
