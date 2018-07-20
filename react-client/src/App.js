import React, { Component } from 'react';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      isLoaded: false,
      user: '',
      password: '',
      token: ''
    };

    this.handleInputChange = this.handleInputChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleInputChange(event) {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    this.setState({
      [name]: value
    });
  }

  handleSubmit(event) {
    fetch('http://localhost:3000/api/token', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        user: this.state.user,
        password: this.state.password
      })
    }).then((responseText) => responseText.json())
    .then((response) => this.setState({
      token: response.token
    }));
    
    event.preventDefault();
  }

  render() {
    const {isLoaded, token } = this.state;
    if(!isLoaded) {
      return (
        <div className="App">
          <h1>MascheDevTools!!!!!!!</h1>
          <img src="img/logo.png" alt="MascheDevTools"/>
          <br/>
          <form onSubmit={this.handleSubmit}>
            <label>
              Usuario: 
              <input type="text" name="user" value={this.state.user} onChange={this.handleInputChange} />
            </label><br/>
            <label>
              Contraseña: 
              <input type="password" name="password" value={this.state.password} onChange={this.handleInputChange} />
            </label><br/>
            <input type="submit" value="Submit" /><br/>
            <label>{this.state.token}</label>
            </form>
        </div>
        );
    } else {   
        return <div>{token}</div> 
    }
  }
}

export default App;
