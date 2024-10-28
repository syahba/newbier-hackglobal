import Destination from "./Destination";

function Itinerary({ data, setId }) {
  return (
    <div>
      {data.map((elm) => (
        <div className="mt-3" key={elm.time}>
          <h5 className="text-base font-bold mb-2">
            {String(elm.time).charAt(0).toUpperCase() + (elm.time).slice(1)}
          </h5>
          <Destination data={elm} setId={setId} />
        </div>
      ))}
    </div>
  );
}

export default Itinerary;
