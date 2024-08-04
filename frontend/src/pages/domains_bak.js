// src/App.js
import React, { useState, useEffect, spacing } from 'react';
import axios from 'axios';
import { Container, Row, Col, Button, Table } from 'react-bootstrap';
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
    <div style={{padding: '1em'}}>
      <div>
      <p><strong><u>Add new domain:</u></strong></p>
      <form id="domain_form" name="domain_form" onSubmit={handleSubmit}>
        <input
            type="text"
            id="target"
            name="target"
            value={target}
            onChange={(e) => setTarget(e.target.value)}
            placeholder="Enter target name"
            required
          />
        <input
          type="text"
          id="domain"
          name="domain"
          value={domain}
          onChange={(e) => setDomain(e.target.value)}
          placeholder="Enter domain name"
          required
        />
        <button type="submit">Save</button>
      </form>
      </div>
      <div>
        {loading ? (
          <div>Loading...</div>
        ) : (
          <>

            <hr/>
            
            <Table striped bordered>
              <tr>
                <th>Target</th>
                <th>Domain</th>
              </tr>
              {domains.map(domain => (
                <tr key={domain.id}>
                  <td>{domain.target}</td>
                  <td>{domain.domain}</td>
                </tr>
              ))}
            </Table>
          </>
        )}
      </div>
    </div>
);

}

export default App;
