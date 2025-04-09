import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Dashboard = () => {
  const [services, setServices] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8080/api/services')
      .then(response => setServices(response.data))
      .catch(error => console.error(error));
  }, []);

  return (
    <div>
      <h1>DevOps Service Dashboard</h1>
      <ul>
        {services.map(service => (
          <li key={service.id}>
            <a href={`/service/${service.id}`}>{service.name}</a>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Dashboard;
