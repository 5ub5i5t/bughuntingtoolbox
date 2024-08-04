import React, { useState, useEffect } from 'react';

const Home = () => {

    useEffect(() => {
        document.title = 'Home';
      }, []);

    return (
        <div>
            <h1>Home</h1>
            <p>Test Change 2</p>
        </div>
    );
  };
  
export default Home;