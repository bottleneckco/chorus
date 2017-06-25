import React from 'react';
import PropTypes from 'prop-types';
import Slider from 'rc-slider';
import 'rc-slider/assets/index.css';

import shareIcon from '../assets/images/share.svg';
import pauseIcon from '../assets/images/pause.svg';
import playIcon from '../assets/images/play.svg';
import skipIcon from '../assets/images/skip.svg';
import '../stylesheets/player.scss';

const Player = (props) => (
  <div className="player">
    <img
      className="player--thumbnail"
      src="http://via.placeholder.com/444x250"
      alt="thumbnail"
    />
    <h2 className="player--song">Hollywood Principle - Firework</h2>
    <div className="player--controls">
      <img className="control control-side" src={shareIcon} alt="share" />
      <img
        className="control control-center"
        src={playIcon}
        alt="play"
        onClick={props.resume}
      />
      <img className="control control-side" src={skipIcon} alt="skip" />
    </div>
    <Slider min={0} max={100} defaultValue={80} />
  </div>
);

Player.propTypes = {
  resume: PropTypes.func.isRequired,
  pause: PropTypes.func.isRequired
};

export default Player;
