import React from "react";
import { API_URL } from "../constants/constants";

class ExcelForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = { submitted: false, response: {} };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }
  handleChange(evt) {
    this.setState({ [evt.target.name]: evt.target.value });
  }
  async handleSubmit(event) {
    event.preventDefault();
    let { s, c, r } = this.state;
    let response = {};
    try {
      response = await (
        await fetch(
          `${API_URL}/api/v1/excel?s=${s.toUpperCase()}&c=${c}&r=${r}`
        )
      ).json();
      this.setState({ submitted: true, s, c, r, response });
    } catch (error) {
      console.error("Error while submiting form", error);
    }
  }
  render() {
    const { submitted, c, r, response = {} } = this.state;
    let tableBody = "";
    if (Object.keys(response).length > 0) {
      for (let i = 0; i < r; i++) {
        tableBody += "<tr>";
        for (let j = 0; j < c; j++) {
          tableBody += "<td>";
          tableBody += `${response[i + j]}`;
          tableBody += "</td>";
        }
        tableBody += "</tr>";
      }
    }
    return (
      <div>
        <form onSubmit={this.handleSubmit} onChange={this.handleChange}>
          <label>
            Column To Start:
            <input type="text" name="s" />
          </label>
          <label>
            No of Columns:
            <input type="text" name="c" />
          </label>
          <label>
            No of Rows:
            <input type="text" name="r" />
          </label>
          <input type="submit" value="Submit" />
        </form>
        {submitted ? (
          <table>
            {Object.keys(tableBody).length > 0 ? (
              <tbody
                className="Container"
                dangerouslySetInnerHTML={{ __html: tableBody }}
              ></tbody>
            ) : (
              <tbody></tbody>
            )}
          </table>
        ) : (
          <div>Enter Url</div>
        )}
      </div>
    );
  }
}

export default ExcelForm;
