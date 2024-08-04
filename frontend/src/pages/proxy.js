import React, { useState, useEffect } from 'react';
import axios from 'axios';
import '../config';

const Proxy = () => {

  useEffect(() => {
    document.title = 'Proxy';
  }, []);

  const handleFetch = () => {
      var element = document.getElementById("proxy");
      element.style.display = "none";
      axios.get('http://' + global.config.config.base.local + '/api/proxy/start')
          .then(() => {})
          .catch(error => console.error('Error:', error));
  };

  return (
    <div id="proxy" name="proxy">
      <button onClick={handleFetch}>Start Proxy</button>
    </div>
  );
};

export default Proxy;