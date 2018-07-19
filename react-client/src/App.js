import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { 
      error: null,isLoaded: true,token: ''};
  }
 
  getToken() {
    fetch('http://localhost:3000/api/Token')
    .then(res => res.json())
    .then((result) => {
      console.log(result);
        this.setState({    isLoaded: false, token : result.token});
    });
  }

  render() {
    const { error, isLoaded, token } = this.state;
    if(isLoaded)
    {
      return (
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <h1 className="App-title">Welcome to React</h1>
          </header>
          <p className="App-intro">
            To get started, edit <code>src/App.js</code> and save to reload.
          </p>
          <button onClick={this.getToken}>Generate Token</button>
          </div>
        );
        } else {   
        if(!isLoaded){
        <div>{this.token}</div>
      }  
  }
}
}

export default App;
