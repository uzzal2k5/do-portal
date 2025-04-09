import React from 'react';
import { Link } from 'react-router-dom';
import '../assets/home.css'; // 👈 import from assets

const Home = () => {
  return (
  <div className="home-container">
      <h1 className="home-heading">Welcome to DOP Services Portal</h1>
      <p className="home-subtext">
         We provide DevOps & Platform Engineering services including CI/CD, Kubernetes, Infrastructure Automation,
         Monitoring, and more.
      </p>
      <div className="home-nav-links">
        <Link to="/about" className="home-button">Who We Do !</Link>
        <Link to="/services" className="home-button">Services </Link>
        <Link to="/dashboard" className="home-button">Dashboard</Link>
        <Link to="/contact" className="home-button-outline">Contact Us</Link>
      </div>
  </div>
  );
};

export default Home;

