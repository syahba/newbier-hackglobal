import CardAction from "./CardAction";
import CardDefault from "./CardDefault";

function Destination({ data, setId, withoutdelete }) {
  return (
    <div className="flex flex-col gap-4 ml-2 border-l-2 border-blue">
      {data.time === "morning" && (
        <div className="relative pl-2">
          <CardDefault />
          {/* <p className="my-3 text-center">300 meters - 20 minutes</p> */}
        </div>
      )}
      {data.Destinations.map((elm, idx) => {
        return (
          <div className="relative pl-2" key={elm.name}>
            <CardAction data={elm} setId={setId} withOutDeleted={withoutdelete} />
            {/* <p className="my-3 text-center">300 meters - 20 minutes</p> */}
          </div>
        );
      })}
    </div>
  );
}

export default Destination;
