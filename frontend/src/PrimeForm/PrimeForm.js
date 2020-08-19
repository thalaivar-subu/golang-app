import React from "react";
import { API_URL } from "../constants/constants";

class PrimeForm extends React.Component {
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
    let { input } = this.state;
    let response = {};
    try {
      response = await (
        await fetch(`${API_URL}/api/v1/primenumber?input=${input}`)
      ).json();
      console.log(response);
      this.setState({ submitted: true, response });
    } catch (error) {
      console.error("Error while submiting form", error);
    }
  }
  render() {
    const { submitted, response } = this.state;
    const heading = "Prime Number";
    return (
      <div>
        <form onSubmit={this.handleSubmit} onChange={this.handleChange}>
          <label>
            Enter Input:
            <input type="number" name="input" />
          </label>
          <input type="submit" value="Submit" />
        </form>
        {submitted ? (
          <table>
            <thead>
              <tr>{heading}</tr>
            </thead>
            <tbody>
              <tr>
                {response && Object.keys(response).length > 0 ? (
                  response.map((x) => <td key={x}>{x}</td>)
                ) : (
                  <td></td>
                )}
              </tr>
            </tbody>
          </table>
        ) : (
          <div></div>
        )}
      </div>
    );
  }
}

export default PrimeForm;
