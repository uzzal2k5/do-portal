import React, { useState, useEffect } from 'react';
import axios from 'axios';

const ServiceDetails = ({ match }) => {
  const [service, setService] = useState(null);

  useEffect(() => {
    axios.get(`http://localhost:8080/api/services/${match.params.id}`)
      .then(response => setService(response.data))
      .catch(error => console.error(error));
  }, [match.params.id]);

  if (!service) return <div>Loading...</div>;

  return (
    <div>
      <h1>{service.name}</h1>
      <p>{service.status}</p>
      <p>{service.logs}</p>
    </div>
  );
}

export default ServiceDetails;
