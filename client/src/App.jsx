import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

import NotFound from './pages/NotFound';
import Index from './pages/Index';
import Home from './pages/Home';
import CardProfile from './components/forms/CardProfile';
import Preference from './pages/Preference';
import Preference2 from './pages/Preference2';
import Preference3 from './pages/preference-3';

function App() {

  return (
    <Router>
      <Routes>
        <Route path="/" element={<Index />} />
        <Route path="/test" element={<CardProfile />} />
        <Route path="/home" element={<Home />} />
        <Route path="/preference/:id" element={<Preference />} />
        <Route path="/preference" element={<Preference2 />} />
        <Route path="/preference-buddy" element={<Preference3 />} />
        {/* Route 404 untuk halaman yang tidak ditemukan */}
        <Route path="*" element={<NotFound />} />
      </Routes>
    </Router>
  )
}

export default App
