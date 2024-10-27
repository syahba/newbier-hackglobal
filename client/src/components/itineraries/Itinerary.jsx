import Destination from "./Destination";

function Itinerary({data}) {
  
  return (
    <div>
      {
        data.map(elm => (
          <div className="mt-3" key={elm.time}>
            <h5 className="text-sm font-bold">{elm.time}</h5>
            <Destination data={elm} />
          </div>
        ))
      }
    </div>
  );
}

export default Itinerary;
