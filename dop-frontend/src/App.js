//import logo from './logo.svg';
//import './App.css';
//
//function App() {
//  return (
//    <div className="App">
//      <header className="App-header">
//        <img src={logo} className="App-logo" alt="logo" />
//        <p>
//          Edit <code>src/App.js</code> and save to reload.
//        </p>
//        <a
//          className="App-link"
//          href="https://reactjs.org"
//          target="_blank"
//          rel="noopener noreferrer"
//        >
//          Learn React
//        </a>
//      </header>
//    </div>
//  );
//}
//
//export default App;
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
//import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import About from './pages/About';
import Dashboard from './pages/Dashboard';
import Contact from './pages/Contact';
import ServiceDetails from './pages/ServiceDetails';
import Portfolio from './pages/Portfolio';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/dashboard" element={Dashboard} />
        <Route path="/service/:id" element={ServiceDetails} />
        <Route path="/contact" element={<Contact />} />
        <Route path="/portfolio" element={Portfolio} />
      </Routes>
    </Router>
  );
}

export default App;

