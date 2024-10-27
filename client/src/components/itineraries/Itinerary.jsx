import Destination from "./Destination";

function Itinerary() {
  return (
    <div>
      <div>
        <h5 className="text-sm font-bold">Morning</h5>
        <Destination></Destination>
      </div>

      <div>
        <h5 className="text-sm font-bold">Afternoon</h5>
        <Destination></Destination>
      </div>
    </div>
  );
}

export default Itinerary;
