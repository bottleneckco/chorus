import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { searchMusic, clearSearchResults } from '../actions/action-search';
import { getChannelId } from '../reducers/reducer-channel';
import { getSearchIsFetching, getResults } from '../reducers/reducer-search';

import SearchItem from '../components/SearchItem';
import '../stylesheets/search.scss';

class SearchBar extends Component {
  constructor(props) {
    super(props);

    this.defaultQuery = 'Search for a song...';
    this.state = {
      query: this.defaultQuery,
      searchItems: []
    };

    this.onTextboxChange = this.onTextboxChange.bind(this);
    this.onTextboxFocus = this.onTextboxFocus.bind(this);
    this.onTextboxBlur = this.onTextboxBlur.bind(this);
    this.onTextboxKeydown = this.onTextboxKeydown.bind(this);
    this.renderItems = this.renderItems.bind(this);
    this.addMusicToQueue = this.addMusicToQueue.bind(this);
    this.handleClickOutside = this.handleClickOutside.bind(this);
    this.destroy = this.destroy.bind(this);

    this.handleClickOutside = this.handleClickOutside.bind(this);
  }

  componentWillMount() {
    this.props.clearSearchResults();
  }

  componentDidMount() {
    window.addEventListener('click', this.handleClickOutside, true);
  }

  componentWillReceiveProps(nextProps) {
    console.log(nextProps);
  }

  componentWillUnmount() {
    window.removeEventListener('click', this.handleClickOutside, true);
  }

  onTextboxChange(e) {
    this.setState({ query: e.target.value });
  }

  onTextboxFocus() {
    if (this.state.query === this.defaultQuery) {
      this.setState({ query: '' });
    }
  }

  onTextboxBlur() {
    if (this.state.query === '') {
      this.setState({ query: this.defaultQuery });
    }

    this.props.clearSearchResults();
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
      // window.removeEventListener('click', this.handleClickOutside.bind(this), true);
      console.log('hi');
      this.destroy();
    }
  }

  destroy() {
    this.props.clearSearchResults();
    this.props.hideSearchBar();
  }

  addMusicToQueue(url) {
    console.log(url)
  }
  
  renderItems() {
    const { results } = this.props;

    return results.map((item) => (
      <SearchItem item={item} key={item.url} addMusic={this.addMusicToQueue} />
    ));
  }

  render() {
    const { query } = this.state;

    return (
      <div className="searchbar">
        <input
          type="text"
          className="searchbar-textbox"
          value={query}
          onChange={this.onTextboxChange}
          onFocus={this.onTextboxFocus}
          onBlur={this.onTextboxBlur}
          onKeyDown={this.onTextboxKeydown}
          disabled={this.props.searchIsFetching}
        />
        <div>
          {this.renderItems()}
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
  hideSearchBar: PropTypes.func.isRequired
};

const mapStateToProps = (state) => ({
  results: getResults(state),
  searchIsFetching: getSearchIsFetching(state),
  channelId: getChannelId(state)
});

const mapDispatchToProps = (dispatch) => (
  bindActionCreators({
    searchMusic,
    clearSearchResults
  }, dispatch)
);

export default connect(mapStateToProps, mapDispatchToProps)(SearchBar);
