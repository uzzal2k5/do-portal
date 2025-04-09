// src/pages/Portfolio.js

import React, { useEffect, useState } from 'react';
import axios from 'axios';
import '../assets/portfolio.css'; // Add your custom styling here

const Portfolio = () => {
  const [portfolio, setPortfolio] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  // Fetch portfolio data
  useEffect(() => {
    axios.get('http://localhost:9090/api/portfolio')
      .then(response => {
        setPortfolio(response.data);
        setLoading(false);
      })
      .catch(error => {
        setError('Error fetching portfolio content');
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div className="portfolio-container">
      <h1 className="portfolio-title">{portfolio.title}</h1>
      <p className="portfolio-description">{portfolio.description}</p>

      <div className="skills">
        <h2>Skills:</h2>
        <p>{portfolio.skill_set}</p>
      </div>

      <div className="experience">
        <h2>Experience:</h2>
        <p>{portfolio.experience}</p>
      </div>
    </div>
  );
};

export default Portfolio;
