import React from "react";
import { API_URL } from "../constants/constants";
import "./WordCountForm.css";

class WordCountForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = { value: "", submitted: false, wordJSON: {} };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }
  handleChange(evt) {
    this.setState({ [evt.target.name]: evt.target.value });
  }
  async handleSubmit(event) {
    event.preventDefault();
    let { value } = this.state;
    let response = {};
    try {
      response = await (
        await fetch(`${API_URL}/api/v1/wordcounter?url=${value}`)
      ).json();
      this.setState({ submitted: true, wordJSON: response });
    } catch (error) {
      console.error("Error while submiting form", error);
    }
  }
  render() {
    const { submitted, wordJSON } = this.state;
    const tableHeadings = ["Word", "Count"].map((v) => <th key={v}> {v} </th>);
    const tableValues = Object.keys(wordJSON).map((k) => {
      return (
        <tr key={k}>
          <td> {k} </td>
          <td> {wordJSON[k]} </td>
        </tr>
      );
    });
    return (
      <div>
        <form onSubmit={this.handleSubmit} onChange={this.handleChange}>
          <label>
            Url:
            <input type="text" name="url" />
          </label>
          <input type="submit" value="Submit" />
        </form>
        {submitted ? (
          <table>
            <thead>
              <tr>{tableHeadings}</tr>
            </thead>
            <tbody>{tableValues}</tbody>
          </table>
        ) : (
          <div>Enter Url</div>
        )}
      </div>
    );
  }
}

export default WordCountForm;
