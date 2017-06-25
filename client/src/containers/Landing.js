import React from 'react';
import { Link } from 'react-router-dom';
import Header from '../components/Header';
import '../stylesheets/landing.scss';

const Landing = () => (
  <div className="landing">
    <Header large />
    <div className="landing--content">
      <h1 className="heading">Music sounds better with friends.</h1>
      <h3>
        <Link className="button-outline" to="/create">Get started</Link>
      </h3>
    </div>
    <span className="landing--company">H & Co.</span>
  </div>
);

export default Landing;
