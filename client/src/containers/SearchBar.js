import React, { Component } from 'react';
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
    URL: 'xxx',
  }
]

class SearchBar extends Component {
  constructor(props) {
    super(props);
    
    this.defaultQuery = 'Search for a song...';
    this.state = {
      query: this.defaultQuery,
      searchItems: []
    };

    this.onTextboxChange = this.onTextboxChange.bind(this);
    this.onTextboxClick = this.onTextboxClick.bind(this);
    this.renderItems = this.renderItems.bind(this);
  }

  onTextboxChange(e) {
    this.setState({ value: e.target.value });
  }

  onTextboxClick() {
    if (this.state.query === this.defaultQuery) {
      this.setState({ query: '' });
    }
  }

  renderItems() {
    return fakeData.map((item) => (
      <SearchItem item={item} />
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
          onClick={this.onTextboxClick}
        />
        <div>
          {this.renderItems()}
        </div>
      </div>
    );
  }
}

export default SearchBar;
