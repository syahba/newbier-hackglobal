import Destination from "./Destination";

function Itinerary({ data, setId, withoutdelete }) {
  return (
    <div>
      {data.map((elm) => (
        <div className="mt-3" key={elm.time}>
          <h5 className="mb-2 text-base font-bold">
            {String(elm.time).charAt(0).toUpperCase() + (elm.time).slice(1)}
          </h5>
          <Destination data={elm} setId={setId} withoutdelete={withoutdelete} />
        </div>
      ))}
    </div>
  );
}

export default Itinerary;
