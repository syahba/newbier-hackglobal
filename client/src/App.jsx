import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

import NotFound from './pages/NotFound';
import Index from './pages/Index';
import Home from './pages/Home';
import Main from './layouts/Main';
import CardProfile from './components/forms/CardProfile';

function App() {

  return (
    <Router>
      <Routes>
        <Route path="/" element={<Index />} />
        <Route path="/test" element={<CardProfile />} />
        <Route path="/home" element={<Home />} />
        {/* Route 404 untuk halaman yang tidak ditemukan */}
        <Route path="*" element={<NotFound />} />
      </Routes>
    </Router>
  )
}

export default App
