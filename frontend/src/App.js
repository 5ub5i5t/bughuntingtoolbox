import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import Home from './pages';
import Proxy from './pages/proxy';
import Target from './pages/targets';
import Domain from './pages/domains';


function App() {
  return (
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/proxy" element={<Proxy />} />
        <Route path="/targets" element={<Target />} />
        <Route path="/domains" element={<Domain />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;