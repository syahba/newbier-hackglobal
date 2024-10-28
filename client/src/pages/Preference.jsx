import { useState, useEffect, useCallback } from "react";
import ButtonAction from "../components/buttons/ButtonAction";
import FieldDate from "../components/forms/FieldDate";
import FieldStyle from "../components/forms/FieldStyle";
import HeaderLogo from "../components/headers/HeaderLogo";
import Main from "../layouts/Main";
import { useNavigate, useParams } from "react-router-dom";

function Preference() {
  const { id } = useParams();
  const [result, setResult] = useState({});
  const [name, setName] = useState("");
  const [trip, setTrip] = useState("");

  const navigate = useNavigate();

  const wrapperSetTrip = useCallback(
    (val) => {
      setTrip(val);
    },
    [setTrip]
  );

  useEffect(() => {
    try {
      const fetchData = async () => {
        try {
          const response = await fetch(
            `${import.meta.env.VITE_REACT_API_URL}/api/destinations/${id}`
          );
          const data = await response.json();
          setResult(data);
          setName(data.name);
        } catch (error) {
          console.log(error);
        }
      };

      fetchData();
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }, []);

  const action = () => {
    const fetchData = async () => {
      try {
        const response = await fetch(
          `${import.meta.env.VITE_REACT_API_URL}/api/generate-itinerary/destination?destination=${name}&trip=${trip}`
        );
        const data = await response.json();
        console.log(data);
        navigate("/itinerary", {
          state: {
            data: data,
            destination: name,
            trip: trip,
          },
        });
      } catch (err) {
        console.log(err);
      }
    };
    fetchData();
  };

  return (
    <Main>
      <div className="w-full bg-center bg-no-repeat bg-contain bg-preference1 h-3/4">
        <HeaderLogo />
        <div className="px-6 mt-12 mb-4">
          <h5 className="mb-4 text-sm font-bold">
            Hi, <span className="text-blue">User</span>! Where do you want to go
            today?
          </h5>
          <input
            type="text"
            className="w-full h-8 px-3 mb-4 rounded-lg shadow-md"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Enter a destination"
          />
        </div>
      </div>

      <div>
        <div className="absolute bottom-0 flex h-1/2 w-96 flex-col gap-8 rounded-t-3xl bg-white px-6 pt-5 text-sm shadow-[0_-4px_20px_rgba(0,0,0,0.1)]">
          <FieldDate />

          <FieldStyle parentStateSetter={wrapperSetTrip} />

          <div className="self-center">
            <ButtonAction text="Confirm" onClick={() => action()} />
          </div>
        </div>

      </div>
    </Main>
  );
}

export default Preference;
