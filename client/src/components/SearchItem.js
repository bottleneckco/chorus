import React from 'react';
import PropTypes from 'prop-types';
import '../stylesheets/search.scss';

const convertToMinSec = (durationInSec) => {
  const minutes = Math.floor(durationInSec / 60);
  const seconds = durationInSec - (minutes * 60);

  return `${minutes}:${seconds}`;
};

const SearchItem = ({ item, addMusic }) => (
  <div className="searchitem" onClick={() => addMusic(item.url)}>
    <div className="searchitem-content">
      <img src={item.thumbnail} className="searchitem-thumbnail" alt={item.thumbnail} />
      <div className="searchitem-details">
        <div className="searchitem-name">
          {item.name}
        </div>
        <div className="searchitem-duration">
          {convertToMinSec(item.duration)}
        </div>
      </div>
    </div>
  </div>
);

SearchItem.defaultProps = {
  item: {
    name: '',
    thumbnail: '',
    url: '',
    duration: 0
  }
};

SearchItem.propTypes = {
  item: PropTypes.shape({
    name: PropTypes.string.isRequired,
    thumbnail: PropTypes.string.isRequired,
    url: PropTypes.string.isRequired,
    duration: PropTypes.number.isRequired
  })
};

export default SearchItem;
