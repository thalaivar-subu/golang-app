import React from "react";
import { API_URL } from "../constants/constants";

class CrudFrom extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      get: false,
      post: false,
      put: false,
      deletee: false,
      getResponse: {},
      putResponse: {},
      postResponse: {},
      deleteResponse: {},
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleGETSubmit = this.handleGETSubmit.bind(this);
    this.handlePOSTSubmit = this.handlePOSTSubmit.bind(this);
    this.handlePUTSubmit = this.handlePUTSubmit.bind(this);
    this.handleDELETESubmit = this.handleDELETESubmit.bind(this);
  }
  handleChange(evt) {
    this.setState({ [evt.target.name]: evt.target.value });
  }

  async handleGETSubmit(event) {
    event.preventDefault();
    console.log(this.state);
    const { email } = this.state;
    try {
      let response = await (
        await fetch(`${API_URL}/api/v1/crud?email=${email}`)
      ).json();
      this.setState({ get: true, getResponse: response });
    } catch (error) {
      console.error(error);
    }
  }
  async handlePOSTSubmit(event) {
    event.preventDefault();
    const { name, email } = this.state;
    try {
      let response = await fetch(`${API_URL}/api/v1/crud`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          name,
          email,
        }),
      });
      this.setState({ post: true, postResponse: await response.json() });
    } catch (error) {
      console.error(error);
    }
  }
  async handlePUTSubmit(event) {
    event.preventDefault();
    const { id, name, email } = this.state;
    try {
      let response = await fetch(`${API_URL}/api/v1/crud`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          name,
          email,
          id,
        }),
      });
      this.setState({ put: true, putResponse: await response.json() });
    } catch (error) {
      console.error(error);
    }
  }
  async handleDELETESubmit(event) {
    event.preventDefault();
    const { id } = this.state;
    try {
      let response = await fetch(`${API_URL}/api/v1/crud?id=${id}`, {
        method: "DELETE",
      });
      this.setState({ deletee: true, deleteResponse: await response.json() });
    } catch (error) {
      console.error(error);
    }
  }
  render() {
    let {
      get,
      post,
      put,
      deletee,
      getResponse,
      putResponse,
      postResponse,
      deleteResponse,
    } = this.state;
    return (
      <div>
        <div>
          <h2>READ</h2>
          <form onSubmit={this.handleGETSubmit} onChange={this.handleChange}>
            <label>
              Enter Email:
              <input type="text" name="email" />
            </label>
            <input type="submit" value="Submit" />
          </form>
        </div>
        {get && Object.keys(getResponse).length > 0 ? (
          <table>
            <tbody>
              <tr>
                <td>{getResponse.email}</td>
                <td>{getResponse.name}</td>
              </tr>
            </tbody>
          </table>
        ) : (
          <div></div>
        )}
        <div>
          <h2>CREATE</h2>
          <form onSubmit={this.handlePOSTSubmit} onChange={this.handleChange}>
            <label>
              Enter Email:
              <input type="text" name="email" />
            </label>
            <label>
              Enter Name:
              <input type="text" name="name" />
            </label>
            <input type="hidden" name="type" value="POST" />
            <input type="submit" value="Submit" />
          </form>
          {post && Object.keys(postResponse).length > 0 ? (
            <div>{postResponse.message}</div>
          ) : (
            <div></div>
          )}
        </div>
        <div>
          <h2>UPDATE</h2>
          <form onSubmit={this.handlePUTSubmit} onChange={this.handleChange}>
            <label>
              Enter Email:
              <input type="text" name="email" />
            </label>
            <label>
              Enter Name:
              <input type="text" name="name" />
            </label>
            <label>
              Enter Id:
              <input type="number" name="id" />
            </label>
            <input type="submit" value="Submit" />
          </form>
          {put && Object.keys(putResponse).length > 0 ? (
            <div>{putResponse.message}</div>
          ) : (
            <div></div>
          )}
        </div>
        <div>
          <h2>DELETE</h2>
          <form onSubmit={this.handleDELETESubmit} onChange={this.handleChange}>
            <label>
              Enter Id:
              <input type="text" name="id" />
            </label>
            <input type="submit" value="Submit" />
          </form>
          {deletee && Object.keys(deleteResponse).length > 0 ? (
            <div>{deleteResponse.message}</div>
          ) : (
            <div></div>
          )}
        </div>
      </div>
    );
  }
}

export default CrudFrom;
