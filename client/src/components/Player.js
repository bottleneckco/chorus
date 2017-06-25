import React from 'react';
import shareIcon from '../assets/images/share.svg';
import pauseIcon from '../assets/images/pause.svg';
import playIcon from '../assets/images/play.svg';
import skipIcon from '../assets/images/skip.svg';
import '../stylesheets/player.scss';

const Player = () => (
  <div className="player">
    <img
      className="player--thumbnail"
      src="http://via.placeholder.com/444x250"
      alt="thumbnail"
    />
    <h2 className="player--song">Hollywood Principle - Firework</h2>
    <div className="player--controls">
      <img className="control control-side" src={shareIcon} alt="share" />
      <img className="control control-center" src={playIcon} alt="play" />
      <img className="control control-side" src={skipIcon} alt="skip" />
    </div>
  </div>
);

export default Player;
