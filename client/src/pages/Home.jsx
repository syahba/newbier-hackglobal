import { useState } from "react";
import Main from "../layouts/Main";
import { useNavigate } from "react-router-dom";
import HeaderDefault from "../components/headers/HeaderDeafult";

function Home() {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);
  const [showDropdown, setShowDropdown] = useState(false);
  const navigate = useNavigate();

  const handleSearch = async (e) => {
    const searchTerm = e.target.value;
    setQuery(searchTerm);
    setShowDropdown(true);

    if (searchTerm.length > 2) {
      try {
        const response = await fetch(`http://localhost:8000/api/destinations?search=${searchTerm}`);
        const data = await response.json();        
        setResults(data.slice(0, 4));
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    } else {
      setResults([]);
    }
  };

  const handleSelect = (item) => {
    // setQuery(item);
    navigate("/preference/" + item);
    // setShowDropdown(false);
  };

  return (
    <Main>
      <HeaderDefault></HeaderDefault>

      <div className="px-6 mt-12 mb-4">
        <h1 className="text-2xl font-bold">
          Hi, <span className="text-blue">User</span>!
        </h1>
        <h5 className="mb-4 text-sm font-bold">
          Where would you like to go today?
        </h5>
        <div className="relative ">
          <input
            type="text"
            className="w-full h-8 px-3 mb-4 shadow-md rounded-lg"
            value={query}
            onChange={handleSearch}
            placeholder="Enter a destination"
          />
          {showDropdown && results.length > 0 && (
            <div className="absolute w-full p-2 bg-white border rounded top-10 ">
              {results.map((item) => (
                <input
                  type="button"
                  value={item.name}
                  key={item.ID}
                  className="w-full px-1 py-2 text-left cursor-pointer"
                  onClick={() => handleSelect(item.ID)}
                />
              ))}
            </div>
          )}
        </div>

        <div className="flex w-full gap-4">
          <button
            className="px-5 py-2 border rounded shadow-md border-blue"
            onClick={() => navigate("/bridging/1")}
          >
            I’m not sure
          </button>
          <button
            className="flex-auto px-5 py-2 border rounded shadow-md border-yellow"
            onClick={() => navigate("/bridging/2")}
          >
            I’d like to travel with a buddy
          </button>
        </div>
      </div>
      
      <div className="absolute bottom-0 w-96 h-[449px]">
        <div className="h-full bg-center bg-no-repeat bg-contain bg-home w-96"></div>
      </div>
    </Main>
  );
}

export default Home;
