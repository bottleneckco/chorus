import React from 'react';
import '../stylesheets/search.scss';

const SearchItem = (props) => (
  <div className="searchitem">
    <div className="searchitem-content">
      <div className="searchitem-thumbnail" />
      <div className="searchitem-details">
        <div className="searchitem-name">
          Fuk u up
        </div>
        <div className="searchitem-duration">
          3.15
        </div>
      </div>
    </div>
  </div>
)

export default SearchItem;
