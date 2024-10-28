import { FaStar } from "react-icons/fa";
import { FaLocationDot } from "react-icons/fa6";

function CardDefault({ data = { name: 'Lyf Funan', star: '4.9' } }) {
  function formatTitle(title) {
    const len = 20;
    if (title.length > len) {
      return title.slice(0, len) + "...";
    }
    return title;
  };

  return (
    <div className="flex items-center gap-3 pl-2">
      <div className="absolute -left-1.5 h-2.5 w-2.5 rounded-full bg-blue"></div>
      <div className="flex flex-grow gap-3 p-3 rounded-lg shadow-lg">
        <div className="w-16 h-16 rounded-lg bg-blue"></div>
        <div className="flex flex-col gap-1">
          <h5 className="text-sm font-bold">{formatTitle(data.name)}</h5>
          <div className="flex gap-2 mb-1">
            <FaStar className="w-4 h-4 text-yellow" />
            {data.star}
          </div>
          <div className="flex gap-2">
            <FaLocationDot className="w-4 h-4 text-yellow" />
            Singapore
          </div>
        </div>
      </div>
    </div>
  );
}

export default CardDefault;
