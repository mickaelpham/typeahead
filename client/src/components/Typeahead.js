import React from "react";
import axios from "axios";

import "./Typeahead.css";

class Typeahead extends React.Component {
  state = { suggestions: [], q: "" };

  handleChange = event => {
    let q = event.target.value;
    this.fetchSuggestions(q);
    this.setState({ q });
  };

  fetchSuggestions(q) {
    axios.get("/search", { params: { q: q } }).then(res => {
      this.setState({ suggestions: res.data });
    });
  }

  renderSuggestions() {
    return this.state.suggestions.map(s => {
      return <li key={s}>{s}</li>;
    });
  }

  render() {
    return (
      <div className="Typeahead">
        <input type="text" value={this.state.q} onChange={this.handleChange} />
        <ul>{this.renderSuggestions()}</ul>
      </div>
    );
  }
}

export default Typeahead;
