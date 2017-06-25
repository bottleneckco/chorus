import React from 'react';
import PropTypes from 'prop-types';
import Slider from 'rc-slider';
import 'rc-slider/assets/index.css';

import shareIcon from '../assets/images/share.svg';
import pauseIcon from '../assets/images/pause.svg';
import playIcon from '../assets/images/play.svg';
import skipIcon from '../assets/images/skip.svg';
import '../stylesheets/player.scss';

const Player = ({ resume, pause, skip, playing, currentSong, adjustAudioVol }) => {
  const renderPlayPause = () => {
    const icon = playing ? pauseIcon : playIcon;
    const onClick = playing ? pause : resume;

    return (
      <img
        className="control control-center"
        src={icon}
        alt="play"
        role="button"
        tabIndex="0"
        onClick={onClick}
      />
    );
  };

  if (Object.keys(currentSong).length === 0) {
    return (
      <h1 className="player--empty">How about adding some songs?</h1>
    );
  }

  return (
    <div className="player">
      <img
        className="player--thumbnail"
        src={currentSong.video.thumbnail}
        alt="thumbnail"
      />
      <h2 className="player--song">{currentSong.video.name}</h2>
      <div className="player--controls">
        <img className="control control-side" src={shareIcon} alt="share" />
        {renderPlayPause()}
        <img
          className="control control-side"
          src={skipIcon}
          alt="skip"
          onClick={skip}
        />
      </div>
      <Slider min={0} max={100} defaultValue={80} onChange={adjustAudioVol} />
    </div>
  );
};

Player.propTypes = {
  resume: PropTypes.func.isRequired,
  pause: PropTypes.func.isRequired,
  playing: PropTypes.bool.isRequired,
  currentSong: PropTypes.object.isRequired,
  adjustAudioVol: PropTypes.func.isRequired
};

export default Player;
