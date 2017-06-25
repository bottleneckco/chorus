import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { addToQueue } from '../actions/action-queue';
import { searchMusic, clearSearchResults } from '../actions/action-search';
import { getChannelId } from '../reducers/reducer-channel';
import { getSearchIsFetching, getResults } from '../reducers/reducer-search';

import SearchItem from '../components/SearchItem';
import '../stylesheets/search.scss';

class SearchBar extends Component {
  constructor(props) {
    super(props);

    this.state = {
      query: '',
      searchItems: []
    };

    this.onTextboxChange = this.onTextboxChange.bind(this);
    this.onTextboxKeydown = this.onTextboxKeydown.bind(this);
    this.renderItems = this.renderItems.bind(this);
    this.addMusicToQueue = this.addMusicToQueue.bind(this);
    this.handleClickOutside = this.handleClickOutside.bind(this);
    this.destroy = this.destroy.bind(this);
  }

  componentWillMount() {
    this.props.clearSearchResults();
  }

  componentDidMount() {
    window.addEventListener('click', this.handleClickOutside, true);
  }

  componentWillReceiveProps() {
    const main = this;
    document.getElementById('overlay').addEventListener('click', function(e) {
      e = window.event || e;
      if (this === e.target) {
        main.destroy();
      }
    });
  }

  componentWillUnmount() {
    window.removeEventListener('click', this.handleClickOutside, true);
  }

  onTextboxChange(e) {
    this.setState({ query: e.target.value });
  }

  onTextboxKeydown(e) {
    // Press ENTER
    if (e.keyCode === 13) {
      this.props.searchMusic(this.props.channelId, this.state.query);
    }
  }

  handleClickOutside(event) {
    const domNode = ReactDOM.findDOMNode(this);

    if ((!domNode || !domNode.contains(event.target))) {
      this.destroy();
    }
  }

  destroy() {
    this.props.clearSearchResults();
    this.props.hideSearchBar();
  }

  addMusicToQueue(urlStr) {
    const data = { url: urlStr };
    this.props.addToQueue(this.props.channelId, data);
    this.destroy();
  }

  renderItems() {
    const { results } = this.props;

    return results.map((item) => (
      <SearchItem
        item={item}
        key={item.url}
        addMusic={this.addMusicToQueue}
      />
    ));
  }

  render() {
    const { query } = this.state;

    return (
      <div id="overlay" className="overlay">
        <div className="search">
          <div className="search--content">
            <input
              type="text"
              value={query}
              className="search--textbox"
              placeholder="Search for a song..."
              onChange={this.onTextboxChange}
              onKeyDown={this.onTextboxKeydown}
              disabled={this.props.searchIsFetching}
            />
            {this.renderItems()}
          </div>
        </div>
      </div>
    );
  }
}

SearchBar.defaultProps = {
  results: [],
  searchIsFetching: false,
  channelId: ''
};

SearchBar.propTypes = {
  results: PropTypes.array,
  searchIsFetching: PropTypes.bool,
  searchMusic: PropTypes.func.isRequired,
  channelId: PropTypes.string.isRequired,
  clearSearchResults: PropTypes.func.isRequired,
  hideSearchBar: PropTypes.func.isRequired,
  addToQueue: PropTypes.func.isRequired
};

const mapStateToProps = (state) => ({
  results: getResults(state),
  searchIsFetching: getSearchIsFetching(state),
  channelId: getChannelId(state)
});

const mapDispatchToProps = (dispatch) => (
  bindActionCreators({
    searchMusic,
    clearSearchResults,
    addToQueue
  }, dispatch)
);

export default connect(mapStateToProps, mapDispatchToProps)(SearchBar);
