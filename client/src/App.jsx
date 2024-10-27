import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

import NotFound from './pages/NotFound';
import Index from './pages/Index';
import Home from './pages/Home';
import CardProfile from './components/forms/CardProfile';
import Preference from './pages/Preference';
import Preference2 from './pages/Preference2';
import Preference3 from './pages/Preference3';
import Bridging1 from './pages/Bridging1';
import Bridging3 from './pages/Bridging3';
import Bridging2 from './pages/Bridging2';
import ItineraryPage from './pages/Itinerary';
import Market from './pages/Market';
import Transaction from './pages/Transaction';

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
        <Route path="/bridging/1" element={<Bridging1 />} />
        <Route path="/bridging/2" element={<Bridging2 />} />
        <Route path="/bridging/3" element={<Bridging3 />} />
        <Route path="/itinerary" element={<ItineraryPage />} />
        <Route path="/market" element={<Market />} />
        <Route path="/transaction" element={<Transaction />} />
        {/* Route 404 untuk halaman yang tidak ditemukan */}
        <Route path="*" element={<NotFound />} />
      </Routes>
    </Router>
  )
}

export default App
