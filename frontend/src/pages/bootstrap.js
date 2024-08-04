// src/App.js
import React, { useState, useEffect, spacing } from 'react';
import axios from 'axios';
import { Container, Row, Col, Button } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  const [target, setTarget] = useState('');
  const [domain, setDomain] = useState('');
  const [domains, setDomains] = useState([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    setLoading(true)
    fetch("http://localhost:8000/api/domains/")
      .then(response => response.json())
      .then(json => setDomains(json))
      .finally(() => {
        setLoading(false)
      })
  }, [])

  useEffect(() => {
    document.title = 'Domains';
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const data = { "target": target, "domain": domain };
      const response = await axios.post('http://localhost:8000/api/domain/add', data);
      console.log('Domain saved:', response.data);
    } catch (error) {
      console.error('Error saving domain:', error);
    }
    var allInputs = document.querySelectorAll('input');
    allInputs.forEach(singleInput => singleInput.value = '');
  };

  return (
    <Container className="text-center my-5">
        <Row>
            <Col md={6} className="bg-primary text-white p-4">
                <h2>Column 1</h2>
                <p>This is a visually appealing layout using the grid system approach.</p>
                <Button variant="light">Click me</Button>
            </Col>
            <Col md={6} className="bg-secondary text-white p-4">
                <h2>Column 2</h2>
                <p>Responsive and stunning design to enhance user experience.</p>
                <Button variant="light">Explore</Button>
            </Col>
        </Row>
    </Container>
);

}

export default App;
