import React, { Component } from 'react';
import Button from 'react-bootstrap/lib/Button'
import logo from './logo.svg';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      isLoaded: false,
      token: ''
    };

  }
  getToken() {
    fetch('http://localhost:3000/api/Token')
    .then(res => res.json())
    .then(({token}) => this.setState({
      token,
      isLoaded: true     
    }));
  }

  render() {
    const {isLoaded, token } = this.state;
    if(!isLoaded) {
      return (
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <h1 className="App-title">Welcome to React</h1>
          </header>
          <p className="App-intro">
            To get started, edit <code>src/App.js</code> and save to reload.
          </p>
          <Button onClick={this.getToken}>Generate Token</Button>
          </div>
        );
    } else {   
        return <div>{token}</div> 
    }
  }
}

export default App;
