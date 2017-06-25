import React from 'react';
import PropTypes from 'prop-types';
import crossIcon from '../assets/images/x.svg';
import '../stylesheets/queue.scss';

const Queue = ({ queue, skipQueueItem }) => {
  const renderSongsInQueue = () => (
    queue.map((item, index) => (
      <div key={item.id} className="item">
        <span className="col--song">{item.video.name}</span>
        <span className="col--user">{item.user.nickname}</span>
        <img
          className="col--skip"
          src={crossIcon}
          alt="skip"
          role="button"
          tabIndex="0"
          onClick={() => skipQueueItem(index)}
        />
      </div>
    ))
  );

  return (
    <div className="queue">
      <input
        type="button"
        value="Add song"
        className="button-solid button-orange-10 queue--add"
      />
      <div className="queue--list">
        <div className="queue--headings">
          <span className="col--song">Next</span>
          <span className="col--user">Added by</span>
        </div>
        {renderSongsInQueue()}
      </div>
    </div>
  );
};

Queue.propTypes = {
  queue: PropTypes.array.isRequired,
  skipQueueItem: PropTypes.func.isRequired
};

export default Queue;
