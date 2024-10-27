import { useState } from "react";
import Main from "../layouts/Main"
import { useNavigate } from "react-router-dom";

function Home() {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]);
  const [showDropdown, setShowDropdown] = useState(false);
  const navigate = useNavigate()

  const handleSearch = async (e) => {
    const searchTerm = e.target.value;
    setQuery(searchTerm);
    setShowDropdown(true);

    if (searchTerm.length > 2) { 
      try {
        const response = await fetch(`http://localhost:8000/api/destinations?search=${searchTerm}`);
        const data = await response.json();
        console.log(data)
        setResults(data);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    } else {
      setResults([]);
    }
  };

  const handleSelect = (item) => {
    // setQuery(item); 
    navigate("/")
    // setShowDropdown(false);
  };

  return (
    <Main>
    <div className="px-6 mt-12 mb-4">
      <h1 className="text-2xl font-bold">Hi, <span className="text-blue">User</span>!</h1>
      <h5 className="mb-4 text-sm font-bold">Where would you like to go today?</h5>
      <div className="relative ">
        <input type="text"  className="mb-4 h-8 w-full px-3 shadow-md" value={query} onChange={handleSearch} placeholder="Enter a destination" />
        {
          showDropdown && results.length > 0 && (
            <div className="absolute top-10 bg-white p-2 rounded border w-full ">
              {
                results.map(item => (
                  <input type="button" value={item.name} key={item.ID} className="py-2 px-1 cursor-pointer w-full text-left" onClick={() => handleSelect(item.name)} />
                ))
              }
            </div>
          )
        }
      </div>

      <div className="flex w-full gap-4">
        <button className="rounded border border-blue px-5 py-2 shadow-md">I’m not sure</button>
        <button className="rounded border flex-auto border-yellow px-5 py-2 shadow-md">I’d like to travel with a buddy</button>
      </div>
    </div>
    <div className="absolute bottom-0 w-96 h-[449px]">
      <div className="bg-home h-full w-96 bg-contain bg-center bg-no-repeat"></div>
    </div>
    </Main>
  )
}

export default Home