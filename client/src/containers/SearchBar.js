import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { searchMusic } from '../actions/action-search';
import { getChannelId } from '../reducers/reducer-channel';
import { getSearchIsFetching, getResults } from '../reducers/reducer-search';

import SearchItem from '../components/SearchItem';
import '../stylesheets/search.scss';

const fakeData = [
  {
    Name: "Mome - Playground",
    ThumbnailURL: "sadsqwewq",
    Duration: 1000,
    URL: 'xxx',
  }, {
    Name: "XXX",
    ThumbnailURL: "sadsqwewq",
    Duration: 1000,
    URL: 'zxcs',
  }
]

const samplejson = `{"status":"ok","count":10,"results":[{"name":"Møme - Take Off (Coca-Cola Summer 2017) ft. Romuald","thumbnail":"https://i.ytimg.com/vi/TxLmO0h9GO0/maxresdefault.jpg","duration":160,"url":"https://www.youtube.com/watch?v=TxLmO0h9GO0"},{"name":"Møme - Playground","thumbnail":"https://i.ytimg.com/vi/CehAKQL463M/maxresdefault.jpg","duration":272,"url":"https://www.youtube.com/watch?v=CehAKQL463M"},{"name":"Møme - Aloha (Official Music Video) ft. Merryn Jeann","thumbnail":"https://i.ytimg.com/vi/1YRW1QRKTBc/maxresdefault.jpg","duration":223,"url":"https://www.youtube.com/watch?v=1YRW1QRKTBc"},{"name":"Møme - On the Top session [EIFFEL TOWER] Live 2016 @Cercle","thumbnail":"https://i.ytimg.com/vi/bGoVFZck_NY/maxresdefault.jpg","duration":4339,"url":"https://www.youtube.com/watch?v=bGoVFZck_NY"},{"name":"Møme Vs Midnight To Monaco - Alive (Official Video)","thumbnail":"https://i.ytimg.com/vi/oBGKpUZzp6A/maxresdefault.jpg","duration":269,"url":"https://www.youtube.com/watch?v=oBGKpUZzp6A"},{"name":"Møme - Why Is It (Official Video) ft. Merryn Jeann","thumbnail":"https://i.ytimg.com/vi/rIrAWDG_vBs/maxresdefault.jpg","duration":190,"url":"https://www.youtube.com/watch?v=rIrAWDG_vBs"},{"name":"Møme - Hold On (Official Music Video) ft. Dylan Wright","thumbnail":"https://i.ytimg.com/vi/witzw1Zijmc/maxresdefault.jpg","duration":172,"url":"https://www.youtube.com/watch?v=witzw1Zijmc"},{"name":"Møme - On The Top Session [SYDNEY]","thumbnail":"https://i.ytimg.com/vi/dJUOsm6jrGk/maxresdefault.jpg","duration":2715,"url":"https://www.youtube.com/watch?v=dJUOsm6jrGk"},{"name":"Møme - Delta","thumbnail":"https://i.ytimg.com/vi/de9ehpyjT6w/maxresdefault.jpg","duration":229,"url":"https://www.youtube.com/watch?v=de9ehpyjT6w"},{"name":"Møme - Cyclope (Official Music Video)","thumbnail":"https://i.ytimg.com/vi/xmTz7BdsJp4/maxresdefault.jpg","duration":266,"url":"https://www.youtube.com/watch?v=xmTz7BdsJp4"}]}`

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
  }

  componentWillReceiveProps(nextProps) {
    console.log(nextProps);
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
  }

  onTextboxKeydown(e) {
    // Press ENTER
    if (e.keyCode === 13) {
      this.props.searchMusic(this.props.channelId, this.state.query);
    }
  }

  addMusicToQueue(url) {
    console.log(url)
  }
  
  renderItems() {
    const { results } = this.props;

    console.log(results);
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
  channelId: PropTypes.string.isRequired
};

const mapStateToProps = (state) => ({
  results: getResults(state),
  searchIsFetching: getSearchIsFetching(state),
  channelId: getChannelId(state)
});

const mapDispatchToProps = (dispatch) => (
  bindActionCreators({
    searchMusic
  }, dispatch)
);

export default connect(mapStateToProps, mapDispatchToProps)(SearchBar);
