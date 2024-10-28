import { useEffect, useState } from "react";
import ButtonAction from "../components/buttons/ButtonAction";
import HeaderLogo from "../components/headers/HeaderLogo";
import Itinerary from "../components/itineraries/Itinerary";
import Scroll from "../layouts/Scroll";
import { FaPlus } from "react-icons/fa";
import { useLocation, useNavigate } from "react-router-dom";

function ItineraryPage() {
  const [modal, setModal] = useState(false);
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);
  const [showDropdown, setShowDropdown] = useState(false);
  const [id, setId] = useState(null);
  const navigate = useNavigate();
  
  
  const { state: {data, activity, trip, destination} } = useLocation();
  const [itinerary, setItinerary] = useState(data);

  // useEffect(() => {
  //   if (id) {
  //     // find selected id
  //     let item = null;
  //     let i = 0;
  //     while (item === null) {
  //       const dest = itinerary[i].Destinations.find(v => v.ID === id);
  //       if (dest) {
  //         item = dest;
  //         itinerary[i].Destinations = itinerary[i].Destinations.filter(v => v.ID !== dest.ID);
  //         console.log(itinerary);
  //         // const newDest = itinerary[i].Destinations.filter(v => v.ID !== dest.ID);
  //         // console.log(newDest);
  //         // delete itinerary[i];
  //         setItinerary(itinerary);
  //         // itinerary[i].Destinations.splice(item, 1);
  //       } else {
  //         i++;
  //       };
  //     };
  //   }
  // }, [id]);

  const handleSearch = async (e) => {
    const searchTerm = e.target.value;
    setQuery(searchTerm);
    setShowDropdown(true);

    if (searchTerm.length > 2) {
      try {
        const response = await fetch(
          `http://localhost:8000/api/destinations?search=${searchTerm}`
        );
        const data = await response.json();
        console.log(data);
        setResults(data.slice(0, 4));
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    } else {
      setResults([]);
    }
  };

  function formatTitle(title) {
    const len = 25;
    if (title.length > len) {
      return title.slice(0, len) + "...";
    }
    return title;
  }

  const handleSelect = (item) => {
    // setQuery(item);
    console.log(item);
    // navigate("/preference/"+item)
    // setShowDropdown(false);
  };

  const nextButton = () => {
    navigate("/buddy", {state: {
      itinerary: data,
      destination: destination || "",
      activity: activity || "",
      trip,
    }})
  }

  return (
    <Scroll>
      <HeaderLogo />

      <div className="px-6 mt-2 mb-4">
        <div className="flex items-center justify-between mb-4">
          <h1 className="text-2xl font-bold">Itinerary</h1>
          <button
            className="flex items-center justify-center w-8 h-8 rounded-full bg-blue"
            onClick={() => setModal(true)}
          >
            <FaPlus className="w-5 h-5 text-white" />
          </button>
        </div>

        <Itinerary data={itinerary} setId={setId} />

        <div className="flex items-center justify-center w-full h-full pt-4 pb-8">
          <ButtonAction text="Next" onClick={nextButton} />
        </div>
      </div>

      {modal && (
        <div className="absolute top-0 z-10 h-full w-96 bg-text/45">
          <div className="flex items-center justify-center w-full h-full">
            <div className="flex h-auto w-5/6 flex-col rounded bg-white p-3 gap-2.5">
              <button
                className="self-end text-sm"
                onClick={() => setModal(false)}
              >
                x
              </button>
              <h5 className="text-sm font-bold">
                Which destination would you like to add?
              </h5>
              <div className="flex flex-col gap-2">
                <select
                  name="time"
                  id="time"
                  className="p-1 rounded shadow-md bg-yellow"
                >
                  <option value="morning">Morning</option>
                  <option value="afternoon">Afternoon</option>
                  <option value="evening">Evening</option>
                  <option value="night">Night</option>
                </select>
                <input
                  type="text"
                  placeholder="Insert something..."
                  value={query}
                  onChange={handleSearch}
                  className="px-2 py-1 shadow-md"
                />
              </div>
              {showDropdown && results.length > 0 && (
                <div>
                  {results.map((item) => {
                    return (
                      <div
                        className="flex flex-grow gap-3 p-3 rounded shadow-md"
                        key={item.ID}
                        onClick={() => handleSelect(item)}
                      >
                        <div className="w-16 h-16 rounded bg-blue"></div>
                        <div className="flex flex-col gap-1">
                          <h5 className="text-sm font-bold">
                            {formatTitle(item.name)}
                          </h5>
                          <div className="flex gap-2">
                            <div className="inline-block w-4 h-4 bg-yellow"></div>
                            {item.star}
                          </div>
                          <div className="flex gap-2">
                            <div className="inline-block w-4 h-4 bg-yellow"></div>
                            Singapore
                          </div>
                        </div>
                      </div>
                    );
                  })}
                </div>
              )}
            </div>
          </div>
        </div>
      )}
    </Scroll>
  );
}

export default ItineraryPage;
