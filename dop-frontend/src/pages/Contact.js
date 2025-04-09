import React, { useState } from 'react';
import axios from 'axios';
//import ReCAPTCHA from 'react-google-recaptcha'; // Google reCAPTCHA
import '../assets/contact.css';
import whatsappLogo from '../assets/whatsapp-icon.png';

const Contact = () => {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    mobile: '',
    message: '',
  });

  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  // Handle input changes
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  // Form submission
  const handleSubmit = async (e) => {
    e.preventDefault();

    // Basic Validation
    const { name, email, mobile, message } = formData;
    if (!name || !email || !mobile || !message ) {
      setError('Please fill in all the fields and complete CAPTCHA');
      return;
    }

    // Validate Email Format
    const emailPattern = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/;
    if (!email.match(emailPattern)) {
      setError('Please enter a valid email address.');
      return;
    }

    // Validate Mobile Number
    const mobilePattern = /^[0-9]{10,15}$/;
    if (!mobile.match(mobilePattern)) {
      setError('Please enter a valid mobile/WhatsApp number.');
      return;
    }

    try {
      // Send data to backend API (Go + PostgreSQL)
      const response = await axios.post(`${process.env.REACT_APP_API_URL}/api/contact`, formData);


      if (response.data.success) {
        setSuccess('Thank you! Your message has been sent.');
        setFormData({ name: '', email: '', mobile: '', message: '' });
      } else {
        setError('There was an issue submitting your query. Please try again later.');
      }
    } catch (error) {
      setError('An error occurred. Please try again.');
    }
  };

  return (
    <div className="contact-container">
      <h1 className="contact-title">Get in Touch</h1>
      <p className="contact-description">Have a question? Fill out the form below, and we'll get back to you shortly!</p>

      <form className="contact-form" onSubmit={handleSubmit}>
        <input
          type="text"
          name="name"
          placeholder="Your Name"
          value={formData.name}
          onChange={handleChange}
        />
        <input
          type="email"
          name="email"
          placeholder="Your Email Address"
          value={formData.email}
          onChange={handleChange}
        />
        <input
          type="text"
          name="mobile"
          placeholder="Your Mobile/WhatsApp Number"
          value={formData.mobile}
          onChange={handleChange}
        />
        <textarea
          name="message"
          placeholder="Your Message"
          rows="5"
          value={formData.message}
          onChange={handleChange}
        ></textarea>


        {/* Submit Button */}
        <button type="submit">Send Message</button>
      </form>

      {error && <div className="error-message">{error}</div>}
      {success && <div className="success-message">{success}</div>}

      {/* WhatsApp Support */}
      <div className="whatsapp-support">
        <h2>Quick Support on WhatsApp</h2>
        <a
          href={`https://wa.me/${process.env.REACT_APP_WHATSAPP_NUMBER}?text=Hi%2C%20I%20need%20support%20with%20your%20DevOps%20services`}
          target="_blank"
          rel="noopener noreferrer"
          className="whatsapp-link"
        >
          <img src={whatsappLogo} alt="WhatsApp" className="whatsapp-logo" />
          {process.env.REACT_APP_WHATSAPP_NUMBER}
        </a>
      </div>
    </div>
  );
};

export default Contact;
