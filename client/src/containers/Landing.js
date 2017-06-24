import React from 'react';
import { Link } from 'react-router-dom';
import logo from '../assets/images/logo.svg';
import '../stylesheets/landing.scss';

const Landing = () => (
  <div className="landing">
    <img src={logo} alt="logo" />
    <h1>Music sounds better with friends.</h1>
    <Link to="/onboard">Get started</Link>
  </div>
);

export default Landing;
