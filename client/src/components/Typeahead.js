import React from "react";
import axios from "axios";

import "./Typeahead.css";

class Typeahead extends React.Component {
  state = { suggestions: [] };

  componentDidMount() {
    axios.get("/search").then(res => {
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
        <input type="text" name="q" />
        <ul>{this.renderSuggestions()}</ul>
      </div>
    );
  }
}

export default Typeahead;
